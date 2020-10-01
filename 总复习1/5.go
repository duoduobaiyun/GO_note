package main

import "fmt"

func main() {
	b:=[]int{9,5,2,10,3,6}
	reverse1(b)
	fmt.Println(reverse1)
	fmt.Printf("%T\n",b)

	for key,val:=range b {
		fmt.Printf("%d,%d\n",key,val)
	}
}

func reverse1(arr []int){//因为切片的关系,是浅拷贝，所以，返回值类型可以不写,浅拷贝的特点是,如果你改了值,剩下的值全都改变

	for index:=0; index<len(arr)/2;index++  {
		arr[index],arr[len(arr)-1-index]=arr[len(arr)-1-index],arr[index]//因为len(arr)-1-index，index初始值是0,又因为在[],也是下标的意思
	}
}
//[]里面,奇数偶数都可以