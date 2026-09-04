# Dorong

[한국어](#lang-ko) | [English](#lang-en) | [日本語](#lang-ja) | [中文](#lang-zh)

---

<a id="lang-ko"></a>
## 한국어

**Dorong(도롱)** 은 Go와 Win32 API로 만든 개인용 Windows 데스크톱 펫입니다. Python, .NET 같은 별도 런타임 없이 단일 EXE로 실행할 수 있습니다.

> 현재 버전: **v0.6.1**

### 다운로드

[⬇ **Dorong v0.6.1 EXE 다운로드**](https://github.com/pashkite/Dorong/releases/download/v0.6.1/Dorong-v0.6.1.exe)

Windows 64비트용 단일 실행 파일입니다. 저장소가 비공개인 동안에는 GitHub에 로그인한 계정만 릴리스 파일을 받을 수 있습니다.

### v0.6.1 변경 사항

- 머리 더블클릭 쓰다듬기 횟수를 설정 파일에 영구 저장
- 누적 쓰다듬기 횟수에 따라 5단계 호감도 추가: 낯가림 → 익숙함 → 친구 → 단짝 → 최고의 친구
- 호감도 단계에 따라 쓰다듬기 반응 대사가 달라짐
- 10 / 30 / 60 / 100회 도달 시 호감도 상승 전용 반응 표시
- 우클릭/트레이 메뉴에 `호감도 보기`를 추가하여 현재 단계, 누적 횟수, 다음 단계 기준 확인 가능
- 기존 v0.6.0의 설정창, 시스템 트레이, 자동 실행과 v0.5.x의 물리/멀티 모니터 기능 유지

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

**Dorong** is a personal Windows desktop pet built with Go and the Win32 API. The prebuilt version runs as a single standalone EXE without requiring Python or .NET.

> Current version: **v0.6.1**

### Download

[⬇ **Download Dorong v0.6.1 EXE**](https://github.com/pashkite/Dorong/releases/download/v0.6.1/Dorong-v0.6.1.exe)

This is a standalone Windows 64-bit executable. While the repository is private, you need to be signed in to an authorized GitHub account to download the release asset.

### What's new in v0.6.1

- Head-petting count now persists in the settings file across restarts
- Added five affection stages based on lifetime pets: Shy → Familiar → Friend → Bestie → Best friend
- Head-pat dialogue changes with the current affection stage
- Dedicated affection-level-up reactions at 10 / 30 / 60 / 100 pets
- Added `Show affection` to the right-click/tray menu to display the current stage, lifetime count, and next threshold
- Preserved the v0.6.0 settings/tray/startup features and v0.5.x physics/multi-monitor behavior

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

**Dorong（ドロン）** は Go と Win32 API で作られた個人用 Windows デスクトップペットです。ビルド済み EXE は Python や .NET を追加インストールせずに単体で実行できます。

> 現在のバージョン: **v0.6.1**

### ダウンロード

[⬇ **Dorong v0.6.1 EXE をダウンロード**](https://github.com/pashkite/Dorong/releases/download/v0.6.1/Dorong-v0.6.1.exe)

Windows 64bit 用の単体実行ファイルです。リポジトリが非公開の間は、アクセス権のある GitHub アカウントでログインする必要があります。

### v0.6.1 の変更点

- 頭をダブルクリックして撫でた累計回数を設定ファイルへ保存
- 累計回数に応じた 5 段階の親密度を追加
- 親密度に応じて撫でたときのセリフが変化
- 10 / 30 / 60 / 100 回到達時にレベルアップ専用リアクションを表示
- 右クリック／トレイメニューから現在の親密度、累計回数、次の基準を確認可能
- v0.6.0 の設定・トレイ・自動起動と v0.5.x の物理・マルチモニター機能を維持

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

**Dorong** 是一个使用 Go 和 Win32 API 制作的个人 Windows 桌面宠物。预编译版本可作为单个 EXE 运行，无需另外安装 Python 或 .NET。

> 当前版本: **v0.6.1**

### 下载

[⬇ **下载 Dorong v0.6.1 EXE**](https://github.com/pashkite/Dorong/releases/download/v0.6.1/Dorong-v0.6.1.exe)

这是 Windows 64 位单文件程序。在仓库保持私有期间，需要使用具有访问权限的 GitHub 账号登录后才能下载 Release 文件。

### v0.6.1 更新

- 双击头部抚摸的累计次数现在会永久保存到设置文件
- 根据累计抚摸次数新增 5 个好感度阶段
- 抚摸时的台词会随当前好感度阶段变化
- 达到 10 / 30 / 60 / 100 次时显示专属好感度提升反应
- 可从右键／托盘菜单查看当前好感度、累计次数和下一阶段条件
- 保留 v0.6.0 的设置、托盘、自动启动以及 v0.5.x 的物理和多显示器功能

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

The Windows build is generated by GitHub Actions from the value in `VERSION`. Successful builds publish `Dorong-v<version>.exe` to the matching GitHub Release automatically.

See [`docs/ROADMAP.md`](docs/ROADMAP.md) for the development roadmap.
