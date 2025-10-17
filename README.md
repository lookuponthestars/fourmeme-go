# fourmeme-go

So far I only made `CreateToken`. At some point I might finish the whole package (buying, selling, etc)
Uses the fourmeme api to create token


```Go
package main

import (
	"fmt"

	"github.com/lookuponthestars/fourmeme-go"
)

func main() {
	client, err := fourmeme.NewClient(
		"https://rpc.48.club",
		"<YOUR_PRIVATE_KEY>",
	)

	if err != nil {
		panic(err)
	}

	txid, err := client.CreateToken(
		"token_name",
		"token_symbol",
		"token_description",
		"path/to/image.png|https://path.to/image.png",
		"https://token.website",
		"https://x.com/token_twitter",
		"https://t.me/token_telegram",
		"0.01", // initial buy amount, can be 0
		fourmeme.RaisedTokenBNB, // bnb only works probably for now, idk if i'll make it work for other coins at some point
		fourmeme.LabelMeme, // label for token
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction ID:", txid)
}
```