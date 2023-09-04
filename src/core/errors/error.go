package errors

const (
	InfoLevel  = iota
	DebugLevel = iota
	WarnLevel  = iota
	ErrorLevel = iota
	FatalLevel = iota
)

type Error interface {
	error
	File() string
	Line() int
	LogLevel() int
	FriendlyMessage() string
	Equals(err Error) bool
}