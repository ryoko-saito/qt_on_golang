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

	gridLayout := widgets.NewQGridLayout2()

	//AddWidgetは第一引数にボタンのウィジェット、第二引数に列(row)、
	//第三引数に行(col)、第四引数に整列規則(Alignment)を入れる
	button1 := widgets.NewQPushButton2("one", nil)
	gridLayout.AddWidget(button1, 0, 0, 0)

	button2 := widgets.NewQPushButton2("two", nil)
	gridLayout.AddWidget(button2, 0, 1, 0)

	//AddWidget3は第三引数まではAddWdidgetと同じ、第四引数にどれくらい列をまたぎたいか？
	//第五引数に行をまたぎたいか？第六引数に整列規則を入れる
	button3 := widgets.NewQPushButton2("three", nil)
	gridLayout.AddWidget3(button3, 1, 0, 2, 0, 0)

	button4 := widgets.NewQPushButton2("four", nil)
	gridLayout.AddWidget(button4, 2, 0, 0)

	button5 := widgets.NewQPushButton2("five", nil)
	gridLayout.AddWidget(button5, 2, 1, 0)

	//widget
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(gridLayout)
	window.SetCentralWidget(widget)

	window.Show()

	app.Exec()
}
