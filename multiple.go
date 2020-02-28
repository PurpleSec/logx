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

import "os"

// Multi is a type of Log that is an alias for an array where each Log
// function will affect each Log instance in the array.
type Multi struct {
	logs []Log
}

// Add appends the specified Log 'l' the Stack array.
func (m *Multi) Add(l Log) {
	if l == nil {
		return
	}
	m.logs = append(m.logs, l)
}

// Multiple returns a Stack struct that contains the Log instances
// specified in the 'l' vardict.
func Multiple(l ...Log) *Multi {
	return &Multi{logs: l}
}

// SetLevel changes the current logging level of this Log instance.
func (m *Multi) SetLevel(l Level) {
	for i := range m.logs {
		m.logs[i].SetLevel(l)
	}
}

// SetPrefix changes the current logging prefox of this Log instance.
func (m *Multi) SetPrefix(p string) {
	for i := range m.logs {
		m.logs[i].SetPrefix(p)
	}
}

// Info writes a informational message to the Global logger.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
func (m *Multi) Info(s string, v ...interface{}) {
	for i := range m.logs {
		if x, ok := m.logs[i].(LogWriter); ok {
			x.Log(Info, 1, s, v...)
		} else {
			m.logs[i].Info(s, v...)
		}
	}
}

// Error writes a error message to the Global logger.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
func (m *Multi) Error(s string, v ...interface{}) {
	for i := range m.logs {
		if x, ok := m.logs[i].(LogWriter); ok {
			x.Log(Error, 1, s, v...)
		} else {
			m.logs[i].Error(s, v...)
		}
	}
}

// Fatal writes a fatal message to the Global logger. This function will result in the program
// exiting with a non-zero error code after being called, unless the logx.FatalExits' setting is 'false'.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
func (m *Multi) Fatal(s string, v ...interface{}) {
	for i := range m.logs {
		if x, ok := m.logs[i].(LogWriter); ok {
			x.Log(Fatal, 1, s, v...)
		} else {
			// Write as Error here to prevent the non-flexable logger from exiting the program
			// before all logs can be written.
			m.logs[i].Error(s, v...)
		}
	}
	if FatalExits {
		os.Exit(1)
	}
}

// Trace writes a tracing message to the Global logger.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
func (m *Multi) Trace(s string, v ...interface{}) {
	for i := range m.logs {
		if x, ok := m.logs[i].(LogWriter); ok {
			x.Log(Trace, 1, s, v...)
		} else {
			m.logs[i].Trace(s, v...)
		}
	}
}

// Debug writes a debugging message to the Global logger.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
func (m *Multi) Debug(s string, v ...interface{}) {
	for i := range m.logs {
		if x, ok := m.logs[i].(LogWriter); ok {
			x.Log(Debug, 1, s, v...)
		} else {
			m.logs[i].Debug(s, v...)
		}
	}
}

// Warning writes a warning message to the Global logger.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
func (m *Multi) Warning(s string, v ...interface{}) {
	for i := range m.logs {
		if x, ok := m.logs[i].(LogWriter); ok {
			x.Log(Warning, 1, s, v...)
		} else {
			m.logs[i].Warning(s, v...)
		}
	}
}
