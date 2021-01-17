package sqlWeb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"time"
)

type SqlService struct {
	Route string
	Port  string
	Id int
	db *sql.DB
}
type User struct {
	Id int
	Name string
	Pwd string
	CreateTime string
}

type SampleServer interface {
	RunServer()
	SayHelloService(http.ResponseWriter, *http.Request)
}

var UserById = make(map[int]*User)
var UserByName =make(map[string][]*User)  //这里保存时个User的指针的切片

func (s *SqlService) RunServer(){
	http.HandleFunc("/login", s.login)
	//http.HandleFunc("/info", sql.userInfo)
	http.HandleFunc("/info",s.userInfoInMysql)
	s.db,_ = sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/sqlgo?charset=utf8")

	if err := http.ListenAndServe(s.Port,nil);err !=nil{
		log.Fatal("ListenAndServe is meeting :",err)
	}
}

func (s *SqlService) login(w http.ResponseWriter, r *http.Request){

	if r.Method == "GET"{
		t,_:=template.ParseFiles("login.html")
		log.Print(t.Execute(w,nil))
	}else{
		_=r.ParseForm()
		fmt.Printf("%s/%s\n",r.Form.Get("username"),r.Form.Get("pwd"))
		s.Id+=1
		user := User{s.Id,r.Form.Get("username"),r.Form.Get("pwd"),
			time.Now().Format("2006/01/02 15:04:06")}

		storeInMem(user)
		storeInMysql(s,user)
	}

}

func storeInMem(user User){
	UserById[user.Id]=&user
	UserByName[user.Name]=append(UserByName[user.Name],&user)
	fmt.Printf("store user %s is ok\n",user.Name)
}
func checkErr(err error){
	if err!=nil{
		fmt.Println("meet err is: ",err)
	}
}

func storeInMysql(s *SqlService,user User){
	stmt,err := s.db.Prepare("INSERT INTO user SET id=?,name=?,pwd=?,create_time=?")
	checkErr(err)
	_, err = stmt.Exec(user.Id, user.Name, user.Pwd, user.CreateTime)
	checkErr(err)
	stmt.Close()
}

func (sql *SqlService) userInfoInMem(w http.ResponseWriter,r *http.Request){
	//解析请求的数据
	_=r.ParseForm()
	for _,user:=range UserByName[r.Form.Get("name")]{
		fmt.Fprintf(w," %v",user)
	}
}

func (s  *SqlService) userInfoInMysql(w http.ResponseWriter,r *http.Request){
	r.ParseForm()

	name := r.Form.Get("name")

	res := s.queryByName(name)
	for _,user:=range res{
		//req =append(req,fmt.Fprintf("",))
		fmt.Fprintf(w,"%v\n",user)
	}
}
func (s *SqlService) queryByName(name string) []User{

	var user []User
	stmt,err:=s.db.Prepare("select  * from user where name=?")
	checkErr(err)
	rows,_:=stmt.Query(name)

	fmt.Println("query record list as follow:")

	for rows.Next(){
		var id int
		var name string
		var pwd string
		var createTime string

		err := rows.Scan(&id,&name,&pwd,&createTime)
		fmt.Printf("id:%d,name:%s,pwd:%s,createTime:%s",id,name,pwd,createTime)
		if err!=nil{
			fmt.Println("result is error")
			break
		}
		user  = append(user, User{id,name,pwd,createTime})
	}
	stmt.Close()
	return user
}

