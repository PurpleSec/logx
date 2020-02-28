// Copyright (C) 2020 PurpleSec Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
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
	// Info writes a informational message to the Global logger.
	// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
	// a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	Info(string, ...interface{})
	// Error writes a error message to the Global logger.
	// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
	// a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	Error(string, ...interface{})
	// Fatal writes a fatal message to the Global logger. This function will result in the program
	// exiting with a non-zero error code after being called, unless the logx.FatalExits' setting is 'false'.
	// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
	// a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	Fatal(string, ...interface{})
	// Trace writes a tracing message to the Global logger.
	// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
	// a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	Trace(string, ...interface{})
	// Debug writes a debugging message to the Global logger.
	// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
	// a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	Debug(string, ...interface{})
	// Warning writes a warning message to the Global logger.
	// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
	// a string that can contain formatting characters. The second argument is a vardict of
	// interfaces that can be omitted or used in the supplied format string.
	Warning(string, ...interface{})
}

// LogWriter is an interface that is used inline with logging operations. This interface defines a single function
// 'Log' which takes a Level to log, the additional stack depth (can be zero), and the message and optional arguments
// to be logged. This function is used as a "quick-logging" helper and is preferred by the Multi logging struct.
//
// This funcion must ONLY log and not preform any other operations such as exiting on a Fatal logging level. Use the
// higher level 'Fatal' function for those additional operations.
//
// The higher level calls may use this function to simplify logging.
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
	}
	return ""
}
