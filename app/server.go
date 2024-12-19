package app

import (
	"goTgTemplate/app/di"
	"goTgTemplate/app/nats"
)

func Start() {
	u := di.Init()
	nats.Init(u)

	select {}
}
