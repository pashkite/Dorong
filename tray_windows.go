//go:build windows

package main

import (
	"syscall"
	"unsafe"
)

var (
	shell32              = syscall.NewLazyDLL("shell32.dll")
	procShellNotifyIconW = shell32.NewProc("Shell_NotifyIconW")
	procLoadIconW        = user32.NewProc("LoadIconW")
)

const (
	wmTrayIcon = 0x8001

	nimAdd    = 0x00000000
	nimDelete = 0x00000002

	nifMessage = 0x00000001
	nifIcon    = 0x00000002
	nifTip     = 0x00000004

	idiApplication = 32512
	trayIconID      = 1
)

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type NOTIFYICONDATA struct {
	CbSize           uint32
	HWnd             uintptr
	UID              uint32
	UFlags           uint32
	UCallbackMessage uint32
	HIcon            uintptr
	SzTip            [128]uint16
	DwState          uint32
	DwStateMask      uint32
	SzInfo           [256]uint16
	UVersion         uint32
	SzInfoTitle      [64]uint16
	DwInfoFlags      uint32
	GuidItem         GUID
	HBalloonIcon     uintptr
}

var trayIconAdded bool

func copyUTF16(dst []uint16, s string) {
	v, err := syscall.UTF16FromString(s)
	if err != nil {
		return
	}
	if len(v) > len(dst) {
		v = v[:len(dst)]
		v[len(v)-1] = 0
	}
	copy(dst, v)
}

func addTrayIcon() bool {
	if pet.hwnd == 0 || trayIconAdded {
		return trayIconAdded
	}
	hIcon, _, _ := procLoadIconW.Call(0, idiApplication)
	nid := NOTIFYICONDATA{
		CbSize:           uint32(unsafe.Sizeof(NOTIFYICONDATA{})),
		HWnd:             pet.hwnd,
		UID:              trayIconID,
		UFlags:           nifMessage | nifIcon | nifTip,
		UCallbackMessage: wmTrayIcon,
		HIcon:            hIcon,
	}
	copyUTF16(nid.SzTip[:], AppName+" v"+AppVersion)
	ok, _, _ := procShellNotifyIconW.Call(nimAdd, uintptr(unsafe.Pointer(&nid)))
	trayIconAdded = ok != 0
	return trayIconAdded
}

func removeTrayIcon() {
	if !trayIconAdded || pet.hwnd == 0 {
		return
	}
	nid := NOTIFYICONDATA{
		CbSize: uint32(unsafe.Sizeof(NOTIFYICONDATA{})),
		HWnd:   pet.hwnd,
		UID:    trayIconID,
	}
	procShellNotifyIconW.Call(nimDelete, uintptr(unsafe.Pointer(&nid)))
	trayIconAdded = false
}
