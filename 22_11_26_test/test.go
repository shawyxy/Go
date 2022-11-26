//package main
//
//import "fmt"
//
//func Adder(x int) func(int) int {
//	sum := 0
//	return func(x int) int {
//		sum += x
//		return sum
//	}
//}
//func main() {
//	a := 6
//	for i := 0; i < a; i++ {
//		fmt.Println(Adder(a))
//	}
//	fmt.Println(a)
//}
package main

import "fmt"

func foo(x int) int {
	ret := 0
	for i := 0; i < x; i++ {
		ret += i
	}
	return ret
}

func main() {
	n := 100
	fmt.Println(foo(n))
}