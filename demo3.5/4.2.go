package main

import "fmt"

func main() {
	var num int
	fmt.Println("请输入数字")
	fmt.Scanln(&num)
	for i := 1; i <= num;  i++  {//行 ,但是有种特殊的用法，它可以控制图型的上下的位置， i 值越大，越往上，i 值越小，越往下
		//fmt.Print(i)
		for j := 0; j < num - i; j++ {//列, 但是有种特殊的用法，它可以控制图型的左右的位置, j 值越大, 越往左，j值越小 ，越往右
			fmt.Print(" ")
		}
		for k := 0; k < 2 * i - 1 ; k++ {//这个k值越大，就会把图形的最上面吞噬掉，然后就是整个图像往下退，然后图像又会重新出现，最后彻底消失
			fmt.Print("*")          //假设输入值为5
			                            //草稿2 * ( 1, 草稿2 , 3 ,4 , 5 ) - 1  都减上1
	 	}                               // k < 草稿2 , k < 4 , k < 6 , k < 8 , k < 10 你画个图就知道了
		fmt.Println()
	}
	  for i := 1; i < num ; i++ {

		for j := 0; j < i  ; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2 * num  - 2 * i - 1 ; k++ { //假设输入值为5
			fmt.Print("*")                      //  10  -  草稿2 * ( 1, 草稿2 , 3 ,4 , 5 ) - 1  都减上1
		}                                           //   10 - 1 = 9 , 10 - 3 = 7 , 10 - 5 = 5 , 10 - 7 = 3 , 10 - 9 = 1 你自己画图就知道了
		fmt.Println()
	}
}
