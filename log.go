// Copyright 2021 - 2022 PurpleSec Team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package logx

import (
	"io"
	"runtime"
	"sync"
	"time"
)

const (
	// Trace is the Tracing log level, everything logged is passed to the Log
	// (it might get noisy).
	Trace Level = iota
	// Debug prints slightly less information than Trace, but is primarily useful
	// for debugging messages.
	Debug
	// Info is the standard informational log Level, this can be used to print
	// out handy information messages.
	Info
	// Warning is exactly how it sounds, any event that occurs of notice, but
	// is not major. This is the default logging
	// level.
	Warning
	// Error is similar to the Warning level, except more serious.
	Error
	// Fatal means the program cannot continue when this event occurs. Normally
	// the program will exit after this.
	//
	// Set 'logx.FatalExits' to 'false' to disable exiting when a Fatal log entry
	// is triggered.
	Fatal
	// Panic is a special level only used by the 'Panic*' functions. This will
	// act similar to the 'Fatal' level but will trigger a Go 'panic()' call
	// instead once log writing is complete.
	//
	// The 'logx.FatalExits' variable does not apply to this level.
	Panic
	// Print is a special level only used for LogWriters. This level instructs
	// the LogWriter to use it's inbuilt print level for the specified message.
	//
	// Attempting to use this out of this context will fail.
	Print
	invalidLevel
)

const (
	// FlagDate instructs the Logger to include the local date in the logging output.
	//
	// Same as 'log.Ldate' or 'Ldate'.
	FlagDate uint8 = 1 << iota
	// FlagTime instructs the Logger to include the local time in the logging output.
	//
	// Same as 'log.Ltime' or 'Ltime'.
	FlagTime
	// FlagMicroseconds instructs the Logger to include the local time (in milliseconds)
	// in the logging output. Implies 'FlagTime'.
	//
	// Same as 'log.Lmicroseconds' or 'Lmicroseconds'.
	FlagMicroseconds
	// FlagFileLong instructs the Logger to include the full source file name and
	// line number in the logging output.
	//
	// Same as 'log.Llongfile' or 'Llongfile'.
	FlagFileLong
	// FlagFileShort instructs the Logger to include the shortened source file
	// name (file name only) and line number in the logging output. Overrides
	// 'FlagFileLong'.
	//
	// Same as 'log.Lshortfile' or 'Lshortfile'.
	FlagFileShort
	// FlagTimeUTC instructs the Logger to use the UTC time zone as the log output
	// date/time instead of the local time zone. Only takes affect if 'FlagDate',
	// 'FlagTime' or 'FlagMicroseconds' is implied.
	//
	// Same as 'log.LUTC' or 'LUTC'.
	FlagTimeUTC

	// FlagStandard is the standard logging flags used as the setting for the default
	// logger. This is the same as 'FlagDate | FlagTime'.
	//
	// Same as 'log.LstdFlags' or 'LstdFlags'.
	FlagStandard = FlagDate | FlagTime
)

// Flag values that mirror the ones in the 'log' package.
//
// NOTE: 'Lmsgprefix' has no effect.
const (
	Ldate               = FlagDate
	Ltime               = FlagTime
	Lmicroseconds       = FlagMicroseconds
	Llongfile           = FlagFileLong
	Lshortfile          = FlagFileShort
	LUTC                = FlagTimeUTC
	Lmsgprefix    uint8 = 0
	LstdFlags           = FlagStandard
)

// FatalExits is a boolean setting that determines if a call to Fatal or LogFatal
// will exit the program using 'os.Exit(1)'. If this is set to false, a call to
// Fatal or LogFatal will continue program execution after being called.
//
// The default value is true.
var FatalExits = true

// Level is an alias of a byte that represents the current Log level.
type Level uint8

// Log is an interface for any type of struct that supports standard Logging
// functions.
type Log interface {
	// SetLevel changes the current logging level of this Log.
	SetLevel(Level)
	// SetPrefix changes the current logging prefix of this Log.
	SetPrefix(string)
	// SetPrintLevel sets the logging level used when 'Print*' statements are called.
	SetPrintLevel(Level)
	// Print writes a message to the logger.
	//
	// The function arguments are similar to 'fmt.Sprint' and 'fmt.Print'. The only
	// argument is a vardict of interfaces that can be used to output a string value.
	//
	// This function is affected by the setting of 'SetPrintLevel'. By default,
	// this will print as an 'Info' logging message.
	Print(...interface{})
	// Panic writes a panic message to the logger.
	//
	// This function will result in the program exiting with a Go 'panic()' after
	// being called. The function arguments are similar to 'fmt.Sprint' and 'fmt.Print.'
	// The only argument is a vardict of interfaces that can be used to output a
	// string value.
	Panic(...interface{})
	// Println writes a message to the logger.
	//
	// The function arguments are similar to fmt.Sprintln and fmt.Println. The only
	// argument is a vardict of interfaces that can be used to output a string value.
	//
	// This function is affected by the setting of 'SetPrintLevel'. By default,
	// this will print as an 'Info' logging message.
	Println(...interface{})
	// Panicln writes a panic message to the logger.
	//
	// This function will result in the program exiting with a Go 'panic()' after
	// being called. The function arguments are similar to 'fmt.Sprintln' and
	// 'fmt.Println'. The only argument is a vardict of interfaces that
	// can be used to output a string value.
	Panicln(...interface{})
	// Info writes an informational message to the logger.
	//
	// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
	// first argument is a string that can contain formatting characters. The second
	// argument is a vardict of interfaces that can be omitted or used in the supplied
	// format string.
	Info(string, ...interface{})
	// Error writes an error message to the logger.
	//
	// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
	// first argument is a string that can contain formatting characters. The second
	// argument is a vardict of interfaces that can be omitted or used in the supplied
	// format string.
	Error(string, ...interface{})
	// Fatal writes a fatal message to the logger.
	//
	// This function will result in the program exiting with a non-zero error code
	// after being called, unless the 'logx.FatalExits' setting is 'false'. The
	// function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The first
	// argument is a string that can contain formatting characters. The second argument
	// is a vardict of interfaces that can be omitted or used in the supplied format
	// string.
	Fatal(string, ...interface{})
	// Trace writes a tracing message to the logger.
	//
	// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
	// first argument is a string that can contain formatting characters. The second
	// argument is a vardict of interfaces that can be omitted or used in the supplied
	// format string.
	Trace(string, ...interface{})
	// Debug writes a debugging message to the logger.
	//
	// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
	// first argument is a string that can contain formatting characters. The second
	// argument is a vardict of interfaces that can be omitted or used in the supplied
	// format string.
	Debug(string, ...interface{})
	// Printf writes a message to the logger.
	//
	// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
	// first argument is a string that can contain formatting characters. The second
	// argument is a vardict of interfaces that can be omitted or used in the supplied
	// format string.
	//
	// This function is affected by the setting of 'SetPrintLevel'. By default,
	// this will print as an 'Info' logging message.
	Printf(string, ...interface{})
	// Panicf writes a panic message to the logger.
	//
	// This function will result in the program exiting with a Go 'panic()' after
	// being called. The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'.
	// The first argument is a string that can contain formatting characters. The
	// second argument is a vardict of interfaces that can be omitted or used in
	// the supplied format string.
	Panicf(string, ...interface{})
	// Warning writes a warning message to the logger.
	//
	// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
	// first argument is a string that can contain formatting characters. The second
	// argument is a vardict of interfaces that can be omitted or used in the supplied
	// format string.
	Warning(string, ...interface{})
}
type logger struct {
	m sync.Mutex
	w io.Writer
	p []byte
	f uint8
}

// LogWriter is an interface that is used inline with logging operations. This
// interface defines a single function 'Log' which takes a Level to log, the
// additional stack depth (can be zero), and the message and optional arguments
// to be logged.
//
// This function is used as a "quick-logging" helper and is preferred by the
// Multi logging struct.
//
// This funcion must ONLY log and not preform any other operations such as exiting
// on a Fatal or Panic logging level.
//
// Use the higher level 'Fatal' and 'Panic' functions for those additional operations.
//
// The higher level calls may use this function to simplify logging and comply
// with the logx Multi-logger.
type LogWriter interface {
	Log(Level, int, string, ...interface{})
}

// String returns the textual name of the Level.
func (l Level) String() string {
	switch l {
	case Trace:
		return "TRACE"
	case Debug:
		return "DEBUG"
	case Info:
		return " INFO"
	case Warning:
		return " WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	case Panic:
		return "PANIC"
	}
	return "INVAL"
}
func (l *logger) SetPrefix(p string) {
	if l.m.Lock(); len(p) == 0 {
		l.p = nil
	} else {
		l.p = []byte(p)
	}
	l.m.Unlock()
}
func itoa(b *[28]byte, p, i, w int) int {
	var (
		o [20]byte
		s = 19
	)
	for i >= 10 || w > 1 {
		w--
		q := i / 10
		o[s] = byte('0' + i - q*10)
		s--
		i = q
	}
	o[s] = byte('0' + i)
	return p + copy((*b)[p:], o[s:])
}

// Normal will attempt to normalize the requested log level. This will check the
// supplied integer and will return it as a valid log level if in bounds of the
// supported log levels. If not, the specified normal log level will be returned
// instead.
func Normal(req int, normal Level) Level {
	if req < 0 {
		return normal
	}
	if req > int(Panic) {
		return normal
	}
	return Level(req)
}

// NormalUint will attempt to normalize the requested log level. This will check
// the supplied integer and will return it as a valid log level if in bounds of
// the supported log levels. If not, the specified normal log level will be
// returned instead. This function is made to specifically work on unsigned
// integers instead.
func NormalUint(req uint, normal Level) Level {
	if req > uint(Panic) {
		return normal
	}
	return Level(req)
}
func (l *logger) Output(d int, s string) error {
	var (
		f string
		p int
	)
	if l.f&(FlagFileLong|FlagFileShort) != 0 {
		var ok bool
		if _, f, p, ok = runtime.Caller(d); !ok {
			f, p = "??", 0
		}
	}
	var (
		b [28]byte
		n int
	)
	l.m.Lock()
	if t := time.Now(); l.f&(FlagDate|FlagTimeUTC|FlagTime|FlagMicroseconds) != 0 {
		if l.f&FlagTimeUTC != 0 {
			t = t.UTC()
		}
		if l.f&FlagDate != 0 {
			y, m, d := t.Date()
			n = itoa(&b, n, y, 4)
			b[n] = '/'
			n = itoa(&b, n+1, int(m), 2)
			b[n] = '/'
			n = itoa(&b, n+1, d, 2)
			b[n] = ' '
			n++
		}
		if l.f&(FlagTime|FlagMicroseconds) != 0 {
			h, m, v := t.Clock()
			n = itoa(&b, n, h, 2)
			b[n] = ':'
			n = itoa(&b, n+1, m, 2)
			b[n] = ':'
			n = itoa(&b, n+1, v, 2)
			if l.f&FlagMicroseconds != 0 {
				b[n] = '.'
				n = itoa(&b, n+1, t.Nanosecond()/1e3, 6)
			}
			b[n] = ' '
			n++
		}
	}
	o := make([]byte, n+len(s)+len(f)+len(l.p)+30)
	if copy(o, b[:n]); l.f&(FlagFileLong|FlagFileShort) != 0 {
		if l.f&FlagFileShort != 0 {
			for i := len(f) - 1; i > 0; i-- {
				if f[i] == '/' {
					f = f[i+1:]
					break
				}
			}
		}
		n += copy(o[n:], f)
		o[n] = ':'
		n++
		c := itoa(&b, 0, p, -1)
		b[c] = ' '
		n += copy(o[n:], b[:c+1])
	}
	if len(l.p) > 0 {
		n += copy(o[n:], l.p)
		o[n] = ' '
		n++
	}
	if n += copy(o[n:], s); len(s) == 0 || o[n-1] != '\n' {
		o[n] = '\n'
		n++
	}
	_, err := l.w.Write(o[:n])
	l.m.Unlock()
	o = nil
	return err
}
