package main

import (
	"ScenarioPlayground/internal/poller"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/client"
	"net/http"
	"os"
)

func main() {
	endPoint := os.Getenv("END_POINT")
	if endPoint == "" {
		panic("END_POINT environment variable not set")
	}

	p := poller.NewPoller(endPoint, 3)

	p.Start()

	app := fiber.New()

	app.Get("/liveness", func(ctx fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	})

	app.Get("/readiness", func(ctx fiber.Ctx) error {
		if err := ready(endPoint); err != nil {
			ctx.Status(http.StatusFailedDependency)
			return err
		}

		return ctx.SendStatus(http.StatusOK)
	})

	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}
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
