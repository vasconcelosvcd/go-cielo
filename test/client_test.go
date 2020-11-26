package test

import (
	"fmt"
	"github.com/vasconcelosvcd/go-cielo"
	"os"
	"testing"
)

func getNewClient(success bool, t *testing.T) (*cielo.Client, error) {
	t.Helper()
	if success {
		return cielo.NewClient(os.Getenv("MERCHANT_ID"), os.Getenv("MERCHANT_KEY"), cielo.SandboxEnvironment)
	}
	return cielo.NewClient("", "", cielo.SandboxEnvironment)
}

func Test_NewClient(t *testing.T) {
	t.Run("SUCCESS", func(t *testing.T) {
		client, err := getNewClient(true, t)
		if err != nil {
			t.Fatalf("Errors was not expected. Err: %v", err.Error())
		}

		expected := "*cielo.Client"
		got := fmt.Sprintf("%T", client)

		if got != expected {
			t.Errorf("Was expected '%v', but got '%v'", expected, got)
		}
	})

	t.Run("FAIL", func(t *testing.T) {
		client, err := getNewClient(false, t)
		if err == nil {
			t.Error("Errors was expected.")
		}

		if client != nil {
			t.Error("Client must be nil")
		}
	})
}
