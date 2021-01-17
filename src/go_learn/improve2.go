/*
 * @Description: 
 * @Author: ccj
 * @Date: 2020-12-28 21:34:19
 * @LastEditTime: 2020-12-28 21:44:23
 * @LastEditors:  
 */


 package basic

 import(
	 "fmt"
	 "time"
 )

 func Learn4(){

	ch1 := make(chan int)
	ch2 := make(chan int)
	go send(ch1,0)
	go send(ch2,10)

	time.Sleep(time.Second)

	for{
		select{
		case val:= <- ch1:
			fmt.Printf("get value %d from ch1\n",val)
		case val:= <- ch2:
			fmt.Printf("get value %d from ch2\n",val)
		case <- time.After(22 * time.Second):
			fmt.Printf("time out\n")
			return
		}
		
	}
 }
 func send(ch chan int, begin int){
	 for i:=begin;i<begin+10;i++{
		 ch <- i
	 }
 }


 
 