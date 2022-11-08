# go-1inch

A golang API wrapper for the 1inch API

## Getting started

```go
package main

import (
    "context"
    "fmt"
	"log"

    go1inch "github.com/jon4hz/go-1inch"
)

func main(){
    // create a new client
    client := go1inch.NewClient()

    // check the health status from the API pointing to the ethereum network
    health, _, err := client.Healthcheck(context.Background(), "eth")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(health)
}

```

## Ressources

- [Official API Docs](https://docs.1inch.io/docs/aggregation-protocol/api/swagger/)
- [Go Docs](https://pkg.go.dev/github.com/jon4hz/go-1inch)
