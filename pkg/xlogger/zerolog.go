package xlogger

import (
	"fmt"

	"github.com/rs/zerolog"
)

type Logger interface {
	/*
		Fields is a helper function to use a map or slice to set fields using type assertion.
		[]any must alternate string keys and arbitrary values, and extraneous ones are ignored. i.e:

		With Request HTTP Context:

			first := "first value"
			second := "second value"
			xlogger.FromReqCtx(ctx).Trace("hello", "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Trace(xlogger.Msgf("hello %s", "world!"), "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Trace("oh snap! got error", "error", fmt.Sprint("%+v", err))
	*/
	Trace(msg string, fields ...any)

	/*
		Fields is a helper function to use a map or slice to set fields using type assertion.
		[]any must alternate string keys and arbitrary values, and extraneous ones are ignored. i.e:

		With Request HTTP Context:

			first := "first value"
			second := "second value"
			xlogger.FromReqCtx(ctx).Debug("hello", "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Debug(xlogger.Msgf("hello %s", "world!"), "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Debug("oh snap! got error", "error", fmt.Sprint("%+v", err))
	*/
	Debug(msg string, fields ...any)

	/*
		Fields is a helper function to use a map or slice to set fields using type assertion.
		[]any must alternate string keys and arbitrary values, and extraneous ones are ignored. i.e:

		With Request HTTP Context:

			first := "first value"
			second := "second value"
			xlogger.FromReqCtx(ctx).Info("hello", "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Info(xlogger.Msgf("hello %s", "world!"), "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Info("oh snap! got error", "error", fmt.Sprint("%+v", err))
	*/
	Info(msg string, fields ...any)

	/*
		Fields is a helper function to use a map or slice to set fields using type assertion.
		[]any must alternate string keys and arbitrary values, and extraneous ones are ignored. i.e:

		With Request HTTP Context:

			first := "first value"
			second := "second value"
			xlogger.FromReqCtx(ctx).Warn("hello", "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Warn(xlogger.Msgf("hello %s", "world!"), "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Warn("oh snap! got error", "error", fmt.Sprint("%+v", err))
	*/
	Warn(msg string, fields ...any)

	/*
		Fields is a helper function to use a map or slice to set fields using type assertion.
		[]any must alternate string keys and arbitrary values, and extraneous ones are ignored. i.e:

		With Request HTTP Context:

			first := "first value"
			second := "second value"
			xlogger.FromReqCtx(ctx).Error("hello", "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Error(xlogger.Msgf("hello %s", "world!"), "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Error("oh snap! got error", "error", fmt.Sprint("%+v", err))
	*/
	Error(msg string, fields ...any)

	/*
		Fields is a helper function to use a map or slice to set fields using type assertion.
		[]any must alternate string keys and arbitrary values, and extraneous ones are ignored. i.e:

		With Request HTTP Context:

			first := "first value"
			second := "second value"
			xlogger.FromReqCtx(ctx).Fatal("hello", "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Fatal(xlogger.Msgf("hello %s", "world!"), "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Fatal("oh snap! got error", "error", fmt.Sprint("%+v", err))
	*/
	Fatal(msg string, fields ...any)

	/*
		Fields is a helper function to use a map or slice to set fields using type assertion.
		[]any must alternate string keys and arbitrary values, and extraneous ones are ignored. i.e:

		With Request HTTP Context:

			first := "first value"
			second := "second value"
			xlogger.FromReqCtx(ctx).Panic("hello", "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Panic(xlogger.Msgf("hello %s", "world!"), "first", first, "second", second)
			xlogger.FromReqCtx(ctx).Panic("oh snap! got error", "error", fmt.Sprint("%+v", err))
	*/
	Panic(msg string, fields ...any)
}

func Msgf(msg string, args ...any) string {
	return fmt.Sprintf(msg, args...)
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
