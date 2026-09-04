# Dorong

**Dorong** is a tiny personal Windows desktop pet written in Go and the Win32 API.
It runs as a standalone Windows executable and does not require Python, .NET, or any separate runtime installation.

> Current version: **v0.5.0**

## Features

- Transparent animated desktop pet
- Drag Dorong anywhere on the desktop
- Release after dragging to apply gravity
- Land on the top edge of ordinary application windows
- Walk along an application window and occasionally hang from its edge
- Double-click the head area to pet Dorong
- Idle reaction that looks toward the mouse cursor
- Idle / walk / sleep / happy / held / focus / fall / hang animation states
- Random speech bubbles
- 25-minute focus timer
- 10-minute alarm
- Always-on-top toggle
- Standalone 64-bit Windows executable
- No network connection or account required

## Controls

| Input | Action |
| --- | --- |
| Left drag | Pick up and move Dorong |
| Release after dragging | Drop Dorong; gravity is applied |
| Double-click head | Pet Dorong |
| Right-click | Open interaction menu |

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
├─ sprite_data_windows.go
├─ sprite_data_0_windows.go … sprite_data_7_windows.go
├─ go.mod
├─ assets/
│  ├─ spritesheet*.gif     # source/derived animation sheets
│  ├─ preview.gif
│  ├─ animation_preview.gif
│  ├─ contact_sheet.png / .jpg
│  └─ icon.png
├─ .github/workflows/
│  └─ build.yml
└─ docs/
   └─ ROADMAP.md
```

## Notes

This repository is currently a **private personal project**. If the repository is made public later, review the redistribution rights for character artwork before publishing the assets.

## Roadmap

See [`docs/ROADMAP.md`](docs/ROADMAP.md).
