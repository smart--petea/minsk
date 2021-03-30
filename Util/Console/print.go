package Console

import (
    "fmt"
)

func Print(text string) {
    left, top := GetCursorPos()

    fmt.Printf(text)

    SetCursorPos(top, left)
}
