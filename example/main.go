package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/suntoucha/kassaclient"
)

func main() {
	// Универсальный клиент для всех игр
	cli := kassaclient.KassaClient{Token: "hellohello", BaseUrl: "https://api.kassa.games"}
	fmt.Printf("Universal client: [%#v]\n", cli)

	fmt.Printf("\n=== STEAM TEST ===\n")
	// Валидируем Steam аккаунт
	valid, serverId, serverName, err := cli.Validate("steam", "world_dragon_508")
	fmt.Printf("Steam validate results: valid - %v, serverId - %v, serverName - %v, error: %v\n", valid, serverId, serverName, err)

	// Создаем корзину для Steam
	uu, err := uuid.NewUUID()
	id := "no-id"
	if err == nil {
		id = uu.String()
	}

	cart, err := cli.Cart(id, "steam", "world_dragon_508", 2, "https://croc.kassa.games/callback")
	fmt.Printf("Steam cart result: %#v, error: %v\n", cart, err)

	fmt.Printf("\n=== GENSHIN IMPACT TEST ===\n")
	// Валидируем Genshin аккаунт
	validGenshin, serverId, serverName, errGenshin := cli.Validate("genshin", "783437191")
	fmt.Printf("Genshin validate results: valid - %v, serverId - %v, serverName - %v, error: %v\n", validGenshin, serverId, serverName, errGenshin)

	// Получаем продукты для Genshin
	productsGenshin, errGenshin2 := cli.Products("genshin")
	fmt.Printf("Genshin products result: %#v, error: %v\n", productsGenshin, errGenshin2)

	fmt.Printf("\n=== HONKAI STAR RAIL TEST ===\n")
	// Валидируем HSR аккаунт
	validHSR, serverId, serverName, errHSR := cli.Validate("honkai-star-rail", "722354753")
	fmt.Printf("HSR validate result: %v, error: %v\n", validHSR, errHSR)

	// Получаем продукты для HSR
	productsHSR, errHSR2 := cli.Products("honkai-star-rail")
	fmt.Printf("HSR products results: valid - %#v, serverId - %v, serverName - %v, error: %v\n", productsHSR, serverId, serverName, errHSR2)

	fmt.Printf("\n=== ZENLESS ZONE ZERO TEST ===\n")
	// Валидируем ZZZ аккаунт
	validZZZ, serverId, serverName, errZZZ := cli.Validate("zenless-zone-zero", "1506168129")
	fmt.Printf("ZZZ validate results: valid - %v, serverId - %v, serverName - %v, error: %v\n", validZZZ, serverId, serverName, errZZZ)

	// Получаем продукты для ZZZ
	productsZZZ, errZZZ2 := cli.Products("zenless-zone-zero")
	fmt.Printf("ZZZ products result: %#v, error: %v\n", productsZZZ, errZZZ2)

	http.HandleFunc("/callback", hello)
	fmt.Println("Listen and serve: ", http.ListenAndServe(":9999", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	cli := kassaclient.KassaClient{}

	cb, err := cli.ParseCallback(r)
	fmt.Printf("Callback: %#v, error: %v\n", cb, err)
}
