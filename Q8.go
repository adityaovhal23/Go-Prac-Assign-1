/*WAP in go language to illustrate pointer to pointer concept. */

package main

import "fmt"

func main() {
	var a int = 10
	var ptr *int = &a
	var ptrPtr **int = &ptr

	fmt.Println("Value of a:", a)
	fmt.Println("Value of ptr:", *ptr)
	fmt.Println("Value of ptrPtr:", **ptrPtr)

	// Change the value of a using ptrPtr
	**ptrPtr = 20
	fmt.Println("Value of a after change:", a)
}
