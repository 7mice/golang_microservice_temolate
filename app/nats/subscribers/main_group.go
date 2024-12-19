package nats

import (
	"github.com/nats-io/nats.go"
	"goTgTemplate/app/di"
	"goTgTemplate/pkg/logging"
)

func SetupMainGroup(nc *nats.Conn, u *di.Initialization) {
	logging.DefaultLogger.Info().Msg("Setting up main nats group...")

	if _, err := nc.Subscribe("check_subscribe", u.TestController.TestCallback); err != nil {
		return
	}
}
