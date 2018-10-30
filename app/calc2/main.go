package main

import (
	"os"
	"strconv"

	"github.com/ryoko-saito/tax_withholding/tax"
	"github.com/therecipe/qt/widgets"
)

//packageスコープ内に変数を定義しておくと、mainの中、外でも変数を使うことができる
var salalyInput *widgets.QLineEdit
var bonusInput *widgets.QLineEdit
var dayInput *widgets.QLineEdit
var trafficInput *widgets.QLineEdit
var traffictotalInput *widgets.QLineEdit
var healthinsuranceInput *widgets.QLineEdit
var paymenttotalInput *widgets.QLineEdit
var employmentinshulanceInput *widgets.QLineEdit
var careinsuranceInput *widgets.QLineEdit
var pensionInput *widgets.QLineEdit
var socialinsulancesumInput *widgets.QLineEdit
var incometaxInput *widgets.QLineEdit
var deductionInput *widgets.QLineEdit
var supportInput *widgets.QLineEdit
var ageInput *widgets.QLineEdit
var kouorotuBox *widgets.QComboBox
var beforeincomeInput *widgets.QLineEdit
var followBox *widgets.QComboBox
var municipaltaxInput *widgets.QLineEdit

func main() {

	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(400, 250)
	window.SetWindowTitle("Salaly calculator")

	gridLayout := widgets.NewQGridLayout2()
	widget := widgets.NewQWidget(nil, 0)

	nameLabel := widgets.NewQLabel2("氏名", nil, 0)
	gridLayout.AddWidget(nameLabel, 0, 0, 0)

	nameInput := widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(nameInput, 1, 0, 0)

	ageLabel := widgets.NewQLabel2("40歳以上64歳未満の扶養人数", nil, 0)
	gridLayout.AddWidget(ageLabel, 0, 2, 0)

	ageInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(ageInput, 1, 2, 0)

	kouorotuLabel := widgets.NewQLabel2("甲または乙の選択", nil, 0)
	gridLayout.AddWidget(kouorotuLabel, 0, 4, 0)

	kouorotuBox = widgets.NewQComboBox(nil)
	kouorotuBox.AddItems([]string{"甲", "乙"})

	gridLayout.AddWidget(kouorotuBox, 1, 4, 0)

	followLabel := widgets.NewQLabel2("従たる給与についての扶養控除申告書の提出", nil, 0)
	gridLayout.AddWidget(followLabel, 0, 5, 0)

	followBox = widgets.NewQComboBox(nil)
	followBox.AddItems([]string{"有り", "無し"})

	gridLayout.AddWidget(followBox, 1, 5, 0)

	supportLabel := widgets.NewQLabel2("扶養人数", nil, 0)
	gridLayout.AddWidget(supportLabel, 2, 4, 0)

	supportInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(supportInput, 3, 4, 0)

	beforeincomeLabel := widgets.NewQLabel2("2年前所得", nil, 0)
	gridLayout.AddWidget(beforeincomeLabel, 2, 5, 0)

	beforeincomeInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(beforeincomeInput, 3, 5, 0)

	salalyLabel := widgets.NewQLabel2("給与", nil, 0)
	gridLayout.AddWidget(salalyLabel, 2, 0, 0)

	salalyInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(salalyInput, 3, 0, 0)

	bonusLabel := widgets.NewQLabel2("賞与", nil, 0)
	gridLayout.AddWidget(bonusLabel, 2, 2, 0)

	bonusInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(bonusInput, 3, 2, 0)

	dayLabel := widgets.NewQLabel2("日数", nil, 0)
	gridLayout.AddWidget(dayLabel, 4, 0, 0)

	trafficfeeLabel := widgets.NewQLabel2("交通費", nil, 0)
	gridLayout.AddWidget(trafficfeeLabel, 4, 2, 0)

	traffictotleLabel := widgets.NewQLabel2("交通費合計", nil, 0)
	gridLayout.AddWidget(traffictotleLabel, 4, 4, 0)

	dayInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(dayInput, 5, 0, 0)

	kakeruLabel := widgets.NewQLabel2("×", nil, 0)
	gridLayout.AddWidget(kakeruLabel, 5, 1, 0)

	trafficInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(trafficInput, 5, 2, 0)

	ikoruLabel := widgets.NewQLabel2("=", nil, 0)
	gridLayout.AddWidget(ikoruLabel, 5, 3, 0)

	traffictotalInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(traffictotalInput, 5, 4, 0)

	paymenttotalLabel := widgets.NewQLabel2("支給合計", nil, 0)
	gridLayout.AddWidget(paymenttotalLabel, 4, 5, 0)

	paymenttotalInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(paymenttotalInput, 5, 5, 0)

	healthinsuranceLabel := widgets.NewQLabel2("健康保険", nil, 0)
	gridLayout.AddWidget(healthinsuranceLabel, 6, 0, 0)

	healthinsuranceInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(healthinsuranceInput, 7, 0, 0)

	careinsuranceLabel := widgets.NewQLabel2("介護保険", nil, 0)
	gridLayout.AddWidget(careinsuranceLabel, 6, 2, 0)

	careinsuranceInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(careinsuranceInput, 7, 2, 0)

	pensionLabel := widgets.NewQLabel2("厚生年金", nil, 0)
	gridLayout.AddWidget(pensionLabel, 6, 4, 0)

	pensionInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(pensionInput, 7, 4, 0)

	employmentinshulanceLabel := widgets.NewQLabel2("雇用保険", nil, 0)
	gridLayout.AddWidget(employmentinshulanceLabel, 8, 0, 0)

	employmentinshulanceInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(employmentinshulanceInput, 9, 0, 0)

	incometaxLabel := widgets.NewQLabel2("所得税", nil, 0)
	gridLayout.AddWidget(incometaxLabel, 10, 0, 0)

	incometaxInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(incometaxInput, 11, 0, 0)

	municipaltaxLabel := widgets.NewQLabel2("住民税", nil, 0)
	gridLayout.AddWidget(municipaltaxLabel, 10, 2, 0)

	municipaltaxInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(municipaltaxInput, 11, 2, 0)

	socialinsulancesumLabel := widgets.NewQLabel2("社会保険合計", nil, 0)
	gridLayout.AddWidget(socialinsulancesumLabel, 8, 5, 0)

	socialinsulancesumInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(socialinsulancesumInput, 9, 5, 0)

	deductiontotalLabel := widgets.NewQLabel2("控除合計", nil, 0)
	gridLayout.AddWidget(deductiontotalLabel, 10, 5, 0)

	deductionInput = widgets.NewQLineEdit(nil)
	gridLayout.AddWidget(deductionInput, 11, 5, 0)

	beforeincomeInput.ConnectEditingFinished(calcAndInsert)
	supportInput.ConnectEditingFinished(calcAndInsert)
	ageInput.ConnectEditingFinished(calcAndInsert)
	dayInput.ConnectEditingFinished(calcAndInsert)
	trafficInput.ConnectEditingFinished(calcAndInsert)
	salalyInput.ConnectEditingFinished(calcAndInsert)
	bonusInput.ConnectEditingFinished(calcAndInsert)
	traffictotalInput.ConnectEditingFinished(calcAndInsert)
	paymenttotalInput.ConnectEditingFinished(calcAndInsert)
	careinsuranceInput.ConnectEditingFinished(calcAndInsert)
	socialinsulancesumInput.ConnectEditingFinished(calcAndInsert)
	healthinsuranceInput.ConnectEditingFinished(calcAndInsert)
	pensionInput.ConnectEditingFinished(calcAndInsert)
	deductionInput.ConnectEditingFinished(calcAndInsert)

	//box系のイベント
	followBox.ConnectCurrentIndexChanged2(func(text string) {
		calcAndInsert()
	})

	kouorotuBox.ConnectCurrentIndexChanged2(func(text string) {
		calcAndInsert()
	})

	widget.SetLayout(gridLayout)
	window.SetCentralWidget(widget)

	window.Show()
	app.Exec()
}

//各Inputの計算
func calcAndInsert() {
	dayV := convert(dayInput)
	trafficV := convert(trafficInput)
	res := dayV * trafficV
	insert(traffictotalInput, res)

	bonusAmount := convert(bonusInput)

	//1行目と2行目の合算
	salalyAmount := convert(salalyInput)
	trafficTotalAmount := convert(traffictotalInput)
	res = salalyAmount + trafficTotalAmount + bonusAmount
	insert(paymenttotalInput, res)

	paymentTotal := convert(paymenttotalInput)
	res = paymentTotal * 3 / 1000
	insert(employmentinshulanceInput, res)

	//健康保険の金額
	supportnumber := convert(supportInput)
	beforeincomeAmout := convert(beforeincomeInput)
	familynumber := supportnumber + 1
	healthinsuranceAmout := tax.CalcHealthInsurance(beforeincomeAmout, familynumber)
	insert(healthinsuranceInput, healthinsuranceAmout)

	//介護保険の計算
	carenumber := convert(ageInput)
	careinsuranceAmount := tax.CalcCareInsurance(carenumber)
	insert(careinsuranceInput, careinsuranceAmount)

	//厚生年金の計算
	pensionAmout := convert(pensionInput)
	employmentinshulanceAmout := convert(employmentinshulanceInput)

	//社会保険の合計
	socialinsulancesumAmount := healthinsuranceAmout + careinsuranceAmount + pensionAmout + employmentinshulanceAmout
	insert(socialinsulancesumInput, socialinsulancesumAmount)

	//所得（給与ー社会保険合計）
	incomeAmout := salalyAmount - socialinsulancesumAmount

	kouorotuIndex := kouorotuBox.CurrentIndex()

	//従たる給与についての扶養控除申告書の提出
	followIndex := followBox.CurrentIndex()

	//所得税の計算
	incometaxAmount := tax.CalcTax(incomeAmout, kouorotuIndex, supportnumber, followIndex)
	insert(incometaxInput, incometaxAmount)

}

func convert(e *widgets.QLineEdit) int {
	str := e.Text()
	if len(str) == 0 {
		return 0
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return i
}

func insert(e *widgets.QLineEdit, i int) {
	h := strconv.Itoa(i)
	e.Clear()
	e.Insert(h)
}
