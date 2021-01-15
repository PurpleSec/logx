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

// Global is the default Global loging instance. This can be used instead of passing
// around a logging handle. All standard Log* functions or functions with a nil struct
// will go to this loging instance.
var Global = Console()

// LogInfo writes a informational message to the Global logger.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
//
// This function is used only as a handy quick usage solution. It is recommended to use a
// direct function call on a logger or the Global logger instead.
func LogInfo(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Info(m, v)
}

// LogError writes a error message to the Global logger.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
//
// This function is used only as a handy quick usage solution. It is recommended to use a
// direct function call on a logger or the Global logger instead.
func LogError(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Error(m, v)
}

// LogFatal writes a fatal message to the Global logger. This function will result in the program
// exiting with a non-zero error code after being called, unless the logx.FatalExits' setting is 'false'.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
//
// This function is used only as a handy quick usage solution. It is recommended to use a
// direct function call on a logger or the Global logger instead.
func LogFatal(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Fatal(m, v)
}

// LogTrace writes a tracing message to the Global logger.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
//
// This function is used only as a handy quick usage solution. It is recommended to use a
// direct function call on a logger or the Global logger instead.
func LogTrace(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Trace(m, v)
}

// LogDebug writes a debugging message to the Global logger.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
//
// This function is used only as a handy quick usage solution. It is recommended to use a
// direct function call on a logger or the Global logger instead.
func LogDebug(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Debug(m, v)
}

// LogPrint writes a message to the Global logger.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
// This function is affected by the setting of 'SetPrintLevel'. By default, this will print as an 'Info'
// logging message.
func LogPrint(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Printf(m, v)
}

// LogPanic writes a panic message to the Global logger. This function will result in the program
// exiting with a Go 'panic()' after being called. The function arguments are similar to fmt.Sprintf and
// fmt.Printf. The first argument is a string that can contain formatting characters. The second argument
// is a vardict of interfaces that can be omitted or used in the supplied format string.
func LogPanic(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Panicf(m, v)
}

// LogWarning writes a warning message to the Global logger.
// The function arguments are similar to fmt.Sprintf and fmt.Printf. The first argument is
// a string that can contain formatting characters. The second argument is a vardict of
// interfaces that can be omitted or used in the supplied format string.
//
// This function is used only as a handy quick usage solution. It is recommended to use a
// direct function call on a logger or the Global logger instead.
func LogWarning(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Warning(m, v)
}
