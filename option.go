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

import "log"

// Append is a loging setting that instructs the Log to override the default log file truncation
// behavior. When this is used in the options for creating a file backed log instance, the new logged data
// will be appended to any previous data that the file contains.
//
// This setting has no effect on non-file backed logging instances.
const Append = settingAppend(true)

// DefaultFlags is the default bitwise flag value that is used for new logging instances that are not
// given an flag options setting when created.
//
// This flag number may be changed before running to affect creation of new logging instances.
// NOTE: 'log.Lmsgprefix' requires Golang 1.14.1
var DefaultFlags = log.Ldate | log.Ltime | log.Lmsgprefix

const (
	setLevel setting = iota
	setFlags
	setAppend
	setPrefix
)

type setting uint8
type settingFlags int
type settingAppend bool
type settingPrefix string

// Option is an interface that allows for passing a vardict of potential
// settings that can be used during creation of a logging instance. This interface
// type will only be fulfilled by interanal functions.
type Option interface {
	setting() setting
}

// Flags will create an Option interface that will set the provided flag value on
// the underlying log instance when created.
//
// Valid values for this can be referenced from the 'log' package.
func Flags(f int) Option {
	return settingFlags(f)
}

// Prefix will create an Option interface that will set the provided prefix on
// the logging instance when created.
func Prefix(p string) Option {
	return settingPrefix(p)
}
func (l Level) setting() setting {
	return setLevel
}
func (settingFlags) setting() setting {
	return setFlags
}
func (settingAppend) setting() setting {
	return setAppend
}
func (settingPrefix) setting() setting {
	return setPrefix
}
