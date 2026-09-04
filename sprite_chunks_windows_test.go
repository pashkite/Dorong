//go:build windows

package main

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestInspectSpriteChunks(t *testing.T) {
	chunks := []string{spriteData0, spriteData1, spriteData2, spriteData3, spriteData4, spriteData5, spriteData6, spriteData7}
	for i, chunk := range chunks {
		tail := chunk
		if len(tail) > 20 {
			tail = tail[len(tail)-20:]
		}
		firstEq := strings.IndexByte(chunk, '=')
		_, err := base64.StdEncoding.DecodeString(chunk)
		t.Logf("chunk %d: len=%d mod4=%d firstEq=%d tail=%q individualDecodeErr=%v", i, len(chunk), len(chunk)%4, firstEq, tail, err)
	}
}
