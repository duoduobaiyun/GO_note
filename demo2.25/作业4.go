package main

import "fmt"

func main() {
	const (
		A = iota
		B =iota
	)
	fmt.Println(A,B)

	const (
		Pi float64  = 3.141596257
		Ai	float32 = 3.01
	)
   fmt.Println(Pi)
	fmt.Println(Ai)

	const (
		size  int8   = 121
		eof   int16  = 32766
		guf   int32  = 2147483647
		dof   int64  = 9223372036854775806
	)
	fmt.Println(size)
	fmt.Println(eof)
	fmt.Println(guf)
	fmt.Println(dof)

	const (
	 a  bool = true

	b = false

	c = ( 1 == 2 )

	)
	fmt.Printf("a的类型是:%T,结果是:%t,b的类型是:%T,结果是:%t,c的类型是:%T,结果是:%t",a,a,b,b,c,c)

}
