/*
 * @Description: 
 * @Author: ccj
 * @Date: 2020-05-02 22:42:42
 * @LastEditTime: 2020-12-28 22:18:12
 * @LastEditors:  
 */
 package main
 
 import
 (
	 "fmt"
	 "go_learn"
	//  "time"
	 "sync"
 )
 
//  定义全局sync变量
 var syncMap sync.Map
 var waitGroup sync.WaitGroup

func main(){

	fmt.Println("Hello world!")
	basic.Learn1()
	basic.Learn2()
	person := &basic.Person{
		"ccj","1995-04",25,
	}
	person.PrintPerson()

	basic.Learn3()
	// time.Sleep(time.Second*20)
	// go basic.Learn4()

	// mark: sync 同步包 Mutex,RWMutex,Map,WaitGroup
	goSize:=5

	
	// 初始化同步等待函数
	waitGroup.Add(goSize)
	
	for i:=0;i<goSize;i++{
		go addNumber(i*10)
	}

	// 阻塞主线程
	waitGroup.Wait()

	var size int

	syncMap.Range(func(key,value interface{}) bool{
		size++
		fmt.Println("key:vaule is ",key,":",value," ")
		return true
	})
	fmt.Printf("syncMap size is %d\n",size)
	value ,ok:=syncMap.Load(0)
	if ok{
		fmt.Println("key 0 has value is",value,"  ")
	}
}


func addNumber(start int){
	for i:=start;i<start+3;i++{
		syncMap.Store(i,i)
	}
	// 通知本次调用结束
	waitGroup.Done()
}