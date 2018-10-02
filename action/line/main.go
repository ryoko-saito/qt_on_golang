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

	layout := widgets.NewQBoxLayout(0, nil)

	//左側
	input1 := widgets.NewQLineEdit(nil)

	//右側
	input2 := widgets.NewQLineEdit(nil)

	//event 左のQLineEditに入力した値を、右のQLineEditに反映する
	input1.ConnectEditingFinished(func() {
		str := input1.Text()
		input2.Insert(str)
	})

	layout.AddWidget(input1, 0, 0)
	layout.AddWidget(input2, 1, 0)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)
	window.SetCentralWidget(widget)

	window.Show()

	app.Exec()
}
