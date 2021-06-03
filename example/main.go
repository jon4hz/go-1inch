package main

import (
	"context"
	"fmt"
	"log"

	go1inch "github.com/jon4hz/go-1inch"
)

func main() {
	client := go1inch.NewClient()
	res, _, err := client.ApproveCalldata(context.Background(), "eth", 500000000, true, "0x6b175474e89094c44da98b954eedeac495271d0f")
	if err != nil {
		log.Fatal("Error while getting a approve calldata ", err)
	}
	fmt.Println(res)
}
