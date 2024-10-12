package logger

import (
	"github.com/krabiworld/lamarr/internal/cfg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strconv"
)

func Init() {
	// Enable pretty logging if debug mode is enabled
	if cfg.Get().Debug {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	// Enable unix time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Enable caller
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	log.Logger = log.With().Caller().Logger()

	// Set log level
	logLevel := cfg.Get().LogLevel

	if logLevel == "" {
		log.Fatal().Msg("Set LOG_LEVEL variable")
	} else {
		var zeroLogLevel zerolog.Level
		if err := zeroLogLevel.UnmarshalText([]byte(logLevel)); err != nil {
			log.Fatal().Err(err).Msg("Failed to unmarshal log level")
		}

		zerolog.SetGlobalLevel(zeroLogLevel)
	}

	log.Info().Msg("Logger successfully initialized")
}
