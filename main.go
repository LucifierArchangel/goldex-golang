package GoldEx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	BaseURLV2 = "https://stage.goldex.space/api/v2"
)

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type successResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BaseURLV2,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) GetCurrencies(ctx context.Context) (*Currencies, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/currency", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	var res Currencies

	if err = c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

func (c *Client) GetExchangeRates(ctx context.Context) (*ExchangeRates, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/exchange_rate", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	var res ExchangeRates

	if err = c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil

}

func (c *Client) GetAllRequests(ctx context.Context) (*Requests, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/request", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	var res Requests

	if err = c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetRequestById(ctx context.Context, requestId string) (*Requests, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/request/%s", c.BaseURL, requestId), nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	var res Requests

	if err = c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateRequest(ctx context.Context, request CreateRequest) (*Request, error) {
	jsonData, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/request/", c.BaseURL), bytes.NewReader(jsonData))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	var res Request

	if err = c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateRequest(ctx context.Context, requestId int, request UpdateRequest) (*Request, error) {
	jsonData, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/request/%d", c.BaseURL, requestId), bytes.NewReader(jsonData))

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	var res Request

	if err = c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetBalance(ctx context.Context) (*UserBalance, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/user/balance", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	var res UserBalance

	if err = c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ReloadApiKey(ctx context.Context) (*User, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/user/reload_key", c.BaseURL), nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	var res User

	if err = c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("auth-token", c.apiKey)

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse

		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			fmt.Println("Err with error response")
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	fullResponse := successResponse{
		Data: v,
	}

	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		fmt.Println("Err with success response")
		return err
	}

	return nil
}
