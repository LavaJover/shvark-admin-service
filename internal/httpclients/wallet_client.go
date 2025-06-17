package httpclients

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type WalletHTTPClient struct {

}

func NewWalletHTTPClient() (*WalletHTTPClient, error){
	return &WalletHTTPClient{}, nil
}

type CreateWalletRequest struct {
	TraderID string `json:"traderId"`
}

type CreateWalletResponse struct {
	Address string `json:"address"`
}

func (c *WalletHTTPClient) CreateTraderWallet(request *CreateWalletRequest) (*CreateWalletResponse, error) {
	proxyRequestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	proxyResponse, err := http.Post("http://localhost:3000/wallets/create", "application/json", bytes.NewBuffer(proxyRequestBody))
	if err != nil {
		return nil, err
	}
	defer proxyResponse.Body.Close()

	proxyResponseBody, err := io.ReadAll(proxyResponse.Body)
	if err != nil {
		return nil, err
	}

	if proxyResponse.StatusCode >= 200 && proxyResponse.StatusCode < 300 {
		var response CreateWalletResponse
		if err := json.Unmarshal(proxyResponseBody, &response); err != nil {
			return nil, err
		}
		return &response, nil
	}

	return nil, err
}
