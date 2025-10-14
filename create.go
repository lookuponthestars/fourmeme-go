package fourmeme

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

const (
	RaisedTokenBNB = iota
	RaisedTokenCAKE
	RaisedTokenUSDT
	RaisedTokenUSD1
	RaisedTokenASTER
)

const (
	LabelMeme = iota
	LabelAi
	LabelDefi
	LabelGames
)

func (c *Client) CreateToken(
	name, symbol, description string,
	imagePath string,
	websiteUrl string,
	twitterUrl string,
	telegramUrl string,
	buyAmount string,
	raisedToken, tokenLabel int,
) (string, error) {
	createParams, err := c.fourmemeApi.GetCreateParameters(
		name, symbol, description,
		imagePath,
		websiteUrl,
		twitterUrl,
		telegramUrl,
		buyAmount,
		raisedToken, tokenLabel,
	)

	if err != nil {
		return "", err
	}

	auth := bind.NewKeyedTransactor(c.privateKey, big.NewInt(56))

	args := common.Hex2Bytes(strings.TrimPrefix(createParams.Data.CreateArg, "0x"))
	signature := common.Hex2Bytes(strings.TrimPrefix(createParams.Data.Signature, "0x"))

	buyAmountBig, _ := new(big.Float).SetString(buyAmount)
	bnbWei := new(big.Float).Mul(new(big.Float).SetFloat64(0.01), weiPerBnb)
	buyAmountBig.Mul(buyAmountBig, weiPerBnb)
	bnbWei.Add(bnbWei, buyAmountBig)
	bnbValue, _ := bnbWei.Int(nil)

	auth.Value = bnbValue

	transaction, err := c.fourmemeAbi.CreateToken(auth, args, signature)

	if err != nil {
		return "", fmt.Errorf("error creating token: %w", err)
	}

	return transaction.Hash().Hex(), nil
}
