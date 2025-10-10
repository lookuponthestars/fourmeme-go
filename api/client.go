package api

import "net/http"

type ApiClient struct {
	httpClient *http.Client
	authToken  string
}

func NewApiClient() *ApiClient {
	return &ApiClient{
		httpClient: &http.Client{},
	}
}
