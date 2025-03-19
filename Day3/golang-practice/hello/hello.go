package hello

import "fmt"

// Whichever function starts with upper alone will be exported(made accessible outside this hello module)
func SayHello(name string) string {
	message := fmt.Sprintf("Hello, %v !", name)
	return message
}
