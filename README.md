# Dorong

[한국어](#lang-ko) | [English](#lang-en) | [日本語](#lang-ja) | [中文](#lang-zh)

---

<a id="lang-ko"></a>
## 한국어

**Dorong(도롱)** 은 Go와 Win32 API로 만든 개인용 Windows 데스크톱 펫입니다. Python, .NET 같은 별도 런타임 없이 실행할 수 있습니다.

> 현재 버전: **v0.6.2**

### 다운로드

[⬇ **Dorong v0.6.2 설치 프로그램 다운로드 (권장)**](https://github.com/pashkite/Dorong/releases/download/v0.6.2/Dorong-Setup-v0.6.2.exe)

[⬇ **Dorong v0.6.2 포터블 EXE 다운로드**](https://github.com/pashkite/Dorong/releases/download/v0.6.2/Dorong-v0.6.2.exe)

설치 프로그램은 `%LOCALAPPDATA%\Programs\Dorong`에 Dorong을 설치하고 시작 메뉴 및 제거 프로그램을 등록합니다. 설치 중 바탕화면 바로가기도 선택할 수 있습니다. 포터블 EXE는 설치 없이 파일 자체를 바로 실행하는 버전입니다.

### v0.6.2 긴급 수정

- v0.6.1에서 실행 직후 종료되던 손상 스프라이트 데이터 문제 수정
- 손상된 외부 스프라이트 대신 앱 내부에서 생성하는 안전한 Dorong 폴백 스프라이트 시트 사용
- 폴백 스프라이트의 크기, 투명도, 디코딩을 자동 테스트
- GitHub Release 배포 전 Windows에서 Dorong을 실제로 실행하고 8초 이상 유지되는지 검사
- 설치 프로그램을 실제로 무인 설치한 뒤 설치된 `Dorong.exe`까지 실행 검사
- 실행 검사를 통과하지 못하면 Release가 생성되지 않도록 배포 파이프라인 강화
- v0.6.1의 누적 쓰다듬기/호감도와 v0.6.0의 설정·트레이·자동실행 기능 유지

> 현재 v0.6.2의 캐릭터 그림은 손상된 원본 이미지 대신 안전하게 생성되는 내장 폴백 그래픽을 사용합니다. 원본 아트 복원은 이후 업데이트에서 별도로 진행할 수 있습니다.

### 주요 기능

- 투명 배경 애니메이션 데스크톱 펫
- 드래그 후 놓으면 중력이 적용되어 낙하
- 여러 모니터의 Windows 작업 영역과 작업표시줄 인식
- 일반 프로그램 창 상단에 착지
- 창 위를 좌우로 이동하고 창 위치/크기 변화 추적
- 창 가장자리 매달림 후 자연스럽게 낙하
- 머리 더블클릭 쓰다듬기 반응
- 누적 쓰다듬기 횟수와 5단계 호감도 영구 저장
- 마우스 커서를 바라보는 대기 반응
- 랜덤 말풍선
- 설정 가능한 집중 타이머 / 알람
- 시스템 트레이 아이콘 및 설정창
- Windows 시작 시 자동 실행 옵션
- 항상 위에 표시 옵션
- 앱 UI 한국어 / English 선택 및 설정 저장
- 네트워크 연결 및 계정 불필요

### 조작

| 조작 | 기능 |
| --- | --- |
| 왼쪽 버튼 드래그 | Dorong을 집어서 이동 |
| 드래그 후 놓기 | 중력을 적용해 떨어뜨리기 |
| 머리 더블클릭 | 쓰다듬기 |
| 오른쪽 클릭 | 메뉴 열기 |

### 빌드

Go 1.23 이상을 권장합니다.

```powershell
go build -ldflags="-H=windowsgui -s -w" -o Dorong.exe .
```

---

<a id="lang-en"></a>
## English

**Dorong** is a personal Windows desktop pet built with Go and the Win32 API. It does not require Python or .NET at runtime.

> Current version: **v0.6.2**

### Download

[⬇ **Download Dorong v0.6.2 Setup (recommended)**](https://github.com/pashkite/Dorong/releases/download/v0.6.2/Dorong-Setup-v0.6.2.exe)

[⬇ **Download Dorong v0.6.2 Portable EXE**](https://github.com/pashkite/Dorong/releases/download/v0.6.2/Dorong-v0.6.2.exe)

The installer places Dorong in `%LOCALAPPDATA%\Programs\Dorong`, creates Start Menu/uninstall registration, and offers an optional desktop shortcut. The portable EXE runs directly without installation.

### v0.6.2 emergency fix

- Fixes the immediate startup crash caused by corrupted sprite data in v0.6.1
- Uses a safe procedurally generated built-in Dorong fallback sprite sheet instead of the corrupted payload
- Adds automated sprite size, transparency, and decode coverage
- Runs Dorong on a real Windows CI runner for at least 8 seconds before publishing a Release
- Silently installs the packaged Setup build and launches the installed `Dorong.exe` before publishing
- A failed startup check now blocks Release publication
- Preserves v0.6.1 persistent petting/affection and v0.6.0 settings, tray, and startup features

> v0.6.2 currently uses safe built-in fallback artwork instead of the corrupted original sprite asset. Restoring the original artwork can be handled in a later update.

### Features

- Transparent animated desktop pet
- Gravity after drag-and-drop
- Per-monitor Windows work-area and taskbar awareness
- Landing on ordinary application-window tops
- Walking on application windows while following moved/resized support windows
- Hanging from a window edge and then falling naturally
- Head-petting reaction on double-click
- Persistent lifetime pet count with five affection stages
- Cursor-following idle reaction
- Random speech bubbles
- Configurable focus timer / alarm
- System tray icon and settings window
- Optional startup with Windows
- Always-on-top option
- Korean / English runtime UI selection with persistent preference
- No network connection or account required by the app itself

### Build

Go 1.23+ is recommended.

```powershell
go build -ldflags="-H=windowsgui -s -w" -o Dorong.exe .
```

---

<a id="lang-ja"></a>
## 日本語

**Dorong（ドロン）** は Go と Win32 API で作られた個人用 Windows デスクトップペットです。Python や .NET の追加インストールは不要です。

> 現在のバージョン: **v0.6.2**

### ダウンロード

[⬇ **Dorong v0.6.2 セットアップをダウンロード（推奨）**](https://github.com/pashkite/Dorong/releases/download/v0.6.2/Dorong-Setup-v0.6.2.exe)

[⬇ **Dorong v0.6.2 ポータブル EXE をダウンロード**](https://github.com/pashkite/Dorong/releases/download/v0.6.2/Dorong-v0.6.2.exe)

### v0.6.2 緊急修正

- v0.6.1 の破損したスプライトデータによる起動直後の終了を修正
- 破損データの代わりに安全な内蔵フォールバックスプライトを生成して使用
- スプライトのサイズ・透明度・デコードを自動テスト
- Release 前に Windows 上で実際に起動し、8 秒以上動作することを確認
- セットアップ版も実際にインストールして、インストール済み `Dorong.exe` を起動確認
- 起動確認に失敗した場合は Release を公開しないよう CI を強化
- v0.6.1 の親密度機能と v0.6.0 の設定・トレイ・自動起動機能を維持

> v0.6.2 は破損した元スプライトの代わりに安全な内蔵フォールバック画像を使用します。

### 主な機能

- 透明背景のアニメーションデスクトップペット
- ドラッグして離すと重力で落下
- 複数モニターごとの Windows 作業領域とタスクバーを認識
- アプリウィンドウ上部への着地
- ウィンドウ上を左右に歩行
- ウィンドウ端へのぶら下がりと自然な落下
- 頭をダブルクリックすると撫で反応
- 累計撫で回数と 5 段階の親密度を永続保存
- マウスカーソルを見る待機反応
- ランダム吹き出し
- 設定可能な集中タイマー / アラーム
- システムトレイアイコンと設定画面
- Windows 起動時の自動実行オプション
- 常に最前面表示
- アプリ UI は現在 한국어 / English を選択可能

---

<a id="lang-zh"></a>
## 中文

**Dorong** 是一个使用 Go 和 Win32 API 制作的个人 Windows 桌面宠物。运行时无需另外安装 Python 或 .NET。

> 当前版本: **v0.6.2**

### 下载

[⬇ **下载 Dorong v0.6.2 安装程序（推荐）**](https://github.com/pashkite/Dorong/releases/download/v0.6.2/Dorong-Setup-v0.6.2.exe)

[⬇ **下载 Dorong v0.6.2 便携版 EXE**](https://github.com/pashkite/Dorong/releases/download/v0.6.2/Dorong-v0.6.2.exe)

### v0.6.2 紧急修复

- 修复 v0.6.1 因损坏精灵数据导致的启动后立即退出问题
- 不再使用损坏数据，改为生成安全的内置 Dorong 备用精灵图
- 自动测试精灵图尺寸、透明度和解码
- 发布 Release 前会在 Windows 环境中实际启动 Dorong 并确认至少运行 8 秒
- 安装程序也会被实际静默安装，并启动安装后的 `Dorong.exe` 进行检查
- 启动检查失败时不会发布 Release
- 保留 v0.6.1 好感度功能以及 v0.6.0 设置、托盘和自动启动功能

> v0.6.2 当前使用安全的内置备用图形代替已损坏的原始精灵素材。

### 主要功能

- 透明背景动画桌面宠物
- 拖动松手后应用重力
- 识别多显示器各自的 Windows 工作区域和任务栏
- 可落在普通应用窗口顶部
- 可在窗口顶部左右移动并跟随窗口位置/尺寸变化
- 可在窗口边缘悬挂后自然掉落
- 双击头部触发抚摸反应
- 永久保存累计抚摸次数和 5 阶段好感度
- 待机时会看向鼠标光标
- 随机气泡消息
- 可设置的专注计时器 / 闹钟
- 系统托盘图标和设置窗口
- Windows 启动时自动运行选项
- 置顶选项
- 应用界面当前支持 한국어 / English

---

## Repository notes

The Windows build is generated by GitHub Actions from the value in `VERSION`. A Release is published only after automated tests, a Windows startup smoke test, and an install-and-launch smoke test for the packaged Setup build pass successfully.

See [`docs/ROADMAP.md`](docs/ROADMAP.md) for the development roadmap.
