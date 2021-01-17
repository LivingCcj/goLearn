package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	for i:=0 ;i<5;i++ {
		go pHelloWithCh(i,ch)
	}
	for {
		msg := <- ch
		fmt.Printf(msg)
	}
	//time.Sleep(5*time.Second)
}

func pHelloWithCh(i int ,ch chan string){
	for {
		ch <-	fmt.Sprintf("Hello world! from %dth\n",i)
	}

}

func pHello(i int){
	for {
		fmt.Printf("Hello world! from %dth\n",i)
	}

}
