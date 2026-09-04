//go:build windows

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"testing"
)

func skipGIFSubBlocks(data []byte, off int) (int, error) {
	for {
		if off >= len(data) {
			return off, fmt.Errorf("unexpected EOF at sub-block size")
		}
		n := int(data[off])
		off++
		if n == 0 {
			return off, nil
		}
		if off+n > len(data) {
			return off, fmt.Errorf("sub-block length %d exceeds data", n)
		}
		off += n
	}
}

func scanGIFStructure(data []byte) (int, error) {
	if len(data) < 13 || (string(data[:6]) != "GIF87a" && string(data[:6]) != "GIF89a") {
		return 0, fmt.Errorf("invalid GIF header")
	}
	off := 13
	packed := data[10]
	if packed&0x80 != 0 {
		off += 3 * (1 << ((packed & 0x07) + 1))
	}
	if off > len(data) {
		return off, fmt.Errorf("global color table exceeds data")
	}
	for off < len(data) {
		blockStart := off
		switch data[off] {
		case 0x3B:
			return off + 1, nil
		case 0x21:
			off += 2 // extension introducer + label
			var err error
			off, err = skipGIFSubBlocks(data, off)
			if err != nil {
				return blockStart, err
			}
		case 0x2C:
			if off+10 > len(data) {
				return blockStart, fmt.Errorf("short image descriptor")
			}
			localPacked := data[off+9]
			off += 10
			if localPacked&0x80 != 0 {
				off += 3 * (1 << ((localPacked & 0x07) + 1))
			}
			if off >= len(data) {
				return blockStart, fmt.Errorf("missing LZW code size")
			}
			off++ // LZW minimum code size
			var err error
			off, err = skipGIFSubBlocks(data, off)
			if err != nil {
				return blockStart, err
			}
		default:
			return blockStart, fmt.Errorf("unknown block 0x%02x", data[off])
		}
	}
	return off, fmt.Errorf("missing GIF trailer")
}

func TestLocateAndRecoverSpriteCharacter(t *testing.T) {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	before := spriteData0 + spriteData1 + spriteData2 + spriteData3 + spriteData4
	after := spriteData6 + spriteData7

	// Put a placeholder at the end of the short chunk. This restores base64
	// alignment after spriteData5; the GIF structural failure then lands close
	// to the actual missing character inside spriteData5.
	baseline := before + spriteData5 + "A" + after
	baselineBytes, err := base64.StdEncoding.DecodeString(baseline)
	if err != nil {
		t.Fatalf("decode baseline: %v", err)
	}
	failureByte, scanErr := scanGIFStructure(baselineBytes)
	if scanErr == nil {
		t.Fatalf("baseline unexpectedly has valid GIF structure")
	}
	approxGlobalChar := failureByte * 4 / 3
	approx := approxGlobalChar - len(before)
	start := approx - 600
	end := approx + 100
	if start < 0 {
		start = 0
	}
	if end > len(spriteData5) {
		end = len(spriteData5)
	}
	t.Logf("baseline structure failed near byte=%d (%v); searching spriteData5 positions %d..%d", failureByte, scanErr, start, end)

	suffixFrom5 := spriteData5 + after
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(before)+len(spriteData5)+1+len(after)))
	var matches []string
	for pos := start; pos <= end; pos++ {
		candidate := make([]byte, 0, len(before)+len(spriteData5)+1+len(after))
		candidate = append(candidate, before...)
		candidate = append(candidate, spriteData5[:pos]...)
		insertIndex := len(candidate)
		candidate = append(candidate, 'A')
		candidate = append(candidate, suffixFrom5[pos:]...)
		for i := 0; i < len(alphabet); i++ {
			candidate[insertIndex] = alphabet[i]
			n, err := base64.StdEncoding.Decode(decoded, candidate)
			if err != nil {
				continue
			}
			if _, err := scanGIFStructure(decoded[:n]); err != nil {
				continue
			}
			img, format, err := image.Decode(bytes.NewReader(decoded[:n]))
			if err != nil || format != "gif" {
				continue
			}
			if img.Bounds().Dx() != spriteSourceSize*spriteColumns || img.Bounds().Dy() != spriteSourceSize*10 {
				continue
			}
			match := fmt.Sprintf("pos=%d char=%c", pos, alphabet[i])
			matches = append(matches, match)
			t.Logf("valid recovered candidate: %s", match)
		}
	}
	if len(matches) == 0 {
		t.Fatalf("no valid recovery candidates found in search window")
	}
	t.Logf("recovery candidates: %v", matches)
}
