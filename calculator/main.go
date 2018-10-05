package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(800, 400)
	window.SetWindowTitle("Salaly calculator")

	group1 := widgets.NewQGroupBox2("項目", nil)
	layout := widgets.NewQBoxLayout(1, nil)

	label := widgets.NewQLabel2("時間", nil, 0)
	layout.AddWidget(label, 0, 0)

	label2 := widgets.NewQLabel2("時給", nil, 0)
	layout.AddWidget(label2, 0, 0)

	label3 := widgets.NewQLabel2("給料", nil, 0)
	layout.AddWidget(label3, 0, 0)

	group1.SetLayout(layout)

	mainLayout := widgets.NewQGridLayout2()
	mainLayout.AddWidget(group1, 0, 0, 0)

	group2 := widgets.NewQGroupBox2("数値", nil)
	layout = widgets.NewQBoxLayout(1, nil)

	input := widgets.NewQLineEdit(nil)
	layout.AddWidget(input, 0, 0)

	label = widgets.NewQLabel2("=", nil, 0)
	layout.AddWidget(label, 0, 0)

	input2 := widgets.NewQLineEdit(nil)
	layout.AddWidget(input2, 0, 0)

	label2 = widgets.NewQLabel2("×", nil, 0)
	layout.AddWidget(label2, 0, 0)

	input3 := widgets.NewQLineEdit(nil)
	layout.AddWidget(input3, 0, 0)
	group2.SetLayout(layout)

	mainLayout.AddWidget(group2, 1, 0, 0)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(mainLayout)
	window.SetCentralWidget(widget)

	window.Show()

	app.Exec()

}
