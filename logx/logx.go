package logx

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	// LTrace is the Tracing level, everything is printed to the Log (it might get noisy).
	LTrace Level = iota
	// LDebug prints slightly less information, but can still be useful for debugging.
	LDebug
	// LInfo is the standard (and default) log Level, prints out handy information messages.
	LInfo
	// LWarning is exactly how it sounds, any event that occurs of notice, but is not major.
	LWarning
	// LError is similar to warning, except more serious.
	LError
	// LFatal means the program cannot continue when this event occurs. Normally the program will exit after this.
	LFatal

	stackDepth uint8 = 3
)

var (
	// Console is a pointer to the output that all the console Log structs will use. This can be set to any
	// type of stream that can be implemented as a Writer, including NUL.
	Console io.Writer = os.Stderr
	// Defaults is the default bitwise number that is used for new Log structs that are not
	// given an options number when created. This option number may be changed before running to affect
	// runtime functions.
	Defaults = log.Ldate | log.Ltime
)

// Level is an alias of a byte that represents the current Log level.
type Level uint8
type file struct {
	file string
	stream
}

// Log is an interface for any type of struct that supports standard Logging functions.
type Log interface {
	// SetLevel changes the current logging level of this Log instance.
	SetLevel(Level)
	// SetPrefix changes the current logging prefix of this Log instance.
	SetPrefix(string)
	// Info writes a information message to the Log instance.
	Info(string, ...interface{})
	// Error writes a error message to the Log instance.
	Error(string, ...interface{})
	// Fatal writes a fatal message to the Log instance. This function
	// will result in the program exiting with a non-zero error code after being called.
	Fatal(string, ...interface{})
	// Trace writes a tracing message to the Log instance.
	Trace(string, ...interface{})
	// Debug writes a debugging message to the Log instance.
	Debug(string, ...interface{})
	// Printf is kept for compatibility reasons. Printf statements are info logs.
	Printf(string, ...interface{})
	// Warning writes a warning message to the Log instance.
	Warning(string, ...interface{})
}
type stream struct {
	w   *log.Logger
	lvl Level
}
type handler interface {
	Level() Level
	Writer() *log.Logger
}

// NewConsole returns a console logger that uses the Console writer.
func NewConsole(l Level) Log {
	return NewWriterOptions(l, Defaults, Console)
}
func (l *stream) Level() Level {
	return l.lvl
}

// String returns the name of the current Level.
func (l Level) String() string {
	switch l {
	case LTrace:
		return "TRACE"
	case LDebug:
		return "DEBUG"
	case LInfo:
		return " INFO"
	case LWarning:
		return " WARN"
	case LError:
		return "ERROR"
	case LFatal:
		return "FATAL"
	}
	return ""
}
func (l *stream) SetLevel(n Level) {
	l.lvl = n
}
func (l *stream) SetPrefix(p string) {
	l.w.SetPrefix(p)
}
func (l *stream) Writer() *log.Logger {
	return l.w
}

// NewWriter returns a Log instance based on the Writer 'w' for the logging output.
func NewWriter(l Level, w io.Writer) Log {
	return NewWriterOptions(l, Defaults, w)
}

// NewConsoleOptions returns a console logger using the Console file for console output and
// allows specifying non-default Logging options.
func NewConsoleOptions(l Level, opts int) Log {
	return NewWriterOptions(l, opts, Console)
}

// NewFile will attempt to create a File backed Log instance that will write to file 's'.
// This function will truncate the file before starting a new Log. If you need to append to a existing log file.
// use the NewWriter function.
func NewFile(l Level, file string) (Log, error) {
	return NewFileOptions(l, Defaults, true, file)
}
func (l *stream) Info(m string, v ...interface{}) {
	writeToLog(l.w, l.lvl, LInfo, stackDepth, m, v)
}
func (l *stream) Error(m string, v ...interface{}) {
	writeToLog(l.w, l.lvl, LError, stackDepth, m, v)
}
func (l *stream) Fatal(m string, v ...interface{}) {
	writeToLog(l.w, l.lvl, LFatal, stackDepth, m, v)
	os.Exit(1)
}
func (l *stream) Trace(m string, v ...interface{}) {
	writeToLog(l.w, l.lvl, LTrace, stackDepth, m, v)
}
func (l *stream) Debug(m string, v ...interface{}) {
	writeToLog(l.w, l.lvl, LDebug, stackDepth, m, v)
}
func (l *stream) Printf(m string, v ...interface{}) {
	writeToLog(l.w, l.lvl, LInfo, stackDepth, m, v)
}
func (l *stream) Warning(m string, v ...interface{}) {
	writeToLog(l.w, l.lvl, LWarning, stackDepth, m, v)
}

// NewWriterOptions returns a Log instance based on the Writer 'w' for the logging output and
// allows specifying non-default Logging options.
func NewWriterOptions(l Level, opts int, w io.Writer) Log {
	return &stream{w: log.New(w, "", opts), lvl: l}
}

// NewFileOptions will attempt to create a File backed Log instance that will write to file specified.
// This function will truncate the file before starting a new Log. If you need to append to a existing log file.
// use the NewWriter function. This function allows specifying non-default Logging options.
func NewFileOptions(l Level, opts int, append bool, filepath string) (Log, error) {
	i := &file{file: filepath}
	i.lvl = l
	p := os.O_RDWR | os.O_CREATE
	if append {
		p |= os.O_TRUNC
	}
	w, err := os.OpenFile(filepath, p, 0644)
	if err != nil {
		return nil, err
	}
	i.w = log.New(w, "", opts)
	return i, nil
}
func writeToLog(i *log.Logger, c Level, l Level, d uint8, m string, v []interface{}) {
	if c > l {
		return
	}
	if stackDepth <= 2 {
		d = stackDepth
	}
	i.Output(int(d), fmt.Sprintf("%s: %s\n", l.String(), fmt.Sprintf(m, v...)))
}
