package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/DarManuals/clean-arch/config"
)

// simple logger example
type Logger interface {
	Info(msg string)
	Error(msg string, err error)
}

func New(cfg config.Config) Logger {
	return logger{
		prefix: cfg.ServiceName,
		l:      log.New(os.Stdout, "", log.LstdFlags),
	}
}

type logger struct {
	prefix string
	l      *log.Logger
}

func (l logger) Info(msg string) {
	l.l.Println(fmt.Sprintf("%s: %s: %s", l.prefix, "INFO", msg))
}

func (l logger) Error(msg string, err error) {
	l.l.Println(fmt.Sprintf("%s: %s: %s: %v", l.prefix, "ERROR", msg, err))
}
