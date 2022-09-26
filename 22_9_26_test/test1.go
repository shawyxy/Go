//    *(0,4)
//  *   *(1,2)(1,6)
// *     *(2,1)(2,7)
//  *   *(3,2)(3,6)
//    *(4,4)
//虽然但是，上面的图形好像找不出规律
//---------------
//   *
//  * *
// *   *
//  * *
//   *
package main
import "fmt"

func main() {

	var n int = 3
	
	for i := 1; i <= n; i++ {
		for j := 1; j <= n - i; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2 * i - 1; j++ {
			if j == 1 || j == 2 * i - 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n");
	}

	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= n - i; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Print("\n");
	}
}