package logger

import (
	"github.com/krabiworld/lamarr/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strconv"
)

func Init() {
	// Enable pretty logging if debug mode is enabled
	if config.Get().Debug {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	// Enable unix time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Change caller format
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	// Enable caller and pretty logs
	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Set log level
	logLevel := config.Get().LogLevel

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
