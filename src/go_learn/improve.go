/*
 * @Description: 
 * @Author: ccj
 * @Date: 2020-12-28 20:11:58
 * @LastEditTime: 2020-12-28 20:50:29
 * @LastEditors:  
 */
package basic

import(
	"fmt"
	"reflect"
)


func Learn3(){
	// mark: 反射类型和种类
	typeP := reflect.TypeOf(Person{})
	fmt.Printf("the type is %v, the kind is %v\n",typeP,typeP.Kind())

	typePtr := reflect.TypeOf(&Person{})
	fmt.Printf("the type is %v, the kind is %v\n",typePtr,typePtr.Kind())


	// mark: 反射类变量和方法
	for i:=0 ; i<typeP.NumField();i++{
		fmt.Printf("field` name is %s,type is %s,kind is %s\n",
					typeP.Field(i).Name,
					typeP.Field(i).Type,
					typeP.Field(i).Type.Kind())
	}

	// mark: 反射值对象 Elem 相当于解引用,可以访问并修改目标地址
	person := &Person{
		"ccj","1995-04",25,
	}
	valueOfPerson := reflect.ValueOf(person).Elem()

	valueOfName := valueOfPerson.FieldByName("Name")
	if valueOfName.CanSet(){
		valueOfName.Set(reflect.ValueOf("cyj"))
	}
	person.PrintPerson()
	// mark: 通过反射调用方法
	methodHello:=reflect.ValueOf(helloReflect)
	// 
	methodHello.Call([]reflect.Value{})
}

func helloReflect(){
	fmt.Println("hello reflecte!")
}
