package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(400, 200)
	window.SetWindowTitle("アルバイト給与")

	mainLayout := widgets.NewQGridLayout2()
	widget := widgets.NewQWidget(nil, 0)

	group1 := widgets.NewQGroupBox2("1", nil)
	boxLayout := widgets.NewQBoxLayout(3, nil)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Please input number")
	boxLayout.AddWidget(input, 0, 0)

	label := widgets.NewQLabel2("時間", nil, 0)
	boxLayout.AddWidget(label, 0, 0)

	group1.SetLayout(boxLayout)

	group2 := widgets.NewQGroupBox2("2", nil)
	boxLayout = widgets.NewQBoxLayout(3, nil)

	input1 := widgets.NewQLineEdit(nil)
	input1.SetPlaceholderText("Please input number")
	boxLayout.AddWidget(input1, 0, 0)

	label1 := widgets.NewQLabel2("時給", nil, 0)
	boxLayout.AddWidget(label1, 0, 0)

	group2.SetLayout(boxLayout)

	group3 := widgets.NewQGroupBox2("3", nil)
	boxLayout = widgets.NewQBoxLayout(3, nil)

	input2 := widgets.NewQLineEdit(nil)
	boxLayout.AddWidget(input2, 0, 0)

	label2 := widgets.NewQLabel2("給料1", nil, 0)
	boxLayout.AddWidget(label2, 0, 0)

	group3.SetLayout(boxLayout)

	group4 := widgets.NewQGroupBox2("4", nil)
	boxLayout = widgets.NewQBoxLayout(3, nil)

	input3 := widgets.NewQLineEdit(nil)
	input3.SetPlaceholderText("Please input number")
	boxLayout.AddWidget(input3, 0, 0)

	label5 := widgets.NewQLabel2("出勤日数", nil, 0)
	boxLayout.AddWidget(label5, 0, 0)

	group4.SetLayout(boxLayout)

	group5 := widgets.NewQGroupBox2("5", nil)
	boxLayout = widgets.NewQBoxLayout(3, nil)

	input4 := widgets.NewQLineEdit(nil)
	input4.SetPlaceholderText("Please input number")
	boxLayout.AddWidget(input4, 0, 0)

	label6 := widgets.NewQLabel2("交通費", nil, 0)
	boxLayout.AddWidget(label6, 0, 0)

	group5.SetLayout(boxLayout)

	group6 := widgets.NewQGroupBox2("6", nil)
	boxLayout = widgets.NewQBoxLayout(3, nil)

	input5 := widgets.NewQLineEdit(nil)
	boxLayout.AddWidget(input5, 0, 0)

	label7 := widgets.NewQLabel2("交通費合計", nil, 0)
	boxLayout.AddWidget(label7, 0, 0)

	group6.SetLayout(boxLayout)

	group7 := widgets.NewQGroupBox2("7", nil)
	boxLayout = widgets.NewQBoxLayout(3, nil)

	input6 := widgets.NewQLineEdit(nil)
	boxLayout.AddWidget(input6, 0, 0)

	label8 := widgets.NewQLabel2("給与合計", nil, 0)
	boxLayout.AddWidget(label8, 0, 0)

	group7.SetLayout(boxLayout)

	label3 := widgets.NewQLabel2("×", nil, 0)
	label4 := widgets.NewQLabel2("=", nil, 0)
	label9 := widgets.NewQLabel2("×", nil, 0)
	label10 := widgets.NewQLabel2("=", nil, 0)

	mainLayout.AddWidget(group1, 0, 0, 0)
	mainLayout.AddWidget(group2, 0, 2, 0)
	mainLayout.AddWidget(group3, 0, 4, 0)
	mainLayout.AddWidget(group4, 1, 0, 0)
	mainLayout.AddWidget(group5, 1, 2, 0)
	mainLayout.AddWidget(group6, 1, 4, 0)
	mainLayout.AddWidget(group7, 2, 4, 0)
	mainLayout.AddWidget(label3, 0, 1, 0)
	mainLayout.AddWidget(label4, 0, 3, 0)
	mainLayout.AddWidget(label9, 1, 1, 0)
	mainLayout.AddWidget(label10, 1, 3, 0)

	widget.SetLayout(mainLayout)
	window.SetCentralWidget(widget)

	window.Show()
	app.Exec()
}
