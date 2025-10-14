package fourmeme

import (
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	fourmeme_abi "github.com/lookuponthestars/fourmeme-go/abi"
	"github.com/lookuponthestars/fourmeme-go/api"
)

const FOURMEME_CONTRACT = "0x5c952063c7fc8610FFDB798152D69F0B9550762b"

var weiPerBnb = big.NewFloat(1e18)

type Client struct {
	rpcClient   *ethclient.Client
	privateKey  *ecdsa.PrivateKey
	fourmemeApi *api.ApiClient
	fourmemeAbi *fourmeme_abi.FourMeme
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

	abiClient, err := fourmeme_abi.NewFourMeme(
		common.HexToAddress(FOURMEME_CONTRACT),
		c.rpcClient,
	)

	if err != nil {
		return nil, err
	}

	c.fourmemeAbi = abiClient

	return c, nil
}
