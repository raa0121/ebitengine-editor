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
	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("ファイル(F)")
	FileMenu.AddText("マップ新規作成", nil, openFile)
	FileMenu.AddText("マップ読み込み", nil, openFile)
	FileMenu.AddText("マップ上書き保存", keys.CmdOrCtrl("s"), openFile)
	FileMenu.AddText("名前を付けて保存", keys.OptionOrAlt("a"), openFile)
	FileMenu.AddSeparator()
	FileMenu.AddText("ゲームデータの作成", keys.OptionOrAlt("g"), openFile)
	FileMenu.AddSeparator()
	FileMenu.AddText("エディターの終了", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(a.ctx)
	})
	EditMenu := AppMenu.AddSubmenu("編集(E)")
	EditMenu.AddText("イベント切り取り", keys.Key("X"), openFile)
	EditMenu.AddText("イベントコピー", keys.Key("C"), openFile)
	EditMenu.AddText("イベント貼り付け", keys.Key("V"), openFile)
	EditMenu.AddText("削除", keys.Key("DELETE"), openFile)
	EditMenu.AddText("マップの基本設定", nil, openFile)
	return AppMenu
}

func openFile(_ *menu.CallbackData) {
}
