package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/suntoucha/kassaclient"
	"net/http"
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

	cart, err := cli.SteamCart(id, "sergekuzmenko", 2, "https://croc.kassa.games/callback")
	fmt.Printf("Cart result: %#v, error: %v\n", cart, err)

	http.HandleFunc("/callback", hello)
	fmt.Println("Listen and serve: ", http.ListenAndServe(":9999", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	cli := kassaclient.KassaClient{}

	cb, err := cli.SteamCallback(r)
	fmt.Printf("Callback: %#v, error: %v\n", cb, err)
}
