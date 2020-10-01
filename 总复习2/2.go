package main

import (
	"errors"
	"fmt"
)

func main() {
	person := Person{
		name:"王大锤",
		age :-1,
		address:"北京...",
	}
	//1、多变量接收
	//草稿2、使用结构体对象调用对应的方法（方法接收值的限定）
	isYoung, err := person.checkAge()
	if err != nil {//处理error
		fmt.Println(err)
		return
	}
	if isYoung{
		fmt.Println("已成年")
	}else {
		fmt.Println("未成年")
	}

	fmt.Println(" =============== ")
	arr := [6]string{"a","b","c","d","e","f"}
	for index := 0;index <10;index++{
		if index < len(arr){
			fmt.Println(arr[index])
		}else {
			panic("数组越界")//index out of range ...
		}
	}

}

/**
  有结构体Person，包含name、age、address等属性。
  定义一个函数接收结构体类型，判断传入的person的年龄是否已经成年，
	1、如果已成年，返回true，---> bool
    草稿2、如果未成年，返回false；---> bool
	3、如果年龄是0或者负数，使用error返回年龄不合法等提示信息。
*/

type Person struct {
	name string
	age int
	address string
}

func (per Person) checkAge()(bool,error){
	if per.age >= 18{//20
		return true, nil
	}else if per.age > 0 && per.age < 18{// 0 - 18 之间：15
		return false,nil
	}
	if per.age <= 0{// -1
		return false, errors.New("person的年龄不合法")
	}
	//默认的返回值
	return false,nil
}


