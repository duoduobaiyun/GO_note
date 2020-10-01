package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    a :=[]string{"W","U","H","A","N","J","I","A","Y","O","U","Z","H","O","N","G","G","U","O","J","I","A","Y","O","U","I","L","O","V","E","C","H","I","N","A"}
    count:=0
    rand.Seed(time.Now().UnixNano())
    a[0]="W"
    a[1]="U"
    a[2]="H"
    a[3]="A"
    a[4]="N"
    a[5]="J"
    a[6]="I"
	a[7]="Y"
	a[8]="O"
	a[9]="G"
	a[10]="L"
	a[11]="E"
	a[12]="Z"
	a[13]="V"
	a[14]="C"
	for i:=0;i<len(a) ;i++  {
		n:=a[i+1]
		if a[i] == n {
			count++
		}
	a[i]= n
	}
	fmt.Println(a)

}
