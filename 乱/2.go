package main

import "fmt"

func main() {
	/*fmt.Println("请输入一个整数")
	 var num1 int
	 fmt.Scanln(&num1)
	 fmt.Println("请输入一个整数")
	 var num2 int
	 fmt.Scanln(&num2)
	 fmt.Println("请输入一个操作:+,-,*,/,++,--")
	 var sym  string
	 fmt.Scanln(&sym)

	 switch sym {
	 case "+" :
     fmt.Printf("%d + %d = %d\n",num1,num2,num1 + num2)
	 case "-":
	 fmt.Printf("%d - %d = %d\n",num1,num2,num1 - num2)
	 case "*":
	 fmt.Printf("%d * %d = %d\n",num1,num2,num1 * num2)
	 case "/":
	 fmt.Printf("%d / %d = %d\n",num1,num2,num1 / num2)
	 case "++":
	 num1 ++
	 num2 ++
	 fmt.Println("num1和num2的自增值是:",num1,num2)
	 case "--":
	 num1 --
	 num2 --
	 fmt.Println("num1和num2的自减值是:",num1,num2)


	  }*/

	num3 := 0
	num4 := 0
	sun  := ""
	fmt.Println("请输入一个整数")
	fmt.Scanln(&num3)
	fmt.Println("请在输入一次整数")
	fmt.Scanln(&num4)
	fmt.Println("请输入一个操作:+,-,*,/")
	fmt.Scanln(&sun)
	switch sun {
	case "+" :
		fmt.Printf("%d + %d = %d\n",num3,num4,num3 + num4)
	case "-":
		fmt.Printf("%d - %d = %d\n",num3,num4,num3 - num4)
	case "*":
		fmt.Printf("%d * %d = %d\n",num3,num4,num3 * num4)
	case "/":
		fmt.Printf("%d / %d = %d\n",num3,num4,num3 / num4)
	}

}
