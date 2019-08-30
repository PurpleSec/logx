package main

import "github.com/iDigitalFlame/logx/logx"

func main() {
	l := logx.NewConsole(logx.LTrace)

	l.Trace("Trace Log Entry!")
	l.Debug("Debug Log Entry!")
	l.Info("Information Log Entry!")
	l.Warning("Warning LOg Entry!")
	l.Fatal("Fatal Log Entry!")
}
