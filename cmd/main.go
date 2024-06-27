package main

import (
	"crm_api_gateway/api"
	"crm_api_gateway/config"
	"crm_api_gateway/pkg/grpc_client"
	"crm_api_gateway/pkg/logger"
)

var (
	log        logger.Logger
	cfg        config.Config
	grpcClient *grpc_client.GrpcClient
)

func initDeps() {
	var err error
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "crm_api_gateway")

	grpcClient, err = grpc_client.New(cfg)
	if err != nil {
		log.Error("grpc dial error", logger.Error(err))
	}
}

func main() {
	initDeps()

	server := api.New(api.Config{
		Logger:     log,
		GrpcClient: grpcClient,
		Cfg:        cfg,
	})

	server.Run(":8080")
}

// go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn"
