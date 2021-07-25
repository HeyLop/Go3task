package main

// @Description  go3task week-02
// @Author playclouds
// @Update    2021/7/25 9:26

//我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
//应该使用warp简单包装加入报错的方法，具体判断有业务使用errors.is判断是否返回为空，进行解析。

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

var (
	userName string = "root"
	password string = "111111"
	ipAdder  string = "127.0.0.1"
	port     int    = 3306
	dbName   string = "go3task"
	charset  string = "utf8"
)
var Db *sql.DB

// use mysql driver init db;
func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAdder, port, dbName, charset)
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer Db.Close()
}

// FindNameByNum Simulation of business calls
func handleFindNameByNum(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	num := query.Get("num")
	if num == "" {
		writer.Write([]byte(fmt.Sprintf("arg num is Empty")))

	}
	selectNum, err := strconv.Atoi(num)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf("arg num is illegal characters")))
		return
	}
	name, err := QueryUserByNum(selectNum)

	if errors.Is(err, sql.ErrNoRows) {

		writer.Write([]byte(fmt.Sprintf("no name res with query num：%s", num)))
		return
	}
	writer.Write([]byte(fmt.Sprintf("Find num %s name：%s", num, name)))
	return

}

func QueryUserByNum(num int) (name string, err error) {

	query := "select name from user WHERE num = ?"
	err = Db.QueryRow(query, num).Scan(&name)

	if errors.Is(err, sql.ErrNoRows) {
		errors.Wrap(err, "Func: QueryUserByNum")
		return

	}
	return
}

func main() {
	//name, err := QueryUserByNum(1)
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		fmt.Println("no name res with query num ")
	//		return
	//
	//	}
	//}
	//fmt.Println("name:", name)

	// http://127.0.0.1:8080/user?num=0
	http.HandleFunc("/user", handleFindNameByNum)
	http.ListenAndServe(":8080", nil)

}
