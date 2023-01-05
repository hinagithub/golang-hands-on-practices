package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 異なる構造体をインターフェースでまとめる P153〜

// Data is interface for Mydata.
type Data interface {
	SetValue(vals map[string]string)
	PrintData()
}

// Mydata is structure.
type Mydata struct {
	Name string
	Data []int
}

// PrintData is Mydata method.
func (md *Mydata) PrintData() {
	if md != nil {
		fmt.Println("Name: ", md.Name)
		fmt.Println("Data: ", md.Data)
	} else {
		fmt.Println("** This is Nil value. **")
	}
}

// SetValue is Mydata method.
func (md *Mydata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	valt := strings.Split(vals["data"], " ")
	vali := []int{}
	for _, i := range valt {
		n, _ := strconv.Atoi(i)
		vali = append(vali, n)
	}
	md.Data = vali
}

// Yourdata is structure
type Yourdata struct {
	Name string
	Mail string
	Age  int
}

// PrintData is Yourdata method.
func (md *Yourdata) PrintData() {
	if md != nil {
		fmt.Printf("I'am %s. (%d).\n", md.Name, md.Age)
		fmt.Printf("Mail: %s.", md.Mail)
	} else {
		fmt.Println("** This is Nil value. **")
	}
}

// SetValue is Yourdata method.
func (md *Yourdata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	md.Mail = vals["mail"]
	n, _ := strconv.Atoi(vals["age"])
	md.Age = n
}

func main() {

	// 同じインターフェース(Data)実装するtypeであれば配列で同じData型として扱うことができる
	// ただし引数と戻り値の型とメソッド名が完全に一致していなければならない

	// 配列にMydata型の値を代入
	ob := [2]Data{}
	ob[0] = new(Mydata)
	ob[0].SetValue(map[string]string{
		"name": "Sachiko",
		"data": "55, 66, 77",
	})

	// 配列にYourdata型の値を代入
	ob[1] = new(Yourdata)
	ob[1].SetValue(map[string]string{
		"name": "Sachiko",
		"mail": "mami@mume.mo",
		"age":  "34",
	})
	for _, d := range ob {
		d.PrintData()
		fmt.Println()
	}

	// Name:  Sachiko
	// Data:  [0 0 77]

	// I'am Sachiko. (34).
	// Mail: mami@mume.mo.

	// 以下nilレシーバについて
	// メソッドの初期が失敗したり、変数宣言だけして値をまだ入れていない状態だとそこには「nil」が入る。
	// nilの状態でもメソッド自体は実行できる（エラー終了するわけではない）ので注意！

	var ob2 *Mydata
	ob2.PrintData()
	ob2 = &Mydata{}
	ob2.SetValue(map[string]string{
		"name": "Jiro",
		"data": "123 456 789",
	})
	ob2.PrintData()
	// 	** This is Nil value. **
	// Name:  Jiro
	// Data:  [123 456 789]

}
