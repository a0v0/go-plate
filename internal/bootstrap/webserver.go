package bootstrap

import (
	"context"
	"os"
	"runtime"
	"strings"
	"time"

	"go_plate/app/middleware"
	"go_plate/app/router"
	"go_plate/internal/bootstrap/database"
	"go_plate/pkg/config"
	"go_plate/pkg/response"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	futils "github.com/gofiber/fiber/v2/utils"
	"go.uber.org/fx"
	"go.uber.org/zap"
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

func Start(lifecycle fx.Lifecycle, cfg *config.Config, fiber *fiber.App, router *router.Router, middlewares *middleware.Middleware, database *database.Database, log *zap.Logger) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Register middlewares & routes
				middlewares.Register()
				router.Register()

				// Custom Startup Messages
				host, port := config.ParseAddr(cfg.App.Port)
				if host == "" {
					if fiber.Config().Network == "tcp6" {
						host = "[::1]"
					} else {
						host = "0.0.0.0"
					}
				}

				// ASCII Art
				ascii, err := os.ReadFile("./storage/ascii_art.txt")
				if err != nil {
					log.Debug("An unknown error occurred when to print ASCII art!", zap.Error(err))
				}

				for _, line := range strings.Split(futils.UnsafeString(ascii), "\n") {
					log.Info(line)
				}

				// Information message
				log.Info(fiber.Config().AppName + " is running at the moment!")

				// Debug informations
				if !cfg.App.Production {
					prefork := "Enabled"
					procs := runtime.GOMAXPROCS(0)
					if !cfg.App.Prefork {
						procs = 1
						prefork = "Disabled"
					}

					log.Debug("",
						zap.String("Version", "-"),
						zap.String("Host", host),
						zap.String("Port", port),
						zap.String("Prefork", prefork),
						zap.Uint32("Handlers", fiber.HandlersCount()),
						zap.Int("Processes", procs),
						zap.Int("PID", os.Getpid()),
					)
				}

				// Listen the app (with TLS Support)
				if cfg.App.TLS.Enable {
					log.Debug("TLS support was enabled.")

					if err := fiber.ListenTLS(cfg.App.Port, cfg.App.TLS.CertFile, cfg.App.TLS.KeyFile); err != nil {

					}
				}

				go func() {
					if err := fiber.Listen(cfg.App.Port); err != nil {
						log.Error("An unknown error occurred when to run server!", zap.Error(err))
					}
				}()

				database.ConnectDatabase()
				database.MigrateModels()
				database.SeedModels()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info("Shutting down the app...")
				if err := fiber.Shutdown(); err != nil {
					log.Panic("", zap.Error(err))
				}
				log.Info("Running cleanup tasks...")
				log.Info("1- Shutdown the database")
				database.ShutdownDatabase()
				log.Info(cfg.App.Name + " was successful shutdown.")
				log.Info("\u001b[96msee you again👋\u001b[0m")

				return nil
			},
		},
	)
}
