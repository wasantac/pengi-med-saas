package main

import (
	"os"
	"pengi-med-saas/core/database"
	"pengi-med-saas/core/logger"
	"pengi-med-saas/features/health"
	i18n_middleware "pengi-med-saas/i18n/middleware"
	"pengi-med-saas/migrations"
	"pengi-med-saas/routes"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var DB_CONNECTION *gorm.DB

func main() {
	mode := os.Getenv("GIN_MODE")
	if mode == "release" {
		mode = "production"
	} else {
		mode = "development"
	}

	logger.Init(mode)
	logger.Info("Starting application...", zap.String("env", mode))

	DB_CONNECTION, err := database.Connect()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	err = migrations.RunAllMigrations(DB_CONNECTION)

	if err != nil {
		panic("Failed to run migrations: " + err.Error())
	}

	r := gin.Default()

	r.Use(i18n_middleware.I18nMiddleware(DB_CONNECTION))

	r.GET("/health", health.Health)

	routes.RegisterRoutes(r.Group("/api"), DB_CONNECTION)

	r.Run() // listen and serve on 0.0.0.0:8080
}
