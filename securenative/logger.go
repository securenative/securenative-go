package securenative

type LoggerInterface interface {
	Info(msg string)
	Debug(msg string)
	Warning(msg string)
	Error(msg string)
}

type Logger struct {
	// TODO implement me
}

func NewLogger(logLevel string) *Logger {
	panic("implement me")
}

func (l *Logger) Info(msg string) {

}

func (l *Logger) Debug(msg string) {

}

func (l *Logger) Warning(msg string) {

}

func (l *Logger)  Error(msg string) {

}