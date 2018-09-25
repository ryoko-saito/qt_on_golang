package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	//go言語でqtを呼び出す
	app := widgets.NewQApplication(len(os.Args), os.Args)

	//central画面を作成する
	window := widgets.NewQMainWindow(nil, 0)

	//フレームワークの大きさを決める 第一引数が幅、第二引数が高さの指定
	window.SetMinimumSize2(400, 300)

	//フレームワークにタイトルをつける
	window.SetWindowTitle("Hello Ryoko")

	//空のウィジェットを作成して、上で作成したフレームワークに挿入する
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	//作成したwindowを表す
	window.Show()

	//Qtアプリケーションを実行する
	app.Exec()
}
