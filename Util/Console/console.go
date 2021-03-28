package Console

import (
    "fmt"
    "log"
)

type Colour string

const (
    COLOUR_DARK_RED Colour = "[31m"
    COLOUR_RED Colour = "[31m"
    COLOUR_GREEN Colour = "[32m"
    COLOUR_DARK_YELLOW Colour = "[33m"
    COLOUR_BLUE Colour = "[34m"
    COLOUR_MAGENTA Colour = "[35m"
    COLOUR_CYAN Colour = "[36m"
    COLOUR_WHITE Colour = "[37m"
    COLOUR_GRAY Colour = "[90m"
    COLOUR_YELLOW Colour = "[93m"
)

func ResetColour() {
    log.Printf("Console.ResetColour")
    ForegroundColour(COLOUR_WHITE)
}

func ForegroundColour(colour Colour) {
    log.Printf("Console.ForegroundColor %s", string(colour))
    fmt.Printf("\033%s", colour)
}

func Clear() {
    fmt.Print("\033[H\033[2J")
}
