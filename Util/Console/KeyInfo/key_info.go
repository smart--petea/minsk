package KeyInfo

import (
//    "log"
    "fmt"
)

type KeyInfo struct {
    Bytes []byte
}

const (
    NoModifiers int = 0
    Alt int = 1
    Control int = 2
    Shift int = 4
)

func NewKeyInfo(bytes []byte) *KeyInfo {
    return &KeyInfo{
        Bytes: bytes,
    }
}

func (c *KeyInfo) Kind() ConsoleKey {
    if len(c.Bytes) == 4 {
        if c.Bytes[0] == 27 && c.Bytes[1] == 91 && c.Bytes[3] == 126 {
            switch  c.Bytes[2] {
            case 49:
                return Home
            case 51:
                return Delete
            case 52:
                return End
            case 53:
                return PageUp
            case 54:
                return PageDown
            }
        }
    }

    if len(c.Bytes) == 3 {
        if c.Bytes[0] == 27 && c.Bytes[1] == 91 {
            switch c.Bytes[2] {
            case 67:
                return RightArrow
            case 68:
                return LeftArrow
            case 65:
                return UpArrow
            case 66:
                return DownArrow
            default:
                panic(fmt.Sprintf("Unknown console command %+v", c.Bytes[:3]))
            }
        }
    }

    if len(c.Bytes) == 2 {
        if c.Bytes[0] == 27 && c.Bytes[1] == 10 {
            return AltEnter
        }
    }

    switch c.Bytes[0] {
    case 27:
        return Escape
    case 10:
        return Enter
    case 127:
        return Backspace
    case 9:
        return Tab
    }

    return Symbol
}
