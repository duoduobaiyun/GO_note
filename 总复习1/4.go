package main

import "fmt"

func main() {
 a:=[6]int{9,5,2,10,3,6}
 fmt.Println(reverse(a),len(a))

 for key,val:=range a {
  fmt.Printf("%d,%d\n",key,val)
 }
}

func reverse(arr [6]int)  [6]int{

   for index:=0; index<len(arr)/2;index++  {
   arr[index],arr[len(arr)-1-index]=arr[len(arr)-1-index],arr[index]//因为len(arr)-1-index，index初始值是0,又因为在[],也是下标的意思
 }
 return arr
}
//[]里面,奇数偶数都可以