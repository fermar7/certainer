package main

import (
	"fmt"
	"log"
)

func main() {

	config, err := initConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)

}
