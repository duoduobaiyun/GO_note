package main

import "fmt"

func main()  {
	 q:=5
	 fmt.Println(q)


	 e:=10
	 fmt.Println(e)


	 name:="刘坤"
	 fmt.Println(name)


	 adress:="江西省上饶市潘阳县"
     fmt.Println(adress)

	var a bool
	a = true

	b := false

	c := ( 1 == 2 )

	d := (4==4)

	fmt.Printf("a是 :%T,%t\nb是:%T,%t\nc是:%T,%t\nd是:%T,%t\n",a,a,b,b,c,c,d,d)

	var f1 float32
	f1 = 3.1
	fmt.Printf("f1的type is %T\n", f1)
	fmt.Printf("f1 = %3.2f\n", f1)

	var f2 float64
	f2 = 3
	fmt.Printf("f2的type is %T\n", f2)
	fmt.Printf("f2 = %3.2f", f2)

}
