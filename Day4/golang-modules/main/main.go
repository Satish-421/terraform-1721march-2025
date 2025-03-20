package main

import (
    "fmt"
    "addition"
    "subtraction"
)

func main() {

	x := float32(500.7)
	y := float32(200.5)

	result1  := addition.Add(x,y)
	result2 := subtraction.Subtract(x,y)

	fmt.Println ( "The sum of ", x, " and ", y, " is ", result1 )
	fmt.Println ( "The difference of ", x, " and ", y, " is ", result2 )

}
