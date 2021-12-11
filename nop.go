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

// NOP is a Logger that can consume all types of logs but does not do anything
// (a NOP).
//
// This can be used similar to context.TODO or if logging is not needed and nil
// log instances are not available or allowed.
var NOP = nop{}

type nop struct{}

func (nop) Panic(v ...interface{}) {
	panic(fmt.Sprint(v...))
}
func (nop) Panicln(v ...interface{}) {
	panic(fmt.Sprint(v...))
}
func (nop) Fatal(_ string, _ ...interface{}) {
	if FatalExits {
		os.Exit(1)
	}
}
func (nop) Panicf(m string, v ...interface{}) {
	panic(fmt.Sprintf(m, v...))
}

func (nop) SetLevel(_ Level)                               {}
func (nop) SetPrefix(_ string)                             {}
func (nop) SetPrintLevel(_ Level)                          {}
func (nop) Print(_ ...interface{})                         {}
func (nop) Println(_ ...interface{})                       {}
func (nop) Info(_ string, _ ...interface{})                {}
func (nop) Error(_ string, _ ...interface{})               {}
func (nop) Trace(_ string, _ ...interface{})               {}
func (nop) Debug(_ string, _ ...interface{})               {}
func (nop) Printf(_ string, _ ...interface{})              {}
func (nop) Warning(_ string, _ ...interface{})             {}
func (nop) Log(_ Level, _ int, _ string, _ ...interface{}) {}
