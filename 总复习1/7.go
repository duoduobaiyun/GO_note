package main

import (
	"errors"
	"fmt"
)

func main() {

p:=Person{
	name:    "王小帅",
	age:     20,
	address: "北京市",
}

   peo,err :=  p.fun0()
	if err !=nil {
		fmt.Println(err)
		return
	}
	if peo {
		fmt.Println("已成年")
	}else {
		fmt.Println("未成年")
	}

}

//有结构体Person，包含name、age、address等属性。
//定义一个函数接收结构体类型，判断传入的person的年龄是否已经成年，
//1、如果已成年，返回true，---> bool
//草稿2、如果未成年，返回false；---> bool
//3、如果年龄是0或者负数，使用error返回年龄不合法等提示信息。

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) fun0() (bool, error) {
	if p.age > 20 {
		return true, nil
	} else if p.age < 20 && p.age > 0 {
		return false, nil
	} else if p.age <= 0 {
        return false,errors.New("年龄错误")
	}
	return false,nil
}
