//go:build windows

package main

// The repository asset at assets/sprites_v057.png is not a valid image file.
// Keep the standalone EXE self-contained by rebuilding the original, valid GIF
// sprite sheet from the preserved base64 chunks.
var spriteAssetBase64 = spriteData0 + spriteData1 + spriteData2 + spriteData3 + spriteData4 + spriteData5 + spriteData6 + spriteData7
