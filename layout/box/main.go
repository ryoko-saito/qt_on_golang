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

	label := widgets.NewQLabel2("ryoko", nil, 0)
	widget.Layout().AddWidget(label)

	label2 := widgets.NewQLabel2("usa", nil, 0)
	widget.Layout().AddWidget(label2)

	label3 := widgets.NewQLabel2("kuma", nil, 0)
	widget.Layout().AddWidget(label3)

	window.Show()

	app.Exec()
}
