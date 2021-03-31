package minsk

import (
    "fmt"
    "os"
    "strings"
    "log"

    CA "minsk/CodeAnalysis"
    "minsk/CodeAnalysis/Syntax"
    "minsk/CodeAnalysis/Symbols"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util/Console"
)

func (m *MinskRepl) EvaluateMetaCommand(input string) {
    switch strings.TrimSpace(input) {
    case "#showTree":
        m.showTree = !m.showTree
        if m.showTree {
            fmt.Println("Showing parse trees.")
        } else {
            fmt.Println("Not showing parse trees")
        }
    case "#showProgram":
        m.showProgram = !m.showProgram
        if m.showProgram {
            fmt.Println("Showing bound tree.")
        } else {
            fmt.Println("Not showing bound tree.")
        }
    case "#cls":
        Console.Clear()
    case "#reset":
        m.previous = nil
        m.variables = make(map[*Symbols.VariableSymbol]interface{})
    default:
        m.Repl.EvaluateMetaCommand(input)
    }
}


type MinskRepl struct {
    Repl

    previous *CA.Compilation
    showTree bool
    showProgram bool
    variables map[*Symbols.VariableSymbol]interface{}
}

func (m *MinskRepl) RenderLine(line string) {
    tokens := Syntax.ParseTokens(line)
    log.Printf("RenderLine line=%s len(tokens)=%d", line, len(tokens))
    for _, token := range tokens {
        isKeyword := strings.HasSuffix(string(token.Kind()), "Keyword")
        isNumber :=  token.Kind() == SyntaxKind.NumberToken
        isIdentifier :=  token.Kind() == SyntaxKind.IdentifierToken

        if isKeyword {
            Console.ForegroundColour(Console.COLOUR_BLUE)
        } else if isIdentifier {
            Console.ForegroundColour(Console.COLOUR_DARK_YELLOW)
        } else if isNumber {
            Console.ForegroundColour(Console.COLOUR_CYAN)
        } else {
            Console.ForegroundColour(Console.COLOUR_GRAY)
        }

        fmt.Print(string(token.Runes))

        Console.ResetColour()
    }

    Console.Print(strings.Repeat(" ", Console.WindowWidth() - len(line)))
}

func (m *MinskRepl) EvaluateSubmission(text string) {
    syntaxTree := Syntax.SyntaxTreeParse(text)

    var compilation *CA.Compilation
    if m.previous == nil {
        compilation = CA.NewCompilation(syntaxTree)
    } else {
        compilation = m.previous.ContinueWith(syntaxTree)
    }

    if m.showTree {
        Console.ForegroundColour(Console.COLOUR_GRAY)
        Syntax.WriteTo(os.Stdout, syntaxTree.Root)
        Console.ResetColour()
    } 

    if m.showProgram {
        Console.ForegroundColour(Console.COLOUR_GRAY)
        compilation.EmitTree(os.Stdout)
        Console.ResetColour()
    } 

    result := compilation.Evaluate(m.variables)

    if len(result.Diagnostics) == 0  {
        Console.ForegroundColour(Console.COLOUR_MAGENTA)
        fmt.Println(result.Value)
        Console.ResetColour()

        m.previous = compilation
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
}

func MinskReplLastTwoLinesAreBlank(text string) bool {
    lines := strings.Split(text, "\n")
    l := len(lines)

    if l < 2 {
        return false
    }

    return lines[l-1] == "" && lines[l-2] == ""
}

func (m *MinskRepl) IsCompleteSubmission(text string) bool {
    if len(text) == 0 {
        return true
    }

    if MinskReplLastTwoLinesAreBlank(text) {
        return true
    }

    syntaxTree := Syntax.SyntaxTreeParse(text)
    if Syntax.SyntaxNodeGetLastToken(syntaxTree.Root.Statement).IsMissing() {
        return false
    }

    return true
}

func NewMinskRepl() *MinskRepl {
    m := MinskRepl{
        variables: make(map[*Symbols.VariableSymbol]interface{}),
    }
    m.Repl.IsCompleteSubmission = m.IsCompleteSubmission
    m.Repl.EvaluateSubmission = m.EvaluateSubmission

    return &m
}
