package main

import (
	"errors"
	"fmt"
)

func main() {
	//作业一：有结构体Person，包含name、age、address等属性。
	//定义一个函数接收结构体类型，判断传入的person的年龄是否已经成年，
	//如果已成年，返回true，如果未成年，返回false；如果年龄是0或者负数，使用error返回年龄不合法等提示信息
	a := Person{"王晓伟", 0, "北京市xxxx"}
	fmt.Println(a)

     adult(a)

    //fmt.Println(a.adult1())
}

type Person struct {
	name    string
	age     int
	address string
}

func adult(a Person) bool {
	var err  error
	if a.age > 20 {
		fmt.Println("已成年")
		return true
	}
	if a.age < 20{
		fmt.Println("未成年")
		return false
	}
	if a.age == 0 || a.age < 0 {
		err = errors.New("年龄不合法")
		fmt.Println(err)
	}
	return true
}

//func (a Person) adult1() error {
//	var err  error
//	if a.age == 0 || a.age < 0 {
//    err = errors.New("年龄不合法")
//	}
//	return err
//}
