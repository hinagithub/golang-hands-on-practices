package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func input(msg string) string {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println(msg)
	sc.Scan()
	return sc.Text()
}

func pointer1() {
	n := 123

	// 自動推論でポインタ型になる
	p := &n

	fmt.Println("number", n)
	fmt.Println("pointer", p)
	fmt.Println("value", *p)

	// number 123
	// pointer 0xc0000ac008
	// value 123

	// ポインタは演算できない
}

func pointer2() {
	n := 123
	p := &n
	q := &p
	m := 10000
	p2 := &m
	q2 := &p2
	fmt.Printf("q value:%d, address:%p \n", **q, *q)
	fmt.Printf("q2 value:%d, address:%p \n", **q2, *q2)
	// q value:10000, address:0xc0000ac010
	// q2 value:10000, address:0xc0000ac008

	pb := p
	p = p2
	p2 = pb
	fmt.Printf("q value:%d, address:%p \n", **q, *q)
	fmt.Printf("q2 value:%d, address:%p \n", **q2, *q)
	// q value:10000, address:0xc0000ac010
	// q2 value:123, address:0xc0000ac010

}

func pointer3() {
	change1 := func(n int) { n *= 2 }
	change2 := func(n *int) { *n *= 2 }

	n := 123
	fmt.Printf("value:%d.\n", n)
	change1(n)
	fmt.Printf("value:%d.\n", n)
	change2(&n)
	fmt.Printf("value:%d.\n", n)
	// value:123.
	// value:123.
	// value:246.
}

// 構造体
func struct1() {

	// 単純な構造体定義
	var mydata struct {
		Name string
		Data []int
	}
	mydata.Name = "Taro"
	mydata.Data = []int{10, 20, 30}
	fmt.Println(mydata)
	// {Taro [10 20 30]}
}

func struct2() {
	// typeを使った構造体
	// Goは外部から利用可能な方や関数はその内容を示すコメントを用意する必要がある
	// Mydata is structure.
	type Mydata struct {
		Name string
		Data []int
	}
	taro := Mydata{"Taro", []int{10, 20, 30}}
	hanako := Mydata{"Hanako", []int{90, 80, 70}}
	fmt.Println(taro)
	fmt.Println(hanako)
	// 	{Taro [10 20 30]}
	// {Hanako [90 80 70]}

}

// 値渡しの構造体
func struct3() {

	// Mydata is structure.
	type Mydata struct {
		Name string
		Data []int
	}

	rev := func(md Mydata) Mydata {
		od := md.Data
		nd := []int{}
		for i := len(od) - 1; i >= 0; i-- {
			nd = append(nd, od[i])
		}
		md.Data = nd
		return md
	}

	taro := Mydata{"Taro", []int{10, 20, 30}}
	fmt.Println(taro)
	// 	{Taro [10 20 30]}

	taro = rev(taro)
	fmt.Println(taro)
	// {Taro [30 20 10]}

}

// 参照渡しの構造体
func struct4() {

	// Mydata is structure.
	type Mydata struct {
		Name string
		Data []int
	}

	rev := func(md *Mydata) {
		od := (*md).Data
		nd := []int{}
		for i := len(od) - 1; i >= 0; i-- {
			nd = append(nd, od[i])
		}
		md.Data = nd
	}

	taro := Mydata{"Taro", []int{10, 20, 30}}
	fmt.Println(taro)
	// 	{Taro [10 20 30]}

	rev(&taro)
	fmt.Println(taro)
	// {Taro [30 20 10]}

}

// newとmakeで初期化する
func struct5() {
	// Mydata is structure.
	type Mydata struct {
		Name string
		Data []int
	}
	taro := new(Mydata)
	fmt.Println(taro)
	// 	&{ []}

	taro.Name = "Taro"
	taro.Data = make([]int, 5, 5)
	fmt.Println(taro)
	// &{Taro [0 0 0 0 0]}

}

// ---------------------------------------------------------------------------

// メソッド(型に組み込まれる関数)を定義する
// Mydata is structure.
type Mydata struct {
	Name string
	Data []int
}

// PrintData is println all data.
// Mydataがれしーばなのでこの関数は「メソッド」である
func (md Mydata) PrintData() {
	fmt.Println("*** Mydata ***")
	fmt.Println("Data: ", md.Data)
	fmt.Println("*** end ***")
}

// メソッド
func struct6() {
	// 関数とメソッドの違いはレシーバの指定を用意しているか否かだけと考えてOK
	hanako := Mydata{
		"Hanako",
		[]int{98, 76, 54, 32, 10},
	}
	hanako.PrintData()
	//  *** Mydata ***
	// Name:  Hanako
	// Data:  [98 76 54 32 10]
	// *** end ***
}

// ---------------------------------------------------------------------------

// intの拡張intpを定義
type intp int

// 素数チェック
func (num intp) IsPrime() bool {
	n := int(num)
	for i := 2; i < (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 素因数分解
func (num intp) PrimeFactor() []int {
	ar := []int{}
	x := int(num)
	n := 2
	for x > n {
		if x%n == 0 {
			x /= n
			ar = append(ar, n)

		} else {
			if n == 2 {
				n++

			} else {
				n += 2
			}
		}
	}
	ar = append(ar, x)
	return ar
}

func (num *intp) doPrime() {
	pf := num.PrimeFactor()
	*num = intp((pf[len(pf)-1]))
}

// 型の拡張
func type1() {
	s := input("type a number")
	n, _ := strconv.Atoi(s)
	x := intp(n)
	// 	type a number
	// 77

	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	// 77 [false].
	// [7 11]

	x.doPrime()
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	// 11 [true].
	// [11]

	x++
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	// 12 [false].
	// [2 2 3]
}

// ---------------------------------------------------------------------------

// Data is interface
type Data interface {
	Initial(name string, data []int)
	PrintData()
}

// Initial is int method.
func (md *Mydata) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

// PrintData is println all data.
func (md *Mydata) PrintlnData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

// check is method.
func (md *Mydata) Check() {
	fmt.Printf("Check! [%s]", md.Name)
}

// Mydataとして扱う
func interface1() {
	// interfaceでメソッド定義を義務化する
	var ob Mydata = Mydata{}
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.PrintlnData()
	// 	Name:  Sachiko
	// Data:  [55 66 77]
}

// Dataとして扱う
func interface2() {

	// newを使って型を代入している
	var ob Data = new(Mydata)
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.PrintData()
	// *** Mydata ***
	// Data:  [55 66 77]
	// *** end ***
}

// Checkする
func interface3() {
	var ob Mydata = Mydata{}
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.Check()
	// Check! [Sachiko]%

	var ob2 Data = new(Mydata)
	ob2.Initial("Sachiko", []int{55, 66, 77})
	// ob2.Check()
	// ob2.Check undefined (type Data has no field or method Check)
}

// ---------------------------------------------------------------------------

func main() {
	pointer1()
	pointer2()
	pointer3()
	struct1()
	struct2()
	struct3() // 値渡し
	struct4() // 参照渡し
	struct5() // newとmake
	struct6() // メソッド
	type1()
	interface1() // Mydataとして扱う
	interface2() // Mydataとして扱う
	interface3() // Check

}
