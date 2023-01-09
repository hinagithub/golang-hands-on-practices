// 並列処理

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	msg := "start"
	prmsg := func(nm string, n int) {
		fmt.Println(nm, msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
	hello := func(n int) {
		const nm string = "hello"
		for i := 0; i < 10; i++ {
			msg += " h" + strconv.Itoa(i)
			prmsg(nm, n)
		}
	}
	main := func(n int) {
		const nm string = "*main"
		for i := 0; i < 5; i++ {
			msg += " m" + strconv.Itoa(i)
			prmsg(nm, 100)
		}
	}
	go hello(60)
	main(100)

	// goを頭につけないとhelloが9回出た後にmainが5回表示される
	// goを頭につけると並列実行されるのでmainとhelloが交互に出る感じになる

	// *main start m0
	// hello start m0 h0
	// hello start m0 h0 h1
	// *main start m0 h0 h1 m1
	// hello start m0 h0 h1 m1 h2
	// hello start m0 h0 h1 m1 h2 h3
	// *main start m0 h0 h1 m1 h2 h3 m2
	// hello start m0 h0 h1 m1 h2 h3 m2 h4
	// *main start m0 h0 h1 m1 h2 h3 m2 h4 m3
	// hello start m0 h0 h1 m1 h2 h3 m2 h4 m3 h5
	// hello start m0 h0 h1 m1 h2 h3 m2 h4 m3 h5 h6
	// *main start m0 h0 h1 m1 h2 h3 m2 h4 m3 h5 h6 m4
	// hello start m0 h0 h1 m1 h2 h3 m2 h4 m3 h5 h6 m4 h7
	// hello start m0 h0 h1 m1 h2 h3 m2 h4 m3 h5 h6 m4 h7 h8

}
