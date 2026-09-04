# Dorong

[한국어](README.md) | [영어](README.en.md) | **일본어** | [중국어](README.zh-CN.md)

**Dorong（ドロン）** は、Go と Win32 API で作られた個人用の Windows デスクトップペットです。
Python や .NET などの追加ランタイムをインストールしなくても、`Dorong.exe` 単体で実行できます。

> 現在のバージョン: **v0.5.5**

## 主な機能

- 透明背景のアニメーションデスクトップペット
- マウスで Dorong をつかんで好きな位置へ移動
- ドラッグ後に離すと重力が適用されて落下
- Windows の作業領域を認識し、タスクバーの上に着地
- 通常のアプリケーションウィンドウの上端に正確に着地
- 狭いウィンドウの端も左足・中央・右足の判定で認識
- アプリケーションウィンドウの上を左右に歩き、ウィンドウが移動・リサイズされても移動範囲を再計算
- ときどきウィンドウの端まで歩いて実際の端にぶら下がる
- ぶら下がっている間にウィンドウが移動・リサイズされると追従し、支えがなくなると自然に落下
- 頭の部分をダブルクリックすると撫でる反応
- 待機中にマウスカーソルを見る反応
- 待機 / 歩行 / 睡眠 / 喜び / 持ち上げ / 集中 / 落下 / ぶら下がりアニメーション
- ランダム吹き出し
- 25分集中タイマー
- 10分アラーム
- 常に手前に表示するオプション
- **アプリUI言語: 韓国語 / English**
- 選択した言語を保存し、次回起動時にも維持
- 64ビット Windows 用の単一実行ファイル
- ネットワーク接続やアカウント不要

## 操作方法

| 操作 | 機能 |
| --- | --- |
| 左ボタンドラッグ | Dorong をつかんで移動 |
| ドラッグ後に離す | Dorong を落とし、重力を適用 |
| 頭をダブルクリック | Dorong を撫でる |
| 右クリック | メニューを開く |

## 言語

この README は、ページ上部のリンクから **韓国語 / 英語 / 日本語 / 中国語** に切り替えられます。

Dorong のアプリ内 UI は現在 **韓国語と英語** に対応しています。Dorong を右クリックし、メニュー下部の `한국어` または `English` を選択してください。
メニュー、吹き出し、集中タイマー、アラーム、システムメッセージはすぐに選択した言語へ切り替わり、設定は Windows のユーザー設定フォルダー内の `Dorong/settings.json` に保存されます。

## 実行

GitHub Actions で生成された Windows ビルド、またはリリース/ローカルビルドの `Dorong.exe` を実行してください。

ビルド済みの実行ファイルを使う場合、**Python、Go、.NET のインストールは不要です。**

## ソースからビルド

Go 1.23 以上を推奨します。

```powershell
go build -ldflags="-H=windowsgui -s -w" -o Dorong.exe .
```

Linux/macOS から Windows 向けにクロスコンパイル:

```bash
GOOS=windows GOARCH=amd64 go build -ldflags='-H=windowsgui -s -w' -o Dorong.exe .
```

## プロジェクト構成

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

## 注意

このリポジトリは現在 **個人用の非公開プロジェクト** です。将来公開リポジトリへ変更する場合は、キャラクター画像の再配布権限を確認してください。

## ロードマップ

[`docs/ROADMAP.md`](docs/ROADMAP.md) を参照してください。
