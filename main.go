package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/fermar7/certainer/acme/infra"
)

func main() {

	config, err := initConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)

	authority, err := url.Parse("https://" + config.Authority)
	if err != nil {
		log.Fatal(err)
	}

	acmeDirectory, err := infra.CreateDirectoryProvider(authority)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(acmeDirectory)

}
