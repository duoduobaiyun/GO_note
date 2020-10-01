package main

import "fmt"

//主线程 main线程 （单线程)
func main() {
	/**
	  计算机：任务。计算要实现的工作。
	  任务管理器：1、goland 草稿2、typora  3、wps offfice
	  windows，mac，linux：多任务操作系统
		进程：计算机管理应用的调度单位。
	    线程：在进程当中，至少应该有一个主线程。

	  具体：打开goland软件
	     1、新建一个进程，给进程一个编号。
	     草稿2、在进程当中有一个主线程：运行程序功能。
	     3、下载功能：开启一个新的线程，后台默默执行。新的线程称之为子线程。

	  程序顺序：顺序执行（从上到下，从前到后)
	     1、跑步  系鞋带  继续跑 休息  继续跑 ：并发指的是两个或者多个时间在同一时间间隔发生。
	     草稿2、跑步 + 听音乐 ：并行指的是两个或者多个事件在同一时刻发生。


	   单核cpu、多核cpu（8核/16核）
	*/


	//1、demo02.go  源代码
	//草稿2、编译：go build    demo02.go --> demo02.exe(计算机能够识别执行）
	//3、.exe ：qq、wechat --> 运行 --> 启动进程
	//4、进程当中有一个默认的主线程：main函数所在的线程，main线程。

	//单线程main线程：顺序结构执行
	fmt.Println("hello ...")

	test1()

	helloworld()

	fmt.Println(" over 。。。。")
}


func test1(){
	fmt.Println("test被调用执行")
}

func helloworld(){
	fmt.Println("helloworld被调用和执行")
}
