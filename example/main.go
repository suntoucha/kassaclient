package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/suntoucha/kassaclient"
)

func main() {
	cli := kassaclient.KassaClient{Token: "hellohello", BaseUrl: "https://api.kassa.games"}

	fmt.Printf("Hello, this is Example [%#v]\n", cli)

	valid, err := cli.SteamValidate("sergekuzmenko")
	fmt.Printf("Validate result: %v, error: %v]\n", valid, err)

	uu, err := uuid.NewUUID()
	id := "no-id"
	if err == nil {
		id = uu.String()
	}

	cart, err := cli.SteamCart(id, "sergekuzmenko", 2, "https://ya.ru")
	fmt.Printf("Cart result: %#v, error: %v\n", cart, err)

}
