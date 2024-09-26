package xlogger

import (
	"github.com/rs/zerolog"
)

type Logger interface {
	/*
		Fields is a helper function to use a map or slice to set fields using type assertion.
		Only map[string]any and []any are accepted. []any must alternate string keys and arbitrary values,
		and extraneous ones are ignored. i.e:

		With Request HTTP Context:

			first := "first value"
			second := "second value"
			xlogger.RequestContext(ctx).Info("hello", "first", first, "second", second)
			xlogger.RequestContext(ctx).Info(fmt.Sprintf("hello %s", "world!"), "first", first, "second", second)
	*/
	Trace(msg string, fields ...any)
	Debug(msg string, fields ...any)
	Info(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Error(msg string, fields ...any)
	Fatal(msg string, fields ...any)
	Panic(msg string, fields ...any)
}

type ZeroLogger struct {
	log *zerolog.Logger
}

func NewZeroLogger(log *zerolog.Logger) Logger {
	return &ZeroLogger{log}
}

// attachFields directly to the zerolog.Event object without creating a new map
func (zl *ZeroLogger) attachFields(event *zerolog.Event, fields []any) *zerolog.Event {
	for i := 0; i < len(fields); i += 2 {
		if i+1 < len(fields) {
			key, ok := fields[i].(string)
			if ok {
				event = event.Interface(key, fields[i+1])
			}
		}
	}
	return event
}

func (zl *ZeroLogger) Trace(msg string, fields ...any) {
	event := zl.log.Trace()
	event = zl.attachFields(event, fields)
	event.Msg(msg)
}

func (zl *ZeroLogger) Debug(msg string, fields ...any) {
	event := zl.log.Debug()
	event = zl.attachFields(event, fields)
	event.Msg(msg)
}

func (zl *ZeroLogger) Info(msg string, fields ...any) {
	event := zl.log.Info()
	event = zl.attachFields(event, fields)
	event.Msg(msg)
}

func (zl *ZeroLogger) Warn(msg string, fields ...any) {
	event := zl.log.Warn()
	event = zl.attachFields(event, fields)
	event.Msg(msg)
}

func (zl *ZeroLogger) Error(msg string, fields ...any) {
	event := zl.log.Error()
	event = zl.attachFields(event, fields)
	event.Msg(msg)
}

func (zl *ZeroLogger) Fatal(msg string, fields ...any) {
	event := zl.log.Fatal()
	event = zl.attachFields(event, fields)
	event.Msg(msg)
}

func (zl *ZeroLogger) Panic(msg string, fields ...any) {
	event := zl.log.Panic()
	event = zl.attachFields(event, fields)
	event.Msg(msg)
}
