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
	//cli := kassaclient.KassaClient{Token: "hellohello", BaseUrl: "http://localhost:10987"}
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
	productGenshin, errGenshin2 := cli.Product("genshin", serverId)
	fmt.Printf("Genshin product result: %#v, error: %v\n", productGenshin, errGenshin2)

	// Создаем корзину для Genshin с server_id
	if validGenshin && serverId != "" {
		uu2, err := uuid.NewUUID()
		genshinCartId := "no-id-genshin"
		if err == nil {
			genshinCartId = uu2.String()
		}
		cartGenshin, errCart := cli.Cart(genshinCartId, "genshin", "783437191", 0.77, "https://croc.kassa.games/callback", serverId)
		fmt.Printf("Genshin cart result: %#v, error: %v\n", cartGenshin, errCart)
	}

	fmt.Printf("\n=== HONKAI STAR RAIL TEST ===\n")
	// Валидируем HSR аккаунт
	validHSR, serverId, serverName, errHSR := cli.Validate("honkai-star-rail", "722354753")
	fmt.Printf("HSR validate results: valid - %#v, serverId - %v, serverName - %v, error: %v\n", validHSR, serverId, serverName, errHSR)
	// Получаем продукты для HSR
	productHSR, errHSR2 := cli.Product("honkai-star-rail", serverId)
	fmt.Printf("HRS product result: %#v, error: %v\n", productHSR, errHSR2)

	fmt.Printf("\n=== ZENLESS ZONE ZERO TEST ===\n")
	// Валидируем ZZZ аккаунт
	validZZZ, serverId, serverName, errZZZ := cli.Validate("zenless-zone-zero", "1506168129")
	fmt.Printf("ZZZ validate results: valid - %v, serverId - %v, serverName - %v, error: %v\n", validZZZ, serverId, serverName, errZZZ)
	// Получаем продукты для ZZZ
	productZZZ, errZZZ2 := cli.Product("zenless-zone-zero", serverId)
	fmt.Printf("ZZZ product result: %#v, error: %v\n", productZZZ, errZZZ2)

	http.HandleFunc("/callback", hello)
	fmt.Println("Listen and serve: ", http.ListenAndServe(":9999", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	cli := kassaclient.KassaClient{}

	cb, err := cli.ParseCallback(r)
	fmt.Printf("Callback: %#v, error: %v\n", cb, err)
}
