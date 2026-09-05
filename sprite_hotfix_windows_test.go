//go:build windows

package main

import (
	"bytes"
	"image/gif"
	"testing"
)

func TestFallbackSpriteSheetDecodes(t *testing.T) {
	data, err := buildFallbackSpriteSheet()
	if err != nil {
		t.Fatalf("build fallback sprite sheet: %v", err)
	}
	img, err := gif.Decode(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("decode fallback sprite sheet: %v", err)
	}
	if got, want := img.Bounds().Dx(), spriteSourceSize*spriteColumns; got != want {
		t.Fatalf("sprite width = %d, want %d", got, want)
	}
	if got, want := img.Bounds().Dy(), spriteSourceSize*10; got != want {
		t.Fatalf("sprite height = %d, want %d", got, want)
	}
	_, _, _, cornerA := img.At(0, 0).RGBA()
	if cornerA != 0 {
		t.Fatalf("top-left background alpha = %d, want 0", cornerA)
	}
	_, _, _, centerA := img.At(56, 62).RGBA()
	if centerA == 0 {
		t.Fatal("idle frame center unexpectedly transparent")
	}
}
