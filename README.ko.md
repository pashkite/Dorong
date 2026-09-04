# Dorong

**한국어** | [영어](README.en.md) | [일본어](README.ja.md) | [중국어](README.zh-CN.md)

**Dorong(도롱)** 은 Go와 Win32 API로 만든 개인용 Windows 데스크톱 펫입니다.
Python, .NET 같은 별도 런타임을 설치하지 않아도 `Dorong.exe` 하나로 실행할 수 있습니다.

> 현재 버전: **v0.5.5**

## 주요 기능

- 투명 배경의 애니메이션 데스크톱 펫
- 마우스로 Dorong을 잡아서 원하는 위치로 이동
- 드래그 후 놓으면 중력이 적용되어 아래로 낙하
- Windows 작업 영역을 인식해 작업표시줄 위에 착지
- 일반 프로그램 창의 상단에 정확히 착지
- 좁은 창 가장자리도 왼발·중앙·오른발 판정으로 인식
- 프로그램 창 위에서 좌우로 이동하며, 창이 이동하거나 크기가 바뀌어도 이동 범위를 다시 계산
- 가끔 창 가장자리까지 걸어가 실제 가장자리에 몸을 걸치고 매달리기
- 매달린 동안 창이 움직이거나 크기가 바뀌면 위치를 따라가며, 창이 사라지면 자연스럽게 낙하
- 머리 부분을 더블클릭하면 쓰다듬기 반응
- 대기 중 마우스 커서를 바라보는 반응
- 대기 / 걷기 / 잠자기 / 기쁨 / 들림 / 집중 / 낙하 / 매달리기 애니메이션
- 랜덤 말풍선
- 25분 집중 타이머
- 10분 알람
- 항상 위에 표시 옵션
- **앱 UI 언어: 한국어 / English**
- 선택한 언어를 저장해 다음 실행에도 유지
- 64비트 Windows 단일 실행 파일
- 네트워크 연결 및 계정 불필요

## 조작 방법

| 조작 | 기능 |
| --- | --- |
| 왼쪽 버튼 드래그 | Dorong을 집어서 이동 |
| 드래그 후 놓기 | Dorong을 떨어뜨림 — 중력 적용 |
| 머리 부분 더블클릭 | Dorong 쓰다듬기 |
| 오른쪽 클릭 | 메뉴 열기 |

## 언어

이 README는 페이지 위쪽의 링크를 통해 **한국어 / 영어 / 일본어 / 중국어**로 볼 수 있습니다.

Dorong 앱 내부 UI는 현재 **한국어와 영어**를 지원합니다. Dorong을 오른쪽 클릭한 뒤 아래쪽의 언어 항목에서 `한국어` 또는 `English`를 선택하면 됩니다.
메뉴, 말풍선, 집중/알람 안내, 시스템 메시지가 즉시 선택한 언어로 바뀌며, 설정은 Windows 사용자 설정 폴더의 `Dorong/settings.json`에 저장되어 다음 실행에도 유지됩니다.

## 실행

GitHub Actions에서 생성된 Windows 빌드 또는 릴리스/로컬 빌드의 `Dorong.exe`를 실행하면 됩니다.

미리 빌드된 실행 파일을 사용하는 경우 **Python, Go, .NET을 설치할 필요가 없습니다.**

## 소스에서 빌드

Go 1.23 이상을 권장합니다.

```powershell
go build -ldflags="-H=windowsgui -s -w" -o Dorong.exe .
```

Linux/macOS에서 Windows용으로 크로스 컴파일:

```bash
GOOS=windows GOARCH=amd64 go build -ldflags='-H=windowsgui -s -w' -o Dorong.exe .
```

## 프로젝트 구조

```text
Dorong/
├─ app_windows.go
├─ core_windows.go
├─ localization.go
├─ settings.go
├─ movement.go
├─ hanging.go
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

## 참고

현재 저장소는 **개인용 비공개 프로젝트**입니다. 나중에 공개 저장소로 전환한다면 캐릭터 이미지의 재배포 권한을 먼저 확인하는 것을 권장합니다.

## 로드맵

[`docs/ROADMAP.md`](docs/ROADMAP.md)를 참고하세요.
