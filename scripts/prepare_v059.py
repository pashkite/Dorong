from pathlib import Path
import re


def require_replace(text: str, old: str, new: str, label: str, count: int = 1) -> str:
    actual = text.count(old)
    if actual != count:
        raise SystemExit(f"{label}: expected {count} match(es), found {actual}")
    return text.replace(old, new, count)


core_path = Path("core_windows.go")
core = core_path.read_text(encoding="utf-8")
core = require_replace(core, 'AppVersion = "0.5.8"', 'AppVersion = "0.5.9"', "AppVersion")
core = require_replace(
    core,
    '\tprocIsWindowVisible       = user32.NewProc("IsWindowVisible")\n',
    '\tprocIsWindowVisible       = user32.NewProc("IsWindowVisible")\n\tprocIsIconic              = user32.NewProc("IsIconic")\n',
    "IsIconic proc",
)
core = require_replace(
    core,
    '''\tok, _, _ := procIsWindow.Call(hwnd)\n\tvis, _, _ := procIsWindowVisible.Call(hwnd)\n\tif ok == 0 || vis == 0 {\n\t\treturn RECT{}, false\n\t}\n''',
    '''\tok, _, _ := procIsWindow.Call(hwnd)\n\tvis, _, _ := procIsWindowVisible.Call(hwnd)\n\ticonic, _, _ := procIsIconic.Call(hwnd)\n\tif ok == 0 || vis == 0 || iconic != 0 {\n\t\treturn RECT{}, false\n\t}\n''',
    "reject minimized support windows",
)
core = require_replace(
    core,
    '''\t\t\t\tpet.supportHwnd = h\n\t\t\t\tsetPos(x, landingY)\n''',
    '''\t\t\t\tpet.supportHwnd = h\n\t\t\t\trememberSupportWindow(h, r)\n\t\t\t\tsetPos(x, landingY)\n''',
    "remember landed support",
)
core_path.write_text(core, encoding="utf-8")

app_path = Path("app_windows.go")
app = app_path.read_text(encoding="utf-8")
app = require_replace(
    app,
    '''\tcase WM_DESTROY:\n\t\tprocKillTimer.Call(hwnd, 1)\n''',
    '''\tcase WM_DESTROY:\n\t\tprocKillTimer.Call(hwnd, 1)\n\t\tprocKillTimer.Call(hwnd, 2)\n''',
    "kill support timer",
)
app = require_replace(
    app,
    '''\tcase WM_TIMER:\n\t\ttick()\n\t\treturn 0\n''',
    '''\tcase WM_TIMER:\n\t\tif wParam == 2 {\n\t\t\tupdateSupportMotion()\n\t\t\treturn 0\n\t\t}\n\t\ttick()\n\t\treturn 0\n''',
    "support timer handler",
)
app = require_replace(
    app,
    '''\tprocSetTimer.Call(hwnd, 1, 40, 0)\n''',
    '''\tprocSetTimer.Call(hwnd, 1, 40, 0)\n\tprocSetTimer.Call(hwnd, 2, 16, 0)\n''',
    "support timer setup",
)
app_path.write_text(app, encoding="utf-8")

support_logic = r'''package main

const supportFootInset int32 = 42

// supportXBounds returns the horizontal range in which Dorong's visible feet
// remain over a supporting application window while also staying on the
// monitor's usable work area.
func supportXBounds(support, area ScreenRect, petWidth int32) (lo, hi int32) {
	lo = support.Left - petWidth/2 + supportFootInset
	hi = support.Right - petWidth/2 - supportFootInset
	if lo < area.Left {
		lo = area.Left
	}
	maxX := area.Right - petWidth
	if hi > maxX {
		hi = maxX
	}
	if hi < lo {
		hi = lo
	}
	return lo, hi
}

func clampSupportX(v, lo, hi int32) int32 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// followSupportX keeps Dorong visually attached to a moving support window.
// Pure translations shift Dorong and an active walking target by the same
// amount. During resize, Dorong keeps its screen X where possible and is only
// clamped back onto the resized window, which avoids sideways jumps when a
// single edge is dragged.
func followSupportX(petX, targetX int32, walking bool, previous, current, area ScreenRect, petWidth int32) (x, nextTarget int32) {
	x = petX
	nextTarget = targetX
	oldWidth := previous.Right - previous.Left
	newWidth := current.Right - current.Left
	if oldWidth == newWidth {
		dx := current.Left - previous.Left
		x += dx
		if walking {
			nextTarget += dx
		}
	}

	lo, hi := supportXBounds(current, area, petWidth)
	x = clampSupportX(x, lo, hi)
	if walking {
		nextTarget = clampSupportX(nextTarget, lo, hi)
	}
	return x, nextTarget
}
'''
Path("support_tracking.go").write_text(support_logic, encoding="utf-8")

support_windows = r'''//go:build windows

package main

var (
	trackedSupportHwnd uintptr
	trackedSupportRect RECT
	trackedSupportOK   bool
)

func rememberSupportWindow(hwnd uintptr, r RECT) {
	trackedSupportHwnd = hwnd
	trackedSupportRect = r
	trackedSupportOK = hwnd != 0
}

func clearSupportWindowTracking() {
	trackedSupportHwnd = 0
	trackedSupportRect = RECT{}
	trackedSupportOK = false
}

// updateSupportMotion runs on a dedicated ~60 Hz timer. The normal 40 ms pet
// tick keeps gameplay/animation timing unchanged, while this lightweight path
// only follows the window Dorong is standing on or hanging from.
func updateSupportMotion() {
	if pet.supportHwnd == 0 {
		clearSupportWindowTracking()
		return
	}
	if pet.dragging || pet.falling {
		return
	}

	r, ok := validSupportWindow(pet.supportHwnd)
	if !ok {
		startFall(0)
		clearSupportWindowTracking()
		return
	}

	if pet.hanging {
		if !syncHangPose() {
			startFall(0)
			clearSupportWindowTracking()
			return
		}
		rememberSupportWindow(pet.supportHwnd, r)
		return
	}

	x, y := currentPos()
	wa := workAreaForWindow(pet.supportHwnd)
	area := ScreenRect{Left: wa.Left, Top: wa.Top, Right: wa.Right, Bottom: wa.Bottom}
	current := ScreenRect{Left: r.Left, Top: r.Top, Right: r.Right, Bottom: r.Bottom}

	if !trackedSupportOK || trackedSupportHwnd != pet.supportHwnd {
		rememberSupportWindow(pet.supportHwnd, r)
		lo, hi := supportXBounds(current, area, PET_W)
		nx := clampSupportX(x, lo, hi)
		ny := r.Top - PET_H
		if nx != x || ny != y {
			setPos(nx, ny)
		}
		return
	}

	previous := ScreenRect{Left: trackedSupportRect.Left, Top: trackedSupportRect.Top, Right: trackedSupportRect.Right, Bottom: trackedSupportRect.Bottom}
	nx, nextTarget := followSupportX(x, pet.targetX, pet.walking, previous, current, area, PET_W)
	if pet.walking {
		pet.targetX = nextTarget
	}
	ny := r.Top - PET_H
	if nx != x || ny != y {
		setPos(nx, ny)
	}
	rememberSupportWindow(pet.supportHwnd, r)
}
'''
Path("support_tracking_windows.go").write_text(support_windows, encoding="utf-8")

support_tests = r'''package main

import "testing"

func TestFollowSupportXPureMove(t *testing.T) {
	area := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040}
	oldRect := ScreenRect{Left: 300, Top: 200, Right: 1100, Bottom: 800}
	newRect := ScreenRect{Left: 420, Top: 260, Right: 1220, Bottom: 860}
	x, target := followSupportX(520, 760, true, oldRect, newRect, area, 236)
	if x != 640 || target != 880 {
		t.Fatalf("pure move = (%d,%d), want (640,880)", x, target)
	}
}

func TestFollowSupportXResizeDoesNotJump(t *testing.T) {
	area := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040}
	oldRect := ScreenRect{Left: 300, Top: 200, Right: 1100, Bottom: 800}
	newRect := ScreenRect{Left: 420, Top: 200, Right: 1100, Bottom: 800}
	x, _ := followSupportX(650, 0, false, oldRect, newRect, area, 236)
	if x != 650 {
		t.Fatalf("left-edge resize moved pet to %d, want 650", x)
	}
}

func TestFollowSupportXResizeClampsToWindow(t *testing.T) {
	area := ScreenRect{Left: 0, Top: 0, Right: 1920, Bottom: 1040}
	oldRect := ScreenRect{Left: 300, Top: 200, Right: 1300, Bottom: 800}
	newRect := ScreenRect{Left: 300, Top: 200, Right: 700, Bottom: 800}
	x, _ := followSupportX(1000, 0, false, oldRect, newRect, area, 236)
	_, hi := supportXBounds(newRect, area, 236)
	if x != hi {
		t.Fatalf("narrow resize x=%d, want clamp %d", x, hi)
	}
}

func TestSupportXBoundsNegativeMonitor(t *testing.T) {
	area := ScreenRect{Left: -1920, Top: 0, Right: 0, Bottom: 1040}
	support := ScreenRect{Left: -1600, Top: 180, Right: -500, Bottom: 900}
	lo, hi := supportXBounds(support, area, 236)
	if lo >= 0 || hi >= 0 || hi <= lo {
		t.Fatalf("negative-monitor bounds = (%d,%d), want ordered negative coordinates", lo, hi)
	}
}
'''
Path("support_tracking_test.go").write_text(support_tests, encoding="utf-8")

Path("VERSION").write_text("0.5.9\n", encoding="utf-8")

roadmap_path = Path("docs/ROADMAP.md")
roadmap = roadmap_path.read_text(encoding="utf-8")
roadmap = require_replace(
    roadmap,
    "- v0.5.8: cross-monitor drag, landing, window walking, hanging, and speech-bubble placement\n",
    "- v0.5.8: cross-monitor drag, landing, window walking, hanging, and speech-bubble placement\n"
    "- v0.5.9: dedicated ~60 Hz support-window tracking without changing gravity/animation timing\n"
    "- v0.5.9: horizontal follow on moved windows, resize clamping, and minimized-window fall handling\n",
    "roadmap v0.5.9",
)
roadmap = require_replace(
    roadmap,
    "- Smoother support-window tracking while windows are moved or resized\n",
    "",
    "remove completed tracking roadmap item",
)
roadmap_path.write_text(roadmap, encoding="utf-8")

readme_path = Path("README.md")
readme = readme_path.read_text(encoding="utf-8")
readme = readme.replace("Current version: **v0.5.8**", "Current version: **v0.5.9**")
readme = readme.replace("현재 버전: **v0.5.8**", "현재 버전: **v0.5.9**")
readme = readme.replace("現在のバージョン: **v0.5.8**", "現在のバージョン: **v0.5.9**")
readme = readme.replace("当前版本: **v0.5.8**", "当前版本: **v0.5.9**")
readme = readme.replace("Dorong v0.5.8 EXE", "Dorong v0.5.9 EXE")
readme = readme.replace("Dorong-v0.5.8.exe", "Dorong-v0.5.9.exe")
readme = readme.replace("/v0.5.8/", "/v0.5.9/")
readme = readme.replace("Download Dorong v0.5.8 EXE", "Download Dorong v0.5.9 EXE")
readme = readme.replace("Dorong v0.5.8 EXE をダウンロード", "Dorong v0.5.9 EXE をダウンロード")
readme = readme.replace("下载 Dorong v0.5.8 EXE", "下载 Dorong v0.5.9 EXE")

sections = [
    (
        "### v0.5.8 변경 사항", "### 주요 기능",
        """### v0.5.9 변경 사항

- 창 위에 서 있거나 매달린 도롱을 별도 약 60Hz 추적 타이머로 더 부드럽게 동기화
- 창을 좌우로 이동하면 도롱의 X 좌표도 같은 거리만큼 즉시 따라가도록 개선
- 도롱이 걷는 중 창이 이동해도 목표 위치가 함께 이동하여 걷기 방향과 목적지가 어긋나지 않음
- 창 크기 변경 시 불필요하게 옆으로 튀지 않고, 새 창 범위를 벗어날 때만 자연스럽게 안쪽으로 보정
- 창 상단 높이가 바뀌면 서 있는 높이를 즉시 다시 맞춤
- 도롱이 올라선 창을 최소화하면 비정상 좌표로 따라가지 않고 자연스럽게 낙하
- v0.5.8의 멀티 모니터 작업 영역 지원 유지

""",
    ),
    (
        "### What's new in v0.5.8", "### Features",
        """### What's new in v0.5.9

- Added a dedicated ~60 Hz support-window tracker for smoother standing and hanging motion
- Dorong now follows horizontal window movement by the same delta instead of lagging behind
- Active walking targets move together with a translated support window
- Resizing preserves Dorong's screen X when possible and only clamps it when the new window bounds require it
- Vertical position immediately follows changes to the support window's top edge
- Minimizing the supporting window now makes Dorong fall naturally instead of following invalid minimized coordinates
- Preserved the v0.5.8 multi-monitor work-area behavior

""",
    ),
    (
        "### v0.5.8 の変更点", "### 主な機能",
        """### v0.5.9 の変更点

- ウィンドウ上に立っている／ぶら下がっている Dorong を専用の約 60Hz タイマーでより滑らかに追従
- ウィンドウを左右に移動すると Dorong の X 座標も同じ距離だけ追従
- 歩行中にウィンドウが移動しても目標位置を同時にずらし、進行方向のずれを防止
- リサイズ時は不要な横跳びを避け、新しいウィンドウ範囲から外れる場合だけ内側へ補正
- ウィンドウ上端の高さ変更に即座に追従
- 支えているウィンドウを最小化すると不正な座標を追わず自然に落下
- v0.5.8 のマルチモニター対応を維持

""",
    ),
    (
        "### v0.5.8 更新", "### 主要功能",
        """### v0.5.9 更新

- 使用独立的约 60Hz 支撑窗口跟踪计时器，让站立和悬挂跟随更顺滑
- 水平移动窗口时，Dorong 的 X 坐标会按相同距离同步移动
- 行走过程中窗口移动时，行走目标也会同步移动，避免方向和目标错位
- 调整窗口大小时尽量保持 Dorong 当前屏幕 X 坐标，仅在超出新窗口范围时向内修正
- 窗口顶部高度变化时立即同步站立高度
- 支撑窗口最小化后 Dorong 会自然下落，不再跟随最小化窗口的异常坐标
- 保留 v0.5.8 的多显示器工作区域支持

""",
    ),
]
for start, end, replacement in sections:
    pattern = re.compile(re.escape(start) + r".*?(?=" + re.escape(end) + r")", re.S)
    readme, n = pattern.subn(replacement, readme, count=1)
    if n != 1:
        raise SystemExit(f"README section not found: {start}")

readme_path.write_text(readme, encoding="utf-8")
