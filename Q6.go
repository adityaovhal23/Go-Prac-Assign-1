/*WAP in go language to print PASCALS triangle. */

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(combination(i, j), " ")
		}
		fmt.Println()
	}
}

func combination(n, r int) int {
	if r > n-r {
		r = n - r
	}
	res := 1
	for i := 0; i < r; i++ {
		res = res * (n - i) / (i + 1)
	}
	return res
}
