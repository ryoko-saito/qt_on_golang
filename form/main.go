package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(400, 300)
	window.SetWindowTitle("Hello Ryoko")

	//ウィジットを作成し、QFormLayoutを作成する
	widget := widgets.NewQWidget(nil, 0)
	formLayout := widgets.NewQFormLayout(nil)

	// QLabelでラベルを作成してみる。QLabelを作成した後、文字列をセットする必要がある。
	label1 := widgets.NewQLabel(nil, 0)
	label1.SetText("name")

	//プレースホルダー付きQLineEditを作成してみる
	input1 := widgets.NewQLineEdit(nil)
	input1.SetPlaceholderText("Please input name's string...")
	formLayout.AddRow(label1, input1)

	//QLabel2でラベルを作成してみる。作成時に文字列を指定することができる
	label2 := widgets.NewQLabel2("reading", nil, 0)
	input2 := widgets.NewQLineEdit(nil)
	formLayout.AddRow(label2, input2)

	// AddRow3でQLineEditを挿入してみる。QLabelを使わなくても、AddRow3で上記と同じように、QLabelとQLineEditを作成できる
	input3 := widgets.NewQLineEdit(nil)
	formLayout.AddRow3("email", input3)

	//AddRow以外でウィジェットを挿入できるか試してみる
	button := widgets.NewQPushButton2("click", nil)
	formLayout.AddWidget(button)

	//フレームワークに上記で作成したレイアウトをセットする
	widget.SetLayout(formLayout)
	window.SetCentralWidget(widget)

	window.Show()

	app.Exec()
}
