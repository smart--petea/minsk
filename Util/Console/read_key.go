package Console

import (
    "os"
    "log"
    KeyInfo "minsk/Util/Console/KeyInfo"
)

func ReadKey() *KeyInfo.KeyInfo {
    b := make([]byte, 8)
    size, _ := os.Stdin.Read(b)
    log.Printf("Console.ReadKey %+v %s", b, string(b))

    return KeyInfo.NewKeyInfo(b[:size])
}
