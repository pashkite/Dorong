//go:build windows

package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"testing"
)

func TestRecoverMissingSpriteCharacter(t *testing.T) {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	prefix := spriteData0 + spriteData1 + spriteData2 + spriteData3 + spriteData4 + spriteData5
	suffix := spriteData6 + spriteData7
	var matches []byte
	for i := 0; i < len(alphabet); i++ {
		candidate := prefix + string(alphabet[i]) + suffix
		data, err := base64.StdEncoding.DecodeString(candidate)
		if err != nil {
			continue
		}
		img, format, err := image.Decode(bytes.NewReader(data))
		if err != nil || format != "gif" {
			continue
		}
		if img.Bounds().Dx() == spriteSourceSize*spriteColumns && img.Bounds().Dy() == spriteSourceSize*10 {
			matches = append(matches, alphabet[i])
			t.Logf("valid candidate %q", alphabet[i])
		}
	}
	if len(matches) != 1 {
		t.Fatalf("expected exactly one recovered character, got %q", string(matches))
	}
}
