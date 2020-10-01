package main

import "fmt"

func main() {
	a:="shijieyexuhenmeihaodanhaixuzijiqufaxian"
	m:=make(map[string]int)
	b:=""
	for i:=0;i<len(a) ; i++ {
		b=a[i:i+1]
		if m[b]!=0 {
			continue
		}
		m[b]=1
		for j:=i+1;j<len(a) ; j++ {
			if b==a[j:j+1]{
				m[b]++
			}
			//fmt.Println(m)
		}
		//fmt.Println(m)
	}
	fmt.Println(m)
}
