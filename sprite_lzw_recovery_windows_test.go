//go:build windows

package main

import (
	"bytes"
	"compress/lzw"
	"encoding/base64"
	"fmt"
	"io"
	"testing"
)

type oneByteCountingReader struct {
	data []byte
	pos  int
}

func (r *oneByteCountingReader) Read(p []byte) (int, error) {
	if r.pos >= len(r.data) {
		return 0, io.EOF
	}
	p[0] = r.data[r.pos]
	r.pos++
	return 1, nil
}

func extractGIFImageData(data []byte) (litWidth int, compressed []byte, sourceOffsets []int, err error) {
	if len(data) < 13 || (string(data[:6]) != "GIF87a" && string(data[:6]) != "GIF89a") {
		return 0, nil, nil, fmt.Errorf("invalid GIF header")
	}
	off := 13
	packed := data[10]
	if packed&0x80 != 0 {
		off += 3 * (1 << ((packed & 0x07) + 1))
	}
	for off < len(data) {
		switch data[off] {
		case 0x21:
			off += 2
			for {
				if off >= len(data) {
					return 0, nil, nil, io.ErrUnexpectedEOF
				}
				n := int(data[off])
				off++
				if n == 0 {
					break
				}
				if off+n > len(data) {
					return 0, nil, nil, io.ErrUnexpectedEOF
				}
				off += n
			}
		case 0x2C:
			if off+10 > len(data) {
				return 0, nil, nil, io.ErrUnexpectedEOF
			}
			localPacked := data[off+9]
			off += 10
			if localPacked&0x80 != 0 {
				off += 3 * (1 << ((localPacked & 0x07) + 1))
			}
			if off >= len(data) {
				return 0, nil, nil, io.ErrUnexpectedEOF
			}
			litWidth = int(data[off])
			off++
			for {
				if off >= len(data) {
					return 0, nil, nil, io.ErrUnexpectedEOF
				}
				n := int(data[off])
				off++
				if n == 0 {
					return litWidth, compressed, sourceOffsets, nil
				}
				if off+n > len(data) {
					return 0, nil, nil, io.ErrUnexpectedEOF
				}
				for i := 0; i < n; i++ {
					compressed = append(compressed, data[off+i])
					sourceOffsets = append(sourceOffsets, off+i)
				}
				off += n
			}
		case 0x3B:
			return 0, nil, nil, fmt.Errorf("GIF has no image")
		default:
			return 0, nil, nil, fmt.Errorf("unknown GIF block 0x%02x at %d", data[off], off)
		}
	}
	return 0, nil, nil, io.ErrUnexpectedEOF
}

func decodeGIFLZW(data []byte) (consumed int, pixels int64, err error) {
	litWidth, compressed, _, err := extractGIFImageData(data)
	if err != nil {
		return 0, 0, err
	}
	r := &oneByteCountingReader{data: compressed}
	zr := lzw.NewReader(r, lzw.LSB, litWidth)
	n, copyErr := io.Copy(io.Discard, zr)
	closeErr := zr.Close()
	if copyErr != nil {
		return r.pos, n, copyErr
	}
	if closeErr != nil {
		return r.pos, n, closeErr
	}
	return r.pos, n, nil
}

func TestRecoverSpriteViaLZW(t *testing.T) {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	before := spriteData0 + spriteData1 + spriteData2 + spriteData3 + spriteData4
	after := spriteData6 + spriteData7

	baseline := before + spriteData5 + "A" + after
	baselineBytes, err := base64.StdEncoding.DecodeString(baseline)
	if err != nil {
		t.Fatalf("decode baseline: %v", err)
	}
	litWidth, compressed, sourceOffsets, err := extractGIFImageData(baselineBytes)
	if err != nil {
		t.Fatalf("extract baseline GIF data: %v", err)
	}
	r := &oneByteCountingReader{data: compressed}
	zr := lzw.NewReader(r, lzw.LSB, litWidth)
	baselinePixels, baselineErr := io.Copy(io.Discard, zr)
	_ = zr.Close()
	if baselineErr == nil {
		t.Fatalf("baseline unexpectedly LZW-decodes (%d pixels)", baselinePixels)
	}
	consumedIndex := r.pos - 1
	if consumedIndex < 0 || consumedIndex >= len(sourceOffsets) {
		t.Fatalf("invalid consumed index %d", consumedIndex)
	}
	failureSourceByte := sourceOffsets[consumedIndex]
	approxGlobalChar := failureSourceByte * 4 / 3
	approx := approxGlobalChar - len(before)
	start := approx - 1200
	end := approx + 200
	if start < 0 {
		start = 0
	}
	if end > len(spriteData5) {
		end = len(spriteData5)
	}
	t.Logf("baseline LZW failed after compressed=%d sourceByte=%d err=%v; searching chunk5 positions %d..%d", r.pos, failureSourceByte, baselineErr, start, end)

	wantPixels := int64(spriteSourceSize * spriteColumns * spriteSourceSize * 10)
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(before)+len(spriteData5)+1+len(after)))
	var matches []string
	for pos := start; pos <= end; pos++ {
		candidate := make([]byte, 0, len(before)+len(spriteData5)+1+len(after))
		candidate = append(candidate, before...)
		candidate = append(candidate, spriteData5[:pos]...)
		insertIndex := len(candidate)
		candidate = append(candidate, 'A')
		candidate = append(candidate, spriteData5[pos:]...)
		candidate = append(candidate, after...)
		for i := 0; i < len(alphabet); i++ {
			candidate[insertIndex] = alphabet[i]
			n, err := base64.StdEncoding.Decode(decoded, candidate)
			if err != nil {
				continue
			}
			_, pixels, err := decodeGIFLZW(decoded[:n])
			if err != nil || pixels != wantPixels {
				continue
			}
			match := fmt.Sprintf("pos=%d char=%c", pos, alphabet[i])
			matches = append(matches, match)
			t.Logf("valid LZW recovery candidate: %s", match)
		}
	}
	if len(matches) == 0 {
		t.Fatalf("no LZW recovery candidates found")
	}
	t.Logf("LZW recovery candidates: %v", matches)
	_ = bytes.MinRead // keep bytes import stable if helpers evolve
}
