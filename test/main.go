package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"

    CA "minsk/CodeAnalysis"
    "minsk/CodeAnalysis/Syntax"
    "minsk/Util/Console"
    "minsk/Util"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    var showTree bool
    variables := make(map[*Util.VariableSymbol]interface{})

    for {
        fmt.Print("> ")
        line, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        line = strings.TrimSpace(line)
        if len(line) == 0 {
            os.Exit(0)
        }

        if line == "#showTree" {
            showTree = !showTree
            if showTree {
                fmt.Println("Showing parse trees.")
            } else {
                fmt.Println("Not showing parse trees")
            }

            continue
        } else if line == "#cls" {
            Console.Clear()
            continue
        }

        syntaxTree := Syntax.ParseSyntaxTree(line)
        compilation := CA.NewCompilation(syntaxTree)
        result := compilation.Evaluate(variables)

        if showTree {
            Console.ForegroundColour(Console.COLOUR_GRAY)
            Syntax.WriteTo(os.Stdout, syntaxTree.Root)
            Console.ResetColour()
        } 

        if len(result.Diagnostics) == 0  {
            fmt.Println(result.Value)
        } else {
            text := syntaxTree.Text

            for _, diagnostic := range result.Diagnostics {
                lineIndex := text.GetLineIndex(diagnostic.Span.Start)
                lineNumber := lineIndex + 1
                character := diagnostic.Span.Start - text.Lines[lineIndex].Start + 1

                Console.ForegroundColour(Console.COLOUR_RED)
                fmt.Printf("(%d, %d): ", lineNumber, character)
                fmt.Println(diagnostic)
                Console.ResetColour()

                span := diagnostic.Span
                prefix := line[:span.Start]
                errorStr := line[span.Start: span.End()]
                suffix := line[span.End():]

                fmt.Printf("    ")
                fmt.Printf("%s", prefix)

                Console.ForegroundColour(Console.COLOUR_RED)
                fmt.Print(errorStr)
                Console.ResetColour()

                fmt.Print(suffix)

                fmt.Println()
            }
        } 
    }
}
