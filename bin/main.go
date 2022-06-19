package main

import (
	"fmt"
	"os"
)

var SECRET_NAME string = os.Getenv("SECRET_NAME")
var SECRET_VALUE string = os.Getenv("SECRET_VALUE")

func main() {
	fmt.Printf("name: %s\nvalue: %s\n", SECRET_NAME, SECRET_VALUE)
}
