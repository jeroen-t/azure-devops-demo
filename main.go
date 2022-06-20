package main

import (
	"fmt"
	"log"

	"github.com/jeroen-t/azure-devops-demo/util"
)

// var SECRET_NAME string = os.Getenv("MY_SECRET_NAME")
// var SECRET_VALUE string = os.Getenv("MY_SECRET_VALUE")

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	fmt.Printf("name: %s\nvalue: %s\n", config.SECRET_NAME, config.SECRET_VALUE)
}
