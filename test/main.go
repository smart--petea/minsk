package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"

    CA "minsk/CodeAnalysis"
    "minsk/CodeAnalysis/Syntax"
    "minsk/CodeAnalysis/Text"
    "minsk/Util/Console"
    "minsk/Util"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    var showTree bool
    variables := make(map[*Util.VariableSymbol]interface{})
    var textBuilder strings.Builder
    var previous *CA.Compilation

    for {
        Console.ForegroundColour(Console.COLOUR_GREEN)
        if textBuilder.Len() == 0 {
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

        if textBuilder.Len() == 0 {
            if isBlank {
                break
            }

            if strings.TrimSpace(input) == "#showTree" {
                showTree = !showTree
                if showTree {
                    fmt.Println("Showing parse trees.")
                } else {
                    fmt.Println("Not showing parse trees")
                }

                continue
            } else if strings.TrimSpace(input) == "#cls" {
                Console.Clear()
                continue
            } else if strings.TrimSpace(input) == "#reset" {
                previous = nil
                continue
            }
        }

        textBuilder.WriteString(input)
        text := textBuilder.String()

        syntaxTree := Syntax.ParseSyntaxTree(text)
        if !isBlank && len(syntaxTree.GetDiagnostics()) > 0 {
            continue
        }

        var compilation *CA.Compilation
        if previous == nil {
            compilation = CA.NewCompilation(syntaxTree)
        } else {
            compilation = previous.ContinueWith(syntaxTree)
        }

        result := compilation.Evaluate(variables)

        if showTree {
            Console.ForegroundColour(Console.COLOUR_GRAY)
            Syntax.WriteTo(os.Stdout, syntaxTree.Root)
            Console.ResetColour()
        } 

        if len(result.Diagnostics) == 0  {
            Console.ForegroundColour(Console.COLOUR_MAGENTA)
            fmt.Println(result.Value)
            Console.ResetColour()

            previous = compilation
        } else {
            for _, diagnostic := range result.Diagnostics {
                lineIndex := syntaxTree.Text.GetLineIndex(diagnostic.Span.Start)
                lineNumber := lineIndex + 1
                line := syntaxTree.Text.Lines[lineIndex]
                character := diagnostic.Span.Start - line.Start + 1

                Console.ForegroundColour(Console.COLOUR_RED)
                fmt.Printf("(%d, %d): ", lineNumber, character)
                fmt.Println(diagnostic)
                Console.ResetColour()

                prefixSpan := Text.NewTextSpan(line.Start,diagnostic.Span.Start - line.Start)
                suffixSpan := Text.NewTextSpan(diagnostic.Span.End(),line.End() - diagnostic.Span.End())

                prefix := syntaxTree.Text.StringBySpan(prefixSpan)
                errorStr := syntaxTree.Text.StringBySpan(diagnostic.Span)
                suffix := syntaxTree.Text.StringBySpan(suffixSpan)

                fmt.Printf("    ")
                fmt.Printf("%s", prefix)

                Console.ForegroundColour(Console.COLOUR_RED)
                fmt.Print(errorStr)
                Console.ResetColour()

                fmt.Print(suffix)

                fmt.Println()
            }
        }

        textBuilder.Reset()
    }
}
