//go:build windows

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	"math/rand"
	"syscall"
	"time"
	"unsafe"
)

// Dorong is a tiny Windows desktop pet.
// The animation sheet is stored as base64 chunks in sprite_data_*_windows.go,
// so the repository remains self-contained while the final Windows build is
// still a single Dorong.exe.
const (
	spriteSourceSize = 112
	spriteColumns    = 2
)

func loadSpriteFromSheet(img image.Image, ox, oy int) (Sprite, error) {
	b := img.Bounds()
	if ox < b.Min.X || oy < b.Min.Y || ox+spriteSourceSize > b.Max.X || oy+spriteSourceSize > b.Max.Y {
		return Sprite{}, fmt.Errorf("sprite rectangle outside sheet: %d,%d", ox, oy)
	}
	var s Sprite
	s.memDC, _, _ = procCreateCompatibleDC.Call(0)
	if s.memDC == 0 {
		return s, fmt.Errorf("CreateCompatibleDC failed")
	}
	bmi := BITMAPINFO{BmiHeader: BITMAPINFOHEADER{BiSize: uint32(unsafe.Sizeof(BITMAPINFOHEADER{})), BiWidth: PET_W, BiHeight: -PET_H, BiPlanes: 1, BiBitCount: 32, BiCompression: BI_RGB}}
	var bits uintptr
	s.bitmap, _, _ = procCreateDIBSection.Call(s.memDC, uintptr(unsafe.Pointer(&bmi)), DIB_RGB_COLORS, uintptr(unsafe.Pointer(&bits)), 0, 0)
	if s.bitmap == 0 || bits == 0 {
		return s, fmt.Errorf("CreateDIBSection failed")
	}
	s.bits = bits
	s.oldBitmap, _, _ = procSelectObject.Call(s.memDC, s.bitmap)
	pix := unsafe.Slice((*byte)(unsafe.Pointer(bits)), PET_W*PET_H*4)
	i := 0
	for y := 0; y < PET_H; y++ {
		sy := oy + y*spriteSourceSize/PET_H
		for x := 0; x < PET_W; x++ {
			sx := ox + x*spriteSourceSize/PET_W
			r, g, bb, a := img.At(sx, sy).RGBA()
			aa := uint8(a >> 8)
			rr := uint8((uint32(r>>8)*uint32(aa) + 127) / 255)
			gg := uint8((uint32(g>>8)*uint32(aa) + 127) / 255)
			bv := uint8((uint32(bb>>8)*uint32(aa) + 127) / 255)
			pix[i], pix[i+1], pix[i+2], pix[i+3] = bv, gg, rr, aa
			i += 4
		}
	}
	return s, nil
}

func initFrames() error {
	data, err := base64.StdEncoding.DecodeString(spriteAssetBase64)
	if err != nil {
		return fmt.Errorf("decode embedded sprite sheet: %w", err)
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("decode sprite image: %w", err)
	}
	states := []string{"idle", "walk", "sleep", "happy", "held", "focus", "fall", "hang", "look_left", "look_right"}
	expectedW, expectedH := spriteSourceSize*spriteColumns, spriteSourceSize*len(states)
	b := img.Bounds()
	if b.Dx() != expectedW || b.Dy() != expectedH {
		return fmt.Errorf("unexpected spritesheet size %dx%d; want %dx%d", b.Dx(), b.Dy(), expectedW, expectedH)
	}
	for row, st := range states {
		for col := 0; col < spriteColumns; col++ {
			s, err := loadSpriteFromSheet(img, col*spriteSourceSize, row*spriteSourceSize)
			if err != nil {
				return err
			}
			frameSets[st] = append(frameSets[st], s)
		}
	}
	return nil
}

func cleanupFrames() {
	for _, arr := range frameSets {
		for _, s := range arr {
			if s.memDC != 0 {
				if s.oldBitmap != 0 {
					procSelectObject.Call(s.memDC, s.oldBitmap)
				}
				if s.bitmap != 0 {
					procDeleteObject.Call(s.bitmap)
				}
				procDeleteDC.Call(s.memDC)
			}
		}
	}
}

func renderAt(x, y int32) {
	arr := frameSets[pet.animState]
	if len(arr) == 0 {
		return
	}
	s := arr[pet.animFrame%len(arr)]
	dst := POINT{x, y}
	src := POINT{0, 0}
	sz := SIZE{PET_W, PET_H}
	blend := BLENDFUNCTION{BlendOp: AC_SRC_OVER, SourceConstantAlpha: 255, AlphaFormat: AC_SRC_ALPHA}
	screen, _, _ := procGetDC.Call(0)
	procUpdateLayeredWindow.Call(pet.hwnd, screen, uintptr(unsafe.Pointer(&dst)), uintptr(unsafe.Pointer(&sz)), s.memDC, uintptr(unsafe.Pointer(&src)), 0, uintptr(unsafe.Pointer(&blend)), ULW_ALPHA)
	procReleaseDC.Call(0, screen)
}

func popup(hwnd uintptr) {
	menu, _, _ := procCreatePopupMenu.Call()
	defer procDestroyMenu.Call(menu)
	procAppendMenuW.Call(menu, MF_STRING, ID_GREET, uintptr(unsafe.Pointer(wchar(tr("menu.greet")))))
	procAppendMenuW.Call(menu, MF_STRING, ID_FOCUS, uintptr(unsafe.Pointer(wchar(tr("menu.focus")))))
	procAppendMenuW.Call(menu, MF_STRING, ID_ALARM, uintptr(unsafe.Pointer(wchar(tr("menu.alarm")))))
	procAppendMenuW.Call(menu, MF_STRING, ID_SLEEP, uintptr(unsafe.Pointer(wchar(tr("menu.sleep")))))
	procAppendMenuW.Call(menu, MF_STRING, ID_DROP, uintptr(unsafe.Pointer(wchar(tr("menu.jump")))))
	procAppendMenuW.Call(menu, MF_SEPARATOR, 0, 0)
	wf := uintptr(MF_STRING)
	if pet.wander {
		wf |= MF_CHECKED
	}
	procAppendMenuW.Call(menu, wf, ID_WANDER, uintptr(unsafe.Pointer(wchar(tr("menu.wander")))))
	tf := uintptr(MF_STRING)
	if pet.topmost {
		tf |= MF_CHECKED
	}
	procAppendMenuW.Call(menu, tf, ID_TOPMOST, uintptr(unsafe.Pointer(wchar(tr("menu.topmost")))))
	procAppendMenuW.Call(menu, MF_STRING, ID_HOME, uintptr(unsafe.Pointer(wchar(tr("menu.home")))))
	procAppendMenuW.Call(menu, MF_SEPARATOR, 0, 0)
	koFlags := uintptr(MF_STRING)
	enFlags := uintptr(MF_STRING)
	if currentLanguage == LangKO {
		koFlags |= MF_CHECKED
	} else {
		enFlags |= MF_CHECKED
	}
	procAppendMenuW.Call(menu, koFlags, ID_LANG_KO, uintptr(unsafe.Pointer(wchar(tr("menu.lang_ko")))))
	procAppendMenuW.Call(menu, enFlags, ID_LANG_EN, uintptr(unsafe.Pointer(wchar(tr("menu.lang_en")))))
	procAppendMenuW.Call(menu, MF_SEPARATOR, 0, 0)
	procAppendMenuW.Call(menu, MF_STRING, ID_EXIT, uintptr(unsafe.Pointer(wchar(tr("menu.exit")))))
	var p POINT
	procGetCursorPos.Call(uintptr(unsafe.Pointer(&p)))
	procSetForegroundWindow.Call(hwnd)
	procTrackPopupMenu.Call(menu, TPM_RIGHTBUTTON, uintptr(p.X), uintptr(p.Y), 0, hwnd, 0)
}

func isHeadPat() bool {
	var p POINT
	procGetCursorPos.Call(uintptr(unsafe.Pointer(&p)))
	var r RECT
	procGetWindowRect.Call(pet.hwnd, uintptr(unsafe.Pointer(&r)))
	rx, ry := p.X-r.Left, p.Y-r.Top
	return rx >= 42 && rx <= PET_W-42 && ry >= 18 && ry <= 132
}

func wndProc(hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case WM_DESTROY:
		procKillTimer.Call(hwnd, 1)
		if pet.bubbleFont != 0 {
			procDeleteObject.Call(pet.bubbleFont)
		}
		if pet.bubble != 0 {
			procDestroyWindow.Call(pet.bubble)
		}
		cleanupFrames()
		procPostQuitMessage.Call(0)
		return 0
	case WM_TIMER:
		tick()
		return 0
	case WM_LBUTTONDOWN:
		pet.dragging = true
		pet.dragMoved = false
		pet.targetX = 0
		pet.walking = false
		pet.edgeTarget = false
		pet.sleepUntil = time.Time{}
		pet.falling = false
		pet.hanging = false
		var p POINT
		procGetCursorPos.Call(uintptr(unsafe.Pointer(&p)))
		pet.dragStart = p
		var r RECT
		procGetWindowRect.Call(hwnd, uintptr(unsafe.Pointer(&r)))
		pet.dragDX, pet.dragDY = p.X-r.Left, p.Y-r.Top
		procSetCapture.Call(hwnd)
		updateAnimation(time.Now(), true)
		return 0
	case WM_MOUSEMOVE:
		if pet.dragging {
			var p POINT
			procGetCursorPos.Call(uintptr(unsafe.Pointer(&p)))
			if abs32(p.X-pet.dragStart.X) > 4 || abs32(p.Y-pet.dragStart.Y) > 4 {
				pet.dragMoved = true
			}
			wa := workArea()
			nx := clamp32(p.X-pet.dragDX, wa.Left-PET_W/2, wa.Right-PET_W/2)
			ny := max32(p.Y-pet.dragDY, wa.Top-PET_H/2)
			setPos(nx, ny)
		}
		return 0
	case WM_LBUTTONUP:
		pet.dragging = false
		procReleaseCapture.Call()
		if pet.dragMoved {
			startFall(0)
		} else {
			pet.happyUntil = time.Now().Add(320 * time.Millisecond)
		}
		updateAnimation(time.Now(), true)
		return 0
	case WM_LBUTTONDBLCLK:
		if isHeadPat() {
			pet.petCount++
			pet.happyUntil = time.Now().Add(2300 * time.Millisecond)
			if pet.petCount%10 == 0 {
				showBubble(tr("pet_count", pet.petCount), 2200*time.Millisecond)
			} else {
				showBubble(tr("happy"), 2*time.Second)
			}
		} else {
			pet.happyUntil = time.Now().Add(900 * time.Millisecond)
			showBubble(tr("ticklish"), 1200*time.Millisecond)
		}
		updateAnimation(time.Now(), true)
		return 0
	case WM_RBUTTONUP:
		popup(hwnd)
		return 0
	case WM_COMMAND:
		switch int(wParam & 0xffff) {
		case ID_GREET:
			pet.happyUntil = time.Now().Add(1300 * time.Millisecond)
			showBubble(tr("greeting"), 3*time.Second)
		case ID_FOCUS:
			pet.focusUntil = time.Now().Add(25 * time.Minute)
			pet.sleepUntil = time.Time{}
			pet.targetX = 0
			pet.walking = false
			pet.edgeTarget = false
			showBubble(tr("focus_started"), 3*time.Second)
		case ID_ALARM:
			pet.alarmUntil = time.Now().Add(10 * time.Minute)
			showBubble(tr("alarm_set"), 3*time.Second)
		case ID_SLEEP:
			pet.sleepUntil = time.Now().Add(10 * time.Second)
			pet.targetX = 0
			pet.walking = false
			pet.edgeTarget = false
			showBubble(tr("sleep"), 1300*time.Millisecond)
		case ID_DROP:
			x, y := currentPos()
			setPos(x, y-18)
			startFall(-6)
			showBubble(tr("jump"), 900*time.Millisecond)
		case ID_WANDER:
			pet.wander = !pet.wander
			if !pet.wander {
				pet.targetX = 0
				pet.walking = false
				pet.edgeTarget = false
			}
		case ID_TOPMOST:
			pet.topmost = !pet.topmost
			pos := HWND_NOTOPMOST
			if pet.topmost {
				pos = HWND_TOPMOST
			}
			procSetWindowPos.Call(hwnd, pos, 0, 0, 0, 0, SWP_NOMOVE|SWP_NOSIZE|SWP_NOACTIVATE|SWP_SHOWWINDOW)
			if pet.bubble != 0 {
				procSetWindowPos.Call(pet.bubble, pos, 0, 0, 0, 0, SWP_NOMOVE|SWP_NOSIZE|SWP_NOACTIVATE)
			}
		case ID_HOME:
			moveHome()
		case ID_LANG_KO:
			setLanguage(LangKO)
			_ = saveSettings()
			showBubble(tr("language_changed_ko"), 1800*time.Millisecond)
		case ID_LANG_EN:
			setLanguage(LangEN)
			_ = saveSettings()
			showBubble(tr("language_changed_en"), 1800*time.Millisecond)
		case ID_EXIT:
			procDestroyWindow.Call(hwnd)
		}
		updateAnimation(time.Now(), true)
		return 0
	}
	r, _, _ := procDefWindowProcW.Call(hwnd, uintptr(msg), wParam, lParam)
	return r
}

func createBubble(hInst uintptr) {
	x, y := currentPos()
	pet.bubble, _, _ = procCreateWindowExW.Call(WS_EX_TOPMOST|WS_EX_TOOLWINDOW|WS_EX_NOACTIVATE, uintptr(unsafe.Pointer(wchar("STATIC"))), uintptr(unsafe.Pointer(wchar(""))), WS_POPUP|WS_BORDER|SS_CENTER|SS_CENTERIMAGE, uintptr(x), uintptr(y-60), 230, 54, 0, 0, hInst, 0)
	face := wchar("Malgun Gothic")
	pet.bubbleFont, _, _ = procCreateFontW.Call(uintptr(^uint32(15-1)), 0, 0, 0, 500, 0, 0, 0, 1, 0, 0, 0, 0, uintptr(unsafe.Pointer(face)))
	if pet.bubbleFont != 0 {
		procSendMessageW.Call(pet.bubble, WM_SETFONT, pet.bubbleFont, 1)
	}
	procShowWindow.Call(pet.bubble, SW_HIDE)
}

func main() {
	loadSettings()
	rand.Seed(time.Now().UnixNano())
	if err := initFrames(); err != nil {
		panic(err)
	}
	hInst, _, _ := procGetModuleHandleW.Call(0)
	className := wchar("DorongDesktopPetClass")
	cb := syscall.NewCallback(wndProc)
	wc := WNDCLASSEX{CbSize: uint32(unsafe.Sizeof(WNDCLASSEX{})), Style: CS_DBLCLKS, LpfnWndProc: cb, HInstance: hInst, LpszClassName: className}
	if r, _, e := procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc))); r == 0 {
		panic(e)
	}
	wa := workArea()
	x, y := wa.Right-PET_W-28, wa.Bottom-PET_H
	hwnd, _, e := procCreateWindowExW.Call(WS_EX_LAYERED|WS_EX_TOPMOST|WS_EX_TOOLWINDOW, uintptr(unsafe.Pointer(className)), uintptr(unsafe.Pointer(wchar(AppName+" v"+AppVersion))), WS_POPUP, uintptr(x), uintptr(y), PET_W, PET_H, 0, 0, hInst, 0)
	if hwnd == 0 {
		panic(e)
	}
	pet.hwnd = hwnd
	pet.supportHwnd = 0
	createBubble(hInst)
	renderAt(x, y)
	procShowWindow.Call(hwnd, SW_SHOW)
	procUpdateWindow.Call(hwnd)
	procSetTimer.Call(hwnd, 1, 40, 0)
	showBubble(tr("greeting"), 1900*time.Millisecond)

	var m MSG
	for {
		r, _, _ := procGetMessageW.Call(uintptr(unsafe.Pointer(&m)), 0, 0, 0)
		if int32(r) <= 0 {
			break
		}
		procTranslateMessage.Call(uintptr(unsafe.Pointer(&m)))
		procDispatchMessageW.Call(uintptr(unsafe.Pointer(&m)))
	}
}
