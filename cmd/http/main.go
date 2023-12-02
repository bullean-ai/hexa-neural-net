package main

import (
	"context"
	"github.com/bullean-ai/hexa-neural-net/config"
	neuralNetJobs "github.com/bullean-ai/hexa-neural-net/domains/neural_net/application/jobs"
	neuralNetServices "github.com/bullean-ai/hexa-neural-net/domains/neural_net/application/services"
	neuralNetHandlers "github.com/bullean-ai/hexa-neural-net/domains/neural_net/handler/http"
	neuralNetRepos "github.com/bullean-ai/hexa-neural-net/domains/neural_net/infrastructure/repository"
	"github.com/bullean-ai/hexa-neural-net/pkg/databases/redis"
	"github.com/bullean-ai/hexa-neural-net/pkg/logger"
	"github.com/bullean-ai/hexa-neural-net/pkg/server"
	"github.com/bullean-ai/hexa-neural-net/pkg/utils/graceful_exit"
	"log"
)

// @title Auth Service
// @version 1.0
// @description Common Auth service broker with REST endpoints
// @contact.email ivanbarayev@hotmail.com
// @BasePath /v1
func main() {
	log.Println("Starting api server")

	cfg, errConfig := config.ParseConfig()
	if errConfig != nil {
		log.Fatal(errConfig)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.APP_VERSION, cfg.Logger.LEVEL, cfg.Server.APP_ENV)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init Clients
	redisDB, err := redis.NewRedisClient(ctx, cfg)
	if err != nil {
		appLogger.Fatal("Error when tyring to connect to Postgresql")
	} else {
		appLogger.Info("Postgresql connected")
	}

	// Init repositories
	redisRepo := neuralNetRepos.NewRedisRepo(ctx, redisDB)

	// Init services
	neuralNetService := neuralNetServices.NewNeuralNetService(cfg, redisRepo, appLogger)

	//Init jobs
	neuraLNetJobs := neuralNetJobs.NewJobRunner(cfg, appLogger, neuralNetService)

	// Interceptors
	//

	servers := server.NewServer(cfg, &ctx, appLogger)

	httpServer, errHttpServer := servers.NewHttpServer()
	if errHttpServer != nil {
		println(errHttpServer.Error())
	}
	versioning := httpServer.Group("/v1")

	// Init handlers for HTTP Server
	neuralNetHandler := neuralNetHandlers.NewHttpHandler(ctx, cfg, neuralNetService, appLogger)

	// Init routes for HTTP Server
	neuralNetHandlers.MapRoutes(neuralNetHandler, versioning)

	//telegram.SendMessage("Send Message to telegram channel")

	//Start Jobs
	go neuraLNetJobs.TrainNeuralNet(ctx)

	// Exit from application gracefully
	graceful_exit.TerminateApp(ctx)

	appLogger.Info("Server Exited Properly")
}
