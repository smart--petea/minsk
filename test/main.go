package main

import (
    "minsk"
)

func main() {
    repl := minsk.NewMinskRepl()
    repl.Run(repl)
}
