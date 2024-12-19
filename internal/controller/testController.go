package controller

import (
	"github.com/nats-io/nats.go"
	"goTgTemplate/internal/service"
)

type TestController interface {
	TestCallback(m *nats.Msg)
}

type TestControllerImpl struct {
	testService service.TestService
}

func TestControllerInit(testService service.TestService) *TestControllerImpl {
	return &TestControllerImpl{
		testService: testService,
	}
}

func (u *TestControllerImpl) TestCallback(m *nats.Msg) {
	return
}
