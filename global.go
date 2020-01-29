// Copyright (C) 2020 iDigitalFlame
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

var (
	// Global is the default Global logger.  This can be used instead of passing
	// around a logging handle. All standard Log functions without a struct will go
	// to this logger.
	Global = NewConsole(LWarning)
)

// Info writes a information message to the Log instance.
func Info(m string, v ...interface{}) {
	Global.Info(m, v)
}

// Error writes a error message to the Log instance.
func Error(m string, v ...interface{}) {
	Global.Error(m, v)
}

// Fatal writes a fatal message to the Log instance. This function
// will result in the program exiting with a non-zero error code after being called.
func Fatal(m string, v ...interface{}) {
	Global.Fatal(m, v)
}

// Trace writes a tracing message to the Log instance.
func Trace(m string, v ...interface{}) {
	Global.Trace(m, v)
}

// Debug writes a debugging message to the Log instance.
func Debug(m string, v ...interface{}) {
	Global.Debug(m, v)
}

// Printf writes a information message to the Log instance.
func Printf(m string, v ...interface{}) {
	Global.Printf(m, v)
}

// Warning writes a warning message to the Log instance.
func Warning(m string, v ...interface{}) {
	Global.Warning(m, v)
}
