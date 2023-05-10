package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) NewMenu() *menu.Menu {
	appMenu := menu.NewMenu()
	fileMenu := appMenu.AddSubmenu("ファイル(F)")
	fileMenu.AddText("マップ新規作成", nil, openFile)
	fileMenu.AddText("マップ読み込み", nil, func(_ *menu.CallbackData) {
		filepath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
			Filters: []runtime.FileFilter{
				{
					DisplayName: "Map",
					Pattern: "*.map",
				},
			},
		})
		if err != nil {
			panic(err)
		}
		runtime.LogInfo(a.ctx, filepath)
	})
	fileMenu.AddText("マップ上書き保存", keys.CmdOrCtrl("s"), openFile)
	fileMenu.AddText("名前を付けて保存", keys.OptionOrAlt("a"), openFile)
	fileMenu.AddSeparator()
	fileMenu.AddText("ゲームデータの作成", keys.OptionOrAlt("g"), openFile)
	fileMenu.AddSeparator()
	fileMenu.AddText("エディターの終了", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(a.ctx)
	})
	editMenu := appMenu.AddSubmenu("編集(E)")
	editMenu.AddText("イベント切り取り", keys.Key("X"), openFile)
	editMenu.AddText("イベントコピー", keys.Key("C"), openFile)
	editMenu.AddText("イベント貼り付け", keys.Key("V"), openFile)
	editMenu.AddText("削除", keys.Key("DELETE"), openFile)
	editMenu.AddText("マップの基本設定", nil, openFile)
	layerMenu := appMenu.AddSubmenu("レイヤー")
	layerMenu.AddRadio("レイヤー1", false, nil, openFile)
	layerMenu.AddRadio("レイヤー2", false, nil, openFile)
	layerMenu.AddRadio("レイヤー3", false, nil, openFile)
	layerMenu.AddRadio("イベント", false, nil, openFile)
	gameSettingMenu := appMenu.AddSubmenu("ゲーム設定")
	gameSettingMenu.AddText("ゲーム基本設定を開く", nil, openFile)
	gameSettingMenu.AddText("コンフィグを開く", nil, openFile)
	gameSettingMenu.AddRadio("デバッグウィンドウ使用", false, nil, openFile)
	optionMenu := appMenu.AddSubmenu("オプション")
	optionMenu.AddText("エディターオプション", nil, openFile)
	helpMenu := appMenu.AddSubmenu("ヘルプ(H)")
	helpMenu.AddText("バージョン情報", nil, func(_ *menu.CallbackData) {
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title: "バージョン情報",
		})
	})
	return appMenu
}


func openFile(_ *menu.CallbackData) {
}
