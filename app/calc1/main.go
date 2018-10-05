package main

import (
	"os"
	"strconv"

	"github.com/therecipe/qt/widgets"
)

//packageスコープ内に変数を定義しておくと、mainの中、外でも変数を使うことができる
var jikyuuInput *widgets.QLineEdit
var jikanInput *widgets.QLineEdit
var goukeiInput *widgets.QLineEdit

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(400, 200)
	window.SetWindowTitle("Salaly calculator")

	gridLayout := widgets.NewQGridLayout2()
	widget := widgets.NewQWidget(nil, 0)

	goukeiLabel := widgets.NewQLabel2("合計", nil, 0)
	gridLayout.AddWidget(goukeiLabel, 0, 4, 0)

	goukeiInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(goukeiInput, 1, 4, 0)

	ikoruLabel := widgets.NewQLabel2("=", nil, 0)
	gridLayout.AddWidget(ikoruLabel, 1, 3, 0)

	jikanLabel := widgets.NewQLabel2("時間", nil, 0)
	gridLayout.AddWidget(jikanLabel, 0, 2, 0)
	jikanInput = widgets.NewQLineEdit(nil)

	jikanInput.ConnectEditingFinished(calcAndInsert) //jikanからフォーカスをはずした後、編集を終了したらcalcAndInsert関数を実行する
	gridLayout.AddWidget(jikanInput, 1, 2, 0)

	kakeruLabel := widgets.NewQLabel2("×", nil, 0)
	gridLayout.AddWidget(kakeruLabel, 1, 1, 0)

	jikyuuLabel := widgets.NewQLabel2("時給", nil, 0)
	gridLayout.AddWidget(jikyuuLabel, 0, 0, 0)

	jikyuuInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(jikyuuInput, 1, 0, 0)

	jikanInput.ConnectEditingFinished(calcAndInsert) //jikyuuからフォーカスをはずした後、編集を終了したらcalcAndInsert関数を実行する
	jikyuuInput.ConnectEditingFinished(calcAndInsert)

	widget.SetLayout(gridLayout)
	window.SetCentralWidget(widget)

	window.Show()
	app.Exec()
}

//jikyuuとjikanのQLineEditから、入力された文字を取り出し、数字に変えて、掛け算した結果をgoukeiに入れる関数
func calcAndInsert() {
	str := jikyuuInput.Text() //jikyuuのQlineEditから文字を取り出す
	i, _ := strconv.Atoi(str) //文字を数字に変換する

	str1 := jikanInput.Text()  //jikanのQlineEditから文字を取り出す
	j, _ := strconv.Atoi(str1) //文字を数字に変換する
	res := i * j
	h := strconv.Itoa(res) //数字を文字に変換する
	goukeiInput.Clear()    //goukeiのQlineEditに入力されている文字を消す
	goukeiInput.Insert(h)  //goukeiに掛け算結果のhを入力する
}
