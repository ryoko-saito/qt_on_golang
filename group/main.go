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

	//フォームレイアウトをグループ化
	group1 := widgets.NewQGroupBox(nil)
	group1.SetTitle("form layout")

	formLayout := widgets.NewQFormLayout(nil)
	input1 := widgets.NewQLineEdit(nil)
	formLayout.AddRow3("string", input1)

	//作成したフォームレイアウトをグループにセットしてみる
	group1.SetLayout(formLayout)

	//グリッドレイアウトをグループ化 NewQGroupBoxでtitleもまとめて作成
	group2 := widgets.NewQGroupBox2("grid layout", nil)

	//グリッドレイアウトを入れてみる
	gridLayout := widgets.NewQGridLayout2()
	label := widgets.NewQLabel2("value", nil, 0)
	gridLayout.AddWidget(label, 0, 0, 0)

	input2 := widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(input2, 0, 1, 0)

	//作成したグリットレイアウトをグループにセット
	group2.SetLayout(gridLayout)

	//QGridLayoutに各グループを追加してみる
	layout := widgets.NewQGridLayout2()
	layout.AddWidget(group1, 0, 0, 0)
	layout.AddWidget(group2, 1, 0, 0)

	//空のウィジェットを作成して、上で作成したフレームワークに挿入する
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	window.SetCentralWidget(widget)

	//作成したwindowを表す
	window.Show()

	//Qtアプリケーションを実行する
	app.Exec()
}
