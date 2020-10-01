//package main
//
//import "fmt"
//
//func main() {
//	//作业三：使用接口定义汽车的标准：① 汽车品牌  ② 汽车能够驱动 并定义结构体，实现卡车，电动汽车、三轮车三种定义，
//	//实现汽车接口标准。分别实例化解放牌卡车，特斯拉电动车，时风三轮车，并调用对应的驱动方法。
//
//	m:=truck{"汽车品牌:解放牌卡车"}
//	fmt.Println(m)
//
//	n:=electromobile{"汽车品牌:特斯拉电动车"}
//	fmt.Println(n)
//
//	l:=tricycle{"汽车品牌:时风三轮车"}
//	fmt.Println(l)
//
//
//	testInterface(m)
//
//}
//
//type Car interface {
//	way()
//	stat()
//}
//
//type truck struct {
//	name string
//
//}
//
//type  electromobile struct {
//	name string
//}
//
//type tricycle struct {
//	name string
//}
//
//func (a truck) way()  {
//	fmt.Println("柴油发动")
//}
//
//func (a truck) start()  {
//   fmt.Println("开始驱动")
//}
//
//func (b electromobile) way()  {
//	fmt.Println("电能驱动")
//}
//func (b electromobile)start()  {
//    fmt.Println("开始驱动")
//}
//
//func (c tricycle) way()  {
//   fmt.Println("机械驱动")
//}
//func (c tricycle) start()  {
//  fmt.Println("开始驱动")
//}
//
//
//func testInterface(car Car)  {
//	Car.stat()
//	Car.way()
//}