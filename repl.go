package minsk

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"

    "minsk/Util/Console"
)

type Repl struct {
    textBuilder strings.Builder

    EvaluateSubmission func(text string) 
    IsCompleteSubmission func(text string) bool
}

func (r *Repl) Run(ir IRepl) {
    reader := bufio.NewReader(os.Stdin)

    for {
        Console.ForegroundColour(Console.COLOUR_GREEN)
        if r.textBuilder.Len() == 0 {
            fmt.Print("» ")
        } else {
            fmt.Print("· ")
        }
        Console.ResetColour()

        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }
        isBlank := len(strings.TrimSpace(input)) == 0

        if r.textBuilder.Len() == 0 {
            if isBlank {
                break
            } else if (strings.HasPrefix(input, "#")) {
                ir.EvaluateMetaCommand(input)
                continue
            }
        }

        r.textBuilder.WriteString(input)
        text := r.textBuilder.String()

        if !ir.IsCompleteSubmission(text) {
            continue
        }

        ir.EvaluateSubmission(text)

        r.textBuilder.Reset()
    }
}

func (r *Repl) EvaluateMetaCommand(input string) {
    Console.ForegroundColour(Console.COLOUR_RED)
    fmt.Printf("Invalid command %s.", input)
    Console.ResetColour()
}
