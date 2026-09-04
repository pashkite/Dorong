//go:build windows

package main

import (
	"compress/lzw"
	"encoding/base64"
	"io"
	"testing"
)

func TestReportSpriteLZWCorruptionOffset(t *testing.T) {
	before := spriteData0 + spriteData1 + spriteData2 + spriteData3 + spriteData4
	baseline := before + spriteData5 + "A" + spriteData6 + spriteData7
	data, err := base64.StdEncoding.DecodeString(baseline)
	if err != nil {
		t.Fatal(err)
	}
	litWidth, compressed, sourceOffsets, err := extractGIFImageData(data)
	if err != nil {
		t.Fatal(err)
	}
	r := &oneByteCountingReader{data: compressed}
	zr := lzw.NewReader(r, lzw.LSB, litWidth)
	pixels, decodeErr := io.Copy(io.Discard, zr)
	_ = zr.Close()
	idx := r.pos - 1
	if idx < 0 || idx >= len(sourceOffsets) {
		t.Fatalf("bad compressed offset %d", idx)
	}
	sourceByte := sourceOffsets[idx]
	approx := sourceByte*4/3 - len(before)
	t.Fatalf("diagnostic: pixelsBeforeFailure=%d compressedBytes=%d sourceByte=%d approximateChunk5Char=%d decodeErr=%v", pixels, r.pos, sourceByte, approx, decodeErr)
}
