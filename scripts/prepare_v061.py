from pathlib import Path
import re


def replace_once(text: str, old: str, new: str, label: str) -> str:
    count = text.count(old)
    if count != 1:
        raise SystemExit(f"{label}: expected 1 match, found {count}")
    return text.replace(old, new, 1)


# app_windows.go: menu, petting behavior, status command, persisted state bootstrap.
path = Path("app_windows.go")
text = path.read_text(encoding="utf-8")
text = replace_once(
    text,
    '\tprocAppendMenuW.Call(menu, MF_STRING, ID_HOME, uintptr(unsafe.Pointer(wchar(tr("menu.home")))))\n\tprocAppendMenuW.Call(menu, MF_SEPARATOR, 0, 0)',
    '\tprocAppendMenuW.Call(menu, MF_STRING, ID_HOME, uintptr(unsafe.Pointer(wchar(tr("menu.home")))))\n\tprocAppendMenuW.Call(menu, MF_STRING, ID_AFFECTION, uintptr(unsafe.Pointer(wchar(tr("menu.affection")))))\n\tprocAppendMenuW.Call(menu, MF_SEPARATOR, 0, 0)',
    "affection menu item",
)
text = replace_once(
    text,
    '''\tcase WM_LBUTTONDBLCLK:\n\t\tif isHeadPat() {\n\t\t\tpet.petCount++\n\t\t\tpet.happyUntil = time.Now().Add(2300 * time.Millisecond)\n\t\t\tif pet.petCount%10 == 0 {\n\t\t\t\tshowBubble(tr("pet_count", pet.petCount), 2200*time.Millisecond)\n\t\t\t} else {\n\t\t\t\tshowBubble(tr("happy"), 2*time.Second)\n\t\t\t}\n\t\t} else {''',
    '''\tcase WM_LBUTTONDBLCLK:\n\t\tif isHeadPat() {\n\t\t\thandleHeadPat()\n\t\t} else {''',
    "head pat handler",
)
text = replace_once(
    text,
    '''\t\tcase ID_LANG_EN:\n\t\t\tsetLanguage(LangEN)\n\t\t\t_ = saveSettings()\n\t\t\tshowBubble(tr("language_changed_en"), 1800*time.Millisecond)\n\t\tcase ID_SETTINGS:''',
    '''\t\tcase ID_LANG_EN:\n\t\t\tsetLanguage(LangEN)\n\t\t\t_ = saveSettings()\n\t\t\tshowBubble(tr("language_changed_en"), 1800*time.Millisecond)\n\t\tcase ID_AFFECTION:\n\t\t\tshowAffectionStatus()\n\t\tcase ID_SETTINGS:''',
    "affection command",
)
text = replace_once(
    text,
    "func main() {\n\tloadSettings()\n\t_ = setStartupEnabled(startupWithWindows)",
    "func main() {\n\tloadSettings()\n\tsyncAffectionFromSettings()\n\t_ = setStartupEnabled(startupWithWindows)",
    "affection startup sync",
)
path.write_text(text, encoding="utf-8")


# core_windows.go: keep source version aligned with VERSION as well as release patching.
path = Path("core_windows.go")
text = path.read_text(encoding="utf-8")
text = re.sub(r'AppVersion\s*=\s*"[^"]+"', 'AppVersion = "0.6.1"', text, count=1)
path.write_text(text, encoding="utf-8")


# localization.go: add menu/status/level/reaction strings in both runtime languages.
path = Path("localization.go")
text = path.read_text(encoding="utf-8")
text = replace_once(
    text,
    '\t\t"menu.home":              "오른쪽 아래로 이동",\n',
    '\t\t"menu.home":              "오른쪽 아래로 이동",\n\t\t"menu.affection":         "호감도 보기",\n',
    "ko affection menu",
)
text = replace_once(
    text,
    '\t\t"pet_count":              "쓰담쓰담 %d번!",\n',
    '''\t\t"pet_count":              "쓰담쓰담 %d번!",\n\t\t"affection.level.0":      "낯가림",\n\t\t"affection.level.1":      "익숙함",\n\t\t"affection.level.2":      "친구",\n\t\t"affection.level.3":      "단짝",\n\t\t"affection.level.4":      "최고의 친구",\n\t\t"affection.level_up":     "호감도 상승! %s",\n\t\t"affection.status":       "호감도 %s · 쓰담 %d회 · 다음 %d회",\n\t\t"affection.status_max":   "호감도 %s · 쓰담 %d회!",\n\t\t"affection.react.0":      "헤헤, 좋아!",\n\t\t"affection.react.1":      "응, 이제 좀 익숙해!",\n\t\t"affection.react.2":      "또 쓰다듬어 줘!",\n\t\t"affection.react.3":      "너랑 있으면 좋아!",\n\t\t"affection.react.4":      "역시 네가 제일 좋아!",\n''',
    "ko affection strings",
)
text = replace_once(
    text,
    '\t\t"menu.home":              "Move to bottom-right",\n',
    '\t\t"menu.home":              "Move to bottom-right",\n\t\t"menu.affection":         "Show affection",\n',
    "en affection menu",
)
text = replace_once(
    text,
    '\t\t"pet_count":              "Pets: %d!",\n',
    '''\t\t"pet_count":              "Pets: %d!",\n\t\t"affection.level.0":      "Shy",\n\t\t"affection.level.1":      "Familiar",\n\t\t"affection.level.2":      "Friend",\n\t\t"affection.level.3":      "Bestie",\n\t\t"affection.level.4":      "Best friend",\n\t\t"affection.level_up":     "Affection up! %s",\n\t\t"affection.status":       "Bond %s · %d pets · next %d",\n\t\t"affection.status_max":   "Bond %s · %d pets!",\n\t\t"affection.react.0":      "Hehe, I like that!",\n\t\t"affection.react.1":      "I'm getting used to you!",\n\t\t"affection.react.2":      "Pet me again!",\n\t\t"affection.react.3":      "I like being with you!",\n\t\t"affection.react.4":      "You're my favorite!",\n''',
    "en affection strings",
)
path.write_text(text, encoding="utf-8")


# VERSION drives the release artifact/tag.
Path("VERSION").write_text("0.6.1\n", encoding="utf-8")


# README: update current release and replace the current-version change sections.
path = Path("README.md")
text = path.read_text(encoding="utf-8").replace("v0.6.0", "v0.6.1")
sections = [
    (
        "### v0.6.1 변경 사항",
        "### 주요 기능",
        '''### v0.6.1 변경 사항\n\n- 머리 더블클릭 쓰다듬기 횟수를 설정 파일에 영구 저장\n- 누적 쓰다듬기 횟수에 따라 5단계 호감도 추가: 낯가림 → 익숙함 → 친구 → 단짝 → 최고의 친구\n- 호감도 단계에 따라 쓰다듬기 반응 대사가 달라짐\n- 10 / 30 / 60 / 100회 도달 시 호감도 상승 전용 반응 표시\n- 우클릭/트레이 메뉴에 `호감도 보기`를 추가하여 현재 단계, 누적 횟수, 다음 단계 기준 확인 가능\n- 기존 v0.6.0의 설정창, 시스템 트레이, 자동 실행과 v0.5.x의 물리/멀티 모니터 기능 유지\n\n''',
    ),
    (
        "### What's new in v0.6.1",
        "### Features",
        '''### What's new in v0.6.1\n\n- Head-petting count now persists in the settings file across restarts\n- Added five affection stages based on lifetime pets: Shy → Familiar → Friend → Bestie → Best friend\n- Head-pat dialogue changes with the current affection stage\n- Dedicated affection-level-up reactions at 10 / 30 / 60 / 100 pets\n- Added `Show affection` to the right-click/tray menu to display the current stage, lifetime count, and next threshold\n- Preserved the v0.6.0 settings/tray/startup features and v0.5.x physics/multi-monitor behavior\n\n''',
    ),
    (
        "### v0.6.1 の変更点",
        "### 主な機能",
        '''### v0.6.1 の変更点\n\n- 頭をダブルクリックして撫でた累計回数を設定ファイルへ保存\n- 累計回数に応じた 5 段階の親密度を追加\n- 親密度に応じて撫でたときのセリフが変化\n- 10 / 30 / 60 / 100 回到達時にレベルアップ専用リアクションを表示\n- 右クリック／トレイメニューから現在の親密度、累計回数、次の基準を確認可能\n- v0.6.0 の設定・トレイ・自動起動と v0.5.x の物理・マルチモニター機能を維持\n\n''',
    ),
    (
        "### v0.6.1 更新",
        "### 主要功能",
        '''### v0.6.1 更新\n\n- 双击头部抚摸的累计次数现在会永久保存到设置文件\n- 根据累计抚摸次数新增 5 个好感度阶段\n- 抚摸时的台词会随当前好感度阶段变化\n- 达到 10 / 30 / 60 / 100 次时显示专属好感度提升反应\n- 可从右键／托盘菜单查看当前好感度、累计次数和下一阶段条件\n- 保留 v0.6.0 的设置、托盘、自动启动以及 v0.5.x 的物理和多显示器功能\n\n''',
    ),
]
for start, end, replacement in sections:
    pattern = re.compile(re.escape(start) + r".*?(?=" + re.escape(end) + r")", re.S)
    text, count = pattern.subn(replacement, text, count=1)
    if count != 1:
        raise SystemExit(f"README section missing: {start}")

text = text.replace(
    "- 머리 더블클릭 쓰다듬기 반응\n",
    "- 머리 더블클릭 쓰다듬기 반응\n- 누적 쓰다듬기 횟수와 5단계 호감도 영구 저장\n",
)
text = text.replace(
    "- Head-petting reaction on double-click\n",
    "- Head-petting reaction on double-click\n- Persistent lifetime pet count with five affection stages\n",
)
text = text.replace(
    "- 頭をダブルクリックすると撫で反応\n",
    "- 頭をダブルクリックすると撫で反応\n- 累計撫で回数と 5 段階の親密度を永続保存\n",
)
text = text.replace(
    "- 双击头部触发抚摸反应\n",
    "- 双击头部触发抚摸反应\n- 永久保存累计抚摸次数和 5 阶段好感度\n",
)
path.write_text(text, encoding="utf-8")


# Roadmap: mark the affection counter as delivered.
path = Path("docs/ROADMAP.md")
text = path.read_text(encoding="utf-8")
text = replace_once(
    text,
    "- v0.6.0: optional startup-with-Windows registration\n",
    "- v0.6.0: optional startup-with-Windows registration\n- v0.6.1: persistent lifetime head-petting counter\n- v0.6.1: five affection stages with stage-specific reactions and status menu\n",
    "roadmap v061",
)
text = text.replace("- Persistent affection / interaction counter\n", "")
path.write_text(text, encoding="utf-8")
