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
	// LogConsoleFile is a pointer to the output that all the console Log structs will use. This can be set to any
	// type of stream that can be implemented as a Writer, including NUL.
	LogConsoleFile io.Writer = os.Stdout
	// LogDefaultOptions is the default bitwise number that is used for new Log structs that are not
	// given an options number when created. This option number may be changed before running to affect
	// runtime functions.
	LogDefaultOptions = log.Ldate | log.Ltime
)

// Stack is a type of Log that is an alias for an array where each Log
// function will affect each Log instance in the array.
type Stack []Log

// Level is an alias of a byte that represents the current Log level.
type Level uint8
type file struct {
	logFile string
	stream
}

// Log is an interface for any type of struct that supports standard Logging functions.
type Log interface {
	// Print is kept for compatibility reasons. Print statements are info logs.
	Print(string)
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
	logLevel  Level
	logWriter *log.Logger
}
type handler interface {
	Level() Level
	Writer() *log.Logger
}

// Add appends the specified Log 'l' the Stack array.
func (s *Stack) Add(l Log) {
	if l == nil {
		return
	}
	*s = append(*s, l)
}

// NewConsole returns a console logger that uses the LogConsoleFile writer.
func NewConsole(l Level) Log {
	return NewWriterOptions(l, LogDefaultOptions, LogConsoleFile)
}

// NewStack returns a Stack struct that contains the Log instances
// specified in the 'l' vardict.
func NewStack(l ...Log) *Stack {
	s := Stack(l)
	return &s
}
func (l *stream) Level() Level {
	return l.logLevel
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

// Print writes a information message to the Log instance.
func (s *Stack) Print(m string) {
	for i := range *s {
		if b, ok := (*s)[i].(handler); ok {
			writeToLog(b.Writer(), b.Level(), LInfo, stackDepth+1, m, nil)
		} else {
			(*s)[i].Info(m)
		}
	}
}
func (l *stream) Print(m string) {
	writeToLog(l.logWriter, l.logLevel, LDebug, stackDepth, m, nil)
}
func (l *stream) Writer() *log.Logger {
	return l.logWriter
}

// NewWriter returns a Log instance based on the Writer 'w' for the logging output.
func NewWriter(l Level, w io.Writer) Log {
	return NewWriterOptions(l, LogDefaultOptions, w)
}

// NewConsoleOptions returns a console logger using the LogConsoleFile file for console output and
// allows specifying non-default Logging options.
func NewConsoleOptions(l Level, opts int) Log {
	return NewWriterOptions(l, opts, LogConsoleFile)
}

// NewFile will attempt to create a File backed Log instance that will write to file 's'.
// This function will truncate the file before starting a new Log. If you need to append to a existing log file.
// use the NewWriter function.
func NewFile(l Level, file string) (Log, error) {
	return NewFileOptions(l, LogDefaultOptions, true, file)
}

// Info writes a information message to the Log instance.
func (s *Stack) Info(m string, v ...interface{}) {
	for i := range *s {
		if b, ok := (*s)[i].(handler); ok {
			writeToLog(b.Writer(), b.Level(), LInfo, stackDepth+1, m, v)
		} else {
			(*s)[i].Info(m, v...)
		}
	}
}

// Error writes a error message to the Log instance.
func (s *Stack) Error(m string, v ...interface{}) {
	for i := range *s {
		if b, ok := (*s)[i].(handler); ok {
			writeToLog(b.Writer(), b.Level(), LError, stackDepth+1, m, v)
		} else {
			(*s)[i].Error(m, v...)
		}
	}
}

// Fatal writes a fatal message to the Log instance. This function
// will result in the program exiting with a non-zero error code after being called.
func (s *Stack) Fatal(m string, v ...interface{}) {
	for i := range *s {
		if b, ok := (*s)[i].(handler); ok {
			writeToLog(b.Writer(), b.Level(), LFatal, stackDepth+1, m, v)
		} else {
			(*s)[i].Fatal(m, v...)
		}
	}
	os.Exit(1)
}

// Trace writes a tracing message to the Log instance.
func (s *Stack) Trace(m string, v ...interface{}) {
	for i := range *s {
		if b, ok := (*s)[i].(handler); ok {
			writeToLog(b.Writer(), b.Level(), LTrace, stackDepth+1, m, v)
		} else {
			(*s)[i].Trace(m, v...)
		}
	}
}

// Printf writes a information message to the Log instance.
func (s *Stack) Printf(m string, v ...interface{}) {
	for i := range *s {
		if b, ok := (*s)[i].(handler); ok {
			writeToLog(b.Writer(), b.Level(), LInfo, stackDepth+1, m, v)
		} else {
			(*s)[i].Info(m, v...)
		}
	}
}

// Debug writes a debugging message to the Log instan
func (s *Stack) Debug(m string, v ...interface{}) {
	for i := range *s {
		if b, ok := (*s)[i].(handler); ok {
			writeToLog(b.Writer(), b.Level(), LDebug, stackDepth+1, m, v)
		} else {
			(*s)[i].Debug(m, v...)
		}
	}
}
func (l *stream) Info(m string, v ...interface{}) {
	writeToLog(l.logWriter, l.logLevel, LInfo, stackDepth, m, v)
}
func (l *stream) Error(m string, v ...interface{}) {
	writeToLog(l.logWriter, l.logLevel, LError, stackDepth, m, v)
}
func (l *stream) Fatal(m string, v ...interface{}) {
	writeToLog(l.logWriter, l.logLevel, LFatal, stackDepth, m, v)
	os.Exit(1)
}
func (l *stream) Trace(m string, v ...interface{}) {
	writeToLog(l.logWriter, l.logLevel, LTrace, stackDepth, m, v)
}
func (l *stream) Debug(m string, v ...interface{}) {
	writeToLog(l.logWriter, l.logLevel, LDebug, stackDepth, m, v)
}
func (l *stream) Printf(m string, v ...interface{}) {
	writeToLog(l.logWriter, l.logLevel, LInfo, stackDepth, m, v)
}

// Warning writes a warning message to the Log instance.
func (s *Stack) Warning(m string, v ...interface{}) {
	for i := range *s {
		if b, ok := (*s)[i].(handler); ok {
			writeToLog(b.Writer(), b.Level(), LWarning, stackDepth+1, m, v)
		} else {
			(*s)[i].Warning(m, v...)
		}
	}
}
func (l *stream) Warning(m string, v ...interface{}) {
	writeToLog(l.logWriter, l.logLevel, LWarning, stackDepth, m, v)
}

// NewWriterOptions returns a Log instance based on the Writer 'w' for the logging output and
// allows specifying non-default Logging options.
func NewWriterOptions(l Level, opts int, w io.Writer) Log {
	return &stream{logLevel: l, logWriter: log.New(w, "", opts)}
}

// NewFileOptions will attempt to create a File backed Log instance that will write to file specified.
// This function will truncate the file before starting a new Log. If you need to append to a existing log file.
// use the NewWriter function. This function allows specifying non-default Logging options.
func NewFileOptions(l Level, opts int, append bool, filepath string) (Log, error) {
	i := &file{logFile: filepath}
	i.logLevel = l
	p := os.O_RDWR | os.O_CREATE
	if append {
		p |= os.O_TRUNC
	}
	w, err := os.OpenFile(filepath, p, 0644)
	if err != nil {
		return nil, err
	}
	i.logWriter = log.New(w, "", opts)
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
