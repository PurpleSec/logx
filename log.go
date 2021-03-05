// Copyright 2021 PurpleSec Team
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

const (
	// Trace is the Tracing log level, everything logged is passed to the Log (it might get noisy).
	Trace Level = iota
	// Debug prints slightly less information than Trace, but is primarily useful for debugging messages.
	Debug
	// Info is the standard informational log Level, this can be used to print out handy information messages.
	Info
	// Warning is exactly how it sounds, any event that occurs of notice, but is not major. This is the default logging
	// level.
	Warning
	// Error is similar to the Warning level, except more serious.
	Error
	// Fatal means the program cannot continue when this event occurs. Normally the program will exit after this.
	// set 'logx.FatalExits' to 'false' to disable exiting when a Fatal log entry is triggered.
	Fatal
	// Panic is a special level only used by the 'Panic*' functions. This will act similar to the 'Fatal' level
	// but will trigger a Go 'panic()' call instead once log writing is complete.
	Panic
	// Print is a special level only used for LogWriters. This level instructs the LogWriter to use it's inbuilt
	// print level for the specified message. Attempting to use this out of this context will fail.
	Print
	invalidLevel
)

// FatalExits is a boolean setting that determines if a call to Fatal or LogFatal will exit the program
// using 'os.Exit(1)'. If this is set to false, a call to Fatal or LogFatal will continue program execution
// after being called. The default value is true.
var FatalExits = true

// Level is an alias of a byte that represents the current Log level.
type Level uint8

// Log is an interface for any type of struct that supports standard Logging functions.
type Log interface {
	// SetLevel changes the current logging level of this Log.
	SetLevel(Level)
	// SetPrefix changes the current logging prefix of this Log.
	SetPrefix(string)
	// SetPrintLevel sets the logging level used when 'Print*' statements are called.
	SetPrintLevel(Level)
	// Print writes a message to the logger.
	// The function arguments are similar to fmt.Sprint and fmt.Print. The only argument is a vardict of
	// interfaces that can be used to output a string value.
	// This function is affected by the setting of 'SetPrintLevel'. By default, this will print as an 'Info'
	// logging message.
	Print(...interface{})
	// Panic writes a panic message to the logger.
	// This function will result in the program exiting with a Go 'panic()' after being called. The function
	// arguments are similar to fmt.Sprint and fmt.Print. The only argument is a vardict of interfaces that can
	// be used to output a string value.
	Panic(...interface{})
	// Println writes a message to the logger.
	// The function arguments are similar to fmt.Sprintln and fmt.Println. The only argument is a vardict of
	// interfaces that can be used to output a string value.
	// This function is affected by the setting of 'SetPrintLevel'. By default, this will print as an 'Info'
	// logging message.
	Println(...interface{})
	// Panicln writes a panic message to the logger.
	// This function will result in the program exiting with a Go 'panic()' after being called. The function
	// arguments are similar to fmt.Sprintln and fmt.Println. The only argument is a vardict of interfaces that
	// can be used to output a string value.
	Panicln(...interface{})
	// Info writes a informational message to the logger.
	// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
	// a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	Info(string, ...interface{})
	// Error writes a error message to the logger.
	// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
	// a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	Error(string, ...interface{})
	// Fatal writes a fatal message to the logger.
	// This function will result in the program exiting with a non-zero error code after being called, unless
	// the logx.FatalExits' setting is 'false'. The function arguments are similar to fmt.Sprintf and fmt.Printf.
	// The first argument is a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	Fatal(string, ...interface{})
	// Trace writes a tracing message to the logger.
	// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
	// a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	Trace(string, ...interface{})
	// Debug writes a debugging message to the logger.
	// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
	// a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	Debug(string, ...interface{})
	// Printf writes a message to the logger.
	// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
	// a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	// This function is affected by the setting of 'SetPrintLevel'. By default, this will print as an 'Info'
	// logging message.
	Printf(string, ...interface{})
	// Panicf writes a panic message to the logger.
	// This function will result in the program exiting with a Go 'panic()' after being called. The function
	// arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is a string that can contain
	// formatting characters. The second argument is a vardict of interfaces that can be omitted or used in
	// the supplied format string.
	Panicf(string, ...interface{})
	// Warning writes a warning message to the logger.
	// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
	// a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	Warning(string, ...interface{})
}

// LogWriter is an interface that is used inline with logging operations. This interface defines a single function
// 'Log' which takes a Level to log, the additional stack depth (can be zero), and the message and optional arguments
// to be logged. This function is used as a "quick-logging" helper and is preferred by the Multi logging struct.
//
// This funcion must ONLY log and not preform any other operations such as exiting on a Fatal or Panic logging level.
// Use the higher level 'Fatal' and 'Panic' functions for those additional operations.
//
// The higher level calls may use this function to simplify logging and comply with the logx Multi-logger.
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

// Normal will attempt to normalize the requested log level. This will check the supplied integer and will return
// it as a valid log level if in bounds of the supported log levels. If not, the specified normal log level will
// be returned instead.
func Normal(req int, normal Level) Level {
	if req < 0 {
		return normal
	}
	if req > int(Panic) {
		return normal
	}
	return Level(req)
}

// NormalUint will attempt to normalize the requested log level. This will check the supplied integer and will return
// it as a valid log level if in bounds of the supported log levels. If not, the specified normal log level will
// be returned instead. This function is made to specifically work on unsigned ints instead.
func NormalUint(req uint, normal Level) Level {
	if req > uint(Panic) {
		return normal
	}
	return Level(req)
}
