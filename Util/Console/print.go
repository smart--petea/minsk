package Console

import (
    "fmt"
    "log"
)

func Print(text string) {
    log.Printf("Console.Print %s", text)
    left, top := GetCursorPos()

    fmt.Printf(text)

    SetCursorPos(top, left)
}
