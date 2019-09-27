package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {

	platform := new(Platform)

	settings := new(Settings)

	settings.Load("./settings.json")

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s",
			settings.DB.Host, settings.DB.Port,
			settings.DB.User, settings.DB.DB, settings.DB.Password),
	)
	if err != nil {

	}
	defer db.Close()

	platform.db = db
	platform.logger = logger.Sugar()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/releases")
	r.GET("/check", platform.checkUpdate)

	g := e.Group("/admin")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
