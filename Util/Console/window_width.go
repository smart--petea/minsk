package Console

import (
    "os"
    "os/exec"
    "strings"
    "strconv"
)

func WindowWidth() int {
    cmd := exec.Command("stty", "size")
    cmd.Stdin = os.Stdin
    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }

    fields := strings.Fields(string(out))
    i, err := strconv.Atoi(fields[1])
    if err != nil {
        panic(err)
    }

    return i
}
