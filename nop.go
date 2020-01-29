// Copyright (C) 2020 iDigitalFlame
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

import "os"

var (
	//Nop is a Log interface that can consume all types of
	// logs but does not do anything (a NOP). This can be used similar
	// to context.TODO or if logging is not needed and nil log instances are not
	// avaliable.
	Nop = nop{}
)

type nop struct{}

func (nop) Fatal(_ string, _ ...interface{}) {
	os.Exit(1)
}
func (nop) SetLevel(_ Level)                   {}
func (nop) SetPrefix(_ string)                 {}
func (nop) Info(_ string, _ ...interface{})    {}
func (nop) Error(_ string, _ ...interface{})   {}
func (nop) Trace(_ string, _ ...interface{})   {}
func (nop) Debug(_ string, _ ...interface{})   {}
func (nop) Printf(_ string, _ ...interface{})  {}
func (nop) Warning(_ string, _ ...interface{}) {}
