/*WAP in go language to print Fibonacci series of n terms. */

package main

import "fmt"

func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}

func main() {
	var n int
	fmt.Print("Enter the number of terms: ")
	fmt.Scan(&n)
	fmt.Println("Fibonacci series of", n, "terms:")
	fibonacci(n)
}
