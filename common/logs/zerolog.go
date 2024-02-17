package logs

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	once sync.Once
)

func Init(isLocalEnv bool) {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	logger := zerolog.New(os.Stderr)

	if isLocalEnv {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	logger = logger.
		Level(zerolog.DebugLevel).
		With().
		Timestamp().
		Logger()

	// set global logger
	once.Do(func() {
		log.Logger = logger
		zerolog.DefaultContextLogger = &logger
	})
}
