//go:build windows

package main

import (
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

const (
	ID_SETTINGS = 1012

	wsCaption         = 0x00C00000
	wsSysMenu         = 0x00080000
	wsChild           = 0x40000000
	wsVisible         = 0x10000000
	wsTabStop         = 0x00010000
	wsExClientEdge    = 0x00000200
	wsExControlParent = 0x00010000

	esNumber        = 0x00002000
	bsAutoCheckbox  = 0x00000003
	bsDefPushButton = 0x00000001

	wmClose    = 0x0010
	bmGetCheck = 0x00F0
	bmSetCheck = 0x00F1
	bstChecked = 1

	idSettingsSave   = 2101
	idSettingsCancel = 2102
	settingsWindowW  = 390
	settingsWindowH  = 250
)

var (
	settingsHwnd       uintptr
	focusEditHwnd      uintptr
	alarmEditHwnd      uintptr
	startupCheckHwnd   uintptr
	settingsClassCB    uintptr
	settingsClassReady bool
	procGetWindowTextW = user32.NewProc("GetWindowTextW")
)

func registerSettingsClass(hInst uintptr) bool {
	if settingsClassReady {
		return true
	}
	settingsClassCB = syscall.NewCallback(settingsWndProc)
	className := wchar("DorongSettingsWindowClass")
	wc := WNDCLASSEX{
		CbSize:        uint32(unsafe.Sizeof(WNDCLASSEX{})),
		LpfnWndProc:   settingsClassCB,
		HInstance:     hInst,
		HbrBackground: 6,
		LpszClassName: className,
	}
	r, _, _ := procRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc)))
	settingsClassReady = r != 0
	return settingsClassReady
}

func createSettingsControl(className, text string, style, exStyle uintptr, x, y, w, h int32, id uintptr) uintptr {
	hwnd, _, _ := procCreateWindowExW.Call(
		exStyle,
		uintptr(unsafe.Pointer(wchar(className))),
		uintptr(unsafe.Pointer(wchar(text))),
		style,
		uintptr(x), uintptr(y), uintptr(w), uintptr(h),
		settingsHwnd,
		id,
		0,
		0,
	)
	if hwnd != 0 && pet.bubbleFont != 0 {
		procSendMessageW.Call(hwnd, WM_SETFONT, pet.bubbleFont, 1)
	}
	return hwnd
}

func readSettingsInt(hwnd uintptr, minValue, maxValue int) (int, bool) {
	var buf [32]uint16
	n, _, _ := procGetWindowTextW.Call(hwnd, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if n == 0 {
		return 0, false
	}
	v, err := strconv.Atoi(syscall.UTF16ToString(buf[:n]))
	if err != nil || v < minValue || v > maxValue {
		return 0, false
	}
	return v, true
}

func showSettingsWindow() {
	if settingsHwnd != 0 {
		procShowWindow.Call(settingsHwnd, SW_SHOW)
		procSetForegroundWindow.Call(settingsHwnd)
		return
	}
	hInst, _, _ := procGetModuleHandleW.Call(0)
	if !registerSettingsClass(hInst) {
		message(AppName, tr("settings.open_error"))
		return
	}

	wa := workArea()
	px, py := currentPos()
	x := px + PET_W/2 - settingsWindowW/2
	y := py + PET_H/2 - settingsWindowH/2
	x = clamp32(x, wa.Left+10, wa.Right-settingsWindowW-10)
	y = clamp32(y, wa.Top+10, wa.Bottom-settingsWindowH-10)

	settingsHwnd, _, _ = procCreateWindowExW.Call(
		WS_EX_TOOLWINDOW|wsExControlParent,
		uintptr(unsafe.Pointer(wchar("DorongSettingsWindowClass"))),
		uintptr(unsafe.Pointer(wchar(tr("settings.title")))),
		wsCaption|wsSysMenu,
		uintptr(x), uintptr(y), settingsWindowW, settingsWindowH,
		pet.hwnd,
		0,
		hInst,
		0,
	)
	if settingsHwnd == 0 {
		message(AppName, tr("settings.open_error"))
		return
	}

	createSettingsControl("STATIC", tr("settings.focus_label"), wsChild|wsVisible, 0, 24, 24, 205, 24, 0)
	focusEditHwnd = createSettingsControl("EDIT", strconv.Itoa(focusMinutes), wsChild|wsVisible|wsTabStop|esNumber, wsExClientEdge, 245, 20, 105, 28, 0)
	createSettingsControl("STATIC", tr("settings.alarm_label"), wsChild|wsVisible, 0, 24, 70, 205, 24, 0)
	alarmEditHwnd = createSettingsControl("EDIT", strconv.Itoa(alarmMinutes), wsChild|wsVisible|wsTabStop|esNumber, wsExClientEdge, 245, 66, 105, 28, 0)
	startupCheckHwnd = createSettingsControl("BUTTON", tr("settings.startup"), wsChild|wsVisible|wsTabStop|bsAutoCheckbox, 0, 24, 116, 326, 28, 0)
	if startupWithWindows {
		procSendMessageW.Call(startupCheckHwnd, bmSetCheck, bstChecked, 0)
	}
	createSettingsControl("BUTTON", tr("settings.save"), wsChild|wsVisible|wsTabStop|bsDefPushButton, 0, 184, 170, 80, 32, idSettingsSave)
	createSettingsControl("BUTTON", tr("settings.cancel"), wsChild|wsVisible|wsTabStop, 0, 274, 170, 80, 32, idSettingsCancel)

	procShowWindow.Call(settingsHwnd, SW_SHOW)
	procUpdateWindow.Call(settingsHwnd)
	procSetForegroundWindow.Call(settingsHwnd)
}

func saveSettingsWindowValues() {
	focus, okFocus := readSettingsInt(focusEditHwnd, 1, maxFocusMinutes)
	alarm, okAlarm := readSettingsInt(alarmEditHwnd, 1, maxAlarmMinutes)
	if !okFocus || !okAlarm {
		message(AppName, tr("settings.invalid"))
		return
	}
	checked, _, _ := procSendMessageW.Call(startupCheckHwnd, bmGetCheck, 0, 0)
	newStartup := checked == bstChecked
	if newStartup != startupWithWindows {
		if err := setStartupEnabled(newStartup); err != nil {
			message(AppName, tr("settings.startup_error"))
			return
		}
	}
	focusMinutes = focus
	alarmMinutes = alarm
	startupWithWindows = newStartup
	if err := saveSettings(); err != nil {
		message(AppName, tr("settings.save_error"))
		return
	}
	showBubble(tr("settings.saved"), 1800*time.Millisecond)
	procDestroyWindow.Call(settingsHwnd)
}

func settingsWndProc(hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case WM_COMMAND:
		switch int(wParam & 0xffff) {
		case idSettingsSave:
			saveSettingsWindowValues()
			return 0
		case idSettingsCancel:
			procDestroyWindow.Call(hwnd)
			return 0
		}
	case wmClose:
		procDestroyWindow.Call(hwnd)
		return 0
	case WM_DESTROY:
		settingsHwnd = 0
		focusEditHwnd = 0
		alarmEditHwnd = 0
		startupCheckHwnd = 0
		return 0
	}
	r, _, _ := procDefWindowProcW.Call(hwnd, uintptr(msg), wParam, lParam)
	return r
}
