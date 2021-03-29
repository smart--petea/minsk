package Console

import (
    "os"
    "fmt"
    "golang.org/x/sys/unix"
    "syscall"
    //"log"
)

func SetCursorPos(left, top int) {
    fmt.Printf("\033[%d;%dH", top, left)
}

func SetCursorVisibile(v bool) {
    if v {
        fmt.Printf("\033]?25h")
    } else {
        fmt.Printf("\033]?25l")
    }
}

func GetCursorPos() (left int, top int) {
    fd := os.Stdin.Fd()
    term, err := unix.IoctlGetTermios(int(fd), unix.TCGETS)
    if err != nil {
        panic(err)
    }

    restore := *term
    term.Lflag &^= (syscall.ICANON | syscall.ECHO)
    err = unix.IoctlSetTermios(int(fd), unix.TCSETS, term)
    if err != nil {
        panic(err)
    }

    _, err = os.Stdout.Write([]byte("\033[6n"))//todo save []byte("\033[6n") in a constant
    if err != nil {
        panic(err)
    }

    var x, y int
    var zeroes int = 1
    var firstNumber bool = true
    b := make([]byte, 1)
    for {
        _, err = os.Stdin.Read(b)
        if err != nil {
            panic(err)
        }

        if b[0] == 'R' {
            break
        }

        if b[0] == ';' {
            firstNumber = false
            zeroes = 1
            continue
        }

        if b[0] >= '0' && b[0] <= '9' {
            if firstNumber {
                x = x * zeroes + int(b[0]) - int('0') //todo constants
            } else {
                y = y * zeroes + int(b[0]) - int('0') //todo constants
            }
            zeroes *= 10
            continue
        }
    }

    err = unix.IoctlSetTermios(int(fd), unix.TCSETS, &restore)
    if err != nil {
        panic(err)
    }

    return x, y
}

