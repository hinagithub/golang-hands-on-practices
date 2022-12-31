package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := Input("type your name")
	fmt.Println(name)
}

func Input(msg string) string {
	fmt.Println("hello.Input() called.")
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println(msg)
	sc.Scan()
	return sc.Text()
}
