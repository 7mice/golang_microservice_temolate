package nats

import (
	"goTgTemplate/app/di"
	nats "goTgTemplate/app/nats/subscribers"
)

func Init(u *di.Initialization) {
	nats.SetupMainGroup(u.Nc, u)
}
