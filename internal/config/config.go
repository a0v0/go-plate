package config

import (
	"log"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/num30/config"
)

type app = struct {
	Name        string        `default:"go_plate" envvar:"APP_NAME"`
	Port        string        `default:":8080" envvar:"APP_PORT"`
	PrintRoutes bool          `default:"false" envvar:"APP_PRINT_ROUTES"`
	Prefork     bool          `default:"true" envvar:"APP_PREFORK"`
	Production  bool          `default:"false" envvar:"APP_PRODUCTION"`
	IdleTimeout time.Duration `default:"5" envvar:"APP_IDLE_TIMEOUT"`

	TLS struct {
		Enable   bool   `default:"false" envvar:"APP_TLS_ENABLE"`
		CertFile string `default:"" envvar:"APP_TLS_CERT_FILE"`
		KeyFile  string `default:"" envvar:"APP_TLS_KEY_FILE"`
	}
}

type email = struct {
	SMTP struct {
		Host     string `validate:"required" envvar:"EMAIL_SMTP_HOST"`
		Port     int    `validate:"required" envvar:"EMAIL_SMTP_PORT"`
		Password string `validate:"required" envvar:"EMAIL_SMTP_PASSWORD"`
		Secure   bool   `default:"false" envvar:"EMAIL_SMTP_SECURE"`
	}
}

type db = struct {
	Postgres struct {
		User     string `envvar:"DB_PG_USER"`
		Password string `envvar:"DB_PG_PASSWORD"`
		Host     string `envvar:"DB_PG_HOST"`
		Port     string `envvar:"DB_PG_PORT"`
		DBName   string `envvar:"DB_PG_DBNAME"`
		SSLMode  string `envvar:"DB_PG_SSLMODE"`
	}
}

type middleware = struct {
	Compress struct {
		Enable bool           `default:"true" envvar:"MIDDLEWARE_COMPRESS_ENABLE"`
		Level  compress.Level `default:"1" envvar:"MIDDLEWARE_COMPRESS_LEVEL"`
	}

	Recover struct {
		Enable bool `default:"true" envvar:"MIDDLEWARE_RECOVER_ENABLE"`
	}

	Monitor struct {
		Enable bool   `default:"true" envvar:"MIDDLEWARE_MONITOR_ENABLE"`
		Path   string `default:"/monitor" envvar:"MIDDLEWARE_MONITOR_PATH"`
	}

	Pprof struct {
		Enable bool `default:"true" envvar:"MIDDLEWARE_PPROF_ENABLE"`
	}

	Limiter struct {
		Enable  bool          `default:"true" envvar:"MIDDLEWARE_LIMITER_ENABLE"`
		Max     int           `default:"20" envvar:"MIDDLEWARE_LIMITER_MAX"`
		ExpSecs time.Duration `default:"60" envvar:"MIDDLEWARE_LIMITER_EXPSECS"`
	}

	Filesystem struct {
		Enable bool   `default:"true" envvar:"MIDDLEWARE_FILESYSTEM_ENABLE"`
		Browse bool   `default:"true" envvar:"MIDDLEWARE_FILESYSTEM_BROWSE"`
		MaxAge int    `default:"3600" envvar:"MIDDLEWARE_FILESYSTEM_MAXAGE"`
		Index  string `default:"index.html" envvar:"MIDDLEWARE_FILESYSTEM_INDEX"`
		Root   string `default:"./storage/public" envvar:"MIDDLEWARE_FILESYSTEM_ROOT"`
	}
}

type Config struct {
	App        app
	Database   db
	Email      email
	Middleware middleware
}

func NewConfig() *Config {
	var c Config
	err := config.NewConfReader("config").WithSearchDirs("./config", ".").Read(&c)
	if err != nil && !fiber.IsChild() {
		log.Fatal(err, "Error reading config")
	}

	return &c
}
