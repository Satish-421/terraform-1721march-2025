package main

import "fmt"

// User-defined data type
type Rectangle struct {
	length int
	width  int
}

func (rect Rectangle) Area() int {
	area := rect.length * rect.width
	return area
}

func main() {
	rectangle := Rectangle{
		length: 100,
		width:  200,
	}

	fmt.Printf("Area of rectange : %d\n", rectangle.Area())
}
