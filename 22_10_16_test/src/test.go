//package main
//
//import "fmt"
//
//	func main() {
//		//声明一个变量
//		var a int
//		var c = 100 //初始化省去类型
//		var str = "hello world"
//
//		d := 100
//		e := "hello"
//
//		fmt.Println(a)
//		fmt.Println("%T", a)
//		fmt.Println(c)
//		fmt.Println(str)
//
//		fmt.Println(d)
//		fmt.Println("%T", d)
//
//		fmt.Println(e)
//		fmt.Println("%T", e)
//
//		str1 := "Hello world"
//		fmt.Printf("%T", str1)
//
//}
//package main

//	func func1(a int, b int) int {
//		fmt.Println(a)
//		fmt.Println(b)
//
//		return a + b
//	}
//
// // 多返回值
// // 匿名返回值
//
//	func func2(a int, b int) (int, int) {
//		return a + b, a - b
//	}
//
// // 有形参名的返回值
//
//	func func3(a int, b int) (r1 int, r2 int) {
//		r1 = a * b
//		r2 = a / b
//		return r1, r2
//	}
//
//	func main() {
//		a := func1(1, 2)
//		b, c := func2(2, 3)
//
//		fmt.Println(a)
//		fmt.Println(b)
//		fmt.Println(c)
//
// }
// ------------------------------------
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func Area(radius float64) float64 {
//	area := math.Pi * radius * radius
//	return area
//}
//func main() {
//	var radius float64
//	fmt.Print("请输入圆的半径：")
//	fmt.Scanf("%v", &radius)
//	fmt.Println("圆的面积为：", Area(radius))
//}

// -------------------------------------
//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func makeSuffix(suffix string) func(string) string {
//	return func(name string) string {
//		if !strings.HasSuffix(name, suffix) {
//			return name + suffix
//		}
//		return name
//	}
//}
//
//func main() {
//	f := makeSuffix(".jpg")
//	fmt.Println("文件名处理后 = ", f("hello"))
//	fmt.Println("文件名处理后 = ", f("world.jpg"))
//	fmt.Println("文件名处理后 = ", f("hiiii.avi"))
//}

// -------------------------------------
package main

import (
	"fmt"
)

func Swap(num1 *float64, num2 *float64) {
	var tmp float64
	tmp = *num1
	*num1 = *num2
	*num2 = tmp
}
func main() {
	var n1 float64
	var n2 float64
	fmt.Print("请输入要交换的数据：")
	fmt.Scanf("%v %v", &n1, &n2)
	Swap(&n1, &n2)
	fmt.Println("交换结果为：n1 =", n1, ",n2 =", n2)
}

---------------------------------
package main

import "fmt"

type Person struct {
	_name  string
	_id    string
	_class int
	_hobby string
}

func (this *Person) IntroduceMe() {
	fmt.Print("我叫：", this._name)
	fmt.Println("，我的学号是：", this._id)
}
func (this *Person) SetPerson(name string, id string, class int, hobby string) {
	this._name = name
	this._id = id
	this._class = class
	this._hobby = hobby
}
func main() {
	var person1 Person
	person1.SetPerson("张三", "2021304892", 1, "睡觉")
	person1.IntroduceMe()
	person2 := Person{_name: "李四", _id: "2021304893", _class: 2, _hobby: "吃饭"}
	person2.IntroduceMe()
}

//package main
//
//import "fmt"
//
//type Animal interface {
//	Eat(food string)
//	Run()
//	Bite(name string)
//}
//
//type Chicken struct {
//	Name string
//}
//
//func (c *Chicken) Eat(food string) {
//	fmt.Println("鸡吃" + food)
//}
//
//func (c *Chicken) Run() {
//	fmt.Println("鸡在跑")
//}
//
//func (c *Chicken) Bite(name string) {
//	fmt.Println()
//}
//
//type Dog struct {
//	Name string
//}
//
//func (d *Dog) Eat(food string) {
//	fmt.Println("狗吃" + food)
//}
//
//func (d *Dog) Run() {
//	fmt.Println("狗在跑")
//}
//
//func (d *Dog) Bite(name string) {
//	fmt.Println(d.Name + "咬" + name)
//}
//
//func main() {
//	var chicken Chicken
//	var animal Animal = &chicken
//	animal.Eat("鸡饲料")
//	animal.Run()
//
//	dog1 := Dog{Name: "Jim"}
//	animal = &dog1
//	animal.Eat("骨头")
//	animal.Run()
//
//	dog := Dog{Name: "Tom"}
//	animal = &dog
//	animal.Eat("骨头")
//	animal.Run()
//	animal.Bite("Jim")
//}

//package main
//
//import "fmt"
//
//type data struct {
//	n float64
//	m float64
//}
//
//// 数据交换
//func swapData(n1 float64, n2 float64) data {
//	res := n1 + n2
//	n1 = res - n1
//	n2 = res - n1
//	return data{
//		n1,
//		n2,
//	}
//}
//
//func main() {
//	var n1 float64
//	var n2 float64
//	fmt.Println("请输入要交换的数据：")
//	fmt.Scanf("%v %v", &n1, &n2)
//	fmt.Println("交换结果为：", swapData(n1, n2))
//}
