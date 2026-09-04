//go:build windows

package main

import "unsafe"

const monitorDefaultToNearest = 0x00000002

var (
	procMonitorFromPoint  = user32.NewProc("MonitorFromPoint")
	procMonitorFromWindow = user32.NewProc("MonitorFromWindow")
	procGetMonitorInfoW   = user32.NewProc("GetMonitorInfoW")
)

type MONITORINFO struct {
	CbSize    uint32
	RcMonitor RECT
	RcWork    RECT
	DwFlags   uint32
}

// primaryWorkArea preserves the old single-monitor fallback for startup and
// for the unlikely case where the monitor APIs cannot resolve a display.
func primaryWorkArea() RECT {
	var r RECT
	ok, _, _ := procSystemParametersInfoW.Call(SPI_GETWORKAREA, 0, uintptr(unsafe.Pointer(&r)), 0)
	if ok != 0 && r.Right > r.Left && r.Bottom > r.Top {
		return r
	}
	sw, _, _ := procGetSystemMetrics.Call(0)
	sh, _, _ := procGetSystemMetrics.Call(1)
	return RECT{0, 0, int32(sw), int32(sh)}
}

func monitorWorkArea(hMonitor uintptr) (RECT, bool) {
	if hMonitor == 0 {
		return RECT{}, false
	}
	mi := MONITORINFO{CbSize: uint32(unsafe.Sizeof(MONITORINFO{}))}
	ok, _, _ := procGetMonitorInfoW.Call(hMonitor, uintptr(unsafe.Pointer(&mi)))
	if ok == 0 || mi.RcWork.Right <= mi.RcWork.Left || mi.RcWork.Bottom <= mi.RcWork.Top {
		return RECT{}, false
	}
	return mi.RcWork, true
}

// workAreaForPoint returns the usable area of the monitor nearest a virtual-
// screen coordinate. Negative X/Y coordinates are valid on Windows when a
// monitor is arranged to the left or above the primary monitor.
func workAreaForPoint(x, y int32) RECT {
	hMonitor, _, _ := procMonitorFromPoint.Call(packPoint(x, y), monitorDefaultToNearest)
	if r, ok := monitorWorkArea(hMonitor); ok {
		return r
	}
	return primaryWorkArea()
}

func workAreaForWindow(hwnd uintptr) RECT {
	if hwnd == 0 {
		return primaryWorkArea()
	}
	hMonitor, _, _ := procMonitorFromWindow.Call(hwnd, monitorDefaultToNearest)
	if r, ok := monitorWorkArea(hMonitor); ok {
		return r
	}
	return primaryWorkArea()
}

func workAreaForPet() RECT {
	if pet.hwnd != 0 {
		return workAreaForWindow(pet.hwnd)
	}
	return primaryWorkArea()
}
