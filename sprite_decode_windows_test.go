//go:build windows

package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"testing"
)

func TestEmbeddedSpriteSheetDecodes(t *testing.T) {
	data, err := base64.StdEncoding.DecodeString(spriteAssetBase64)
	if err != nil {
		t.Fatalf("decode embedded sprite data: %v", err)
	}
	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("decode embedded sprite image: %v", err)
	}
	if format != "gif" {
		t.Fatalf("unexpected sprite format %q; want gif", format)
	}
	wantW := spriteSourceSize * spriteColumns
	wantH := spriteSourceSize * 10
	if img.Bounds().Dx() != wantW || img.Bounds().Dy() != wantH {
		t.Fatalf("unexpected sprite size %dx%d; want %dx%d", img.Bounds().Dx(), img.Bounds().Dy(), wantW, wantH)
	}
}
