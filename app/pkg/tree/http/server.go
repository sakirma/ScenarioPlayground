package http

import (
	"ScenarioPlayground/contract"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

type TreeHTTP struct {
	producer contract.TreeProducer
}

func (t TreeHTTP) SetTreeProducer(producer contract.TreeProducer) {
	t.producer = producer
}

func (t TreeHTTP) Serve() error {
	app := fiber.New()

	app.Get("/healthy", func(ctx fiber.Ctx) error {
		status := t.producer.Healthy()

		if status {
			return ctx.SendStatus(http.StatusOK)
		} else {
			return ctx.SendStatus(http.StatusServiceUnavailable)
		}
	})

	app.Get("/readiness", func(ctx fiber.Ctx) error {
		status := t.producer.Ready()

		if status {
			return ctx.SendStatus(http.StatusOK)
		} else {
			return ctx.SendStatus(http.StatusServiceUnavailable)
		}
	})

	return app.Listen(":8080")
}
