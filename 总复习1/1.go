package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(fun())

	fmt.Println(fun1(100))

	fun2(1,5,6,6,1,6,1,6,1,6,1)

    fmt.Println(fun3("cniowncownoicnw"))

	fun4("bcuuoeblcscwibcobbrowbreo")
}
func fun()int  {
	sun:=0
	for i:=1;i<10 ;i++  {
		sun+=i
	}
	return sun
}
func fun1(n int)int  {
	sun:=0
	for i:=0;i<=n ;i++  {
		sun+=i
	}
	return sun
}
func fun2(n...int) {
    sun:=0
	for i:=0;i<len(n) ;i++  {
		sun+=n[i]
		//fmt.Println(sun)
	}
	fmt.Println(sun)
}

func fun3(n string)map[string]int  {
     m:=make(map[string]int)
     s:=""
	for i:=0;i<len(n) ; i++ {
		s=string(n[i])//取到每个字符,是每一个，单独的！！！！！！！！！！
		if m[s]==0 {
			m[s]=strings.Count(n,s)
		}
	}
	return m
}

func fun4(str  string)string  {
	//m:=make(map[string]int)
	//s:=""
	for i:='a';i<'z' ;i++  {
		a:=string(i)
		n:=strings.Count(str,a)
		if n!=0 {
			fmt.Println("这个字符",a,"一共出现",n)
		}
	}
	return str
}