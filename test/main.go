package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"
//    "unicode"
 //   "strconv"

    "minsk"
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

        parser := minsk.NewParser(line)
        expression := parser.Parse()

        fmt.Print("\033[90m")
        PrettyPrint(expression, "")
        fmt.Print("\033[37m")
    }
}

func PrettyPrint(node minsk.SyntaxNode, indent string) {
    fmt.Printf("%s%s", indent, node.Kind())

    if node.Value() != nil {
        fmt.Print(" ")
        fmt.Print(node.Value())
    }
    fmt.Printf("\n")

    indent = indent + "   "
    for _, child := range node.GetChildren() {
        PrettyPrint(child, indent)
    }
}
