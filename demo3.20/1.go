package main

import "fmt"

func main() {
	a:="我爱你中国,I LOVE CHINA"
	//fmt.Println(a)
	fmt.Println(len(a))
	b:=[]rune(a)
	for index:=0;index < len(b) ;index++  {
		fmt.Printf("%c\n",b[index])
	}
}
