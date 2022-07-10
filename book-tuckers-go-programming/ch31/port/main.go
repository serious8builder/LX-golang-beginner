package main

import (
	"fmt"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "6001"
	}
	fmt.Printf("%T, %v", port, port)
}
