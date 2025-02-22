/*WAP in go Language to print address of a variable*/

package main

import "fmt"

func main() {
	var num int = 10
	fmt.Printf("Address of num variable: %p\n", &num)
}
