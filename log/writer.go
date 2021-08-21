package log

import "go.uber.org/zap/zapcore"

// Writer tht log writer interface
type Writer interface {
	Setup(wc WriterConfig) (zapcore.Core, error)
}

var (
	writers = make(map[string]Writer)
)

// RegisterWriter register the writer
func RegisterWriter(name string, w Writer) {
	writers[name] = w
}

// GetWriter get the writer by name
func GetWriter(name string) Writer {
	return writers[name]
}
