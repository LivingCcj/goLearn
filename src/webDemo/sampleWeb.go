package webDemo

import (
	json2 "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type SampleService struct {
	Route string
	Port  string
}
type User struct {
	Name string
	Habits []string
}

type SampleServer interface {
	RunServer()
	SayHelloService(http.ResponseWriter, *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Write(w http.ResponseWriter, r *http.Request)
}

func (sample *SampleService) RunServer() {
	//handle实现路由
	http.Handle("/", sample)

	//handleFunc实现路由
	http.HandleFunc(sample.Route, sample.Login)
	http.HandleFunc("/write",sample.Write)
	//run Gin Web服务
	sample.GinWeb()
	if err := http.ListenAndServe(sample.Port, nil); err != nil {
		log.Fatal("linsenAndserve: ", err)
	}
	fmt.Println("Server is start.")
}

func (sample *SampleService) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if content, _ := ioutil.ReadFile(".\\login.html"); content == nil {
			w.Write([]byte("resource has missing"))

		}else{
			w.Write(content)
		}
	} else {
		_=r.ParseForm()  //解析请求的参数，默认是不解析的
		fmt.Printf("%v:%v\n", r.Form.Get("username"), r.Form.Get("pwd"))
		if pwd := r.Form.Get("pwd"); pwd == "123456" {
			fmt.Fprintf(w, "Welcome %s", r.Form.Get("username"))
		} else {
			fmt.Fprintf(w, "pwd is error")
		}
	}
}

func (sample *SampleService) Write(w http.ResponseWriter, r *http.Request) {
	//响应头
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("X-Custom-Header","custom")
	//状态码
	w.WriteHeader(201)
	//响应数据

	user := &User{
		Name: "aoho",
		Habits: []string{"eat","sleep","run"},
	}

	json,_:=json2.Marshal(user)
	w.Write(json)
}


func (sample *SampleService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println(r.Form)
	for k, v := range r.Form {
		fmt.Printf("key:%s \t value:%s", k, v)
	}

	fmt.Printf("正在访问...，route：%s,port%s\n", sample.Route, sample.Port)
	//w.Write([]byte("hello there is sample serveice"))
	_, _ = fmt.Fprintf(w, "Hello Web,%v!", r.Form.Get("name"))

}

func (sample *SampleService) GinWeb(){
	router := gin.Default()

	router.GET(sample.Route, func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"welcome ccj home",
		})
	})
	router.Run("8000")
}
