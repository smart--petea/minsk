package main

import (
    "minsk"
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("/tmp/minsk.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(file)
    log.SetFlags(log.Lshortfile)


    repl := minsk.NewMinskRepl()
    repl.Run(repl)
}
