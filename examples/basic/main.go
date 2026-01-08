package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stax-ai/emailengine"
)

func main() {
	token := os.Getenv("EMAILENGINE_TOKEN")
	client, err := emailengine.New(
		emailengine.WithToken(token),
	)
	if err != nil {
		panic(err)
	}
	// Example: ping the service by listing accounts (first page)
	accounts, _, err := client.Raw().AccountAPI.GetV1Accounts(context.Background()).Execute()
	if err != nil {
		panic(err)
	}
	fmt.Printf("accounts: %d\n", len(accounts.GetAccounts()))
}
