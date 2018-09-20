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

	widget := widgets.NewQWidget(nil, 0)
	formLayout := widgets.NewQFormLayout(nil)

	label1 := widgets.NewQLabel(nil, 0)
	label1.SetText("name")

	input1 := widgets.NewQLineEdit(nil)
	input1.SetPlaceholderText("Please input name's string...")
	formLayout.AddRow(label1, input1)

	label2 := widgets.NewQLabel2("reading", nil, 0)
	input2 := widgets.NewQLineEdit(nil)
	formLayout.AddRow(label2, input2)

	input3 := widgets.NewQLineEdit(nil)
	formLayout.AddRow3("email", input3)

	button := widgets.NewQPushButton2("click", nil)
	formLayout.AddWidget(button)

	widget.SetLayout(formLayout)
	window.SetCentralWidget(widget)

	window.Show()

	app.Exec()
}
