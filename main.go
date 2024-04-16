package main

import (
	"go_plate/app/middleware"

	"go_plate/app/module"
	"go_plate/app/router"
	_ "go_plate/docs"
	"go_plate/internal/bootstrap"
	"go_plate/internal/bootstrap/database"
	"go_plate/pkg/config"

	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// @title            go_plate API Specification
// @version         1.0
// @description     This is the API specification for the go_plate backend.
// @contact.name   API Support
// @contact.url    https://www.go_plate.com/support
// @contact.email  developers@go_plate.com
// @host      localhost:8080
// @BasePath
func main() {
	fx.New(
		// Provide patterns
		fx.Provide(
			config.NewConfig,
			database.NewDatabase,
			middleware.NewMiddleware,
			bootstrap.NewFiber,
			bootstrap.NewLogger,
			router.NewRouter,
		),

		// Provide modules
		module.NewModule,

		// Start Application
		fx.Invoke(bootstrap.Start),

		// This will replace the [Fx] messages with messages printed to the logger.
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	).Run()
}
