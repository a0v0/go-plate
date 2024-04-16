package config

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/num30/config"
)

type app = struct {
	Name        string        `default:"go_plate"`
	Port        string        `default:":8080"`
	PrintRoutes bool          `default:"false"`
	Prefork     bool          `default:"true"`
	Production  bool          `default:"false"`
	IdleTimeout time.Duration `default:"5"`

	MaxDisplayPictures int `default:"10"` // Maximum number of display pictures a user can have

	TLS struct {
		Enable   bool   `default:"false"`
		CertFile string `default:""`
		KeyFile  string `default:""`
	}
}

type email = struct {
	SMTP struct {
		Host     string `validate:"required"`
		Port     int    `validate:"required"`
		Password string `default:""`
		Secure   bool   `default:"false"`
	}
}

type db = struct {
	Postgres struct {
		DSN string `validate:"required"`
	}
}

type middleware = struct {
	Compress struct {
		Enable bool
		Level  compress.Level `default:"1"`
	}

	Recover struct {
		Enable bool
	}

	Monitor struct {
		Enable bool
		Path   string `default:"/monitor"`
	}

	Pprof struct {
		Enable bool
	}

	Limiter struct {
		Enable  bool
		Max     int           `default:"20"`
		ExpSecs time.Duration `default:"60"`
	}

	Filesystem struct {
		Enable bool
		Browse bool   `default:"true"`
		MaxAge int    `default:"3600"`
		Index  string `default:"index.html"`
		Root   string `default:"./storage/public"`
	}
}

type Config struct {
	App        app
	DB         db
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

// ParseAddr From https://github.com/gofiber/fiber/blob/master/helpers.go#L305.
func ParseAddr(raw string) (host, port string) {
	if i := strings.LastIndex(raw, ":"); i != -1 {
		return raw[:i], raw[i+1:]
	}
	return raw, ""
}
