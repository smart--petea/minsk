package Console

import (
    "os"
    KeyInfo "minsk/Util/Console/KeyInfo"
)

func ReadKey() *KeyInfo.KeyInfo {
    b := make([]byte, 8)
    size, _ := os.Stdin.Read(b)

    return KeyInfo.NewKeyInfo(b[:size])
}
