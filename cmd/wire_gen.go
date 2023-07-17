// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"callMysql/internal/biz"
	"callMysql/internal/conf"
	"callMysql/internal/data"
	"callMysql/internal/server"
	"callMysql/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData ,logger)
	if err != nil {
		return nil, nil, err
	}
	dataController := data.NewMysqlController(dataData ,logger)
	dataControllerUsecase := biz.NewDataControllerUsecase(dataController,logger)
	callMysqlService := service.NewCallMysqlService(dataControllerUsecase)
	grpcServer := server.NewGRPCServer(confServer, callMysqlService,logger)
	httpServer := server.NewHTTPServer(confServer, callMysqlService,logger)
	// httpServer := NewMuxHttpServer()
	app := newApp(grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
