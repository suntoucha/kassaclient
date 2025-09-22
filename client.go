package kassaclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type KassaClient struct {
	BaseUrl string `json:"base_url"`
	Token   string `json:"token"`
}

func (x KassaClient) Validate(game string, account string) (bool, string, string, error) {
	type Request struct {
		Account string `json:"account"`
	}
	jsonData, err := json.Marshal(Request{Account: account})
	if err != nil {
		return false, "", "", err
	}

	req, err := http.NewRequest("POST", x.BaseUrl+"/"+game+"/validate", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, "", "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", x.Token)

	code, body, err := execute(req)
	if err != nil {
		return false, "", "", err
	}

	if code != 200 {
		return false, "", "", fmt.Errorf("Validate request failed, code: %v", code)
	}

	type Response struct {
		Valid  bool `json:"valid"`
		Server struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"server"`
	}
	var resp Response
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return false, "", "", err
	}

	return resp.Valid, resp.Server.Id, resp.Server.Name, nil
}

type PaymentOption struct {
	ProviderPayId string  `json:"provider_pay_id"`
	Price         float64 `json:"price"`
	Method        string  `json:"method"`
	Url           string  `json:"url"`
}

type Cart struct {
	CartId      string          `json:"cart_id"`
	Account     string          `json:"account"`
	Amount      float64         `json:"amount"`
	CallbackUrl string          `json:"callback_url"`
	Pay         []PaymentOption `json:"pay"`
}

func (x KassaClient) Cart(cartId string, game string, account string, amount float64, callbackUrl string, productId string, serverId string) (Cart, error) {
	type Request struct {
		CartId      string  `json:"cart_id"`
		Account     string  `json:"account"`
		Amount      float64 `json:"amount"`
		ProductId   string  `json:"product_id,omitempty"`
		CallbackUrl string  `json:"callback_url"`
		ServerId    string  `json:"server_id,omitempty"`
	}

	jsonData, err := json.Marshal(Request{
		CartId:      cartId,
		Account:     account,
		Amount:      amount,
		ProductId:   productId,
		CallbackUrl: callbackUrl,
		ServerId:    serverId,
	})
	if err != nil {
		return Cart{}, err
	}

	req, err := http.NewRequest("POST", x.BaseUrl+"/"+game+"/cart", bytes.NewBuffer(jsonData))
	if err != nil {
		return Cart{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", x.Token)

	code, body, err := execute(req)
	if err != nil {
		return Cart{}, err
	}

	type Response struct {
		Status string `json:"status"`
		Error  string `json:"error"`
		Cart
	}
	var resp Response
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return Cart{}, err
	}

	if code != 200 || resp.Status != "success" {
		return Cart{}, fmt.Errorf("Cart request failed, code: %v, status: %v, error: %v", code, resp.Status, resp.Error)
	}

	return resp.Cart, nil
}

type Callback struct {
	ProviderPayId string `json:"provider_pay_id"`
	CartId        string `json:"cart_id"`
	Status        string `json:"status"`
	Error         string `json:"error"`
}

func (x KassaClient) ParseCallback(r *http.Request) (Callback, error) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return Callback{}, err
	}

	cb := Callback{}
	err = json.Unmarshal(body, &cb)
	if err != nil {
		return Callback{}, err
	}

	return cb, nil
}

type Product struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (x KassaClient) Product(game string, serverId string) ([]Product, error) {
	type Request struct {
		ServerId string `json:"server_id"`
	}

	jsonData, err := json.Marshal(Request{
		ServerId: serverId,
	})
	if err != nil {
		return []Product{}, err
	}

	req, err := http.NewRequest("POST", x.BaseUrl+"/"+game+"/product", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", x.Token)

	code, body, err := execute(req)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("GetProducts request failed, code: %v, body: %v", code, body)
	}

	type Response struct {
		Status  string    `json:"status"`
		Error   string    `json:"error"`
		Product []Product `json:"product"`
	}
	var resp Response
	err = json.Unmarshal([]byte(body), &resp)
	if err != nil {
		return nil, err
	}

	if resp.Status != "success" {
		return nil, fmt.Errorf("GetProducts request failed, status: %v, error: %v", resp.Status, resp.Error)
	}

	return resp.Product, nil
}
