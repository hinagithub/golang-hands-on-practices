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

func practice1() {

	// 変数宣言
	// 初期値なし
	var n int
	fmt.Println(n)
	// 0

	// 複数一度に宣言することも可能
	var o, p, q int
	fmt.Println(o, p, q)
	// 0 0 0

	// 初期値を設定
	var v1 string = "a"
	fmt.Println(v1)
	// a

	// 型推論を使って省略した書き方
	v2 := "b"
	fmt.Println(v2)
	// b2

	// 計算して出力
	a, b, c := 100, 200, 300
	// 演算子は基本的にスペースを開けない
	// 演算順番を明示する時は開けたりするらしい
	fmt.Println("total: ", a+b+c)
	// total:  600
}

func practice2() {

	// キャスト
	// int32 -> int64など暗黙可能なものも必ず明示的にキャストするのがGoの書き方

	// int32 -> int64
	var x int32 = 100
	var y int64 = int64(x)

	// int64 -> float32
	var z float32 = float32(y)

	fmt.Println(x, y, z)
	// 100 100 100
}

func practice3() {
	// string -> int
	x := input("type a price")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Error! input: ", x)
		return
	}
	p := float64(n)
	fmt.Println(int(p * 1.1))

	// "100"を入力した場合
	// 110

	// "a" を入力した場合
	// Error! input:  a
}

func practice_const() {

	// 変数で型推論を利用するのは単純に値から型が予測できるから
	// 定数で型推論を利用するのはあえて型を固定しないため

	// 例えばintで宣言したら
	const n int = 100
	// 定数を使って計算しようと思ったときにキャストが必要になる
	m := float64(n) * 1.1
	fmt.Println(m)

	// 型を指定せずに宣言しておけば
	const n2 = 100
	// キャストが不要になる
	m2 := n2 * 1.1
	fmt.Println(m2)
}

func practice_if() {
	x := input("type a number")
	n, err := strconv.Atoi(x)

	// 一般的なif
	if err != nil {
		fmt.Println("Error!")
		return
	}
	fmt.Print(x + "は")
	if n%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// ショートステートメント付きif if 文; 条件{...}
	if n, err := strconv.Atoi(x); err == nil {
		fmt.Print(x + "は")
		if n%2 == 0 {
			fmt.Println("偶数")
		} else {
			fmt.Println("奇数")
		}
	}
}

func practice_switch1() {

	// Goのswitchはヒットしたら自動でbreakされる
	x := input("type a month")
	fmt.Print(x + "は")
	switch n, err := strconv.Atoi(x); n {
	case 0:
		fmt.Println("整数値が得られません")
		fmt.Println(err)
	case 1, 2, 12:
		fmt.Println("冬")
	case 3, 4, 5:
		fmt.Println("春")

	case 6, 7, 8:
		fmt.Println("夏")

	case 9, 10, 11:
		fmt.Println("秋")

	default:
		fmt.Println("月の値ではない")
	}
}

func practice_switch2() {

	x := 5
	switch x {
	case f(1):
		fmt.Println("* first case. *")
	case f(2):
		fmt.Println("* second case. *")
	case f(3):
		fmt.Println("* third case. *")
	default:
		fmt.Println("* default case. *")
	}
	// No,  1
	// No,  2
	// No,  3
	// * default case. *
}
func f(n int) int {
	fmt.Println("No, ", n)
	return n
}

func practice_switch3() {
	fmt.Println("偶数奇数のswitch")
	x := input("type a number")
	if n, err := strconv.Atoi(x); err == nil {
		switch {
		case n%2 == 0:
			fmt.Println("偶数")
		case n%2 == 1:
			fmt.Println("奇数")
		}
	}
}

func practice_switch4() {
	x := input("type 1 - 5")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Print(x + "までの距離は")
	} else {
		return
	}
	t := 0
	switch n {
	case 5:
		t += 5
		fallthrough
	case 4:
		t += 4
		fallthrough
	case 3:
		t += 3
		fallthrough
	case 2:
		t += 2
		fallthrough
	case 1:
		t++
	default:
		fmt.Println("範囲外です")
		return
	}
	fmt.Println(t, "です")

	// type 1 - 5
	// 3
	// 3までの距離は6 です
}

func practice_for1() {
	// 一般的なforの書き方
	x := input("type a number")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Println("1から" + x + "の合計は")
	}
	t := 0
	c := 1
	for c <= n {
		t += c
		c++
	}
	fmt.Println(t, "です")
}

func practice_for2() {

	// 無限ループ
	// for{}

	// continue
	t := 0
	for {
		t++
		if t <= 3 {

			fmt.Println("continue: ", t)
			continue
		} else if t >= 3 {
			fmt.Println("break: ", t)
			break
		}
	}
	// continue:  1
	// continue:  2
	// continue:  3
	// break:  4
}

func practice_goto() {
	t := 0
	n := 5
	for i := 1; i <= n; i++ {
		fmt.Println(i)
		if i == 3 {
			goto err
		}
	}
	fmt.Println("total:", t)
	return

err:
	fmt.Println("goto err called!")

	// 1
	// 2
	// 3
	// goto err called!
}

func practice_arr() {
	// 配列の宣言

	// var 変数名 [要素数] 型
	var arr1 [3]int
	fmt.Println(arr1)
	// [0 0 0]

	// var 変数名 [要素数] 型　{値1, 値2, ...}
	var arr2 [2]string = [2]string{"a", "b"}
	fmt.Println(arr2)
	// [a b]

	// 変数名 := [...] 型　{値1, 値2, ...}
	arr3 := [...]float32{0.1, 0.2}
	fmt.Println(arr3)
	// [0.1 0.2]

	// rangeを使って配列を繰り返し処理する
	for i, v := range arr2 {
		fmt.Println(i, v)
	}
	// a
	// b
}

func practice_slice1() {
	// スライスの宣言

	// いきなりスライスを作成
	// 変数名 := []型　{値1, 値2, ...}
	slice1 := []int{10, 20, 30, 40, 50}
	fmt.Println(slice1)
	// [10 20 30 40 50]

	// 配列からスライスを取りだす
	arr := [5]int{10, 20, 30, 40, 50}
	// 変数名 := 配列[開始:終了]
	slice2 := arr[0:3]
	fmt.Println(slice2)
	// [10 20 30]
}

func practice_slice2() {
	// スライスは配列の参照にすぎない
	// スライスの長さ: スライスに保管されている値の数(配列のlenと同じ)
	// スライスの要領: スライスが参照する配列の大きさ

	a := [5]int{1, 2, 3, 4, 5}
	b := a[0:3]
	fmt.Println("配列: ", a)
	fmt.Println("スライス: ", b)
	// 配列:  [1 2 3 4 5]
	// スライス:  [1 2 3]

	// 元の配列に値を足すとスライスの値も増えている
	a[0] = 10
	fmt.Println("配列: ", a)
	fmt.Println("スライス: ", b)
	// 配列:  [10 2 3 4 5]
	// スライス:  [10 2 3]

	// スライスの値を足した場合も同様で、元の配列の値も増えている
	b[1] = 20
	fmt.Println("配列: ", a)
	fmt.Println("スライス: ", b)
	// 配列:  [10 20 3 4 5]
	// スライス:  [10 20 3]

	// スライスにappendすると配列の対応する番号の値も置き変わる
	b = append(b, 30)
	fmt.Println("配列: ", a)
	fmt.Println("スライス: ", b)
	// 	配列:  [10 20 3 30 5]
	// スライス:  [10 20 3 30]

	// ただし要素数を超えてappendした場合は元の配列は値が増えない
	b = append(b, 40)
	b = append(b, 50)
	b = append(b, 60)
	fmt.Println("配列: ", a)
	fmt.Println("スライス: ", b)
	// 配列:  [10 20 3 30 40]
	// スライス:  [10 20 3 30 40 50 60]

}

func practice_slice3() {

	// Goは配列の要素を先頭に追加するなどの細かい関数はないので自分で実装することになる

	// 末尾に追加
	push := func(a []int, v int) []int { return append(a, v) }

	// 末尾を削除
	pop := func(a []int) []int { return a[:len(a)-1] }

	// 先頭に追加
	unshift := func(a []int, v int) []int { return append([]int{v}, a...) }

	// 先頭を削除
	shift := func(a []int) []int { return a[1:] }

	// 任意の場所に追加
	insert := func(a []int, v int, p int) []int {
		a = append(a, 0)
		fmt.Println("1 ", a)
		a = append(a[:p+1], a[p:len(a)-1]...)
		fmt.Println("2 ", a)
		fmt.Println("3 ", a[p])
		a[p] = v
		return a
	}
	// 任意の場所を削除
	remove := func(a []int, p int) []int { return append(a[:p], a[p+1]) }

	a := []int{10, 20, 30}
	fmt.Println(a)
	// [10 20 30]

	// Push
	a = push(a, 1000)
	fmt.Println(a)
	// [10 20 30 1000]

	// pop
	a = pop(a)
	fmt.Println(a)
	// [10 20 30]

	// unshift
	a = unshift(a, 1000)
	fmt.Println(a)
	// [1000 10 20 30]

	// shift
	a = shift(a)
	fmt.Println(a)

	// insert
	a = insert(a, 1000, 2)
	fmt.Println(a)
	// 	[10 20 1000 30]

	// remove
	a = remove(a, 2)
	fmt.Println(a)
	// [10 20 30]
}

func practice_map() {
	// macの宣言
	// var 変数名 map[キーの型]値の型
	var map1 map[int]string
	fmt.Println(map1)
	// map[]

	map2 := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	fmt.Println(map2)
	// map[1:a 2:b 3:c]

	map2[4] = map2[1] + map2[2] + map2[3]
	fmt.Println(map2)
	// map[1:a 2:b 3:c 4:abc]

	delete(map2, 1)
	fmt.Println(map2)
	// map[2:b 3:c 4:abc]
}

func practice_map_for() {

	// mapはキーの順番を保証しないので注意
	m := map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
	}
	for k, v := range m {
		fmt.Println(k+": ", v)
	}
	// a:  100
	// b:  200
	// c:  300
}

func practice_func1() {
	push := func(a []string, v string) ([]string, int) { return append(a, v), len(a) }
	pop := func(a []string) ([]string, string) { return a[:len(a)-1], a[len(a)-1] }

	m := []string{}
	m, _ = push(m, "apple")
	m, _ = push(m, "banana")
	m, _ = push(m, "orange")

	fmt.Println(m)
	// [apple banana orange]

	m, v := pop(m)
	fmt.Println("get "+v+" -> ", m)
	// get orange ->  [apple banana]
}

func practice_func2() {

	insert := func(a []string, v string, p int) (s []string) {
		s = append(a, "")
		s = append(s[:p+1], s[p:len(s)-1]...)
		s[p] = v
		return
	}
	m := []string{
		"one", "two", "three",
	}
	fmt.Println(m)
	// 	[one two three]

	m = insert(m, "*", 2)
	m = insert(m, "*", 1)
	fmt.Println(m)
	// [one * two * three]

}
func practice_func3() {
	// 可変長引数
	// func 関数名(変数 ...型) 戻り値{...}

	push := func(a []string, v ...string) (s []string) {
		s = append(a, v...)
		return
	}
	m := []string{"one", "two", "three"}
	fmt.Println(m)
	// [one two three]

	m = push(m, "1", "2", "3")
	fmt.Println(m)
	// [one two three 1 2 3]

}

func practice_anonymous_func() {
	f := func(a []string) ([]string, string) { return a[1:], a[0] }
	m := []string{"one", "two", "three"}
	s := ""
	fmt.Println(m)
	// [one two three]

	for len(m) > 0 {
		m, s = f(m)
		fmt.Println(s+" ->", m)
	}
	// one -> [two three]
	// two -> [three]
	// three -> []

}

func practice_func4() {

	modify := func(a []string, f func([]string) []string) []string { return f(a) }
	m := []string{"1st", "2nd", "3rd"}
	fmt.Println(m)
	// [1st 2nd 3rd]

	m1 := modify(m, func([]string) []string { return append(m, m...) })
	fmt.Println(m1)
	// [1st 2nd 3rd 1st 2nd 3rd]

	m2 := modify(m, func([]string) []string { return m[:len(m)-1] })
	fmt.Println(m2)
	// [1st 2nd]

	m3 := modify(m, func([]string) []string { return m[1:] })
	fmt.Println(m3)
	// [2nd 3rd]

}

func practice_closure() {
	modify := func(d string) func() []string {
		m := []string{"1st", "2nd"}
		// 関数の中に値が保持される
		return func() []string { return append(m, d) }
	}
	data := "*新しい値*"
	m1 := modify(data)
	data = "+new data+"
	m2 := modify(data)

	fmt.Println(m1())
	fmt.Println(m2())
}

func practice_printf() {
	n := 123
	b := true
	s := "hello"
	fmt.Printf("number:%d, boot:%t, string:%s", n, b, s)
	// number:123, boot:true, string:hello

	// %t bool
	fmt.Printf("%t", true)
	// true

	// %d 10進数
	fmt.Printf("%d", 256)
	// 256

	// %b 2進数
	fmt.Printf("%b", 256)
	// 100000000

	// %o 8進数　0xなし
	fmt.Printf("%o", 256)

	//400

	// %O 8進数　0あり
	fmt.Printf("%O", 256)
	//0o400

	// %x 16進数 小文字
	fmt.Printf("%x", 10000)
	// 100

	// %X 16進数 大文字
	fmt.Printf("%X", 10000)
	// 100

	// %q 整数
	fmt.Printf("%q", 256)
	//'Ā'

	// %e 実数値 科学記号e始まり
	fmt.Printf("%e", -1.234456e+78)
	//-1.234456e+78

	// %E 実数値 科学記号E始まり
	fmt.Printf("%E", -1.234456e+78)
	//-1.234456e+78

	// %f %F 実数値
	fmt.Printf("%f", -1.234456e+78)
	//-1234456000000000008479131885172305811737921938103077369281321992118750003855360.000000

	fmt.Printf("%F", -1.234456e+78)
	//-1234456000000000008479131885172305811737921938103077369281321992118750003855360.000000

	// %g %G 実数値 大きい桁は科学記号にする
	fmt.Printf("%g", -1.234456e+78)
	//-1.234456e+78

	fmt.Printf("%G", -1.234456e+78)
	//-1.234456e+78

	// %s 文字列
	fmt.Printf("%s", "ABC")
	// "ABC"

	// %q char
	fmt.Printf("%q", 'a')
	// 'a'

	// %p ポインタアドレス
	pointer := "pointer"
	fmt.Printf("%p", &pointer)
	//0xc000108060

}

func main() {
	practice1()
	practice2()
	practice3()
	practice_const()
	practice_if()
	practice_switch1()
	practice_switch2()
	practice_switch3()
	practice_switch4()
	practice_for1()
	practice_for2()
	practice_goto()
	practice_arr()
	practice_slice1()
	practice_slice2()
	practice_slice3()
	practice_map()
	practice_map_for()
	practice_func1()
	practice_func2()
	practice_func3()
	practice_anonymous_func()
	practice_func4()
	practice_closure()
	practice_printf()
}
