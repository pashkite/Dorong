//go:build windows

package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/gif"
)

// The repository sprite asset used by v0.6.1 was corrupted.  Keep the app
// self-contained and bootable by replacing the embedded data at package init
// time with a small, procedurally rendered Dorong sprite sheet.  initFrames
// continues to use the same decoder/frame-loading path, so a restored artwork
// asset can replace this fallback later without touching animation logic.
func init() {
	data, err := buildFallbackSpriteSheet()
	if err != nil {
		return
	}
	spriteAssetBase64 = base64.StdEncoding.EncodeToString(data)
}

const (
	palTransparent uint8 = iota
	palOutline
	palBody
	palHair
	palHairDark
	palEye
	palBlush
	palShadow
	palAccent
)

var fallbackPalette = color.Palette{
	color.NRGBA{R: 0, G: 0, B: 0, A: 0},
	color.NRGBA{R: 67, G: 57, B: 74, A: 255},
	color.NRGBA{R: 250, G: 248, B: 249, A: 255},
	color.NRGBA{R: 241, G: 154, B: 190, A: 255},
	color.NRGBA{R: 213, G: 101, B: 151, A: 255},
	color.NRGBA{R: 120, G: 83, B: 157, A: 255},
	color.NRGBA{R: 255, G: 195, B: 211, A: 255},
	color.NRGBA{R: 207, G: 204, B: 214, A: 255},
	color.NRGBA{R: 233, G: 92, B: 126, A: 255},
}

func buildFallbackSpriteSheet() ([]byte, error) {
	const rows = 10
	img := image.NewPaletted(
		image.Rect(0, 0, spriteSourceSize*spriteColumns, spriteSourceSize*rows),
		fallbackPalette,
	)
	for row := 0; row < rows; row++ {
		for col := 0; col < spriteColumns; col++ {
			drawDorongFrame(img, col*spriteSourceSize, row*spriteSourceSize, row, col)
		}
	}
	var buf bytes.Buffer
	if err := gif.Encode(&buf, img, &gif.Options{NumColors: len(fallbackPalette)}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func drawDorongFrame(img *image.Paletted, ox, oy, state, phase int) {
	cx := 56
	cy := 62
	rx, ry := 37, 41
	bob := 0
	if phase == 1 && state != 2 && state != 6 && state != 7 {
		bob = -2
	}

	// sleep: squash the body down; fall/hang: lift and stretch the pose.
	if state == 2 {
		cy, rx, ry = 72, 39, 28
	}
	if state == 6 {
		cy, rx, ry = 58, 34, 38
		if phase == 1 {
			cx += 2
		}
	}
	if state == 7 {
		cy, rx, ry = 58, 32, 36
		if phase == 1 {
			cx++
		}
	}
	cy += bob

	// Small ground shadow makes the white body readable on bright wallpapers.
	if state != 6 && state != 7 {
		fillEllipse(img, ox+cx, oy+101, 27, 5, palShadow)
	}

	// Tail behind the body.
	if state != 2 && state != 7 {
		fillEllipse(img, ox+cx+34, oy+cy+18, 14, 8, palOutline)
		fillEllipse(img, ox+cx+34, oy+cy+18, 11, 5, palBody)
	}

	// Main soft white body with a dark outline.
	fillEllipse(img, ox+cx, oy+cy, rx, ry, palOutline)
	fillEllipse(img, ox+cx, oy+cy, rx-3, ry-3, palBody)

	// Feet / walking stance.
	leftFootX, rightFootX := cx-18, cx+18
	if state == 1 {
		if phase == 0 {
			leftFootX -= 5
			rightFootX += 2
		} else {
			leftFootX -= 1
			rightFootX += 6
		}
	}
	if state != 7 {
		fillEllipse(img, ox+leftFootX, oy+cy+ry-4, 12, 7, palOutline)
		fillEllipse(img, ox+leftFootX, oy+cy+ry-5, 9, 4, palBody)
		fillEllipse(img, ox+rightFootX, oy+cy+ry-4, 12, 7, palOutline)
		fillEllipse(img, ox+rightFootX, oy+cy+ry-5, 9, 4, palBody)
	}

	// Pink hair cap, rose bun and ribbon — the recognizable Dorong silhouette.
	hairY := cy - 27
	fillEllipse(img, ox+cx, oy+hairY, 32, 20, palOutline)
	fillEllipse(img, ox+cx, oy+hairY, 29, 17, palHair)
	// Fringe tufts.
	fillTriangle(img, ox+cx-23, oy+hairY+8, ox+cx-8, oy+hairY+19, ox+cx-2, oy+hairY+5, palHair)
	fillTriangle(img, ox+cx-3, oy+hairY+7, ox+cx+8, oy+hairY+19, ox+cx+14, oy+hairY+4, palHair)
	fillTriangle(img, ox+cx+10, oy+hairY+6, ox+cx+24, oy+hairY+16, ox+cx+26, oy+hairY+1, palHair)

	bunX, bunY := cx+24, hairY-14
	fillEllipse(img, ox+bunX, oy+bunY, 12, 12, palOutline)
	fillEllipse(img, ox+bunX, oy+bunY, 9, 9, palHairDark)
	fillEllipse(img, ox+bunX-3, oy+bunY-2, 4, 4, palHair)
	fillEllipse(img, ox+bunX+3, oy+bunY+2, 3, 3, palHair)
	fillTriangle(img, ox+bunX-8, oy+bunY+8, ox+bunX-20, oy+bunY+15, ox+bunX-5, oy+bunY+18, palHairDark)
	fillTriangle(img, ox+bunX+7, oy+bunY+8, ox+bunX+18, oy+bunY+15, ox+bunX+4, oy+bunY+18, palHairDark)

	faceY := cy - 3
	// Arms and paws vary by animation state.
	switch state {
	case 4, 6: // held / falling
		drawThickLine(img, ox+cx-29, oy+cy-5, ox+cx-39, oy+cy-27, palOutline, 3)
		drawThickLine(img, ox+cx+29, oy+cy-5, ox+cx+39, oy+cy-27, palOutline, 3)
		fillEllipse(img, ox+cx-40, oy+cy-29, 6, 6, palBody)
		fillEllipse(img, ox+cx+40, oy+cy-29, 6, 6, palBody)
	case 7: // hanging
		drawThickLine(img, ox+cx-18, oy+cy-12, ox+cx-19, oy+10, palOutline, 4)
		drawThickLine(img, ox+cx+18, oy+cy-12, ox+cx+19, oy+10, palOutline, 4)
		fillEllipse(img, ox+cx-19, oy+9, 7, 6, palBody)
		fillEllipse(img, ox+cx+19, oy+9, 7, 6, palBody)
	default:
		fillEllipse(img, ox+cx-31, oy+cy+7, 8, 6, palOutline)
		fillEllipse(img, ox+cx-30, oy+cy+6, 5, 4, palBody)
		fillEllipse(img, ox+cx+31, oy+cy+7, 8, 6, palOutline)
		fillEllipse(img, ox+cx+30, oy+cy+6, 5, 4, palBody)
	}

	// Face.
	if state == 2 { // sleeping
		drawThickLine(img, ox+cx-19, oy+faceY, ox+cx-9, oy+faceY+2, palEye, 2)
		drawThickLine(img, ox+cx+9, oy+faceY+2, ox+cx+19, oy+faceY, palEye, 2)
		drawThickLine(img, ox+cx-3, oy+faceY+9, ox+cx+3, oy+faceY+9, palOutline, 1)
		drawZ(img, ox+cx+34, oy+faceY-19, 7)
		drawZ(img, ox+cx+43, oy+faceY-30, 5)
	} else if state == 3 { // happy
		drawThickLine(img, ox+cx-20, oy+faceY, ox+cx-14, oy+faceY+4, palEye, 2)
		drawThickLine(img, ox+cx-14, oy+faceY+4, ox+cx-8, oy+faceY, palEye, 2)
		drawThickLine(img, ox+cx+8, oy+faceY, ox+cx+14, oy+faceY+4, palEye, 2)
		drawThickLine(img, ox+cx+14, oy+faceY+4, ox+cx+20, oy+faceY, palEye, 2)
		drawThickLine(img, ox+cx-4, oy+faceY+10, ox+cx, oy+faceY+14, palOutline, 1)
		drawThickLine(img, ox+cx, oy+faceY+14, ox+cx+4, oy+faceY+10, palOutline, 1)
		drawHeart(img, ox+cx+38, oy+faceY-21, 8)
	} else {
		pupilShift := 0
		if state == 8 {
			pupilShift = -3
		} else if state == 9 {
			pupilShift = 3
		}
		fillEllipse(img, ox+cx-15, oy+faceY, 6, 8, palEye)
		fillEllipse(img, ox+cx+15, oy+faceY, 6, 8, palEye)
		fillEllipse(img, ox+cx-14+pupilShift, oy+faceY-1, 2, 3, palOutline)
		fillEllipse(img, ox+cx+14+pupilShift, oy+faceY-1, 2, 3, palOutline)
		setPaletted(img, ox+cx-13+pupilShift, oy+faceY-3, palBody)
		setPaletted(img, ox+cx+15+pupilShift, oy+faceY-3, palBody)
		drawThickLine(img, ox+cx-4, oy+faceY+11, ox+cx, oy+faceY+14, palOutline, 1)
		drawThickLine(img, ox+cx, oy+faceY+14, ox+cx+4, oy+faceY+11, palOutline, 1)
	}

	fillEllipse(img, ox+cx-27, oy+faceY+9, 7, 4, palBlush)
	fillEllipse(img, ox+cx+27, oy+faceY+9, 7, 4, palBlush)

	if state == 5 { // focus: small determined brows
		drawThickLine(img, ox+cx-22, oy+faceY-12, ox+cx-10, oy+faceY-9, palOutline, 2)
		drawThickLine(img, ox+cx+10, oy+faceY-9, ox+cx+22, oy+faceY-12, palOutline, 2)
	}
}

func setPaletted(img *image.Paletted, x, y int, idx uint8) {
	if image.Pt(x, y).In(img.Rect) {
		img.SetColorIndex(x, y, idx)
	}
}

func fillEllipse(img *image.Paletted, cx, cy, rx, ry int, idx uint8) {
	if rx <= 0 || ry <= 0 {
		return
	}
	rx2, ry2 := rx*rx, ry*ry
	limit := rx2 * ry2
	for y := cy - ry; y <= cy+ry; y++ {
		dy := y - cy
		for x := cx - rx; x <= cx+rx; x++ {
			dx := x - cx
			if dx*dx*ry2+dy*dy*rx2 <= limit {
				setPaletted(img, x, y, idx)
			}
		}
	}
}

func fillTriangle(img *image.Paletted, x1, y1, x2, y2, x3, y3 int, idx uint8) {
	minX, maxX := x1, x1
	minY, maxY := y1, y1
	for _, v := range []int{x2, x3} {
		if v < minX { minX = v }
		if v > maxX { maxX = v }
	}
	for _, v := range []int{y2, y3} {
		if v < minY { minY = v }
		if v > maxY { maxY = v }
	}
	area := func(ax, ay, bx, by, px, py int) int {
		return (px-ax)*(by-ay) - (py-ay)*(bx-ax)
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			a := area(x1, y1, x2, y2, x, y)
			b := area(x2, y2, x3, y3, x, y)
			c := area(x3, y3, x1, y1, x, y)
			if (a >= 0 && b >= 0 && c >= 0) || (a <= 0 && b <= 0 && c <= 0) {
				setPaletted(img, x, y, idx)
			}
		}
	}
}

func drawThickLine(img *image.Paletted, x0, y0, x1, y1 int, idx uint8, thickness int) {
	dx := absInt(x1 - x0)
	sx := -1
	if x0 < x1 { sx = 1 }
	dy := -absInt(y1 - y0)
	sy := -1
	if y0 < y1 { sy = 1 }
	err := dx + dy
	for {
		fillEllipse(img, x0, y0, thickness, thickness, idx)
		if x0 == x1 && y0 == y1 { break }
		e2 := 2 * err
		if e2 >= dy { err += dy; x0 += sx }
		if e2 <= dx { err += dx; y0 += sy }
	}
}

func drawZ(img *image.Paletted, x, y, size int) {
	drawThickLine(img, x, y, x+size, y, palEye, 1)
	drawThickLine(img, x+size, y, x, y+size, palEye, 1)
	drawThickLine(img, x, y+size, x+size, y+size, palEye, 1)
}

func drawHeart(img *image.Paletted, cx, cy, size int) {
	fillEllipse(img, cx-size/3, cy-size/4, size/3, size/3, palAccent)
	fillEllipse(img, cx+size/3, cy-size/4, size/3, size/3, palAccent)
	fillTriangle(img, cx-size*2/3, cy-size/5, cx+size*2/3, cy-size/5, cx, cy+size, palAccent)
}

func absInt(v int) int {
	if v < 0 { return -v }
	return v
}
