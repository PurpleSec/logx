// Copyright 2021 - 2023 PurpleSec Team
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

// Append is a logging setting that instructs the Log to override the default log
// file truncation behavior. When this is used in the options for creating a file
// backed log instance, the new logged data will be appended to any previous data
// that the file contains.
//
// This setting has no effect on non-file backed logging instances.
const Append = settingAppend(true)

// DefaultFlags is the default bitwise flag value that is used for new logging
// instances that are not given a flag options setting when created.
//
// This flag number may be changed before running to affect creation of new
// logging instances.
var DefaultFlags = FlagStandard

const (
	setLevel setting = iota
	setFlags
	setPrint
	setAppend
	setPrefix
)

type setting uint8
type settingFlags int8
type settingPrint uint8
type settingAppend bool
type settingPrefix string

// Option is an interface that allows for passing a vardict of potential
// settings that can be used during creation of a logging instance.
//
// This interface type will only be fulfilled by internal functions.
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

// PrintLevel will return an Option interface that will set the level used by
// the 'Print*' functions.
//
// This is similar to calling the 'SetPrintLevel' function.
func PrintLevel(l Level) Option {
	return settingPrint(l)
}
func (l Level) setting() setting {
	return setLevel
}
func (settingFlags) setting() setting {
	return setFlags
}
func (settingPrint) setting() setting {
	return setPrint
}
func (settingAppend) setting() setting {
	return setAppend
}
func (settingPrefix) setting() setting {
	return setPrefix
}
