package main

import "fmt"

func main() {
	fmt.Printf("作业一\n")
	var str string = "WUHANJIAYOUZHONGGUOJIAYOUILOVECHINA"
	fmt.Printf("%s\n",str)
	arr :=make(map[string]int)
	var a string = ""
	for i:=0;i<len(str);i++{
		a = str[i:i+1]
		if arr[a]!=0 {
			continue
		}
		for j:=0;j<len(str);j++{
			if a == str[j:j+1]{
				arr[a]++
			}
		}
	}
	fmt.Println(arr)
}
