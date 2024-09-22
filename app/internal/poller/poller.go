package poller

import (
	"errors"
	"github.com/gofiber/fiber/v3/client"
	"github.com/gofiber/fiber/v3/log"
	"net/http"
	"time"
)

type Poller struct {
	endPoint string
	interval time.Duration

	client *client.Client
}

func NewPoller(endPoint string, intervalSec uint32) *Poller {
	cc := client.New()
	cc.SetTimeout(10 * time.Second)

	return &Poller{
		client:   cc,
		interval: time.Duration(intervalSec) * time.Second,
		endPoint: endPoint,
	}
}

func (p *Poller) Start() {
	go func() {
		timer := time.NewTicker(p.interval)
		defer timer.Stop()

		for range timer.C {
			err := p.Poll()
			if err != nil {
				log.Errorf("polling the destination: %v", err.Error())
			}
		}
	}()
}

func (p *Poller) Poll() error {
	resp, err := p.client.Get(p.endPoint)

	if resp != nil && resp.StatusCode() != http.StatusOK {
		return errors.New("status was not OK")
	}
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (p *Poller) SetTimeout(timeout time.Duration) {
	p.client.SetTimeout(timeout)
}
