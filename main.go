package main

import (
	"go_plate/app/middleware"
	"go_plate/app/module"
	"go_plate/app/router"

	"go_plate/internal/config"
	"go_plate/internal/database"
	"go_plate/internal/logger"
	"go_plate/internal/server"

	fxzerolog "github.com/efectn/fx-zerolog"
	_ "github.com/joho/godotenv/autoload"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		// Provide patterns
		fx.Provide(
			config.NewConfig,
			database.NewDatabase,
			middleware.NewMiddleware,
			server.NewFiber,
			logger.NewLogger,
			router.NewRouter,
		),

		// Provide modules
		module.NewModule,

		// Start Application
		fx.Invoke(server.Start),

		// This will replace the [Fx] messages with messages printed to the logger.
		fx.WithLogger(fxzerolog.InitPtr()),
	).Run()
}
