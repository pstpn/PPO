package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"course/config"
	"course/internal/controller/http"
	"course/internal/service"
	storage2 "course/internal/storage"
	mdb "course/internal/storage/mongodb"
	storage "course/internal/storage/postgres"
	"course/pkg/jwt"
	"course/pkg/logger"
	httpserver "course/pkg/server/http"
	"course/pkg/storage/mongodb"
	"course/pkg/storage/postgres"
)

func main() {
	// Read config
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create auth token manager
	tokenManager, err := jwt.NewManager(c.Auth.SigningKey)
	if err != nil {
		log.Fatal(err)
	}

	// Create logger
	loggerFile, err := os.OpenFile(
		c.Logger.File,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		log.Fatal(err)
	}
	l := logger.New(c.Logger.Level, loggerFile)

	// Connect to Postgres
	db, err := postgres.New(fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		c.Database.Postgres.User,
		c.Database.Postgres.Password,
		c.Database.Postgres.Host,
		c.Database.Postgres.Port,
		c.Database.Postgres.Database,
	))
	// Connect to MongoDB
	mongo, err := mongodb.New(
		c.Database.MongoDB.URI,
		c.Database.MongoDB.Database,
		c.Database.MongoDB.Bucket,
	)

	// Create storages
	checkpointStorage := storage.NewCheckpointStorage(db)
	//companyStorage := storage.NewCompanyStorage(db)
	documentStorage := storage.NewDocumentStorage(db)
	employeeStorage := storage.NewEmployeeStorage(db)
	fieldStorage := storage.NewFieldStorage(db)
	infoCardStorage := storage.NewInfoCardStorage(db)
	photoMetaStorage := storage.NewPhotoMetaStorage(db)
	photoDataStorage := mdb.NewPhotoDataStorage(mongo)

	// Create services
	authService := service.NewAuthService(
		l,
		employeeStorage,
		infoCardStorage,
		tokenManager,
		c.Auth.AccessTokenTTL,
		c.Auth.RefreshTokenTTL,
	)
	infoCardService := service.NewInfoCardService(l, infoCardStorage)
	documentService := service.NewDocumentService(l, documentStorage)
	fieldService := service.NewFieldService(l, fieldStorage)
	photoService := service.NewPhotoService(l, storage2.PhotoStorage{
		PhotoDataStorage: photoDataStorage,
		PhotoMetaStorage: photoMetaStorage,
	})
	checkpointService := service.NewCheckpointService(l, checkpointStorage)

	// Create controller
	handler := gin.New()
	controller := http.NewRouter(handler)

	// Set routes
	controller.SetAuthRoute(l, authService)
	controller.SetInfoCardRoute(l, infoCardService, documentService, fieldService, checkpointService, photoService, authService)
	controller.SetProfileRoute(l, infoCardService, documentService, fieldService, authService, photoService)
	controller.SetPassageRoute(l, documentService, checkpointService, authService)

	// Create router
	router := httpserver.New(handler, httpserver.Port(c.HTTP.Port))

	// Starting server
	err = router.Start()
	if err != nil {
		log.Fatal(err)
	}
}
