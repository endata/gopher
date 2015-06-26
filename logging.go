package gopher

func Hello(msg string) {

	GetContext().Logger.Hello(msg)
}
