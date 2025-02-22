/*WAP in go language to swap the number without temporary variable.*/

package main

import "fmt"

func swapNumbers(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	num1 := 10
	num2 := 20

	fmt.Println("Before swapping:")
	fmt.Println("Number 1:", num1)
	fmt.Println("Number 2:", num2)

	swapNumbers(&num1, &num2)

	fmt.Println("After swapping:")
	fmt.Println("Number 1:", num1)
	fmt.Println("Number 2:", num2)
}
