package go1inch

import (
	"context"
	"testing"
)

func TestHealthCheckReponse(t *testing.T) {
	client := NewClient()
	res, _, err := client.Healthcheck(context.Background(), "eth")
	if err != nil {
		t.Error("Error while getting a healthcheck ", err)
	}
	if res.Status != "OK" {
		t.Errorf("healthcheck returned %s, expected OK", res.Status)
	}

}
