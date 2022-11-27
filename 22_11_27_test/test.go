// package main
//
// import (
//
//	"fmt"
//	"math/rand"
//	"time"
//
// )
//
//	func getRandNum(max, min int) int {
//		rand.Seed(time.Now().UnixNano())
//		return rand.Intn(min-max+1) + max
//	}
//
//	func main() {
//		count := 0
//		randNum := getRandNum(100, 10)
//		for {
//			count++
//			num := 0
//
//			fmt.Scanln(&num)
//			if num < randNum {
//				fmt.Println("small ")
//			}
//		}
//	}
//package main

//	func Print(a []int) {
//		for _, value := range a {
//			value++
//			fmt.Println(value)
//		}
//	}
//
//	func main() {
//		a := []int{1, 2, 3}
//		//for i := 0; i < len(a); i++ {
//		//	fmt.Println(a[i])
//		//}
//		//for index, value := range a {
//		//	fmt.Println("index = ", index, ", value = ", value)
//		//}
//		Print(a)
//
//		for _, value := range a {
//			value++
//			fmt.Println(value)
//		}
//	}
//package main
//
//import "fmt"
//
//type Animals interface {
//	Eat(food string)
//	Run()
//	Sleep()
//}
//
//type Dog struct {
//	name string
//}
//
//func (d Dog) Eat(food string) {
//	fmt.Println("狗吃:" + food)
//}
//func (d Dog) Run() {
//	fmt.Println("狗跑")
//}
//func (d Dog) Sleep() {
//	fmt.Println("狗睡")
//}
//func main() {
//	var animal Animals
//
//	animal = Dog{"大黄"}
//	animal.Run()
//	animal.Eat("骨头")
//	animal.Sleep()
//}
//
//package main
//
//import "fmt"
//
//type Person interface {
//	Say(str string)
//	Look(str string)
//	Dream(other string)
//}
//
//type XY struct {
//	Name string
//	Age  int
//}
//
//func (xy XY) Say(str string) {
//	fmt.Println("xy在说:", str)
//}
//
//func (xy XY) Look(str string) {
//	fmt.Println("xy在看:", str)
//}
//func (xy XY) Dream(str string) {
//	fmt.Println("xy的梦想是:", str)
//}
//
//func main() {
//	var person Person
//	person = &XY{"xy", 21}
//	person.Say("hello world")
//	person.Look("sunshine")
//	person.Dream("Live for himself")

package main

import "fmt"

func deferfun() int {
	fmt.Println("deferfun")
	return 0
}
func fun() int {
	fmt.Println("fun")
	return 0
}
func returnfun() int {
	defer deferfun()
	return fun()
}
func main() {
	returnfun()
}
