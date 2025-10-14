package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type raisedToken struct {
	Symbol         string   `json:"symbol"`
	NativeSymbol   string   `json:"nativeSymbol"`
	SymbolAddress  string   `json:"symbolAddress"`
	DeployCost     string   `json:"deployCost"`
	BuyFee         string   `json:"buyFee"`
	SellFee        string   `json:"sellFee"`
	MinTradeFee    string   `json:"minTradeFee"`
	B0Amount       string   `json:"b0Amount"`
	TotalBAmount   string   `json:"totalBAmount"`
	TotalAmount    string   `json:"totalAmount"`
	LogoUrl        string   `json:"logoUrl"`
	TradeLevel     []string `json:"tradeLevel"`
	Status         string   `json:"status"`
	BuyTokenLink   string   `json:"buyTokenLink"`
	ReservedNumber int      `json:"reservedNumber"`
	SaleRate       string   `json:"saleRate"`
	NetworkCode    string   `json:"networkCode"`
	Platform       string   `json:"platform"`
}

type createRequest struct {
	Name         string      `json:"name"`
	ShortName    string      `json:"shortName"`
	Desc         string      `json:"desc"`
	WebUrl       string      `json:"webUrl,omitempty"`
	TwitterUrl   string      `json:"twitterUrl,omitempty"`
	TelegramUrl  string      `json:"telegramUrl,omitempty"`
	TotalSupply  int64       `json:"totalSupply"`
	RaisedAmount string      `json:"raisedAmount"`
	SaleRate     float64     `json:"saleRate"`
	ReserveRate  int         `json:"reserveRate"`
	ImgUrl       string      `json:"imgUrl"`
	RaisedToken  raisedToken `json:"raisedToken"`
	LaunchTime   int64       `json:"launchTime"`
	FunGroup     bool        `json:"funGroup"`
	PreSale      string      `json:"preSale"`
	ClickFun     bool        `json:"clickFun"`
	Symbol       string      `json:"symbol"`
	Label        string      `json:"label"`
	LpTradingFee float64     `json:"lpTradingFee"`
	DexType      string      `json:"dexType"`
	RushMode     bool        `json:"rushMode"`
	OnlyMPC      bool        `json:"onlyMPC"`
}

type createResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TokenId     int64  `json:"tokenId"`
		TotalAmount string `json:"totalAmount"`
		SaleAmount  string `json:"saleAmount"`
		Template    int    `json:"template"`
		LaunchTime  int64  `json:"launchTime"`
		ServerTime  int64  `json:"serverTime"`
		CreateArg   string `json:"createArg"`
		Signature   string `json:"signature"`
		Bamount     string `json:"bamount"`
		Tamount     string `json:"tamount"`
	} `json:"data"`
}

func (c *ApiClient) GetCreateParameters(
	name, symbol, description string,
	imagePath string,
	websiteUrl string,
	twitterUrl string,
	telegramUrl string,
	buyAmount string,
	raisedToken, tokenLabel int,
) (*createResponse, error) {
	if raisedToken < 0 || raisedToken >= len(raisedTokenOptions) {
		return nil, fmt.Errorf("error: invalid raised token option")
	}

	if tokenLabel < 0 || tokenLabel >= len(labels) {
		return nil, fmt.Errorf("error: invalid token label option")
	}

	raisedTokenOption := raisedTokenOptions[raisedToken]

	fourmemeImgUrl, err := c.uploadFile(imagePath)

	if err != nil {
		return nil, fmt.Errorf("error uploading image to fourmeme: %w", err)
	}

	body := createRequest{
		Name:         name,
		ShortName:    name,
		Desc:         description,
		TotalSupply:  1000000000,
		RaisedAmount: raisedTokenOption.TotalBAmount,
		SaleRate:     0.8,
		ReserveRate:  0,
		ImgUrl:       fourmemeImgUrl,
		RaisedToken:  raisedTokenOption,
		LaunchTime:   time.Now().UnixMilli(),
		FunGroup:     false,
		PreSale:      buyAmount,
		ClickFun:     false,
		Symbol:       raisedTokenOption.Symbol,
		Label:        labels[tokenLabel],
		LpTradingFee: 0.0025,
		DexType:      "PANCAKE_SWAP",
		RushMode:     false,
		OnlyMPC:      false,
	}

	if websiteUrl != "" {
		if strings.HasPrefix(websiteUrl, "http://") || strings.HasPrefix(websiteUrl, "https://") {
			body.WebUrl = websiteUrl
		}
	}

	if twitterUrl != "" {
		if (strings.HasPrefix(websiteUrl, "http://") || strings.HasPrefix(websiteUrl, "https://")) && (strings.Contains(twitterUrl, "twitter.com/") || strings.Contains(twitterUrl, "x.com/")) {
			body.TwitterUrl = twitterUrl
		}
	}

	if telegramUrl != "" {
		if (strings.HasPrefix(telegramUrl, "http://") || strings.HasPrefix(telegramUrl, "https://")) && strings.Contains(telegramUrl, "t.me/") {
			body.TelegramUrl = telegramUrl
		}
	}

	requestBody, err := json.Marshal(&body)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, "https://four.meme/meme-api/v1/private/token/create", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	request.Header = headers
	request.Header.Set("meme-web-access", c.authToken)

	response, err := c.httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var jsonResponse createResponse

	err = json.NewDecoder(response.Body).Decode(&jsonResponse)

	if err != nil {
		return nil, err
	}

	if jsonResponse.Code != 0 || jsonResponse.Msg != "success" {
		return nil, fmt.Errorf("error from server when creating token: %s", jsonResponse.Msg)
	}

	return &jsonResponse, nil
}

var (
	raisedTokenOptions = []raisedToken{
		{
			Symbol:        "BNB",
			NativeSymbol:  "BNB",
			SymbolAddress: "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c",
			DeployCost:    "0",
			BuyFee:        "0.01",
			SellFee:       "0.01",
			MinTradeFee:   "0",
			B0Amount:      "8",
			TotalBAmount:  "18",
			TotalAmount:   "1000000000",
			LogoUrl:       "https://static.four.meme/market/fc6c4c92-63a3-4034-bc27-355ea380a6795959172881106751506.png",
			TradeLevel: []string{
				"0.1",
				"0.5",
				"1",
			},
			Status:         "PUBLISH",
			BuyTokenLink:   "https://pancakeswap.finance/swap",
			ReservedNumber: 10,
			SaleRate:       "0.8",
			NetworkCode:    "BSC",
			Platform:       "MEME",
		},
		{
			Symbol:        "CAKE",
			NativeSymbol:  "CAKE",
			SymbolAddress: "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
			DeployCost:    "0",
			BuyFee:        "0.01",
			SellFee:       "0.01",
			MinTradeFee:   "0",
			B0Amount:      "2000",
			TotalBAmount:  "10000",
			TotalAmount:   "1000000000",
			LogoUrl:       "https://static.four.meme/market/349f0ef2-d7be-4f7e-931c-385dc147ccc612887266201720862208.png",
			TradeLevel: []string{
				"50",
				"200",
				"500",
			},
			Status:         "PUBLISH",
			BuyTokenLink:   "https://pancakeswap.finance/swap",
			ReservedNumber: 10,
			SaleRate:       "0.8",
			NetworkCode:    "BSC",
			Platform:       "MEME",
		},
		{
			Symbol:        "USDT",
			NativeSymbol:  "USDT",
			SymbolAddress: "0x55d398326f99059ff775485246999027b3197955",
			DeployCost:    "0",
			BuyFee:        "0.01",
			SellFee:       "0.01",
			MinTradeFee:   "0",
			B0Amount:      "4000",
			TotalBAmount:  "12000",
			TotalAmount:   "1000000000",
			LogoUrl:       "https://static.four.meme/market/fb833cca-71e4-48f6-97dc-d1629cb21c0f1634031926700046438.png",
			TradeLevel: []string{
				"50",
				"250",
				"500",
			},
			Status:         "PUBLISH",
			BuyTokenLink:   "https://pancakeswap.finance/swap?outputCurrency=0x55d398326f99059fF775485246999027B3197955",
			ReservedNumber: 10,
			SaleRate:       "0.8",
			NetworkCode:    "BSC",
			Platform:       "MEME",
		},
		{
			Symbol:        "USD1",
			NativeSymbol:  "USD1",
			SymbolAddress: "0x8d0d000ee44948fc98c9b98a4fa4921476f08b0d",
			DeployCost:    "0",
			BuyFee:        "0.01",
			SellFee:       "0.01",
			MinTradeFee:   "0",
			B0Amount:      "4000",
			TotalBAmount:  "12000",
			TotalAmount:   "1000000000",
			LogoUrl:       "https://static.four.meme/market/e04be5fd-1dce-4398-9aa1-4a2aa8bdfd467257033410940904813.png",
			TradeLevel: []string{
				"50",
				"250",
				"500",
			},
			Status:         "PUBLISH",
			BuyTokenLink:   "https://pancakeswap.finance/swap?outputCurrency=0x8d0D000Ee44948FC98c9B98A4FA4921476f08B0d",
			ReservedNumber: 10,
			SaleRate:       "0.8",
			NetworkCode:    "BSC",
			Platform:       "MEME",
		},
		{
			Symbol:        "ASTER",
			NativeSymbol:  "ASTER",
			SymbolAddress: "0x000Ae314E2A2172a039B26378814C252734f556A",
			DeployCost:    "0",
			BuyFee:        "0.01",
			SellFee:       "0.01",
			MinTradeFee:   "0",
			B0Amount:      "4000",
			TotalBAmount:  "5000",
			TotalAmount:   "1000000000",
			LogoUrl:       "https://static.four.meme/market/7e4a117c-496f-435c-9410-83af2d5b0d8014173403460554224381.png",
			TradeLevel: []string{
				"30",
				"150",
				"300",
			},
			Status:         "PUBLISH",
			BuyTokenLink:   "https://pancakeswap.finance/swap?outputCurrency=0x000Ae314E2A2172a039B26378814C252734f556A",
			ReservedNumber: 10,
			SaleRate:       "0.8",
			NetworkCode:    "BSC",
			Platform:       "MEME",
		},
	}

	labels = []string{
		"Meme",
		"AI",
		"Defi",
		"Games",
	}
)
