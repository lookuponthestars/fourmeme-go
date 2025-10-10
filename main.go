package fourmeme

import (
	"crypto/ecdsa"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lookuponthestars/fourmeme-go/api"
)

type Client struct {
	rpcClient   *ethclient.Client
	privateKey  *ecdsa.PrivateKey
	fourmemeApi *api.ApiClient
	address     string
}

func NewClient(
	rpcUrl string,
	privateKeyStr string,
) (*Client, error) {
	c := &Client{}

	client, err := ethclient.Dial(rpcUrl)

	if err != nil {
		return nil, err
	}

	c.rpcClient = client

	privkey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyStr, "0x"))

	if err != nil {
		return nil, err
	}

	c.privateKey = privkey

	publicKey := c.privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	c.address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	c.fourmemeApi = api.NewApiClient()

	signMessageStr, err := c.fourmemeApi.CreateMessage(c.address)

	if err != nil {
		return nil, err
	}

	signMessageBytes := accounts.TextHash([]byte(signMessageStr))

	signedMessage, err := crypto.Sign(signMessageBytes, c.privateKey)

	if err != nil {
		return nil, err
	}

	signedMessageStr := hexutil.Encode(signedMessage)

	err = c.fourmemeApi.SendMessage(c.address, signedMessageStr)

	if err != nil {
		return nil, err
	}

	return c, nil
}
