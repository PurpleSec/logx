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

import "testing"

func TestLogging(_ *testing.T) {
	l := NewConsole(LTrace)

	l.Trace("Trace Log Entry!")

	l.Debug("Debug Log Entry!")

	l.Info("Information Log Entry!")

	l.Warning("Warning LOg Entry!")

	l.Fatal("Fatal Log Entry!")
}
