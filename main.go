package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cecepsprd/inventory-app/config"
	"github.com/cecepsprd/inventory-app/handler"
	"github.com/cecepsprd/inventory-app/repository"
	"github.com/cecepsprd/inventory-app/service"
	"github.com/cecepsprd/inventory-app/utils/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var cfg = config.NewConfig()

	db, err := cfg.MysqlConnect()
	if err != nil {
		log.Fatal("error connecting to database: ", err.Error())
	}
	defer db.Close()

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
	}))

	e.GET("/api/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service OK!")
	})

	// Starting server
	go func() {
		if cfg.App.HTTPPort == "" {
			cfg.App.HTTPPort = os.Getenv("PORT")
		}

		err := e.Start(":" + cfg.App.HTTPPort)
		if err != nil {
			logger.Log.Error("error starting server: ", err)
		}
	}()

	inventoryRepo := repository.NewInventoryRepository(db)
	inventoryService := service.NewInventoryService(inventoryRepo)
	handler.NewInventoryHandler(e, inventoryService)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Block until a signal is received.
	<-quit

	// gracefully shutdown the server, waiting max 5 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = e.Shutdown(ctx); err != nil {
		logger.Log.Error("error shutting down server: ", err.Error())
	}
}
