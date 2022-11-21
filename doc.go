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

// Package logx is an easy to use logging library for usage in any Golang application.
//
// With LogX you can:
// - Log to Console!
// - Log to Files!
// - Log to multiple logs simultaneously!
// - Log to any io.Writer!
//
// LogX supports all the options of the standard Golang log library with the
// addition of the Trace, Debug, Info, Warning and Error logging levels.
//
// Example using LogX:
//
//
// package main
//
// import "github.com/PurpleSec/logx"
//
// func main() {
//     New Console Log
//     con := logx.Console(logx.Info)
//     con.Error("Testing Error!")
//     con.Info("Informational Numbers: %d %d %d..", 1, 2, 3)
//
//     New File Log
//     fil, err := logx.File("log.log", logx.Debug)
//     if err != nil {
//          panic(err)
//      }
//    fil.Debug("Debugging in progress!")
//    fil.Trace("You shouldn't see this :P")
//    fil.Warning("Objects in mirror are closer than they appear!")
//
//    Multi Log Logging!
//    multi := logx.Multiple(con, fil)
//
//    Disable Exiting on a Fatal call for testing.
//    logx.FatalExits = false
//
//    multi.Debug("This will only appear in the file log :D")
//    multi.Info("Hello World!")
//    multi.Fatal("OMG BAD STUFF")
//    multi.Trace("This won't get logged, yay!")
// }
//
//
package logx
