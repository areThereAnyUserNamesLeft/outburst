package outburst

type Level int

const (
	Trace Level = iota
	Debug
	Info
	Warn
	Error
	Fatal
	Panic
)

func (l Level) String() string {
	return [...]string{"Trace", "Debug", "Info", "Warn", "Error", "Fatal", "Panic"}[l]
}
