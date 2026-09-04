//go:build windows

package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var (
	advapi32            = syscall.NewLazyDLL("advapi32.dll")
	procRegCreateKeyExW = advapi32.NewProc("RegCreateKeyExW")
	procRegOpenKeyExW   = advapi32.NewProc("RegOpenKeyExW")
	procRegSetValueExW  = advapi32.NewProc("RegSetValueExW")
	procRegDeleteValueW = advapi32.NewProc("RegDeleteValueW")
	procRegCloseKey     = advapi32.NewProc("RegCloseKey")
)

const (
	hkeyCurrentUser = uintptr(0x80000001)
	keySetValue     = 0x0002
	regSZ           = 1
)

func setStartupEnabled(enabled bool) error {
	subKey := wchar(`Software\Microsoft\Windows\CurrentVersion\Run`)
	valueName := wchar("Dorong")

	if !enabled {
		var hKey uintptr
		status, _, _ := procRegOpenKeyExW.Call(
			hkeyCurrentUser,
			uintptr(unsafe.Pointer(subKey)),
			0,
			keySetValue,
			uintptr(unsafe.Pointer(&hKey)),
		)
		if status != 0 {
			return nil
		}
		defer procRegCloseKey.Call(hKey)
		status, _, _ = procRegDeleteValueW.Call(hKey, uintptr(unsafe.Pointer(valueName)))
		if status != 0 && status != 2 {
			return fmt.Errorf("RegDeleteValueW failed: %d", status)
		}
		return nil
	}

	exe, err := os.Executable()
	if err != nil {
		return err
	}
	command := "\"" + exe + "\""
	value, err := syscall.UTF16FromString(command)
	if err != nil {
		return err
	}

	var hKey uintptr
	var disposition uint32
	status, _, _ := procRegCreateKeyExW.Call(
		hkeyCurrentUser,
		uintptr(unsafe.Pointer(subKey)),
		0,
		0,
		0,
		keySetValue,
		0,
		uintptr(unsafe.Pointer(&hKey)),
		uintptr(unsafe.Pointer(&disposition)),
	)
	if status != 0 {
		return fmt.Errorf("RegCreateKeyExW failed: %d", status)
	}
	defer procRegCloseKey.Call(hKey)
	status, _, _ = procRegSetValueExW.Call(
		hKey,
		uintptr(unsafe.Pointer(valueName)),
		0,
		regSZ,
		uintptr(unsafe.Pointer(&value[0])),
		uintptr(uint32(len(value)*2)),
	)
	if status != 0 {
		return fmt.Errorf("RegSetValueExW failed: %d", status)
	}
	return nil
}
