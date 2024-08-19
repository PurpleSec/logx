# LogX - Simple Golang Logging Library

LogX is an easy to use logging library for usage in any Golang application.

With LogX you can:

- Log to Console!
- Log to Files!
- Log to multiple logs simultaneously!
- Log to any io.Writer!

LogX supports all the options of the standard Golang log library with the addition of the Trace, Debug, Info, Warning and Error logging levels.

Example using LogX:

```[golang]
package main

import "github.com/PurpleSec/logx"

func main() {
    // New Console Log
    con := logx.Console(logx.Info)
    con.Error("Testing Error!")
    con.Info("Informational Numbers: %d %d %d...", 1, 2, 3)

    // New File Log
    fil, err := logx.File("log.log", logx.Debug)
    if err != nil {
        panic(err)
    }
    fil.Debug("Debugging in progress!")
    fil.Trace("You shouldn't see this :P")
    fil.Warning("Objects in mirror are closer than they appear!")

    // Multi Log Logging!
    multi := logx.Multiple(con, fil)

    // Disable Exiting on a Fatal call for testing.
    logx.FatalExits = false

    multi.Debug("This will only appear in the file log :D")
    multi.Info("Hello World!")
    multi.Fatal("OMG BAD STUFF")
    multi.Trace("This won't get logged, yay!")
}
```

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/Z8Z4121TDS)
