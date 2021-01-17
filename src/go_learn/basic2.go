/*
 * @Description: go 语言基础学习2：结构体，接口，匿名函数，闭包，
 * @Author: ccj
 * @Date: 2020-12-28 15:45:10
 * @LastEditTime: 2020-12-28 17:01:38
 * @LastEditors:  
 */


package basic

import(
	"fmt"
	"time"
)
 
 // mark: 结构体定义
 
 type Person struct{
	 Name string
	 Birth string
	 Age int
 }
 
 // mark: 匿名函数声明
 func proc(input string, processor func(str string)){
	 processor(input)
 }
 // mark: 函数首字母大写，表示包外可以访问的函数
 func Proc(input string){
	 // 
 }
 // mark：接口和嵌套 ,关键字:type 和 interface，其中，接口需要实现
 
 type Tank interface{
	 func1()
	 func2()
 }
 type Plane interface{
	 Fly()
 }
 // 嵌套对应的是：接口的继承
 type TPlane interface{
	 Tank
	 Plane
 }
 
 func Learn2()  {
	 // 匿名函数, 没有函数名字，常用于回调函数
	proc("ccj",func(str string){
		for _,v:= range str{
			fmt.Printf("value is %c\n",v)
		}
	})

	// 函数遍历类型

	currentTime :=func(){
		fmt.Println(time.Now())
	}
	// 调用匿名函数
	for i :=1 ;i<=10;i++{
		currentTime()
	}
	currentTime()

	
	// 调用闭包函数
	c1:=createCounter(-1)
	fmt.Println(c1())
	fmt.Println(createCounter(10)())

	// mark: 结构体初始化和示例
	var p1 Person
	// or p1:=new(Person)
	// or p1:=&Person{}
	// or  p1:=&Person{"ccj","1995-04",25}
	p1.Name="ccj"
	p1.Birth="1995-04"
	p1.Age=25

	p1.PrintPerson()
	p1.changeName("cyj")
	p1.PrintPerson()

	// mark: 结构体接口测试

	catDog := &CatDog{
		"cyj",
	}
	var cat Cat
	cat = catDog //catDog的指针类型付给cat
	cat.CatchMouse()

	var dog Dog
	dog=catDog //catDog 指针赋值给dog
	dog.Bark()

	// mark: 内嵌(组合) 测试
	
	wild := &WildDuck{
		Swim{
			"ccj",
		},
		Fly{
			"ccj",
		},
	}
	wild.Flying()
	wild.Swimming()
	domestic := &DomesticDuck{
		Swim{
			"cyj",
		},
	}
	domestic.Swimming()


 }

 // mark: 结构体方法，带有接收器得函数，类似面向对象得this，self

// 指针类型接收器
func (person *Person) changeName(name string){
	person.Name=name
}


// 非指针类型接收器
func (person Person) PrintPerson(){
	fmt.Printf("person name is %v, birthday is %v, age is %v\n",person.Name,person.Birth,person.Age)
}

// mark: 闭包：携带状态的函数，将函数外部和函数内部连接起来的桥梁，通过闭包我们可以读取函数内的变量

// 闭包实现计时器
func createCounter(start int) func() int{
	if start<0{
		start=0
	}
	

	return func() int{
		start++
		return start
	}
}

// mark: 结构体实现接口 方法签名和接口签名，方法需要全部实现

// create 两个接口
type Cat interface{
	CatchMouse()
}

type Dog interface{
	Bark()
}

type CatDog struct{
	Name string
}

// 结构体接口实现
func (catDog *CatDog) CatchMouse(){
	fmt.Printf("%v start to catch mouse\n",catDog.Name)
}

func (catDog *CatDog) Bark(){
	fmt.Printf("%v start to bark....\n",catDog.Name)
}

// mark: 内嵌和组合

type Swim struct{
	Name string
}

func (swim *Swim) Swimming(){
	fmt.Printf("%v can swim\n",swim.Name)
}

type Fly struct{
	Name string
}

func (fly *Fly) Flying(){
	fmt.Printf("%v cam fly\n",fly.Name)
}

// 内嵌结构

type WildDuck struct{
	Swim
	Fly
}
type DomesticDuck struct{
	Swim
}

