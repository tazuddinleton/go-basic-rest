package logger

import (
	"fmt"
	"io"
	dlog "log"
	"os"

	"github.com/rs/zerolog"
	"github.com/tazuddinleton/go/basicrest/conf"
)

var Log zerolog.Logger

// Configure configures the logger and returns handle to log file if loging to file is enabled
func Configure(conf conf.Configuration) *os.File {

	l := conf.Logging()
	var writers []io.Writer
	if l.ConsoleEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if l.File == "" {
		dlog.Fatal("Log file cannot be found")
	}

	logFile, err := os.Create(l.File)
	if err != nil {
		dlog.Fatal(fmt.Sprintf("Cannot open file %v", l.File))
	}

	writers = append(writers, logFile)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if l.Level == "DEBUG" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	mw := io.MultiWriter(writers...)
	Log = zerolog.New(mw).With().Timestamp().Logger()
	Log.Info().Msg("logger created")
	return logFile
}
