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
	f string
	stream
}
type stream struct {
	l Level
	*log.Logger
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

// Writer returns a Log instance based on the Writer 'w' for the logging output and
// allows specifying non-default Logging options.
func Writer(w io.Writer, o ...Option) Log {
	var (
		f settingFlags = -1
		p settingPrefix
		l Level = invalidLevel
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
	return &stream{l, log.New(w, string(p), int(f))}
}

// File will attempt to create a File backed Log instance that will write to file specified.
// This function will truncate the file before starting a new Log. If you need to append to a existing log file.
// use the NewWriter function. This function allows specifying non-default Logging options.
func File(s string, o ...Option) (Log, error) {
	var (
		f settingFlags = -1
		p settingPrefix
		a settingAppend
		l Level = invalidLevel
		n       = os.O_WRONLY | os.O_CREATE
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
	if a {
		n |= os.O_APPEND
	}
	w, err := os.OpenFile(s, n, 0644)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %q for logging: %w", s, err)
	}
	return &file{s, stream{l, log.New(w, string(p), int(f))}}, nil
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
func (s *stream) Warning(m string, v ...interface{}) {
	if s == nil {
		Global.(LogWriter).Log(Warning, 0, m, v...)
		return
	}
	s.Log(Warning, 0, m, v...)
}
func (s *stream) Log(l Level, c int, m string, v ...interface{}) {
	if s.l > l {
		return
	}
	s.Logger.Output(3+c, fmt.Sprintf("[%s]: %s\n", l.String(), fmt.Sprintf(m, v...)))
}
