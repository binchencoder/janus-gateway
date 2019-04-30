package glog

import (
	"context"
)

const (
	traceFormat = "(trace: %s) "
)

// Switch is a boolean type with a method Context.
// See the documentation of V for more information.
type Switch bool

// Logger is a struct contains MDC for log methods. It provides functions Info,
// Warning, Error, Fatal, plus formatting variants such as Infof.
// See the documentation of github.com/golang/glog for more information.
type Logger struct {
	Ctx      context.Context
	filename interface{}
}

// Verbose is a variant type of Logger. It provides less log methods.
// If the value is nil, it doesn't record log.
type Verboser Logger

// One may write either
//	if log.V(2) { log.Context(ctx).Info("log this") }
// or
//	glog.V(2).Context(ctx).Info("log this")
// The second form is shorter but the first is cheaper if logging is off because it does
// not evaluate its arguments.
//
// See the documentation of github.com/golang/glog.V for more information.
func VL(level Level) Switch {
	return Switch(VDepth(1, level))
}

// Context return a Logger with MDC from ctx.
func Context(ctx context.Context, filename interface{}) *Logger {
	if ctx == nil {
		ctx = context.Background()
	}
	if filename == nil {
		filename = FileName{}
	}
	return &Logger{Ctx: ctx, filename: filename}
}

// Context return a Verboser with MDC from ctx.
func (s Switch) Context(ctx context.Context, filename interface{}) *Verboser {
	if s {
		if ctx == nil {
			ctx = context.Background()
		}
		if filename == nil {
			filename = FileName{}
		}
		return &Verboser{Ctx: ctx, filename: filename}
	}

	return nil
}

func (l *Logger) Flush() {
	Flush()
}
