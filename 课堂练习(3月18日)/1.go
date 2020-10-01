package main

import "fmt"

func main() {
    // a :="xiwanzijiyishebkuaikuailele"
	// str:=""
	// m:=make(map[string]int)
	//for i:=0;i<len(a) ;i++  {
	//	str=a[i:i+1]
	//	if m[str]!=0 {
	//		continue
	//	}
	//	//m[str]=1
	//	for j:=0;j<len(a) ; j++ {
	//		if str == a[j:j+1]{//由于str已经把每个不同的字母打印一遍,a[j:j+1]是从头开始找相同的字母，然后，找到了就++
    //            m[str]++
	//		}
	//	}
	//	fmt.Println(m)
	//}
	////fmt.Println(m)
     //第二种,略有不同,方便程序更快的执行*
	a :="xiwanzijiyishebkuaikuailele"
	str:=""
	m:=make(map[string]int)
	for i:=0;i<len(a) ;i++  {//把不同的字母都打一遍
		str=a[i:i+1]
		if m[str]!=0 {//重复的跳过
			continue
		}
		m[str]=1//每个字母出现记录为1
		for j:=i+1;j<len(a) ; j++ {
			if str == a[j:j+1]{//由于str已经把每个不同的字母打印一遍,a[j:j+1]是从头开始找相同的字母，然后，找到了就++
				m[str]++
			}
		}
		fmt.Println(m)
	}
	//fmt.Println(m)
}
