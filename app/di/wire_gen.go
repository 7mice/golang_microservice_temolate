// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"goTgTemplate/app/connections"
	"goTgTemplate/internal/controller"
	"goTgTemplate/internal/repository"
	"goTgTemplate/internal/service"
)

// Injectors from wire.go:

func Init() *Initialization {
	conn := connections.ConnectToNatsBroker()
	testRepositoryImpl := repository.TestRepositoryInit()
	testServiceImpl := service.TestServiceInit(testRepositoryImpl)
	testControllerImpl := controller.TestControllerInit(testServiceImpl)
	initialization := NewInitialization(conn, testControllerImpl)
	return initialization
}

// wire.go:

var connectionsSet = wire.NewSet(connections.ConnectToNatsBroker, connections.ConnectToRedis)

var testSet = wire.NewSet(repository.TestRepositoryInit, wire.Bind(new(repository.TestRepository), new(*repository.TestRepositoryImpl)), service.TestServiceInit, wire.Bind(new(service.TestService), new(*service.TestServiceImpl)), controller.TestControllerInit, wire.Bind(new(controller.TestController), new(*controller.TestControllerImpl)))
