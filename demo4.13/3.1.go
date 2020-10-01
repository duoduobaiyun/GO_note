package main

import "fmt"

func main() {

    //作业三：使用接口定义汽车的标准：① 汽车品牌  ② 汽车能够驱动 并定义结构体，
	//实现卡车，电动汽车、三轮车三种定义，实现汽车接口标准。
	//分别实例化解放牌卡车，特斯拉电动车，时风三轮车，并调用对应的驱动方法。


	//1.创建Mouse类型
	m1:=truck{"汽车品牌:解放牌卡车"}
	fmt.Println(m1.name)
	//草稿2.创建FlashDisk
	f1:=electromobile{"汽车品牌:特斯拉电动车"}
	fmt.Println(f1.name)

	g1:=tricycle{"汽车品牌:时风三轮车"}

	testInterface(m1)
	testInterface(f1)
	testInterface(g1)


}
//1.定义接口
type Car interface {
	start()
	way()
}

//草稿2.实现类
type truck  struct {
	name string
}

type electromobile struct {
	name string
}

type tricycle struct {
	name string
}

func (m  truck )start() {
	fmt.Println(m.name,"开始驱动")
}

func (m  truck ) way()  {
	fmt.Println(m.name,"柴油驱动")
}

func (f  electromobile) start() {
	fmt.Println(f.name,"开始驱动")
}

func (f electromobile)way()  {
	fmt.Println(f.name,"电能驱动")
}

func (g tricycle)start()  {
	fmt.Println(g.name,"开始驱动")
}

func (g tricycle) way()  {
	fmt.Println(g.name,"机械驱动")
}

//3.测试方法
func testInterface(car Car)  {
	car.start()
	car.way()
}

