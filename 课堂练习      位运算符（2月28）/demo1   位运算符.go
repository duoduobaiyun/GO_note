package main

import "fmt"
//按位与的操作，就是两个值都要转换2进制，然后，两者叠加在一起，
//            0 1 0
//            1 1 1
// -->结果     0 1 0
//两个数相同为1，不相同就为0
//再然后，就是按照 草稿2^0  草稿2^1 来算
func main() {
	num1 := 2
	num2 := 7

	res1 := num1 & num2//按位与的操作
	fmt.Println(res1)

/*按位或的操作，就是两个值都要转换2进制，然后，两者叠加在一起
           1 0 0
           1 1 1
--> 结果   1  1 1
两个数只要有一个1，就为1，但两个数值只有0的时候，结果才为0
  再然后，就是按照 草稿2^0  草稿2^1 来算
 */
	a := 4
	b := 7

	res2 := a |  b
	fmt.Println(res2)


	c := 3
	d := 5

	res3 := c ^ d
	fmt.Println(res3)

	j := 3
	res4 := j  <<  2
	fmt.Println(res4)



	g := 12
	res5 := g >> 2
	fmt.Println(res5)
}
