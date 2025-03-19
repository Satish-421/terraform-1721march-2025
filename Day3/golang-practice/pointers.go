package main

import "fmt"

// Pass by pointer
func sayHello(msgPtr *string) {
	tmp := *msgPtr //The value stored at the address pointed by msgPtr is assigned to tmp string

	//The value stored at address pointed by msgPtr we are changing to "Hello Golang !"
	*msgPtr = tmp + " Golang " + "!"
}

func main() {

	msg := "Hello"

	//The below line will print "Hello"
	fmt.Println("Message before calling sayHello function is ", msg)

	//sayHello function is taking the address of msg string
	sayHello(&msg)

	//The below line will print "Hello Golang !"
	fmt.Println("Message after calling sayHello function is ", msg)

}
