package server

import (
	"context"
	"os"
	"runtime"
	"time"

	"go_plate/app/middleware"
	"go_plate/app/router"
	"go_plate/internal/config"
	"go_plate/internal/database"

	"go_plate/pkg/response"
	"go_plate/pkg/utils"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func NewFiber(cfg *config.Config) *fiber.App {
	// Setup Webserver
	app := fiber.New(fiber.Config{
		ServerHeader:          cfg.App.Name,
		AppName:               cfg.App.Name,
		Prefork:               cfg.App.Prefork,
		ErrorHandler:          response.ErrorHandler,
		IdleTimeout:           cfg.App.IdleTimeout * time.Second,
		EnablePrintRoutes:     cfg.App.PrintRoutes,
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})

	// Pass production config to check it
	response.IsProduction = cfg.App.Production

	return app
}

func Start(lifecycle fx.Lifecycle, cfg *config.Config, fiber *fiber.App, router *router.Router, middlewares *middleware.Middleware, database *database.Database) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Register middlewares & routes
				middlewares.Register()
				router.Register()

				// Custom Startup Messages
				host, port := utils.ParseAddr(cfg.App.Port)
				if host == "" {
					if fiber.Config().Network == "tcp6" {
						host = "[::1]"
					} else {
						host = "0.0.0.0"
					}
				}

				// Information message

				log.Info().Msg(fiber.Config().AppName + " is running at the moment!")
				// Debug informations
				if !cfg.App.Production {
					prefork := "Enabled"
					procs := runtime.GOMAXPROCS(0)
					if !cfg.App.Prefork {
						procs = 1
						prefork = "Disabled"
					}

					log.Debug().Str("Version", "-").Str("Host", host).Str("Port", port).Str("Prefork", prefork).Uint32("Handlers", fiber.HandlersCount()).Int("Processes", procs).Int("PID", os.Getpid())
				}

				// Listen the app (with TLS Support)
				if cfg.App.TLS.Enable {
					log.Debug().Msg("TLS support was enabled.")

					if err := fiber.ListenTLS(cfg.App.Port, cfg.App.TLS.CertFile, cfg.App.TLS.KeyFile); err != nil {
						log.Error().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}

				go func() {
					if err := fiber.Listen(cfg.App.Port); err != nil {
						log.Error().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}()

				err := database.ConnectDatabase()
				if err != nil {
					log.Fatal().Err(err).Msg("An unknown error occurred when to connect the database!")
				} else {

					log.Info().Msg("Database connection was successful.")
				}

				err = database.MigrateModels()
				if err != nil {
					log.Fatal().Err(err).Msg("An unknown error occurred when trying to migrate the database!")
				} else {

					log.Info().Msg("Database migration was successful.")
				}

				err = database.SeedModels()
				if err != nil {
					log.Error().Err(err).Msg("An unknown error occurred when to seed the database!")
				} else {

					log.Info().Msg("Database seeding was successful.")
				}

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info().Msg("Gracefully shutting down the server...")
				if err := fiber.Shutdown(); err != nil {
					log.Panic().Err(err).Msg("An unknown error occurred when to shutdown the server!")
				}

				log.Info().Msg("Running cleanup tasks...")
				log.Info().Msg("1- Shutdown the database")
				err := database.ShutdownDatabase()
				if err != nil {
					log.Panic().Err(err).Msg("An unknown error occurred when to shutdown the database!")
				} else {
					log.Info().Msg(cfg.App.Name + " was successful shutdown.")
				}
				log.Info().Msg("\u001b[96msee you again👋\u001b[0m")

				return nil
			},
		},
	)
}
