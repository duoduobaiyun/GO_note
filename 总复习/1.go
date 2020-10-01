package main

import "fmt"

func main() {
	//i:=0
	//for ;;i++  {
	//	if i>20 {
	//		break
	//	}
	//	fmt.Printf("%d\n",i)
	//}

	//var i  int
	//for   {
	//	if (i>10) {
	//		break
	//	}
	//	fmt.Println(i)
	//	i++
	//}

	//str:="123ABCabc——丁丐"
	//for i,val:=range str {
	//	fmt.Printf("第%d位的ASCLL值=%d,字符是%c\n",i,val,val)
	//}

	//遍历切片中元素
	//arr:=[]int{100,200,300}
	//for i,val:=range arr {
	//	fmt.Println(i,":",val)
	//}

	//求1-100的和
	//sun:=0
	//for i:=1;i<100 ; i++ {
	//	sun+=i
	//	fmt.Println(sun)
	//}
	//fmt.Println(sun)

	//求1-100之间3的倍数的和
	//i:=1
	//sum:=0
	//for i<=100  {
	//	i++
	//	if i%3==0 {
	//		sum+=i
	//		fmt.Print(i)
	//		if i<99 {
	//			fmt.Print("+")
	//		}else {
	//			fmt.Printf("=%d\n",sum)
	//		}
	//	}
	//	//i++
	//}

	//3、截⽵竹竿。32⽶米⽵竹竿，每次截1.5⽶米，最快截⼏次之后能小于4米？
	//count:=0
	//for i:=32;i>4 ; i-=1.5{
    // count++
	//}
	//fmt.Println(count)

	var a  int
	fmt.Println("请输入数字")
	fmt.Scanln(&a)

	//打印左下直角三角形
	for index:=0;index<a ; index++ {//行
		for j:=0;j<index ;j++  {//列
			fmt.Print("❤"+" ")
		}
		fmt.Println()
	}


	//打印矩形
	// b :=0
	//fmt.Println("请输入数字")
	//fmt.Scanln(&b)
	//for i:=0; i<b;i++  {//行
	//	for j:=0;j<b ; j++ {//列
	//		fmt.Print("❤"+" ")
	//	}
	//	fmt.Println()
	//}
	//fmt.Println()

	//打印左上直角三角形
	//c :=0
	//fmt.Println("请输入数字")
	//fmt.Scanln(&c)
	//for i:=0;i<c ; i++ {//行
	//	for j:=0 ;j<c-i ; j++ {
    //       fmt.Print("❤")
	//	}
	//	fmt.Println()
	//}
}
