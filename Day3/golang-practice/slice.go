package main

import "fmt"

func main() {

	//shorter syntax to declare and initialize an array of int with 6 values
	//index            0   1   2 	3  4   5
	intArray := [6]int{10, 20, 30, 40, 50, 60}

	fmt.Println("Array elements are ...")
	fmt.Println(intArray)

	//Slice uses an array internally
	//in this case, the slice is referring to an array from index position 2 to 4
	var mySlice []int = intArray[2:5] // 2 is lower bound and 5 is the upper bound index, index 5 is not included

	fmt.Println("Slice elements are ...")
	fmt.Println(mySlice)

	//Let's modify the slice at certain indices
	//When the slice is modified, it allows affects the original array that is pointed by the slice
	//myslice index starts from zero just like arrays
	mySlice[0] = 100
	mySlice[1] = 200
	mySlice[2] = 300

	fmt.Println("Slice elements are ...")
	fmt.Println(mySlice)

	fmt.Println("Array elements are ...")
	fmt.Println(intArray)

	//When the array pointed by slice can't fit new elements, the slice will get disconnected from the array
	//slice creates a new array to fit the old values and the newly appended values
	//hence any changes to the mySlice will no longer affect the intArray
	mySlice = append(mySlice, 400, 500)

	fmt.Println("Array elements are ...")
	fmt.Println(intArray)

	fmt.Println("Slice elements are ...")
	fmt.Println(mySlice)
}
