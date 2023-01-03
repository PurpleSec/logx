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

// Global is the default Global logging instance. This can be used instead of passing
// around a logging handle.
//
// All standard 'Log*' functions or functions with a nil struct will go to this
// logging instance.
var Global = Console()

// LogInfo writes an informational message to the Global logger.
//
// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
// first argument is a string that can contain formatting characters. The second
// argument is a vardict of interfaces that can be omitted or used in the supplied
// format string.
//
// This function is used only as a handy quick usage solution. It is recommended
// to use a direct function call on a logger or the Global logger instead.
func LogInfo(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Info(m, v)
}

// LogError writes an error message to the Global logger.
//
// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
// first argument is a string that can contain formatting characters. The second
// argument is a vardict of interfaces that can be omitted or used in the supplied
// format string.
//
// This function is used only as a handy quick usage solution. It is recommended
// to use a direct function call on a logger or the Global logger instead.
func LogError(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Error(m, v)
}

// LogFatal writes a fatal message to the Global logger.
//
// This function will result in the program exiting with a non-zero error code
// after being called, unless the 'logx.FatalExits' setting is 'false'. The
// function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The first
// argument is a string that can contain formatting characters. The second argument
// is a vardict of interfaces that can be omitted or used in the supplied format
// string.
//
// This function is used only as a handy quick usage solution. It is recommended
// to use a direct function call on a logger or the Global logger instead.
func LogFatal(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Fatal(m, v)
}

// LogTrace writes a tracing message to the Global logger.
//
// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
// first argument is a string that can contain formatting characters. The second
// argument is a vardict of interfaces that can be omitted or used in the supplied
// format string.
//
// This function is used only as a handy quick usage solution. It is recommended
// to use a direct function call on a logger or the Global logger instead.
func LogTrace(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Trace(m, v)
}

// LogDebug writes a debugging message to the Global logger.
//
// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
// first argument is a string that can contain formatting characters. The second
// argument is a vardict of interfaces that can be omitted or used in the supplied
// format string.
//
// This function is used only as a handy quick usage solution. It is recommended
// to use a direct function call on a logger or the Global logger instead.
func LogDebug(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Debug(m, v)
}

// LogPrint writes a message to the Global logger.
//
// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
// first argument is a string that can contain formatting characters. The second
// argument is a vardict of interfaces that can be omitted or used in the supplied
// format string.
//
// This function is affected by the setting of 'Global.SetPrintLevel'. By default,
// this will print as an 'Info' logging message.
//
// This function is used only as a handy quick usage solution. It is recommended
// to use a direct function call on a logger or the Global logger instead.
func LogPrint(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Printf(m, v)
}

// LogPanic writes a panic message to the Global logger.
//
// This function will result in the program exiting with a Go 'panic()' after
// being called. The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'.
// The first argument is a string that can contain formatting characters. The
// second argument is a vardict of interfaces that can be omitted or used in
// the supplied format string.
//
// This function is used only as a handy quick usage solution. It is recommended
// to use a direct function call on a logger or the Global logger instead.
func LogPanic(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Panicf(m, v)
}

// LogWarning writes a warning message to the Global logger.
//
// The function arguments are similar to 'fmt.Sprintf' and 'fmt.Printf'. The
// first argument is a string that can contain formatting characters. The second
// argument is a vardict of interfaces that can be omitted or used in the supplied
// format string.
//
// This function is used only as a handy quick usage solution. It is recommended
// to use a direct function call on a logger or the Global logger instead.
func LogWarning(m string, v ...interface{}) {
	if Global == nil {
		return
	}
	Global.Warning(m, v)
}
