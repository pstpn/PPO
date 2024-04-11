package main

import (
	"log"
	"os"

	"course/config"
)

func main() {
	// Read config
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create logger
	loggerFile, err := os.Open(c.Logger.File)
	if err != nil {
		log.Fatal(err)
	}
	defer func(loggerFile *os.File) {
		err := loggerFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(loggerFile)
	//l := logger.New(c.Logger.Level, loggerFile)

	// Connect to Postgres
	//db, err := postgres.New(fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
	//	c.Database.Postgres.User,
	//	c.Database.Postgres.Password,
	//	c.Database.Postgres.Host,
	//	c.Database.Postgres.Port,
	//	c.Database.Postgres.Database,
	//))
	//// Connect to MongoDB
	//mongo, err := mongodb.New(
	//	c.Database.MongoDB.URI,
	//	c.Database.MongoDB.Database,
	//	c.Database.MongoDB.Bucket,
	//)

	// Create storages
	//checkpointStorage := storage.NewCheckpointStorage(db)
	//companyStorage := storage.NewCompanyStorage(db)
	//documentStorage := storage.NewDocumentStorage(db)
	//employeeStorage := storage.NewEmployeeStorage(db)
	//fieldStorage := storage.NewFieldStorage(db)
	//infoCardStorage := storage.NewInfoCardStorage(db)
	//photoMetaStorage := storage.NewPhotoMetaStorage(db)
	//
	//photoDataStorage := mdb.NewPhotoDataStorage(mongo)

	// Create services
}
