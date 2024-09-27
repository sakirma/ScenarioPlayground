package main

import (
	"ScenarioPlayground/contract"
	"fmt"
	"github.com/gofiber/fiber/v3/client"
	"net/http"
)

func main() {
	setupTree()

}

func setupTree() (contract.TreeHTTP, contract.TreeProducer) {

}

var cc = client.New()

func ready(endPoint string) error {
	resp, err := cc.Get(endPoint)
	if err != nil {
		return err
	}

	if resp != nil && resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return nil
}
