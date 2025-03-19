package main

import "fmt"

func main() {

	count := 5 //Declares a count variable of type int and assigns a value 5

	for count > 0 {

		fmt.Println(count)
		count-- //equivalent to count = count - 1
		//--count pre-decrement is not supported in golang
		//++count pre-increment is not supported in golang
		//count++ is supported in golang
	}

	fmt.Println("Value of count is", count, " after for loop")

	count = 0 //Variable is already declared in line no 7, we are reinitializing count with 0

	for count = 1; count < 10; count++ {
		fmt.Printf("%d\t", count)
	}

	fmt.Println()
}
