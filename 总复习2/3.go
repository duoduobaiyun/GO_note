package main

import (
	"errors"
	"fmt"
)

func main() {
	//作业一：有结构体Person，包含name、age、address等属性。定义一个函数接收结构体类型，
	//判断传入的person的年龄是否已经成年，如果已成年，返回true，如果未成年，返回false；
	//如果年龄是0或者负数，使用error返回年龄不合法等提示信息
	p := Person{
		name:    "王晓伟",
		age:     20,
		address: "北京市",
	}

	bool1,err:= p.fun()
    if err != nil{
    	fmt.Println(err)
		return
	}
	if bool1 {
		fmt.Println("已成年")
	}else {
		fmt.Println("未成年")
	}
}
type Person struct {
	name string
	age  int
	address string
}

func(p Person) fun()(bool,error) {
	if p.age > 18 {
		return true, nil
	} else if p.age < 18 && p.age > 0 {
		return false, nil
	}
	if p.age <= 0 {
		return false, errors.New("错误")
	}
	return false,nil
}

