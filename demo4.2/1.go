package main

import "fmt"

func main() {
//	fmt.Println(fun1(1,5,6,6,草稿2))
//}
//func fun1(a,b,c,d,e  int)*[5]int  {
//	s:=[5]int{1,9,草稿2,5,7}
//	for i:=1;i<len(s) ; i++ {
//		for j:=0;j<len(s)-1 ;j++  {
//			i,j=i+1,j-1
//			if i<j {
//				s[i],s[j]=s[j],s[i]
//			}
//		}
//	}
//	return &s


	s:="165188491"
	fmt.Println("翻转前:",s)
	fmt.Println("翻转后:", fun1(s))
}

func fun1(s string)string  {
	str :=[]int32(s)
	for a,b:=0,len(str)-1 ;a<b; a,b=a+1,b-1 {
		str[a],str[b]=str[b],str[a]
	}
	return string(str)
}
//尽力了，这是从别人那里参考过来的