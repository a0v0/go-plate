package logger

import (
	"go_plate/internal/config"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger(cfg *config.Config) (*zerolog.Logger, error) {
	// pretty caller location
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	log.Logger = log.With().Caller().Logger()

	if !cfg.App.Production {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC1123Z}
		// output.FormatLevel = func(i interface{}) string {
		// 	return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		// }
		// output.FormatMessage = func(i interface{}) string {
		// 	return fmt.Sprintf("***%s****", i)
		// }
		// output.FormatFieldName = func(i interface{}) string {
		// 	return fmt.Sprintf("%s:", i)
		// }
		// output.FormatFieldValue = func(i interface{}) string {
		// 	return strings.ToUpper(fmt.Sprintf("%s", i))
		// }

		log.Logger = log.Output(output)
		return &log.Logger, nil
	}
	return &log.Logger, nil
}
