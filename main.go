package main

import (
	"fmt"
	funcs "github.com/BelehovEgor/go-fuzz-targets/examples"
)

func main() {
	// Calling the generateJSON function with some sample arguments
	jsonString, err := funcs.GenerateJSON("name", "Alice", "age", 25, "isStudent", true)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("JSON String:", jsonString)
	}
}
