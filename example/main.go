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
	valid, err := cli.Validate("steam", "world_dragon_508")
	fmt.Printf("Steam validate result: %v, error: %v\n", valid, err)

	// Создаем корзину для Steam
	uu, err := uuid.NewUUID()
	id := "no-id"
	if err == nil {
		id = uu.String()
	}

	cart, err := cli.CreateCart("steam", id, "world_dragon_508", 2, "https://croc.kassa.games/callback")
	fmt.Printf("Steam cart result: %#v, error: %v\n", cart, err)

	fmt.Printf("\n=== GENSHIN IMPACT TEST ===\n")
	// Валидируем Genshin аккаунт
	validGenshin, errGenshin := cli.Validate("genshin", "783437191")
	fmt.Printf("Genshin validate result: %v, error: %v\n", validGenshin, errGenshin)

	// Получаем продукты для Genshin
	productsGenshin, errGenshin2 := cli.GetProducts("genshin")
	fmt.Printf("Genshin products result: %#v, error: %v\n", productsGenshin, errGenshin2)

	fmt.Printf("\n=== HONKAI STAR RAIL TEST ===\n")
	// Валидируем HSR аккаунт
	validHSR, errHSR := cli.Validate("honkai-star-rail", "722354753")
	fmt.Printf("HSR validate result: %v, error: %v\n", validHSR, errHSR)

	// Получаем продукты для HSR
	productsHSR, errHSR2 := cli.GetProducts("honkai-star-rail")
	fmt.Printf("HSR products result: %#v, error: %v\n", productsHSR, errHSR2)

	fmt.Printf("\n=== ZENLESS ZONE ZERO TEST ===\n")
	// Валидируем ZZZ аккаунт
	validZZZ, errZZZ := cli.Validate("zenless-zone-zero", "1506168129")
	fmt.Printf("ZZZ validate result: %v, error: %v\n", validZZZ, errZZZ)

	// Получаем продукты для ZZZ
	productsZZZ, errZZZ2 := cli.GetProducts("zenless-zone-zero")
	fmt.Printf("ZZZ products result: %#v, error: %v\n", productsZZZ, errZZZ2)

	http.HandleFunc("/callback", hello)
	fmt.Println("Listen and serve: ", http.ListenAndServe(":9999", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	cli := kassaclient.KassaClient{}

	cb, err := cli.ParseCallback(r)
	fmt.Printf("Callback: %#v, error: %v\n", cb, err)
}
