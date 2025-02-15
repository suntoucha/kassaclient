package kassaclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (x KassaClient) SteamValidate(account string) (bool, error) {
	type Request struct {
		Account string `json:"account"`
	}
	jsonData, err := json.Marshal(Request{Account: account})
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", x.BaseUrl+"/steam/validate", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", x.Token)

	code, body, err := execute(req)
	if err != nil {
		return false, err
	}

	if code != 200 {
		return false, errors.New(fmt.Sprintf("Validate request failed, code: %v", code))
	}

	type Response struct {
		Valid bool `json:"valid"`
	}
	var resp Response
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return false, err
	}

	return resp.Valid, nil
}

type SteamPay struct {
	Price  float64 `json:"price"`
	Method string  `json:"method"`
	Url    string  `json:"url"`
}

type SteamCart struct {
	CartId      string     `json:"cart_id"`
	Account     string     `json:"account"`
	Amount      float64    `json:"amount"`
	CallbackUrl string     `json:"callback_url"`
	Pay         []SteamPay `json:"pay"`
}

func (x KassaClient) SteamCart(cartId string, account string, amount float64, callbackUrl string) (SteamCart, error) {
	type Request struct {
		CartId      string  `json:"cart_id"`
		Account     string  `json:"account"`
		Amount      float64 `json:"amount"`
		CallbackUrl string  `json:"callback_url"`
	}
	jsonData, err := json.Marshal(Request{
		CartId:      cartId,
		Account:     account,
		Amount:      amount,
		CallbackUrl: callbackUrl,
	})
	if err != nil {
		return SteamCart{}, err
	}

	req, err := http.NewRequest("POST", x.BaseUrl+"/steam/cart", bytes.NewBuffer(jsonData))
	if err != nil {
		return SteamCart{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", x.Token)

	code, body, err := execute(req)
	if err != nil {
		return SteamCart{}, err
	}

	type Response struct {
		Status string `json:"status"`
		Error  string `json:"error"`
		SteamCart
	}
	var resp Response
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return SteamCart{}, err
	}

	if code != 200 || resp.Status != "success" {
		return SteamCart{}, errors.New(fmt.Sprintf("Cart request failed, code: %v, status: %v, error: %v", code, resp.Status, resp.Error))
	}

	return resp.SteamCart, nil

}
