package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
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

        if line == "1 + 2 * 3" {
            fmt.Println("7")
        } else {
            fmt.Println("ERROR: Invalid expression!")
        }
    }
}

type SyntaxKind int

type SyntaxToken struct {
    kind SyntaxToken
    position int
    text string
}

func NewSyntaxToken(kind SyntaxToken, position int, text string) *SyntaxToken {
    return &SyntaxToken{
        kind: SyntaxToken,
        position: int,
        text: string,
    }
}

type Lexer struct {
    text string
    position int
}

func NewLexer(text string) *Lexer {
    return &Lexer{
        text: text
    }
}

func (l *Lexer) NextToken() *SyntaxToken {
    //<numbers>
    //+ - * / ( )
    //<whitespace>
}
