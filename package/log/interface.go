package log

type Log interface {
	Info(message string, attr interface{})
	Warning(message string, _err error, attr interface{})
	Error(_err error, message string, attr interface{})
	Panic(_err error, message string, attr interface{})
	PanicOnError(_err error, message string, attr interface{})
}
