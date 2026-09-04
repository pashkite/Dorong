//go:build windows

package main

import (
	"fmt"
	"math/rand"
	"syscall"
	"time"
	"unsafe"
)

const (
	AppName    = "Dorong"
	AppVersion = "0.5.1"
	PET_W      = 236
	PET_H      = 236
)

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	gdi32    = syscall.NewLazyDLL("gdi32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procRegisterClassExW      = user32.NewProc("RegisterClassExW")
	procCreateWindowExW       = user32.NewProc("CreateWindowExW")
	procDefWindowProcW        = user32.NewProc("DefWindowProcW")
	procDestroyWindow         = user32.NewProc("DestroyWindow")
	procShowWindow            = user32.NewProc("ShowWindow")
	procUpdateWindow          = user32.NewProc("UpdateWindow")
	procGetMessageW           = user32.NewProc("GetMessageW")
	procTranslateMessage      = user32.NewProc("TranslateMessage")
	procDispatchMessageW      = user32.NewProc("DispatchMessageW")
	procPostQuitMessage       = user32.NewProc("PostQuitMessage")
	procSetTimer              = user32.NewProc("SetTimer")
	procKillTimer             = user32.NewProc("KillTimer")
	procSetCapture            = user32.NewProc("SetCapture")
	procReleaseCapture        = user32.NewProc("ReleaseCapture")
	procGetCursorPos          = user32.NewProc("GetCursorPos")
	procGetWindowRect         = user32.NewProc("GetWindowRect")
	procSetWindowPos          = user32.NewProc("SetWindowPos")
	procCreatePopupMenu       = user32.NewProc("CreatePopupMenu")
	procAppendMenuW           = user32.NewProc("AppendMenuW")
	procTrackPopupMenu        = user32.NewProc("TrackPopupMenu")
	procDestroyMenu           = user32.NewProc("DestroyMenu")
	procMessageBoxW           = user32.NewProc("MessageBoxW")
	procGetSystemMetrics      = user32.NewProc("GetSystemMetrics")
	procSystemParametersInfoW = user32.NewProc("SystemParametersInfoW")
	procSetForegroundWindow   = user32.NewProc("SetForegroundWindow")
	procUpdateLayeredWindow   = user32.NewProc("UpdateLayeredWindow")
	procGetDC                 = user32.NewProc("GetDC")
	procReleaseDC             = user32.NewProc("ReleaseDC")
	procSetWindowTextW        = user32.NewProc("SetWindowTextW")
	procSendMessageW          = user32.NewProc("SendMessageW")
	procWindowFromPoint       = user32.NewProc("WindowFromPoint")
	procGetAncestor           = user32.NewProc("GetAncestor")
	procIsWindow              = user32.NewProc("IsWindow")
	procIsWindowVisible       = user32.NewProc("IsWindowVisible")
	procGetDesktopWindow      = user32.NewProc("GetDesktopWindow")
	procGetShellWindow        = user32.NewProc("GetShellWindow")

	procCreateCompatibleDC = gdi32.NewProc("CreateCompatibleDC")
	procDeleteDC           = gdi32.NewProc("DeleteDC")
	procCreateDIBSection   = gdi32.NewProc("CreateDIBSection")
	procSelectObject       = gdi32.NewProc("SelectObject")
	procDeleteObject       = gdi32.NewProc("DeleteObject")
	procCreateFontW        = gdi32.NewProc("CreateFontW")

	procGetModuleHandleW = kernel32.NewProc("GetModuleHandleW")
)

const (
	CS_DBLCLKS       = 0x0008
	WS_POPUP         = 0x80000000
	WS_BORDER        = 0x00800000
	WS_EX_LAYERED    = 0x00080000
	WS_EX_TOPMOST    = 0x00000008
	WS_EX_TOOLWINDOW = 0x00000080
	WS_EX_NOACTIVATE = 0x08000000
	SS_CENTER        = 0x00000001
	SS_CENTERIMAGE   = 0x00000200

	SW_HIDE           = 0
	SW_SHOW           = 5
	SW_SHOWNOACTIVATE = 4

	WM_DESTROY       = 0x0002
	WM_TIMER         = 0x0113
	WM_COMMAND       = 0x0111
	WM_LBUTTONDOWN   = 0x0201
	WM_LBUTTONUP     = 0x0202
	WM_MOUSEMOVE     = 0x0200
	WM_LBUTTONDBLCLK = 0x0203
	WM_RBUTTONUP     = 0x0205
	WM_SETFONT       = 0x0030

	MF_STRING       = 0x0000
	MF_SEPARATOR    = 0x0800
	MF_CHECKED      = 0x0008
	TPM_RIGHTBUTTON = 0x0002

	HWND_TOPMOST   = ^uintptr(0)
	HWND_NOTOPMOST = ^uintptr(1)
	SWP_NOMOVE     = 0x0002
	SWP_NOSIZE     = 0x0001
	SWP_NOACTIVATE = 0x0010
	SWP_SHOWWINDOW = 0x0040

	BI_RGB         = 0
	DIB_RGB_COLORS = 0
	AC_SRC_OVER    = 0
	AC_SRC_ALPHA   = 1
	ULW_ALPHA      = 2

	MB_ICONINFORMATION = 0x40
	SPI_GETWORKAREA    = 0x0030
	GA_ROOT            = 2

	ID_GREET   = 1001
	ID_FOCUS   = 1002
	ID_ALARM   = 1003
	ID_WANDER  = 1004
	ID_TOPMOST = 1005
	ID_HOME    = 1006
	ID_SLEEP   = 1007
	ID_DROP    = 1008
	ID_EXIT    = 1009
)

type POINT struct{ X, Y int32 }
type SIZE struct{ CX, CY int32 }
type RECT struct{ Left, Top, Right, Bottom int32 }
type MSG struct {
	Hwnd    uintptr
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}
type WNDCLASSEX struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     uintptr
	HIcon         uintptr
	HCursor       uintptr
	HbrBackground uintptr
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       uintptr
}
type BITMAPINFOHEADER struct {
	BiSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}
type RGBQUAD struct{ Blue, Green, Red, Reserved byte }
type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors [1]RGBQUAD
}
type BLENDFUNCTION struct{ BlendOp, BlendFlags, SourceConstantAlpha, AlphaFormat byte }

type Sprite struct {
	memDC, bitmap, oldBitmap, bits uintptr
}

type PetState struct {
	hwnd, bubble, bubbleFont uintptr

	dragging, dragMoved bool
	dragDX, dragDY      int32
	dragStart           POINT

	wander, topmost bool
	vx, targetX     int32
	edgeTarget      bool

	falling bool
	vy      float64

	supportHwnd uintptr // 0 means desktop work-area floor
	hanging     bool
	hangSide    int
	hangUntil   time.Time

	focusUntil, alarmUntil, nextAction, bubbleUntil time.Time
	sleepUntil, happyUntil                          time.Time

	animState string
	animFrame int
	lastAnim  time.Time
	lookSide  int // -1 left, 0 center, +1 right

	petCount int
}

var pet = PetState{
	wander: true, topmost: true, vx: 2,
	nextAction: time.Now().Add(4 * time.Second), animState: "idle", lastAnim: time.Now(),
}
var frameSets = map[string][]Sprite{}

func wchar(s string) *uint16 { p, _ := syscall.UTF16PtrFromString(s); return p }
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
func max32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
func abs32(v int32) int32 {
	if v < 0 {
		return -v
	}
	return v
}
func clamp32(v, lo, hi int32) int32 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func message(title, text string) {
	procMessageBoxW.Call(pet.hwnd, uintptr(unsafe.Pointer(wchar(text))), uintptr(unsafe.Pointer(wchar(title))), MB_ICONINFORMATION)
}

func workArea() RECT {
	var r RECT
	ok, _, _ := procSystemParametersInfoW.Call(SPI_GETWORKAREA, 0, uintptr(unsafe.Pointer(&r)), 0)
	if ok != 0 && r.Right > r.Left && r.Bottom > r.Top {
		return r
	}
	sw, _, _ := procGetSystemMetrics.Call(0)
	sh, _, _ := procGetSystemMetrics.Call(1)
	return RECT{0, 0, int32(sw), int32(sh)}
}

func currentPos() (int32, int32) {
	var r RECT
	procGetWindowRect.Call(pet.hwnd, uintptr(unsafe.Pointer(&r)))
	return r.Left, r.Top
}

func setPos(x, y int32) {
	procSetWindowPos.Call(pet.hwnd, 0, uintptr(x), uintptr(y), PET_W, PET_H, SWP_NOSIZE|SWP_NOACTIVATE)
	renderAt(x, y)
	positionBubble()
}

func moveHome() {
	wa := workArea()
	pet.falling = false
	pet.hanging = false
	pet.supportHwnd = 0
	pet.targetX = 0
	setPos(wa.Right-PET_W-28, wa.Bottom-PET_H)
}

func showBubble(text string, d time.Duration) {
	if pet.bubble == 0 {
		return
	}
	procSetWindowTextW.Call(pet.bubble, uintptr(unsafe.Pointer(wchar(text))))
	positionBubble()
	procShowWindow.Call(pet.bubble, SW_SHOWNOACTIVATE)
	pet.bubbleUntil = time.Now().Add(d)
}

func hideBubble() {
	if pet.bubble != 0 {
		procShowWindow.Call(pet.bubble, SW_HIDE)
	}
	pet.bubbleUntil = time.Time{}
}

func positionBubble() {
	if pet.bubble == 0 {
		return
	}
	x, y := currentPos()
	bw, bh := int32(230), int32(54)
	bx := x + PET_W/2 - bw/2
	by := y - bh - 4
	wa := workArea()
	if by < wa.Top+4 {
		by = y + PET_H + 4
	}
	bx = clamp32(bx, wa.Left+4, wa.Right-bw-4)
	procSetWindowPos.Call(pet.bubble, HWND_TOPMOST, uintptr(bx), uintptr(by), uintptr(bw), uintptr(bh), SWP_NOACTIVATE)
}

func packPoint(x, y int32) uintptr {
	return uintptr(uint64(uint32(x)) | (uint64(uint32(y)) << 32))
}

func topWindowAt(x, y int32) (uintptr, RECT, bool) {
	h, _, _ := procWindowFromPoint.Call(packPoint(x, y))
	if h == 0 {
		return 0, RECT{}, false
	}
	root, _, _ := procGetAncestor.Call(h, GA_ROOT)
	if root == 0 {
		root = h
	}
	if root == pet.hwnd || root == pet.bubble {
		return 0, RECT{}, false
	}
	desk, _, _ := procGetDesktopWindow.Call()
	shell, _, _ := procGetShellWindow.Call()
	if root == desk || root == shell {
		return 0, RECT{}, false
	}
	ok, _, _ := procIsWindowVisible.Call(root)
	if ok == 0 {
		return 0, RECT{}, false
	}
	var r RECT
	if rr, _, _ := procGetWindowRect.Call(root, uintptr(unsafe.Pointer(&r))); rr == 0 {
		return 0, RECT{}, false
	}
	if r.Right-r.Left < 80 || r.Bottom-r.Top < 60 {
		return 0, RECT{}, false
	}
	return root, r, true
}

func validSupportWindow(hwnd uintptr) (RECT, bool) {
	if hwnd == 0 {
		return workArea(), true
	}
	ok, _, _ := procIsWindow.Call(hwnd)
	vis, _, _ := procIsWindowVisible.Call(hwnd)
	if ok == 0 || vis == 0 {
		return RECT{}, false
	}
	var r RECT
	if rr, _, _ := procGetWindowRect.Call(hwnd, uintptr(unsafe.Pointer(&r))); rr == 0 {
		return RECT{}, false
	}
	return r, true
}

func supportXRange() (int32, int32) {
	wa := workArea()
	if pet.supportHwnd == 0 {
		return wa.Left, wa.Right - PET_W
	}
	r, ok := validSupportWindow(pet.supportHwnd)
	if !ok {
		return wa.Left, wa.Right - PET_W
	}
	// Keep the visual center/feet over the supporting window.
	lo := r.Left - PET_W/2 + 42
	hi := r.Right - PET_W/2 - 42
	lo = max32(lo, wa.Left)
	hi = min32(hi, wa.Right-PET_W)
	if hi < lo {
		hi = lo
	}
	return lo, hi
}

func snapToCurrentSupport() bool {
	if pet.supportHwnd == 0 {
		return true
	}
	r, ok := validSupportWindow(pet.supportHwnd)
	if !ok {
		return false
	}
	x, y := currentPos()
	wanted := r.Top - PET_H
	if abs32(y-wanted) > 1 {
		setPos(x, wanted)
	}
	return true
}

func startFall(initial float64) {
	pet.falling = true
	pet.vy = initial
	pet.supportHwnd = 0
	pet.targetX = 0
	pet.edgeTarget = false
	pet.hanging = false
}

func updateFalling() {
	if !pet.falling || pet.dragging {
		return
	}
	x, y := currentPos()
	oldBottom := y + PET_H
	var dy int32
	pet.vy, dy = nextFallStep(pet.vy)
	ny := y + dy
	newBottom := ny + PET_H
	centerX := x + PET_W/2

	// Detect a real window under Dorong's feet while crossing its top edge.
	if h, r, ok := topWindowAt(centerX, newBottom+2); ok {
		if r.Top >= oldBottom-5 && r.Top <= newBottom+8 {
			pet.falling = false
			pet.vy = 0
			pet.supportHwnd = h
			setPos(x, r.Top-PET_H)
			pet.happyUntil = time.Now().Add(350 * time.Millisecond)
			return
		}
	}

	wa := workArea()
	if landedY, landed := resolveFloorLanding(ny, PET_H, wa.Bottom); landed {
		pet.falling = false
		pet.vy = 0
		pet.supportHwnd = 0
		setPos(clamp32(x, wa.Left, wa.Right-PET_W), landedY)
		pet.happyUntil = time.Now().Add(350 * time.Millisecond)
		return
	}
	setPos(clamp32(x, wa.Left, wa.Right-PET_W), ny)
}

func updateCursorLook() {
	if pet.dragging || pet.falling || pet.hanging || pet.targetX != 0 || time.Now().Before(pet.sleepUntil) || time.Now().Before(pet.happyUntil) || !pet.focusUntil.IsZero() {
		pet.lookSide = 0
		return
	}
	var p POINT
	procGetCursorPos.Call(uintptr(unsafe.Pointer(&p)))
	x, _ := currentPos()
	center := x + PET_W/2
	if p.X < center-55 {
		pet.lookSide = -1
	} else if p.X > center+55 {
		pet.lookSide = 1
	} else {
		pet.lookSide = 0
	}
}

func currentAnimState(now time.Time) string {
	if pet.dragging {
		return "held"
	}
	if pet.falling {
		return "fall"
	}
	if pet.hanging {
		return "hang"
	}
	if now.Before(pet.happyUntil) {
		return "happy"
	}
	if !pet.focusUntil.IsZero() {
		return "focus"
	}
	if now.Before(pet.sleepUntil) {
		return "sleep"
	}
	if pet.targetX != 0 && pet.wander {
		return "walk"
	}
	if pet.lookSide < 0 {
		return "look_left"
	}
	if pet.lookSide > 0 {
		return "look_right"
	}
	return "idle"
}

func animInterval(state string) time.Duration {
	switch state {
	case "walk":
		return 115 * time.Millisecond
	case "held":
		return 90 * time.Millisecond
	case "fall":
		return 85 * time.Millisecond
	case "hang":
		return 150 * time.Millisecond
	case "happy":
		return 120 * time.Millisecond
	case "sleep":
		return 380 * time.Millisecond
	case "focus":
		return 320 * time.Millisecond
	case "look_left", "look_right":
		return 280 * time.Millisecond
	default:
		return 240 * time.Millisecond
	}
}

func updateAnimation(now time.Time, force bool) {
	state := currentAnimState(now)
	if state != pet.animState {
		pet.animState = state
		pet.animFrame = 0
		pet.lastAnim = now
		force = true
	}
	frames := frameSets[state]
	if len(frames) == 0 {
		return
	}
	if force || now.Sub(pet.lastAnim) >= animInterval(state) {
		if !force {
			pet.animFrame = (pet.animFrame + 1) % len(frames)
		}
		pet.lastAnim = now
		x, y := currentPos()
		renderAt(x, y)
	}
}

func randomAction() {
	if !pet.focusUntil.IsZero() || pet.falling || pet.hanging {
		return
	}
	now := time.Now()
	r := rand.Intn(100)
	switch {
	case r < 28:
		phrases := []string{"뭐 하고 있어?", "나 여기 있어.", "물 한 모금 어때?", "오늘도 같이 있자.", "하나씩 해보자!", "쓰담쓰담 대기 중!", "도롱도롱…"}
		showBubble(phrases[rand.Intn(len(phrases))], 2300*time.Millisecond)
	case r < 48:
		pet.sleepUntil = now.Add(time.Duration(5+rand.Intn(5)) * time.Second)
		pet.targetX = 0
		showBubble("조금만 잘게…", 1800*time.Millisecond)
	case r < 90 && pet.wander:
		lo, hi := supportXRange()
		pet.edgeTarget = false
		if pet.supportHwnd != 0 && rand.Intn(100) < 22 {
			// Occasionally walk all the way to an application-window edge and hang.
			if rand.Intn(2) == 0 {
				pet.targetX = lo
			} else {
				pet.targetX = hi
			}
			pet.edgeTarget = true
		} else if hi > lo {
			pet.targetX = lo + int32(rand.Intn(int(hi-lo)+1))
		} else {
			pet.targetX = lo
		}
		x, _ := currentPos()
		if pet.targetX < x {
			pet.vx = -2
		} else {
			pet.vx = 2
		}
	}
	pet.nextAction = now.Add(time.Duration(5+rand.Intn(7)) * time.Second)
}

func updateWalking(now time.Time) {
	if !pet.wander || pet.dragging || pet.falling || pet.hanging || pet.targetX == 0 {
		return
	}
	if pet.supportHwnd != 0 && !snapToCurrentSupport() {
		startFall(0)
		return
	}
	x, y := currentPos()
	if abs32(pet.targetX-x) <= 3 {
		pet.targetX = 0
		if pet.edgeTarget && pet.supportHwnd != 0 {
			pet.edgeTarget = false
			pet.hanging = true
			pet.hangUntil = now.Add(1200 * time.Millisecond)
			lo, hi := supportXRange()
			if abs32(x-lo) < abs32(x-hi) {
				pet.hangSide = -1
			} else {
				pet.hangSide = 1
			}
			showBubble("앗…!", 900*time.Millisecond)
		}
		return
	}
	nx := x + pet.vx
	lo, hi := supportXRange()
	nx = clamp32(nx, lo, hi)
	setPos(nx, y)
}

func tick() {
	now := time.Now()
	if !pet.bubbleUntil.IsZero() && now.After(pet.bubbleUntil) {
		hideBubble()
	}

	if !pet.focusUntil.IsZero() {
		if now.After(pet.focusUntil) {
			pet.focusUntil = time.Time{}
			pet.happyUntil = now.Add(2 * time.Second)
			showBubble("집중 끝! 수고했어.", 5*time.Second)
			message(AppName, "25분 집중이 끝났어! 잠깐 쉬자.")
		} else if now.Second()%30 == 0 && now.Nanosecond() < 80_000_000 {
			rem := time.Until(pet.focusUntil)
			showBubble(fmt.Sprintf("집중 중 · %02d:%02d", int(rem.Minutes()), int(rem.Seconds())%60), 1100*time.Millisecond)
		}
	}
	if !pet.alarmUntil.IsZero() && now.After(pet.alarmUntil) {
		pet.alarmUntil = time.Time{}
		pet.happyUntil = now.Add(2 * time.Second)
		showBubble("알람이야! 시간이 됐어.", 5*time.Second)
		message(AppName, "10분 알람 시간이 됐어!")
	}

	if pet.hanging && now.After(pet.hangUntil) {
		x, y := currentPos()
		if pet.hangSide < 0 {
			x -= 8
		} else {
			x += 8
		}
		setPos(x, y+5)
		startFall(1.5)
	}

	if pet.falling {
		updateFalling()
	} else {
		if pet.supportHwnd != 0 && !pet.dragging && !pet.hanging && pet.targetX == 0 {
			if !snapToCurrentSupport() {
				startFall(0)
			}
		}
		updateWalking(now)
	}

	if now.After(pet.nextAction) && !pet.dragging && !pet.falling && !pet.hanging && now.After(pet.sleepUntil) {
		randomAction()
	}
	updateCursorLook()
	updateAnimation(now, false)
}
