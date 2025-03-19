package main

import "fmt"

func main() {

	fmt.Println(sayHello("Golang"))

}

// This function accepts a string input and returns a string value
func sayHello(msg string) string {
	return "Hello " + msg + " !"
}
