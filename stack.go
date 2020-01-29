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

// Stack is a type of Log that is an alias for an array where each Log
// function will affect each Log instance in the array.
type Stack []Log

// Add appends the specified Log 'l' the Stack array.
func (s *Stack) Add(l Log) {
	if l == nil {
		return
	}
	*s = append(*s, l)
}

// NewStack returns a Stack struct that contains the Log instances
// specified in the 'l' vardict.
func NewStack(l ...Log) *Stack {
	s := Stack(l)
	return &s
}

// SetLevel changes the current logging level of this Log instance.
func (s *Stack) SetLevel(n Level) {
	for i := range *s {
		(*s)[i].SetLevel(n)
	}
}

// SetPrefix changes the current logging prefox of this Log instance.
func (s *Stack) SetPrefix(p string) {
	for i := range *s {
		(*s)[i].SetPrefix(p)
	}
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

// Debug writes a debugging message to the Log instance.
func (s *Stack) Debug(m string, v ...interface{}) {
	for i := range *s {
		if b, ok := (*s)[i].(handler); ok {
			writeToLog(b.Writer(), b.Level(), LDebug, stackDepth+1, m, v)
		} else {
			(*s)[i].Debug(m, v...)
		}
	}
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
