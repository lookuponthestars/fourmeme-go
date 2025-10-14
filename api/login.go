package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var headers http.Header = http.Header{
	"accept":          []string{"application/json, text/plain, */*"},
	"accept-language": []string{"en-US,en;q=0.9"},
	"content-type":    []string{"application/json;charset=UTF-8"},
	"origin":          []string{"https://four.meme"},
	"user-agent":      []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36"},
}

type nonceRequest struct {
	AccountAddress string `json:"accountAddress"`
	VerifyType     string `json:"verifyType"`
	NetworkCode    string `json:"networkCode"`
}

type nonceResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func (c *ApiClient) CreateMessage(address string) (string, error) {
	body := nonceRequest{
		AccountAddress: address,
		VerifyType:     "LOGIN",
		NetworkCode:    "BSC",
	}

	requestBody, err := json.Marshal(&body)

	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(http.MethodPost, "https://four.meme/meme-api/v1/private/user/nonce/generate", bytes.NewBuffer(requestBody))

	if err != nil {
		return "", err
	}

	request.Header = headers

	response, err := c.httpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var jsonResponse nonceResponse

	err = json.NewDecoder(response.Body).Decode(&jsonResponse)

	if err != nil {
		return "", err
	}

	if jsonResponse.Code != 0 || jsonResponse.Msg != "success" {
		return "", fmt.Errorf("error from server when creating sign-in message: %s", jsonResponse.Msg)
	}

	return fmt.Sprintf("You are sign in Meme %s", jsonResponse.Data), nil
}

type verifyInfo struct {
	Address     string `json:"address"`
	NetworkCode string `json:"networkCode"`
	Signature   string `json:"signature"`
	VerifyType  string `json:"verifyType"`
}

type loginRequest struct {
	Region     string     `json:"region"`
	LangType   string     `json:"langType"`
	LoginIp    string     `json:"loginIp"`
	InviteCode string     `json:"inviteCode"`
	VerifyInfo verifyInfo `json:"verifyInfo"`
	WalletName string     `json:"walletName"`
}

type loginResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func (c *ApiClient) SendMessage(address, signature string) error {
	body := loginRequest{
		Region:     "WEB",
		LangType:   "EN",
		LoginIp:    "",
		InviteCode: "",
		VerifyInfo: verifyInfo{
			Address:     address,
			NetworkCode: "BSC",
			Signature:   signature,
			VerifyType:  "LOGIN",
		},
		WalletName: "Metamask",
	}

	requestBody, err := json.Marshal(&body)

	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, "https://four.meme/meme-api/v1/private/user/login/dex", bytes.NewBuffer(requestBody))

	if err != nil {
		return err
	}

	request.Header = headers

	response, err := c.httpClient.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	var jsonResponse loginResponse

	err = json.NewDecoder(response.Body).Decode(&jsonResponse)

	if err != nil {
		return err
	}

	if jsonResponse.Code != 0 || jsonResponse.Msg != "success" {
		return fmt.Errorf("error from server when creating sign-in message: %s", jsonResponse.Msg)
	}

	c.authToken = jsonResponse.Data

	return nil
}
