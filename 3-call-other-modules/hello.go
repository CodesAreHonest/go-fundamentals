package main

import (
	"fmt"

	"fundamentals/CreateModule"
)

func main() {
    // Get a greeting message and print it.
    message := CreateModule.Hello("Gladys")
    fmt.Println(message)
}