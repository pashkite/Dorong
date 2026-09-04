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
	if format != "png" {
		t.Fatalf("unexpected sprite format %q; want png", format)
	}
	if img.Bounds().Dx() != spriteSourceSize*spriteColumns {
		t.Fatalf("unexpected sprite width %d", img.Bounds().Dx())
	}
}
