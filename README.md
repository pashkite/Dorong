# Dorong

**Dorong** is a tiny personal Windows desktop pet written in Go and the Win32 API.
It runs as a standalone Windows executable and does not require Python, .NET, or a separate runtime installation.

> Current version: **v0.5.0**

## Features

- Transparent always-on-top desktop pet
- Embedded animated sprite frames
- Idle / walking / sleeping / happy / held / focus animations
- Drag Dorong anywhere on the desktop
- Release after dragging to make Dorong fall with gravity
- Land on the top edge of ordinary application windows
- Walk along a window and occasionally hang from its edge before falling
- Double-click Dorong's head to pet it
- Cursor-following idle reaction
- Random speech bubbles
- 25-minute focus timer
- 10-minute alarm
- Right-click interaction menu
- No network connection or account required

## Controls

| Input | Action |
| --- | --- |
| Left drag | Pick up and move Dorong |
| Release after dragging | Drop Dorong; gravity is applied |
| Double-click head | Pet Dorong |
| Right-click | Open interaction menu |

## Run

Download or build `Dorong.exe`, then double-click it on 64-bit Windows.

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
├─ main.go
├─ go.mod
├─ frames/          # sprite animation frames embedded into Dorong.exe
├─ assets/          # preview images
└─ docs/
```

## Notes

This repository is for a personal desktop-pet project. The included character artwork is used as the project's private/personal skin. If this repository is made public later, review the rights to any character artwork before redistributing it.

## Roadmap

See [`docs/ROADMAP.md`](docs/ROADMAP.md).
