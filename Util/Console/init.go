package Console

import (
    "os/exec"
    "fmt"
)

func Init() {
    //clean screen
    fmt.Print("\033[2J")
    SetCursorPos(1,1)

    //disable input buffering
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    //do not display entered characters on the screen
    exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

