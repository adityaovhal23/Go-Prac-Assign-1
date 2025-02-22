/* WAP in go language to explain new function.*/

package main

import "fmt"

func main() {
	// declaring a function
	add := func(x int, y int) int {
		return x + y
	}

	// calling the function
	result := add(5, 6)
	fmt.Println("The result is: ", result)
}
