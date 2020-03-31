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

// Package logx is a easy to use logging library for usage in any Golang application.
//
// With LogX you can:
// - Log to Console!
// - Log to Files!
// - Log to multiple logs simultaneously!
// - Log to any io.Writer!
//
// LogX supports all the options of the standard Golang log library with the addition of the Trace, Debug, Info, Warning and Error logging levels.
//
// Example using LogX:
//
// package main
//
// import "github.com/PurpleSec/logx"
//
// func main() {
//     New Console Log
//     con := logx.Console(logx.Info)
//     con.Error("Testing Error!")
//     con.Info("Informational Numbers: %d %d %d...", 1, 2, 3)
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
package logx
