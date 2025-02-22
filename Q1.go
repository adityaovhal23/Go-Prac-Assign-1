/*WAP in go language to print Student name, rollno, division and college name */

package main

import "fmt"

type Student struct {
	Name     string
	RollNo   int
	Division string
	College  string
}

func main() {
	student := Student{
		Name:     "John Doe",
		RollNo:   123,
		Division: "A",
		College:  "XYZ University",
	}

	fmt.Printf("Student Name: %s\n", student.Name)
	fmt.Printf("Roll No: %d\n", student.RollNo)
	fmt.Printf("Division: %s\n", student.Division)
	fmt.Printf("College Name: %s\n", student.College)
}
