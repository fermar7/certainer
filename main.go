package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/fermar7/certainer/acme/acmecont"

	"github.com/fermar7/certainer/acme/acmereq"
	"github.com/fermar7/certainer/acme/infra"
)

func main() {

	config, err := initConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(config)

	authority, err := url.Parse("https://" + config.Authority)
	if err != nil {
		log.Fatal(err)
	}

	acmeDirectory, err := infra.CreateDirectoryProvider(authority)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(acmeDirectory)

	accountKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	newAcctRequest, err := acmereq.CreateAccountCreationRequest(acmeDirectory, accountKey, []string{"mailto:klaus@abgfjireg.com"})
	if err != nil {
		log.Fatal(err)
	}

	newAcctResponse, err := acmereq.Execute(newAcctRequest)
	if err != nil {
		log.Fatal(err)
	}

	var acc acmecont.Account

	err = json.Unmarshal(newAcctResponse.Content, &acc)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(newAcctResponse.Header.Get("Location"))
	log.Println(acc)

	log.Println("--------------------------------------------------")

	kid := fmt.Sprintf("%s/acme/acct/%d", authority, acc.ID)

	log.Println(kid)

	newOrderRequest, err := acmereq.CreateNewOrderRequest(acmeDirectory, accountKey, kid, []string{"abgfjireg.com", "api.abgfjireg.com", "sales.abgfjireg.com"})
	if err != nil {
		log.Fatal(err)
	}

	newOrderResponse, err := acmereq.Execute(newOrderRequest)
	if err != nil {
		log.Fatal(err)
	}

	var order acmecont.Order

	err = json.Unmarshal(newOrderResponse.Content, &order)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(newOrderResponse.Header.Get("Location"))
	log.Println(order)

}
