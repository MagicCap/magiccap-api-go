package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/minio/minio-go/v6"
	"go.uber.org/zap"
)

func init() {
	flag.BoolVar(&generateConfig, "g", false, "Generate config file")
	flag.Parse()
}

var generateConfig bool

func main() {

	platform := new(Platform)

	settings := new(Settings)

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	if generateConfig {
		settings.Save("settings.example.json")
		sugar.Info("Generated config and written to settings.example.json")
		os.Exit(0)
	}

	settings.Load("./settings.json")

	platform.logger = sugar

	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			settings.DB.Host, settings.DB.Port,
			settings.DB.User, settings.DB.DB, settings.DB.Password,
			settings.DB.SSLMode),
	)
	if err != nil {
		platform.logger.Fatal("Unable to connect to DB: ", err)
	}
	defer db.Close()

	db.AutoMigrate(&Release{})
	db.AutoMigrate(&Admin{})

	// Storage Client
	endpoint := settings.Spaces.Endpoint
	accessKeyID := settings.Spaces.AccessKeyID
	secretAccessKey := settings.Spaces.SecretAccessKey
	sugar.Debugf("Use SSL: %b", settings.Spaces.UseSSL)
	storageClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, settings.Spaces.UseSSL)
	if err != nil {
		platform.logger.Fatal("Unable to configure Spaces", err)
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
	g.Use(middleware.KeyAuth(platform.checkKey))
	g.POST("/release", platform.addRelease)

	port := settings.Port
	if port == "" {
		port = "1323"
	}

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
