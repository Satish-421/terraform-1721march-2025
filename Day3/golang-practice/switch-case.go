package main

import "fmt"

func main() {

	var direction string //declares a variable named direction of type string

	//Valid values are east, west, north, south
	fmt.Print("Enter some direction :")
	fmt.Scanln(&direction)

	switch direction {
	case "east":
		fmt.Println("You entered direction ", direction)
	case "west":
		fmt.Println("You entered direction ", direction)
	case "north":
		fmt.Println("You entered direction ", direction)
	case "south":
		fmt.Println("You entered direction ", direction)
	default:
		fmt.Println("Invalid direction, possible values are east, west,north,south")
	}
}
