package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
	"os"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var count int

type Post struct {
	No string
	Name string
	H    string
	A    string
	B    string
	C    string
	D    string
	S    string
	Sum  string
}

func Eucjp_to_utf8(str string) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.EUCJP.NewDecoder()))
	if err != nil {
		return "", err
	}
	return string(ret), err
}

func main() {
	var data [9]string
	var counter int
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	doc, err := goquery.NewDocument("https://yakkun.com/sm/status_list.htm")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	doc.Find("table").Find("td").Each(func(_ int, s *goquery.Selection) {
		text, _ := Eucjp_to_utf8(s.Text())
		data[counter] = text
		counter++
		if counter == 9 {
			stmtIns, errs := db.Prepare(fmt.Sprintf("INSERT INTO value (No,Name,H,A,B,C,D,S,Sum) VALUES (?,?,?,?,?,?,?,?,?)"))
			if errs != nil {
				panic(err.Error())
			}
			defer stmtIns.Close()
			_, err = stmtIns.Exec(data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8])
			counter = 0
		}
	})
}
