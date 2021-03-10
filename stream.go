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

import (
	"fmt"
	"io"
	"log"
	"os"
)

// DefaultConsole is a pointer to the output that all the console Log structs will use when created.
// This can be set to any io.Writer. The default is the Stderr console.
var DefaultConsole io.Writer = os.Stderr

type file struct {
	stream
	f string
}
type stream struct {
	*log.Logger
	l Level
	p Level
}

// Console returns a console logger that uses the Console writer.
func Console(o ...Option) Log {
	return Writer(DefaultConsole, o...)
}
func (s *stream) SetLevel(n Level) {
	s.l = n
}
func (s *stream) SetPrefix(p string) {
	s.Logger.SetPrefix(p)
}
func (s *stream) SetPrintLevel(n Level) {
	s.p = n
}

func (s *stream) Print(v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(s.p, 0, "", v...)
		return
	}
	s.Log(s.p, 0, "", v...)
}
func (s *stream) Panic(v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(Panic, 0, "", v...)
	} else {
		s.Log(Panic, 0, "", v...)
	}
	panic(fmt.Sprintln(v...))
}
func (s *stream) Println(v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(s.p, 0, "", v...)
		return
	}
	s.Log(s.p, 0, "", v...)
}
func (s *stream) Panicln(v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(Panic, 0, "", v...)
	} else {
		s.Log(Panic, 0, "", v...)
	}
	panic(fmt.Sprintln(v...))
}

// Writer returns a Log instance based on the Writer 'w' for the logging output and
// allows specifying non-default Logging options.
func Writer(w io.Writer, o ...Option) Log {
	var (
		f    settingFlags = -1
		p    settingPrefix
		l, k Level = invalidLevel, invalidLevel
	)
	for i := range o {
		if o[i] == nil {
			continue
		}
		switch o[i].setting() {
		case setLevel:
			l, _ = o[i].(Level)
		case setFlags:
			f, _ = o[i].(settingFlags)
		case setPrint:
			k, _ = o[i].(Level)
		case setPrefix:
			p, _ = o[i].(settingPrefix)
		}
	}
	if f == -1 {
		f = settingFlags(DefaultFlags)
	}
	if l == invalidLevel {
		l = Warning
	}
	if k == invalidLevel {
		k = Info
	}
	return &stream{l: l, p: k, Logger: log.New(w, string(p), int(f))}
}

// File will attempt to create a File backed Log instance that will write to file specified.
// This function will truncate the file before starting a new Log. If you need to append to a existing log file.
// use the NewWriter function. This function allows specifying non-default Logging options.
func File(s string, o ...Option) (Log, error) {
	var (
		f    settingFlags = -1
		p    settingPrefix
		a    settingAppend
		l, k Level = invalidLevel, invalidLevel
		n          = os.O_WRONLY | os.O_CREATE
	)
	for i := range o {
		if o[i] == nil {
			continue
		}
		switch o[i].setting() {
		case setLevel:
			l, _ = o[i].(Level)
		case setFlags:
			f, _ = o[i].(settingFlags)
		case setPrint:
			k, _ = o[i].(Level)
		case setAppend:
			a, _ = o[i].(settingAppend)
		case setPrefix:
			p, _ = o[i].(settingPrefix)
		}
	}
	if f == -1 {
		f = settingFlags(DefaultFlags)
	}
	if l == invalidLevel {
		l = Warning
	}
	if k == invalidLevel {
		k = Info
	}
	if a {
		n |= os.O_APPEND
	}
	w, err := os.OpenFile(s, n, 0644)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %q for logging: %w", s, err)
	}
	return &file{f: s, stream: stream{l: l, p: k, Logger: log.New(w, string(p), int(f))}}, nil
}
func (s *stream) Info(m string, v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(Info, 0, m, v...)
		return
	}
	s.Log(Info, 0, m, v...)
}
func (s *stream) Error(m string, v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(Error, 0, m, v...)
		return
	}
	s.Log(Error, 0, m, v...)
}
func (s *stream) Fatal(m string, v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(Fatal, 0, m, v...)
	} else {
		s.Log(Fatal, 0, m, v...)
	}
	if FatalExits {
		os.Exit(1)
	}
}
func (s *stream) Trace(m string, v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(Trace, 0, m, v...)
		return
	}
	s.Log(Trace, 0, m, v...)
}
func (s *stream) Debug(m string, v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(Debug, 0, m, v...)
		return
	}
	s.Log(Debug, 0, m, v...)
}
func (s *stream) Printf(m string, v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(s.p, 0, m, v...)
		return
	}
	s.Log(s.p, 0, m, v...)
}
func (s *stream) Panicf(m string, v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(Panic, 0, m, v...)
	} else {
		s.Log(Panic, 0, m, v...)
	}
	panic(fmt.Sprintf(m, v...))
}
func (s *stream) Warning(m string, v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(Warning, 0, m, v...)
		return
	}
	s.Log(Warning, 0, m, v...)
}
func (s *stream) Log(l Level, c int, m string, v ...interface{}) {
	if l == Print {
		// Duplicate code here to prevent inf loops.
		if s.l > s.p {
			return
		}
		if len(m) == 0 {
			s.Logger.Output(3+c, fmt.Sprintf("[%s]: %s\n", s.p.String(), fmt.Sprint(v...)))
			return
		}
		s.Logger.Output(3+c, fmt.Sprintf("[%s]: %s\n", s.p.String(), fmt.Sprintf(m, v...)))
		return
	}
	if s.l > l {
		return
	}
	if len(m) == 0 {
		s.Logger.Output(3+c, fmt.Sprintf("[%s]: %s\n", l.String(), fmt.Sprint(v...)))
		return
	}
	s.Logger.Output(3+c, fmt.Sprintf("[%s]: %s\n", l.String(), fmt.Sprintf(m, v...)))
}
