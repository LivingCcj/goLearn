package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sqlWeb"

	"time"
)

func main() {
	sqlService:=sqlWeb.SqlService{Route: "/", Port: ":8080"}
	sqlService.RunServer()
	//server := webDemo.SampleService{Route: "/login", Port: ":8080"}
	//
	//server.RunServer()
	//db,_ := sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/sqlgo?charset=utf8")
	//stmt,err:=db.Prepare("SELECT * FROM user ")
	//
	//if err !=nil{
	//	panic("select is error")
	//}
	//rows,_:=stmt.Query()
	//
	//fmt.Println(rows)

	now := time.Now().Format("2006/01/02 15:04:06")
	fmt.Println(now)


}
