package main

import "fmt"

// General is all type data.
type General interface{}

//GData is holding General value
type GData interface {
	Set(nm string, g General)
	Print()
}

// GDataImpl is structure.
type GDataImpl struct {
	Name string
	Data General
}

// Set is GDataImpl method.
func (gd *GDataImpl) Print() {
	fmt.Printf("<<%s>>", gd.Name)
	fmt.Println(gd.Data)

}

func main() {

	// 空のGeneralインターフェース P158〜

	type General interface{}
	var v General
	v = 123
	v = 0.01
	v = "Hello"
	v = true
	fmt.Println("どんな型を入れてもエラーにならない", v)
	// どんな型を入れてもエラーにならない true

	var data = []GDataImpl{}
	data = append(data, GDataImpl{"Taro", 123})
	data = append(data, GDataImpl{"Hanako", "hello"})
	data = append(data, GDataImpl{"Sachiko", []int{1, 2, 3}})
	for _, ob := range data {
		ob.Print()
	}
	// <<Taro>>123
	// <<Hanako>>hello
	// <<Sachiko>>[1 2 3]

}
