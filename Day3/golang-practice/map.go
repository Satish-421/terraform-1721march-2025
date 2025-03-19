package main

import "fmt"

func main() {

	//Declares a map with string type key and string type value
	//key and values can be of different types as well
	toolsPath := map[string]string{
		"java_home": "/usr/lib/jdk-11",
		"mvn_home":  "/usr/share/maven",
	}

	//Given a key, map will be able to retrieve the value
	fmt.Println("Java Home Directory : ", toolsPath["java_home"])

	//add a key,value pair into an existing map
	toolsPath["go_home"] = "/usr/go"

	//iterating a map and printing its values
	for key, value := range toolsPath {
		fmt.Println(key, value)
	}

	//delete a key-value pair from an existing map
	delete(toolsPath, "go_home")
	fmt.Println(toolsPath)

}
