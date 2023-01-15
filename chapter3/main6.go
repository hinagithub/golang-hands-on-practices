// チャンネル

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func channel1() {
	total := func(n int, c chan int) {
		t := 0
		for i := 1; i <= n; i++ {
			t += i
		}
		c <- t
	}

	c := make(chan int)
	go total(1000, c)
	go total(100, c)
	go total(10, c)
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z)
	// 55 500500 5050

	// ↑処理が終わったものから順に取り出せる仕様なので順番に注意すること

}

func channel2() {
	total := func(c chan int) {
		n := <-c
		fmt.Println("n: ", n)
		t := 0
		for i := 1; i <= n; i++ {
			t += i

		}
		fmt.Println("total:", t)

	}
	c := make(chan int)

	// ゴルーチン実行前に値を設定することはできない。
	// c <- 100
	go total(c)
	c <- 100
	time.Sleep(100 * time.Millisecond)

	// 	n:  100
	// total: 5050
}

func channel3() {
	prmsg := func(n int, s string) {
		fmt.Println(s)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	first := func(n int, c chan string) {
		const nm string = "first-"
		for i := 0; i < 10; i++ {
			s := nm + strconv.Itoa(i)
			prmsg(n, s)
			c <- s
		}

	}

	second := func(n int, c chan string) {
		for i := 0; i < 10; i++ {
			prmsg(n, "second:["+<-c+"]")
		}

	}

	c := make(chan string)
	go first(10, c)
	second(10, c)
	// first-0
	// first-1
	// second:[first-0]
	// second:[first-1]
	// first-2
	// second:[first-2]
	// first-3
	// second:[first-3]
	// first-4
	// second:[first-4]
	// first-5
	// second:[first-5]
	// first-6
	// first-7
	// second:[first-6]
	// first-8
	// second:[first-7]
	// first-9
	// second:[first-8]
	// second:[first-9]

}

// チャネルで双方向でやり取りする
func channel4() {
	total := func(cs chan int, cr chan int) {
		n := <-cs
		fmt.Println("n: ", n)
		t := 0
		for i := 1; i <= n; i++ {
			t += i

		}
		cr <- t
	}

	cs := make(chan int)
	cr := make(chan int)
	go total(cs, cr)
	cs <- 100
	fmt.Println("total: ", <-cr)
	// n:  100
	// total:  5050
}

func channel5() {
	count := func(n int, s int, c chan int) {
		for i := 1; i <= n; i++ {
			c <- i
			time.Sleep(time.Duration(s) * time.Millisecond)
		}
	}
	n1, n2, n3 := 3, 5, 10
	m1, m2, m3 := 100, 75, 50

	c1 := make(chan int)
	go count(n1, m1, c1)

	c2 := make(chan int)
	go count(n2, m2, c2)

	c3 := make(chan int)
	go count(n3, m3, c3)

	for i := 0; i < n1+n2+n3; i++ {
		select {
		case re := <-c1:
			fmt.Println("* first ", re)

		case re := <-c2:
			fmt.Println("* second ", re)

		case re := <-c3:
			fmt.Println("* third ", re)
		}
	}
	fmt.Println("*** finish ***")
	// * third  1
	// * first  1
	// * second  1
	// * third  2
	// * second  2
	// * third  3
	// * first  2
	// * third  4
	// * second  3
	// * first  3
	// * third  5
	// * second  4
	// * third  6
	// * second  5
	// * third  7
	// * third  8
	// * third  9
	// * third  10
	// *** finish ***

}

func channel6() {

	// StrData is structure
	type StrData struct {
		msg string
		mux sync.Mutex
	}

	sd := StrData{msg: "Start"}
	prmsg := func(nm string, n int) {
		fmt.Println(nm, sd.msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	main := func(n int) {
		const nm string = "*main"
		sd.mux.Lock() // 　　　　　　＊
		for i := 0; i < 5; i++ {
			sd.msg += " m" + strconv.Itoa(i)
			prmsg(nm, 100)
		}
		sd.mux.Unlock() // 　		*
	}

	hello := func(n int) {
		const nm string = "hello"
		sd.mux.Lock() // 			*
		for i := 0; i < 5; i++ {
			sd.msg += " h" + strconv.Itoa(i)
			prmsg(nm, n)
		}
		sd.mux.Unlock() // 			*
	}

	go main(100)
	go hello(50)
	time.Sleep(5 * time.Second)

	// lockとunlockをしないと並列処理されてランダムに表示されるが使っているので綺麗に順番に表示される
	// *main Start m0
	// *main Start m0 m1
	// *main Start m0 m1 m2
	// *main Start m0 m1 m2 m3
	// *main Start m0 m1 m2 m3 m4
	// hello Start m0 m1 m2 m3 m4 h0
	// hello Start m0 m1 m2 m3 m4 h0 h1
	// hello Start m0 m1 m2 m3 m4 h0 h1 h2
	// hello Start m0 m1 m2 m3 m4 h0 h1 h2 h3
	// hello Start m0 m1 m2 m3 m4 h0 h1 h2 h3 h4
}

func main() {
	// channel1()
	// channel2()
	// channel3()
	// channel4()
	// channel5()
	channel6()
}
