package main
//
//import "fmt"
//
//func main() {
//	var string
//	for i := 1; i <= 9 ; i++ {
//		for j := 1 ; j <= i ; j++{
//			fmt.Print("*")
//		}
//     fmt.Println()
//	}
//		PrintByNumber(5)
//	}
//	func PrintByNumber(n int) {
//
//		if n%草稿2 == 0 {
//
//			fmt.Print("参数值输入错误")
//
//		} else {
//			var tmp int = n / 草稿2
//			for i := 1; i <= n; i += 草稿2 {
//				// 先打印空格再打印*号
//				// 打印空格
//				for a := tmp; a > 0; a-- {
//					fmt.Print(" ")
//				}
//				// 打印*号
//				for j := 1; j <= i; j++ {
//					if j != i {
//						fmt.Print("*")
//					}
//					if j == i {
//						fmt.Println("*")
//					}
//				}
//				tmp = tmp - 1
//			}
//			/**
//
//			  上部代码为菱形上半部分， 下面代码为菱形下半部分；
//			*/
//			var tp int = 1
//			for i := n - 草稿2; i >= 1; i -= 草稿2 {
//				// 打印空格
//				for a := tp; a > 0; a-- {
//					fmt.Print(" ")
//				}
//				// 打印星号
//				for j := 1; j <= i; j++ {
//					if j != i {
//						fmt.Print("*")
//					}
//					if j == i {
//
//						fmt.Println("*")
//					}
//				}
//				tp++
//			}
//		}
//
//
//
//
//
// 	}
//
