# Dorong

[한국어](README.md) | [영어](README.en.md) | [일본어](README.ja.md) | **중국어**

**Dorong（多龙）** 是一个使用 Go 和 Win32 API 制作的个人用 Windows 桌面宠物。
无需安装 Python、.NET 等额外运行环境，只需运行 `Dorong.exe` 即可使用。

> 当前版本：**v0.5.5**

## 主要功能

- 透明背景的动画桌面宠物
- 使用鼠标拖动 Dorong 到任意位置
- 拖动后松开时应用重力并向下坠落
- 识别 Windows 工作区域，并停在任务栏上方
- 能准确落在普通应用窗口的顶部
- 即使是较窄的窗口边缘，也会通过左脚、中间、右脚三个判定点进行识别
- 可在应用窗口上左右行走，并在窗口移动或缩放后重新计算活动范围
- 偶尔会走到窗口边缘并悬挂在真正的边缘位置
- 悬挂期间如果窗口移动或缩放，Dorong 会跟随；如果支撑窗口消失，则会自然坠落
- 双击头部区域可触发抚摸反应
- 待机时会注视鼠标指针
- 待机 / 行走 / 睡眠 / 开心 / 被提起 / 专注 / 坠落 / 悬挂动画
- 随机对话气泡
- 25 分钟专注计时器
- 10 分钟提醒
- 始终置顶选项
- **应用界面语言：韩语 / English**
- 保存所选语言，并在下次启动时继续使用
- 单个 64 位 Windows 可执行文件
- 不需要网络连接或账号

## 操作方式

| 操作 | 功能 |
| --- | --- |
| 按住左键拖动 | 抓起并移动 Dorong |
| 拖动后松开 | 放下 Dorong，并应用重力 |
| 双击头部 | 抚摸 Dorong |
| 右键点击 | 打开交互菜单 |

## 语言

本 README 可通过页面顶部的链接切换为 **韩语 / 英语 / 日语 / 中文**。

Dorong 应用内界面目前支持 **韩语和英语**。右键点击 Dorong，然后在菜单底部选择 `한국어` 或 `English`。
菜单、对话气泡、专注计时器、提醒信息和系统消息会立即切换为所选语言。语言设置会保存在 Windows 用户配置目录下的 `Dorong/settings.json` 中，并在下次启动时继续使用。

## 运行

可以运行 GitHub Actions 生成的 Windows 构建文件，或发行版/本地构建中的 `Dorong.exe`。

如果使用已经构建好的可执行文件，**无需安装 Python、Go 或 .NET。**

## 从源代码构建

建议使用 Go 1.23 或更高版本。

```powershell
go build -ldflags="-H=windowsgui -s -w" -o Dorong.exe .
```

在 Linux/macOS 上交叉编译 Windows 版本：

```bash
GOOS=windows GOARCH=amd64 go build -ldflags='-H=windowsgui -s -w' -o Dorong.exe .
```

## 项目结构

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

## 说明

此仓库目前是一个 **个人使用的私有项目**。如果以后改为公开仓库，建议先确认角色图片的再分发权限。

## 路线图

请参阅 [`docs/ROADMAP.md`](docs/ROADMAP.md)。
