from pathlib import Path
import re


def rep(text, old, new, label):
    n = text.count(old)
    if n != 1:
        raise SystemExit(f"{label}: expected 1 match, found {n}")
    return text.replace(old, new, 1)

app_path = Path("app_windows.go")
app = app_path.read_text(encoding="utf-8")
app = rep(app, 'wchar(tr("menu.focus"))', 'wchar(tr("menu.focus", focusMinutes))', "focus menu")
app = rep(app, 'wchar(tr("menu.alarm"))', 'wchar(tr("menu.alarm", alarmMinutes))', "alarm menu")
app = rep(app, '\tprocAppendMenuW.Call(menu, MF_SEPARATOR, 0, 0)\n\tprocAppendMenuW.Call(menu, MF_STRING, ID_EXIT, uintptr(unsafe.Pointer(wchar(tr("menu.exit")))))', '\tprocAppendMenuW.Call(menu, MF_SEPARATOR, 0, 0)\n\tprocAppendMenuW.Call(menu, MF_STRING, ID_SETTINGS, uintptr(unsafe.Pointer(wchar(tr("menu.settings")))))\n\tprocAppendMenuW.Call(menu, MF_STRING, ID_EXIT, uintptr(unsafe.Pointer(wchar(tr("menu.exit")))))', "settings menu")
app = rep(app, '\tcase WM_DESTROY:\n\t\tprocKillTimer.Call(hwnd, 1)', '\tcase WM_DESTROY:\n\t\tremoveTrayIcon()\n\t\tprocKillTimer.Call(hwnd, 1)', "tray cleanup")
app = rep(app, '\tcase WM_RBUTTONUP:\n\t\tpopup(hwnd)\n\t\treturn 0\n\tcase WM_COMMAND:', '\tcase WM_RBUTTONUP:\n\t\tpopup(hwnd)\n\t\treturn 0\n\tcase wmTrayIcon:\n\t\tevent := uint32(lParam & 0xffff)\n\t\tif event == WM_RBUTTONUP {\n\t\t\tpopup(hwnd)\n\t\t} else if event == WM_LBUTTONDBLCLK {\n\t\t\tshowSettingsWindow()\n\t\t}\n\t\treturn 0\n\tcase WM_COMMAND:', "tray callback")
app = rep(app, 'pet.focusUntil = time.Now().Add(25 * time.Minute)', 'pet.focusUntil = time.Now().Add(time.Duration(focusMinutes) * time.Minute)', "focus duration")
app = rep(app, 'showBubble(tr("focus_started"), 3*time.Second)', 'showBubble(tr("focus_started", focusMinutes), 3*time.Second)', "focus text")
app = rep(app, 'pet.alarmUntil = time.Now().Add(10 * time.Minute)', 'pet.alarmUntil = time.Now().Add(time.Duration(alarmMinutes) * time.Minute)', "alarm duration")
app = rep(app, 'showBubble(tr("alarm_set"), 3*time.Second)', 'showBubble(tr("alarm_set", alarmMinutes), 3*time.Second)', "alarm text")
app = rep(app, '\t\tcase ID_LANG_EN:\n\t\t\tsetLanguage(LangEN)\n\t\t\t_ = saveSettings()\n\t\t\tshowBubble(tr("language_changed_en"), 1800*time.Millisecond)\n\t\tcase ID_EXIT:', '\t\tcase ID_LANG_EN:\n\t\t\tsetLanguage(LangEN)\n\t\t\t_ = saveSettings()\n\t\t\tshowBubble(tr("language_changed_en"), 1800*time.Millisecond)\n\t\tcase ID_SETTINGS:\n\t\t\tshowSettingsWindow()\n\t\tcase ID_EXIT:', "settings command")
app = rep(app, 'func main() {\n\tloadSettings()\n\trand.Seed', 'func main() {\n\tloadSettings()\n\t_ = setStartupEnabled(startupWithWindows)\n\trand.Seed', "startup sync")
app = rep(app, '\tprocUpdateWindow.Call(hwnd)\n\tprocSetTimer.Call(hwnd, 1, 40, 0)', '\tprocUpdateWindow.Call(hwnd)\n\taddTrayIcon()\n\tprocSetTimer.Call(hwnd, 1, 40, 0)', "tray add")
app_path.write_text(app, encoding="utf-8")

loc_path = Path("localization.go")
loc = loc_path.read_text(encoding="utf-8")
for old, new, label in [
    ('"menu.focus":          "25분 집중 시작",', '"menu.focus":          "%d분 집중 시작",', "ko focus menu"),
    ('"menu.alarm":          "10분 알람 설정",', '"menu.alarm":          "%d분 알람 설정",', "ko alarm menu"),
    ('"focus_started":       "25분 집중 시작!",', '"focus_started":       "%d분 집중 시작!",', "ko focus start"),
    ('"alarm_set":           "10분 뒤에 알려줄게.",', '"alarm_set":           "%d분 뒤에 알려줄게.",', "ko alarm set"),
    ('"focus_dialog":        "25분 집중이 끝났어! 잠깐 쉬자.",', '"focus_dialog":        "집중 시간이 끝났어! 잠깐 쉬자.",', "ko focus dialog"),
    ('"alarm_dialog":        "10분 알람 시간이 됐어!",', '"alarm_dialog":        "알람 시간이 됐어!",', "ko alarm dialog"),
    ('"menu.focus":          "Start 25-minute focus",', '"menu.focus":          "Start %d-minute focus",', "en focus menu"),
    ('"menu.alarm":          "Set 10-minute alarm",', '"menu.alarm":          "Set %d-minute alarm",', "en alarm menu"),
    ('"focus_started":       "25-minute focus started!",', '"focus_started":       "%d-minute focus started!",', "en focus start"),
    ('"alarm_set":           "I\'ll tell you in 10 minutes.",', '"alarm_set":           "I\'ll tell you in %d minutes.",', "en alarm set"),
    ('"focus_dialog":        "Your 25-minute focus session is over. Take a short break!",', '"focus_dialog":        "Your focus session is over. Take a short break!",', "en focus dialog"),
    ('"alarm_dialog":        "Your 10-minute alarm is up!",', '"alarm_dialog":        "Your alarm is up!",', "en alarm dialog"),
]:
    loc = rep(loc, old, new, label)
loc = rep(loc, '\t\t"menu.exit":           "Dorong 종료",', '\t\t"menu.settings":       "설정...",\n\t\t"menu.exit":           "Dorong 종료",', "ko settings menu")
loc = rep(loc, '\t\t"menu.exit":           "Exit Dorong",', '\t\t"menu.settings":       "Settings...",\n\t\t"menu.exit":           "Exit Dorong",', "en settings menu")
loc = rep(loc, '\t\t"language_changed_en": "영어로 바꿨어!",', '\t\t"language_changed_en": "영어로 바꿨어!",\n\t\t"settings.title":      "Dorong 설정",\n\t\t"settings.focus_label": "집중 시간 (분, 1~180)",\n\t\t"settings.alarm_label": "알람 시간 (분, 1~1440)",\n\t\t"settings.startup":    "Windows 시작 시 Dorong 자동 실행",\n\t\t"settings.save":       "저장",\n\t\t"settings.cancel":     "취소",\n\t\t"settings.saved":      "설정을 저장했어!",\n\t\t"settings.invalid":    "시간은 표시된 범위 안의 숫자로 입력해줘.",\n\t\t"settings.startup_error": "Windows 자동 실행 설정을 변경하지 못했어.",\n\t\t"settings.save_error": "설정 파일을 저장하지 못했어.",\n\t\t"settings.open_error": "설정창을 열지 못했어.",', "ko settings strings")
loc = rep(loc, '\t\t"language_changed_en": "Switched to English!",', '\t\t"language_changed_en": "Switched to English!",\n\t\t"settings.title":      "Dorong Settings",\n\t\t"settings.focus_label": "Focus minutes (1-180)",\n\t\t"settings.alarm_label": "Alarm minutes (1-1440)",\n\t\t"settings.startup":    "Start Dorong with Windows",\n\t\t"settings.save":       "Save",\n\t\t"settings.cancel":     "Cancel",\n\t\t"settings.saved":      "Settings saved!",\n\t\t"settings.invalid":    "Enter a number within the shown range.",\n\t\t"settings.startup_error": "Could not change the Windows startup setting.",\n\t\t"settings.save_error": "Could not save the settings file.",\n\t\t"settings.open_error": "Could not open the settings window.",', "en settings strings")
loc_path.write_text(loc, encoding="utf-8")

Path("VERSION").write_text("0.6.0\n", encoding="utf-8")

roadmap_path = Path("docs/ROADMAP.md")
roadmap = roadmap_path.read_text(encoding="utf-8")
roadmap = rep(roadmap, '# Dorong Roadmap\n\n## v0.5.x — current\n', '# Dorong Roadmap\n\n## v0.6.x — current\n\n- v0.6.0: configurable focus/alarm duration UI with persistent settings\n- v0.6.0: native Windows settings window and system-tray integration\n- v0.6.0: optional startup-with-Windows registration\n\n## v0.5.x — completed\n', "roadmap heading")
roadmap = roadmap.replace('- Configurable focus/alarm duration UI\n', '')
roadmap = roadmap.replace('- System tray icon and settings window\n', '')
roadmap = roadmap.replace('- Startup-with-Windows option\n', '')
roadmap_path.write_text(roadmap, encoding="utf-8")

readme_path = Path("README.md")
readme = readme_path.read_text(encoding="utf-8").replace("v0.5.9", "v0.6.0")
sections = [
("### v0.6.0 변경 사항", "### 주요 기능", """### v0.6.0 변경 사항

- 네이티브 Windows 설정창 추가
- 집중 타이머 시간을 1~180분 범위에서 직접 설정 가능
- 알람 시간을 1~1440분 범위에서 직접 설정 가능
- 집중/알람 시간과 자동 실행 설정을 기존 언어 설정과 함께 영구 저장
- Windows 시스템 트레이 아이콘 추가: 우클릭으로 기존 메뉴, 더블클릭으로 설정창 열기
- Windows 시작 시 Dorong 자동 실행 옵션 추가
- 기존 단일 EXE 구조, 멀티 모니터, 창 추적, 중력/애니메이션 동작 유지

"""),
("### What's new in v0.6.0", "### Features", """### What's new in v0.6.0

- Added a native Windows settings window
- Focus duration is configurable from 1 to 180 minutes
- Alarm duration is configurable from 1 to 1440 minutes
- Timer and startup preferences persist together with the language setting
- Added a system-tray icon: right-click opens the existing menu and double-click opens Settings
- Added an optional Start Dorong with Windows setting
- Preserved the standalone EXE architecture, multi-monitor behavior, support-window tracking, gravity, and animations

"""),
("### v0.6.0 の変更点", "### 主な機能", """### v0.6.0 の変更点

- Windows ネイティブの設定ウィンドウを追加
- 集中タイマーを 1〜180 分の範囲で設定可能
- アラームを 1〜1440 分の範囲で設定可能
- タイマー時間と自動起動設定を言語設定と一緒に保存
- システムトレイアイコンを追加：右クリックで既存メニュー、ダブルクリックで設定画面
- Windows 起動時に Dorong を自動実行するオプションを追加
- 単一 EXE、マルチモニター、ウィンドウ追従、重力・アニメーション動作を維持

"""),
("### v0.6.0 更新", "### 主要功能", """### v0.6.0 更新

- 新增 Windows 原生设置窗口
- 专注计时可在 1～180 分钟范围内自定义
- 闹钟可在 1～1440 分钟范围内自定义
- 计时器和开机启动设置会与语言设置一起持久保存
- 新增系统托盘图标：右键打开原有菜单，双击打开设置窗口
- 新增 Windows 启动时自动运行 Dorong 的选项
- 保持单 EXE、多显示器、窗口跟随、重力与动画行为不变

"""),
]
for start, end, body in sections:
    pattern = re.compile(re.escape(start) + r'.*?(?=' + re.escape(end) + r')', re.S)
    readme, n = pattern.subn(body, readme, count=1)
    if n != 1:
        raise SystemExit(f"README section missing: {start}")
readme = readme.replace('- 25분 집중 타이머 / 10분 알람\n', '- 설정 가능한 집중 타이머 / 알람\n- 시스템 트레이 아이콘 및 설정창\n- Windows 시작 시 자동 실행 옵션\n')
readme = readme.replace('- 25-minute focus timer / 10-minute alarm\n', '- Configurable focus timer / alarm\n- System tray icon and settings window\n- Optional startup with Windows\n')
readme = readme.replace('- 25 分集中タイマー / 10 分アラーム\n', '- 設定可能な集中タイマー / アラーム\n- システムトレイアイコンと設定画面\n- Windows 起動時の自動実行オプション\n')
readme = readme.replace('- 25 分钟专注计时器 / 10 分钟闹钟\n', '- 可设置的专注计时器 / 闹钟\n- 系统托盘图标和设置窗口\n- Windows 启动时自动运行选项\n')
readme_path.write_text(readme, encoding="utf-8")
