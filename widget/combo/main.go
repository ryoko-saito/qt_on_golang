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

	//フレームワークにQTBoxLayoutをはめ込む
	//第一引数で0は左から右、1は右から左、2は上から下、3は下から上
	widget.SetLayout(widgets.NewQBoxLayout(3, nil))
	window.SetCentralWidget(widget)

	comboBox := widgets.NewQComboBox(nil)
	comboBox.SetWindowTitle("Combo Box")

	items := []string{"Some", "Combo", "Box", "Items"}
	comboBox.AddItems(items)

	comboBox.ConnectCurrentIndexChanged2(func(text string) {
		widgets.QMessageBox_Information(nil, "OK", text, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})

	widget.Layout().AddWidget(comboBox)

	window.Show()

	app.Exec()
}
