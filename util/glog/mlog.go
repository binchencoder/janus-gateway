package glog

import (
	"fmt"

	"golang.org/x/net/context"

	"binchencoder.com/letsgo/trace"
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

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *Logger) Info(args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}

	InfoDepth(2, l.filename, fmt.Sprint(args...))
}

// InfoDepth acts as Info but uses depth to determine which call frame to log.
// InfoDepth(0, "msg") is the same as Info("msg").
func (l *Logger) InfoDepth(depth int, args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}

	InfoDepth(depth+1, l.filename, fmt.Sprint(args...))
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (l *Logger) Infoln(args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}
	args = append(args, "\n")

	InfoDepth(2, l.filename, fmt.Sprint(args...))
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *Logger) Infof(format string, args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		format = traceFormat + format
		args = append([]interface{}{tid}, args...)
	}

	InfoDepth(2, l.filename, fmt.Sprintf(format, args...))
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *Logger) Warning(args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}

	WarningDepth(2, l.filename, fmt.Sprint(args...))
}

// WarningDepth acts as Warning but uses depth to determine which call frame to log.
// WarningDepth(0, "msg") is the same as Warning("msg").
func (l *Logger) WarningDepth(depth int, args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}

	WarningDepth(depth+1, l.filename, fmt.Sprint(args...))
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (l *Logger) Warningln(args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}
	args = append(args, "\n")

	WarningDepth(2, l.filename, fmt.Sprint(args...))
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *Logger) Warningf(format string, args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		format = traceFormat + format
		args = append([]interface{}{tid}, args...)
	}

	WarningDepth(2, l.filename, fmt.Sprintf(format, args...))
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *Logger) Error(args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}

	ErrorDepth(2, l.filename, fmt.Sprint(args...))
}

// ErrorDepth acts as Error but uses depth to determine which call frame to log.
// ErrorDepth(0, "msg") is the same as Error("msg").
func (l *Logger) ErrorDepth(depth int, args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}

	ErrorDepth(depth+1, l.filename, fmt.Sprint(args...))
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (l *Logger) Errorln(args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}
	args = append(args, "\n")

	ErrorDepth(2, l.filename, fmt.Sprint(args...))
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *Logger) Errorf(format string, args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		format = traceFormat + format
		args = append([]interface{}{tid}, args...)
	}

	ErrorDepth(2, l.filename, fmt.Sprintf(format, args...))
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *Logger) Fatal(args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}

	FatalDepth(2, l.filename, fmt.Sprint(args...))
}

// FatalDepth acts as Fatal but uses depth to determine which call frame to log.
// FatalDepth(0, "msg") is the same as Fatal("msg").
func (l *Logger) FatalDepth(depth int, args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}

	FatalDepth(depth+1, l.filename, fmt.Sprint(args...))
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func (l *Logger) Fatalln(args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}
	args = append(args, "\n")

	FatalDepth(2, l.filename, fmt.Sprint(args...))
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		format = traceFormat + format
		args = append([]interface{}{tid}, args...)
	}

	FatalDepth(2, l.filename, fmt.Sprintf(format, args...))
}

// Exit logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func (l *Logger) Exit(args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}

	ExitDepth(2, l.filename, fmt.Sprint(args...))
}

// ExitDepth acts as Exit but uses depth to determine which call frame to log.
// ExitDepth(0, "msg") is the same as Exit("msg").
func (l *Logger) ExitDepth(depth int, args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}

	ExitDepth(depth+1, l.filename, fmt.Sprint(args...))
}

// Exitln logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
func (l *Logger) Exitln(args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		args = append([]interface{}{fmt.Sprintf(traceFormat, tid)}, args...)
	}
	args = append(args, "\n")

	ExitDepth(2, l.filename, fmt.Sprint(args...))
}

// Exitf logs to the FATAL, ERROR, WARNING, and INFO logs, then calls os.Exit(1).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func (l *Logger) Exitf(format string, args ...interface{}) {
	if tid := trace.GetTraceIdOrEmpty(l.Ctx); tid != "" {
		format = traceFormat + format
		args = append([]interface{}{tid}, args...)
	}

	ExitDepth(2, l.filename, fmt.Sprintf(format, args...))
}

func (l *Logger) Flush() {
	Flush()
}

// Info is equivalent to the Logger.Info function, guarded by the value of v.
// See the documentation of V for usage.
func (v *Verboser) InfoV(args ...interface{}) {
	if v != nil {
		(*Logger)(v).InfoDepth(1, args...)
	}
}

// Infoln is equivalent to the Logger.Infoln function, guarded by the value of v.
// See the documentation of V for usage.
func (v *Verboser) InfoVln(args ...interface{}) {
	if v != nil {
		args = append(args, "\n")
		(*Logger)(v).InfoDepth(1, args...)
	}
}

// Infof is equivalent to the Logger.Infof function, guarded by the value of v.
// See the documentation of V for usage.
func (v *Verboser) InfoVf(format string, args ...interface{}) {
	if v != nil {
		(*Logger)(v).InfoDepth(1, fmt.Sprintf(format, args...))
	}
}
