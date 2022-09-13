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
	"fmt"
	"os"
)

// Multi is a type of Log that is an alias for an array where each Log function
// will affect each Log instance in the array.
type Multi []Log

// Add appends the specified Log 'l' the Stack array.
func (m *Multi) Add(l Log) {
	if l == nil {
		return
	}
	*m = append(*m, l)
}

// Multiple returns a Stack struct that contains the Log instances
// specified in the 'l' vardict.
func Multiple(l ...Log) *Multi {
	m := Multi(l)
	return &m
}

// SetLevel changes the current logging level of this Log instance.
func (m Multi) SetLevel(l Level) {
	for i := range m {
		m[i].SetLevel(l)
	}
}

// SetPrefix changes the current logging prefox of this Log instance.
func (m Multi) SetPrefix(p string) {
	for i := range m {
		m[i].SetPrefix(p)
	}
}

// SetPrintLevel sets the logging level used when 'Print*' statements are called.
func (m Multi) SetPrintLevel(n Level) {
	for i := range m {
		m[i].SetPrintLevel(n)
	}
}

// Print writes a message to the logger.
//
// The function arguments are similar to 'fmt.Sprint' and 'fmt.Print'. The only
// argument is a vardict of interfaces that can be used to output a string value.
//
// This function is affected by the setting of 'SetPrintLevel'. By default,
// this will print as an 'Info' logging message.
func (m Multi) Print(v ...interface{}) {
	for i := range m {
		if x, ok := m[i].(LogWriter); ok {
			x.Log(Print, 1, "", v...)
		} else {
			m[i].Print(v...)
		}
	}
}

// Panic writes a panic message to the logger.
//
// This function will result in the program exiting with a Go 'panic()' after
// being called. The function arguments are similar to 'fmt.Sprint' and 'fmt.Print.'
// The only argument is a vardict of interfaces that can be used to output a
// string value.
func (m Multi) Panic(v ...interface{}) {
	for i := range m {
		if x, ok := m[i].(LogWriter); ok {
			x.Log(Panic, 1, "", v...)
		} else {
			// NOTE(dij): Write as Error here to prevent the non-flexable logger
			//            from exiting the program before all logs can be written.
			m[i].Error("", v...)
		}
	}
	panic(fmt.Sprint(v...))
}

// Println writes a message to the logger.
//
// The function arguments are similar to fmt.Sprintln and fmt.Println. The only
// argument is a vardict of interfaces that can be used to output a string value.
//
// This function is affected by the setting of 'SetPrintLevel'. By default,
// this will print as an 'Info' logging message.
func (m Multi) Println(v ...interface{}) {
	for i := range m {
		if x, ok := m[i].(LogWriter); ok {
			x.Log(Print, 1, "", v...)
		} else {
			m[i].Println(v...)
		}
	}
}

// Panicln writes a panic message to the logger.
//
// This function will result in the program exiting with a Go 'panic()' after
// being called. The function arguments are similar to 'fmt.Sprintln' and
// 'fmt.Println'. The only argument is a vardict of interfaces that
// can be used to output a string value.
func (m Multi) Panicln(v ...interface{}) {
	for i := range m {
		if x, ok := m[i].(LogWriter); ok {
			x.Log(Panic, 1, "", v...)
		} else {
			// NOTE(dij): Write as Error here to prevent the non-flexable logger
			//            from exiting the program before all logs can be written.
			m[i].Error("", v...)
		}
	}
	panic(fmt.Sprint(v...))
}

// Info writes n informational message to the logger.
//
// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
// first argument is a string that can contain formatting characters. The second
// argument is a vardict of interfaces that can be omitted or used in the supplied
// format string.
func (m Multi) Info(s string, v ...interface{}) {
	for i := range m {
		if x, ok := m[i].(LogWriter); ok {
			x.Log(Info, 1, s, v...)
		} else {
			m[i].Info(s, v...)
		}
	}
}

// Error writes an error message to the logger.
//
// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
// first argument is a string that can contain formatting characters. The second
// argument is a vardict of interfaces that can be omitted or used in the supplied
// format string.
func (m Multi) Error(s string, v ...interface{}) {
	for i := range m {
		if x, ok := m[i].(LogWriter); ok {
			x.Log(Error, 1, s, v...)
		} else {
			m[i].Error(s, v...)
		}
	}
}

// Fatal writes a fatal message to the logger.
//
// This function will result in the program exiting with a non-zero error code
// after being called, unless the 'logx.FatalExits' setting is 'false'. The
// function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The first
// argument is a string that can contain formatting characters. The second argument
// is a vardict of interfaces that can be omitted or used in the supplied format
// string.
func (m Multi) Fatal(s string, v ...interface{}) {
	for i := range m {
		if x, ok := m[i].(LogWriter); ok {
			x.Log(Fatal, 1, s, v...)
		} else {
			// NOTE(dij): Write as Error here to prevent the non-flexable logger
			//            from exiting the program before all logs can be written.
			m[i].Error(s, v...)
		}
	}
	if FatalExits {
		os.Exit(1)
	}
}

// Trace writes a tracing message to the logger.
//
// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
// first argument is a string that can contain formatting characters. The second
// argument is a vardict of interfaces that can be omitted or used in the supplied
// format string.
func (m Multi) Trace(s string, v ...interface{}) {
	for i := range m {
		if x, ok := m[i].(LogWriter); ok {
			x.Log(Trace, 1, s, v...)
		} else {
			m[i].Trace(s, v...)
		}
	}
}

// Debug writes a debugging message to the logger.
//
// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
// first argument is a string that can contain formatting characters. The second
// argument is a vardict of interfaces that can be omitted or used in the supplied
// format string.
func (m Multi) Debug(s string, v ...interface{}) {
	for i := range m {
		if x, ok := m[i].(LogWriter); ok {
			x.Log(Debug, 1, s, v...)
		} else {
			m[i].Debug(s, v...)
		}
	}
}

// Printf writes a message to the logger.
//
// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
// first argument is a string that can contain formatting characters. The second
// argument is a vardict of interfaces that can be omitted or used in the supplied
// format string.
//
// This function is affected by the setting of 'SetPrintLevel'. By default,
// this will print as an 'Info' logging message.
func (m Multi) Printf(s string, v ...interface{}) {
	for i := range m {
		if x, ok := m[i].(LogWriter); ok {
			x.Log(Print, 1, s, v...)
		} else {
			m[i].Printf(s, v...)
		}
	}
}

// Panicf writes a panic message to the logger.
//
// This function will result in the program exiting with a Go 'panic()' after
// being called. The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'.
// The first argument is a string that can contain formatting characters. The
// second argument is a vardict of interfaces that can be omitted or used in
// the supplied format string.
func (m Multi) Panicf(s string, v ...interface{}) {
	for i := range m {
		if x, ok := m[i].(LogWriter); ok {
			x.Log(Panic, 1, s, v...)
		} else {
			// NOTE(dij): Write as Error here to prevent the non-flexable logger
			//            from exiting the program before all logs can be written.
			m[i].Error(s, v...)
		}
	}
	panic(fmt.Sprintf(s, v...))
}

// Warning writes a warning message to the logger.
//
// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
// first argument is a string that can contain formatting characters. The second
// argument is a vardict of interfaces that can be omitted or used in the supplied
// format string.
func (m Multi) Warning(s string, v ...interface{}) {
	for i := range m {
		if x, ok := m[i].(LogWriter); ok {
			x.Log(Warning, 1, s, v...)
		} else {
			m[i].Warning(s, v...)
		}
	}
}
