package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/minio/minio-go/v6"
	"go.uber.org/zap"
)

func main() {

	platform := new(Platform)

	settings := new(Settings)

	settings.Load("./settings.json")

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	platform.logger = logger.Sugar()

	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s",
			settings.DB.Host, settings.DB.Port,
			settings.DB.User, settings.DB.DB, settings.DB.Password),
	)
	if err != nil {
		platform.logger.Panic("Unable to connect to DB: ", err)
	}
	defer db.Close()

	// Storage Client
	endpoint := settings.Spaces.Endpoint
	accessKeyID := settings.Spaces.AccessKeyID
	secretAccessKey := settings.Spaces.SecretAccessKey
	storageClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, true)
	if err != nil {
		platform.logger.Panic("Unable to configure Spaces", err)
	}

	platform.storage = storageClient
	platform.db = db
	platform.settings = settings

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/releases")
	r.GET("/check", platform.checkUpdate)

	g := e.Group("/admin")
	g.POST("/release", platform.addRelease)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
