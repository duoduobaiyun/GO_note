package main

import "fmt"

func main() {
	a :="WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA"
	map1:=make(map[string]int)

	 m  := ""
    //k:=0
	for i:=0;i<len(a) ; i++ {
		m =a[i:i+1]
		fmt.Println(map1[m])
		if map1[m]!= 0{
			continue
		}
		//fmt.Println(a)
		//k++
		//fmt.Println("——————————————————————————————",k)
		for j:=0;j<len(a) ; j++ {
			if m == a[j:j+1] {
				map1[m]++
			}
		}
	}
   fmt.Println(map1)
}
