package di

import (
	"github.com/nats-io/nats.go"
	"goTgTemplate/internal/controller"
)

type Initialization struct {
	Nc             *nats.Conn
	TestController controller.TestController
}

func NewInitialization(nc *nats.Conn,
	testController controller.TestController) *Initialization {
	init := &Initialization{
		Nc:             nc,
		TestController: testController,
	}

	return init
}
