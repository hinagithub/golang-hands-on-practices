package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/PuerkitoBio/goquery"
	_ "github.com/mattn/go-sqlite3"
)

func write() {

	// write text function.
	wt := func(f *os.File, s string) {
		_, er := f.WriteString(s + "\n")
		if er != nil {
			panic(er)
		}
	}

	fn := "data.txt"
	f, er := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if er != nil {
		panic(er)
	}

	// defer close.
	defer f.Close()

	fmt.Println("*** start ***")
	wt(f, "*** start ***")
	for {
		s := Input("type message")
		if s == "" {
			break
		}
		wt(f, s)
	}
	wt(f, "*** end ***\n\n")
	fmt.Println("*** end ***")
	er = f.Close()
	if er != nil {
		fmt.Println(er)
	}
}

func readAll() {
	// read text function.
	rt := func(f *os.File) {
		s, er := ioutil.ReadAll(f)
		if er != nil {
			panic(er)
		}
		fmt.Println(string(s))
	}

	fn := "data.txt"

	f, er := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
	if er != nil {
		panic(er)
	}

	// defet close.
	defer f.Close()

	fmt.Println("<< start >>")
	rt(f)
	fmt.Println("<< end >>")
}

func readLines() {
	// read text function.
	rt := func(f *os.File) {
		r := bufio.NewReaderSize(f, 4096)
		for i := 1; true; i++ {
			s, _, er := r.ReadLine()
			if er != nil {
				break
			}
			fmt.Println(i, ":", string(s))
		}
	}

	fn := "data.txt"

	f, er := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
	if er != nil {
		panic(er)
	}

	// defet close.
	defer f.Close()

	fmt.Println("<< start >>")
	rt(f)
	fmt.Println("<< end >>")
	// << start >>
	// 1 : *** start ***
	// 2 : Hello!
	// 3 : this is test message.
	// 4 : *** end ***
	// << end >>
}

func file() {
	fs, er := ioutil.ReadDir(".")
	if er != nil {
		panic(er)
	}

	for _, f := range fs {
		fmt.Println(f.Name(), "(", f.Size(), ")")
	}
	// data.txt ( 54 )
	// main.go ( 1782 )
}

func http_call() {
	p := "https://golang.org"
	re, er := http.Get(p)
	if er != nil {
		panic(er)
	}
	defer re.Body.Close()

	s, er := ioutil.ReadAll(re.Body)
	if er != nil {
		panic(er)
	}
	fmt.Println(string(s))
}

func go_query() {
	p := "https://golang.org"
	doc, er := goquery.NewDocument(p)
	if er != nil {
		panic(er)
	}
	doc.Find("a").Each(func(n int, sel *goquery.Selection) {

		lk, _ := sel.Attr("href")
		println(n, sel.Text(), "(", lk, ")")
	})
	// 0( / )
	// 1 Why Go arrow_drop_down( # )
	// 2 Case Studies( /solutions#case-studies )
	// 3 Use Cases( /solutions#use-cases )
	// 4 Security Policy( /security/policy/ )
	// ...略
}

func jsondata() {
	p := "https://tuyano-dummy-data.firebaseio.com/mydata.json"
	re, er := http.Get(p)
	if er != nil {
		panic(er)
	}
	defer re.Body.Close()

	s, er := ioutil.ReadAll(re.Body)
	if er != nil {
		panic(er)
	}
	var data []interface{}
	er = json.Unmarshal(s, &data)
	if er != nil {
		panic(er)
	}

	fmt.Println(data)

	for i, im := range data {
		m := im.(map[string]interface{})
		fmt.Println(
			i,
			m["name"].(int),
			m["mail"].(string),
			m["tel"].(string),
		)
	}
	// 0 taro zero@zero 000-000
	// 1 tuyano syoda@tuyano.com 999-999
	// 2 hanako hanako@flower 888-888
	// 3 sachiko sachiko@happy 777-777

}

func jsondata2() {
	p := "https://tuyano-dummy-data.firebaseio.com/mydata.json"
	re, er := http.Get(p)
	if er != nil {
		panic(er)
	}
	defer re.Body.Close()

	s, er := ioutil.ReadAll(re.Body)
	if er != nil {
		panic(er)
	}

	var item []Mydata
	er = json.Unmarshal(s, &item)
	if er != nil {
		panic(er)
	}
	for i, im := range item {
		println(i, im.Str())
	}
	// 0 <"taro"zero@zero,000-000>
	// 1 <"tuyano"syoda@tuyano.com,999-999>
	// 2 <"hanako"hanako@flower,888-888>
	// 3 <"sachiko"sachiko@happy,777-777>
}

func getSqliteData() {
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	q := "select * from mydata;"
	rs, er := con.Query(q)
	if er != nil {
		panic(er)
	}
	for rs.Next() { // カーソルを順次スキャンし次がなくなったらfalseが返却される
		var md User
		er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		fmt.Println(md.Str())
	}
	// <"1"Taro,taro@yamada,39>
	// <"2"Hanako,hanako@flower,28>
	// <"3"Sachiko,sachiko@happy,17>
	// <"4"Jiro,jiro@change,6>
}

func selectById() {
	var qry string = "select * from mydata where id = ?"
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	for {
		// ●begin
		s := Input("id")
		if s == "" {
			break
		}
		n, er := strconv.Atoi(s)
		if er != nil {
			panic(er)
		}
		rs, er := con.Query(qry, n)
		// ●end
		if er != nil {
			panic(er)
		}

		for rs.Next() {
			var md User
			er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
			if er != nil {
				panic(er)
			}
			fmt.Println(md.Str())
		}
	}
	fmt.Println("***end***")
	// id
	// 1
	// <"1"Taro,taro@yamada,39>
	// id
	// 2
	// <"2"Hanako,hanako@flower,28>
	// id
	// 3
	// <"3"Sachiko,sachiko@happy,17>
	// id
	// 4
	// <"4"Jiro,jiro@change,6>
	// id
	// 5
	// id

	// ***end***
}

func findOne() {
	var qry string = "select * from mydata where id = ?"
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	for {
		// ●begin
		s := Input("id")
		if s == "" {
			break
		}
		n, er := strconv.Atoi(s)
		if er != nil {
			panic(er)
		}
		rs := con.QueryRow(qry, n)

		var md User
		er = rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		fmt.Println(md.Str())

	}
	fmt.Println("***end***")
}

func like() {
	var qry string = "select * from mydata where name like ? or mail like ?"
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	for {
		// ●begin
		s := Input("find")
		if s == "" {
			break
		}
		rs := con.QueryRow(qry, "%"+s+"%", "%"+s+"%")

		var md User
		er = rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		fmt.Println(md.Str())
	}
	fmt.Println("***end***")
}

func insert() {

	mydatafmRws := func(rs *sql.Rows) *User {
		var md User
		er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		return &md
	}

	showRecord := func(con *sql.DB) {
		qry := "select * from mydata"
		rs, _ := con.Query(qry)
		for rs.Next() {
			fmt.Println(mydatafmRws(rs).Str())
		}
	}

	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	nm := Input("name")
	ml := Input("mail")
	age := Input("age")
	ag, _ := strconv.Atoi(age)

	qry := "insert into mydata (name, mail, age) values (?,?,?)"
	con.Exec(qry, nm, ml, ag)
	showRecord(con)
	// name
	// tetsuko
	// mail
	// testuko@happy
	// age
	// 15
	// <"1"Taro,taro@yamada,39>
	// <"2"Hanako,hanako@flower,28>
	// <"3"Sachiko,sachiko@happy,17>
	// <"4"Jiro,jiro@change,6>
	// <"5"tetsuko,testuko@happy,15>
}

func update() {
	mydatafmRws := func(rs *sql.Rows) *User {
		var md User
		er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		return &md
	}

	mydatafmRw := func(rs *sql.Row) *User {
		var md User
		er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		return &md
	}

	showRecord := func(con *sql.DB) {
		qry := "select * from mydata"
		rs, _ := con.Query(qry)
		for rs.Next() {
			fmt.Println(mydatafmRws(rs).Str())
		}
	}
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	ids := Input("update ID")
	id, _ := strconv.Atoi(ids)
	qry := "select * from mydata where id = ?"
	rw := con.QueryRow(qry, id)
	tgt := mydatafmRw(rw)
	ae := strconv.Itoa(tgt.Age)
	nm := Input("name(" + tgt.Name + ")")
	ml := Input("mail(" + tgt.Mail + ")")
	ge := Input("age(" + ae + ")")
	ag, _ := strconv.Atoi(ge)

	if nm == "" {
		nm = tgt.Name
	}
	if ml == "" {
		ml = tgt.Mail
	}
	if ge == "" {
		ag = tgt.Age
	}

	qry = "update mydata set name=?, mail=?, age=? where id=?"
	con.Exec(qry, nm, ml, ag, id)
	showRecord(con)

}

func delete() {
	mydatafmRws := func(rs *sql.Rows) *User {
		var md User
		er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		return &md
	}

	mydatafmRw := func(rs *sql.Row) *User {
		var md User
		er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}
		return &md
	}

	showRecord := func(con *sql.DB) {
		qry := "select * from mydata"
		rs, _ := con.Query(qry)
		for rs.Next() {
			fmt.Println(mydatafmRws(rs).Str())
		}
	}
	con, er := sql.Open("sqlite3", "data.sqlite3")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	ids := Input("delete ID")
	id, _ := strconv.Atoi(ids)
	qry := "select * from mydata where id=?"
	rw := con.QueryRow(qry, id)
	tgt := mydatafmRw(rw)
	fmt.Println(tgt.Str())
	f := Input("delete it? (y/n)")
	if f == "y" {
		qry = "delete from mydata where id=?"
		con.Exec(qry, id)
	}
	showRecord(con)
}

func main() {
	// write()
	// readAll()
	// readLines()
	// file()
	// http_call()
	// go_query()
	// jsondata()
	// jsondata2()
	// getSqliteData()
	// selectById()
	// findOne()
	// like()
	// insert()
	// update()
	// delete()
	Markdown()
}

func Input(msg string) string {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println(msg)
	sc.Scan()
	return sc.Text()
}

// User is json structure.
type User struct {
	ID   int
	Name string
	Mail string
	Age  int
}

// Str get string value
func (u *User) Str() string {
	return "<\"" + strconv.Itoa(u.ID) + "\"" + u.Name + "," + u.Mail + "," + strconv.Itoa(u.Age) + ">"
}

// Mydata is json structure.
type Mydata struct {
	Name string
	Mail string
	Tel  string
}

// Str get string value
func (m *Mydata) Str() string {
	return "<\"" + m.Name + "\"" + m.Mail + "," + m.Tel + ">"
}
