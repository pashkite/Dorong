//go:build windows

package main

import (
	"encoding/base64"
	_ "embed"
)

// v0.5.7 uses the user-approved Dorong character sprite sheet packed as a
// normal repository asset. initFrames keeps the existing decoder path, while
// this small bridge lets the final EXE remain a single standalone file.
//go:embed assets/sprites_v057.png
var spriteAssetBytes []byte

var spriteAssetBase64 = base64.StdEncoding.EncodeToString(spriteAssetBytes)
