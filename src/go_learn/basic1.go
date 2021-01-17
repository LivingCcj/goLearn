/*
 * @Description: go 语言基础学习1：fmt，基本类型，数组，切片，链表，字典，for-range，
 * @Author: ccj
 * @Date: 2020-05-02 22:42:42
 * @LastEditTime: 2020-12-28 16:12:56
 * @LastEditors:  
 */
 package basic
 
 import
 (
	 "container/list"
	 "fmt"
	 "time"
	 "math"
 )
 

 // 定义一个int的类型别名

func Learn1(){

	time.Sleep(time.Second)


	// 多重赋值
	var a int = 1
	var b int = 2

	b,a=a,b
	// 匿名变量
	firstName,_:=getName()
	_,secondName:=getName() 

	fmt.Printf("first name :%v, second name %v \n",firstName,secondName)


	// 整型 int8 - int64  uint8 - int64 

	// 浮点型 float32 float64
	fmt.Printf("%.3f\n",math.E)

	// 布尔类型 bool ： true 和 false 无法转化为整型
	var flag bool = true
	fmt.Printf("%v\n",flag)
	// 字符串型：byte(int8,只能表示ASCII) 和 rune(int32,可以表示任意字符:Unicode)

	f := "Golang学习"

	for _,h :=range f{
		fmt.Printf("%c ",h)
	}

	// 指针 new
	str := new(string) //指向string 的类
	*str = "cuichaojie"
	fmt.Printf("\n type is %T , value is %v ,addresss is %p",str,str,&str)


	// 常量 const: 声明之后，不可以改变

	const (
		name = "cui"
		age = 25
	)
	fmt.Printf("\n %v 's age is %v\n",name,age)

	// Go的容器： 数组，切片，列表，字典，遍历
	

	// 1.数组 
	// var nameArray [3]string
	// nameArray :=[...] string{"ccj","cyj"}
	nameArray := new([3]string)
	nameArray[0]="ccj"
	nameArray[1]="cyj"

	fmt.Println(nameArray[0])

	// 2.切片 变长数组：array,len，cap>=len,常用函数copy，

	// var nameSlide []string
	// nameslide := []string{"ccj","cyj"}
	
	// sli = make([]int , 2, 4)

	// fmt.Printf("value is %v , len is %v, cap is %v \n", sli, len(sli), cap(sli))

	// 3.列表(双向链表)
	
	// var number list.List
	// or number:=list.New()
	number:=list.New()

	for i:=1 ;i<=10;i++{
		number.PushBack(i)
	}

	first :=number.PushFront(0)
	number.Remove(first)

	for l := number.Front();l!=nil;{
		fmt.Print(l.Value," ")
		l=l.Next()
	}


	fmt.Println()

	// 4. 字典
	nameMap := make(map[int]string)
	// or nameMap := map[int]string{ 初始值 }
	nameMap[0]="ccj"
	nameMap[1]="cyj"
	
	mate,ok:=nameMap[0]
	
	fmt.Printf("key is %v, isHere is %v\n",mate,ok)
	
	// 5.容器遍历 for-range

	for k,v := range nameMap{
		fmt.Println(k,v," ")
	}
	
}



func getName()(string,string){
	return "cui","chaojie"
}

