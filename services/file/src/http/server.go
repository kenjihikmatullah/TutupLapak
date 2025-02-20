package httpServer

import (
	"fmt"

	"github.com/TimDebug/TutupLapak/File/src/cache"
	"github.com/TimDebug/TutupLapak/File/src/config"
	"github.com/TimDebug/TutupLapak/File/src/http/middleware/errorHandler"
	"github.com/TimDebug/TutupLapak/File/src/http/middleware/identifier"
	"github.com/TimDebug/TutupLapak/File/src/logger"
	"github.com/TimDebug/TutupLapak/File/src/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HttpServer struct {
	DB *pgxpool.Pool
}

func (s *HttpServer) Listen() {
	app := fiber.New(fiber.Config{
		ServerHeader: "TIM-DEBUG",
		ErrorHandler: errorHandler.ErrorHandler,
	})
	app.Use(identifier.RequestID)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: false,
		AllowMethods:     "POST",
	}))
	app.Use(logger.ResponseLogger)

	appConfig := config.GetConfig()
	var storageClient StorageClient
	if appConfig.IsProduction {
		storageClient = NewS3StorageClient()
	} else {
		storageClient = NewMockS3StorageClient()
	}
	repo := repo.NewFileRepository(s.DB)
	redis := cache.NewRedisClient()
	service := NewFileService(repo, storageClient, redis)
	defer service.Shutdown()
	controller := NewFileController(service)

	routes := app.Group("/v1")
	routes.Post("/file", controller.Upload)

	app.Listen(fmt.Sprintf("%s:%s", "0.0.0.0", appConfig.Port))
}
