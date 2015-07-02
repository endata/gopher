package providers

import (
	"fmt"

	log "github.com/gopherlabs/gopher/vendor/_nuts/github.com/Sirupsen/logrus"
)

type LogProvider struct {
	log log.Logger
}

func (l LogProvider) Register(config map[string]interface{}) interface{} {
	l.log = *log.New()
	l.log.Formatter = &log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: config["FullTimestamp"].(bool),
	}
	return l
}

func (l LogProvider) GetKey() string {
	return "LOGGER"
}

func (l LogProvider) Info(msg string, args ...interface{}) {
	l.log.Info(sprintf(msg, args...))
}

func (l LogProvider) Debug(msg string, args ...interface{}) {
	log.Debug(sprintf(msg, args...))
}

func (l LogProvider) Warn(msg string, args ...interface{}) {
	log.Warn(sprintf(msg, args...))
}

func (l LogProvider) Error(msg string, args ...interface{}) {
	log.Error(sprintf(msg, args...))
}

// Calls os.Exit(1) after logging
func (l LogProvider) Fatal(msg string, args ...interface{}) {
	log.Fatal(sprintf(msg, args...))
}

// Calls panic() after logging
func (l LogProvider) Panic(msg string, args ...interface{}) {
	log.Panic(sprintf(msg, args...))
}

func sprintf(msg string, args ...interface{}) string {
	if args == nil {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}
