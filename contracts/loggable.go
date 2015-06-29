package contracts

type Loggable interface {
	NewLogger() Loggable

	Info(msg string, args ...interface{})

	Debug(msg string, args ...interface{})

	Warn(msg string, args ...interface{})

	Error(msg string, args ...interface{})

	Fatal(msg string, args ...interface{})

	Panic(msg string, args ...interface{})
}
