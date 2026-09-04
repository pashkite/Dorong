# Dorong

**English** | [한국어](README.ko.md)

**Dorong** is a tiny personal Windows desktop pet written in Go and the Win32 API.
It runs as a standalone Windows executable and does not require Python, .NET, or any separate runtime installation.

> Current version: **v0.5.4**

## Features

- Transparent animated desktop pet
- Drag Dorong anywhere on the desktop
- Release after dragging to apply gravity
- Respect the Windows work area so Dorong lands above a docked taskbar
- Land accurately on ordinary application-window tops, including narrow edges under either foot
- Walk left and right on application windows while adapting to moved or resized support windows
- Occasionally hang from an application-window edge
- Double-click the head area to pet Dorong
- Idle reaction that looks toward the mouse cursor
- Idle / walk / sleep / happy / held / focus / fall / hang animation states
- Random speech bubbles
- 25-minute focus timer
- 10-minute alarm
- Always-on-top toggle
- **Korean / English language selection**
- Language preference persists across restarts
- Standalone 64-bit Windows executable
- No network connection or account required

## Controls

| Input | Action |
| --- | --- |
| Left drag | Pick up and move Dorong |
| Release after dragging | Drop Dorong; gravity is applied |
| Double-click head | Pet Dorong |
| Right-click | Open interaction menu |

## Language

Right-click Dorong and choose either `한국어` or `English` near the bottom of the menu.
Menus, speech bubbles, focus/alarm messages, and system dialogs immediately use the selected language. The preference is saved in the user's Windows config directory under `Dorong/settings.json`.

## Run

Download the Windows build artifact from GitHub Actions, or use the packaged build supplied with a tagged/local release bundle.

No Python, Go, or .NET installation is required **to run the prebuilt executable**.

## Build from source

Go 1.23+ is recommended.

```powershell
go build -ldflags="-H=windowsgui -s -w" -o Dorong.exe .
```

Cross-compiling from Linux/macOS:

```bash
GOOS=windows GOARCH=amd64 go build -ldflags='-H=windowsgui -s -w' -o Dorong.exe .
```

## Project structure

```text
Dorong/
├─ app_windows.go
├─ core_windows.go
├─ localization.go
├─ settings.go
├─ movement.go
├─ physics.go
├─ *_test.go
├─ sprite_data_windows.go
├─ sprite_data_0_windows.go … sprite_data_7_windows.go
├─ go.mod
├─ assets/
├─ .github/workflows/
│  └─ build.yml
└─ docs/
   └─ ROADMAP.md
```

## Notes

This repository is currently a **private personal project**. If the repository is made public later, review the redistribution rights for character artwork before publishing the assets.

## Roadmap

See [`docs/ROADMAP.md`](docs/ROADMAP.md).
