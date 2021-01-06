package Console

import (
    "fmt"
)

type Colour string

const (
    COLOUR_DARK_RED Colour = "[31m"
    COLOUR_WHITE Colour = "[37m"
    COLOUR_RED Colour = "[31m"
    COLOUR_GRAY Colour = "[90m"
)

func ResetColour() {
    ForegroundColour(COLOUR_WHITE)
}

func ForegroundColour(colour Colour) {
    fmt.Printf("\033%s", colour)
}

func Clear() {
    fmt.Print("\033[H\033[2J")
}
