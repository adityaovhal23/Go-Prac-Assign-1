/*WAP in go to print table of given number.*/

package main

import "fmt"

func printTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)
	fmt.Printf("Table of %d:\n", num)
	printTable(num)
}
