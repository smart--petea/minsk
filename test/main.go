package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"

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
        PrettyPrint(expression, "", true)
        fmt.Print("\033[37m")

        if len(parser.Diagnostics) > 0  {
            fmt.Print("\033[31m")

            for _, diagnostic := range parser.Diagnostics {
                fmt.Println(diagnostic)
            }

            fmt.Print("\033[37m")
        }
    }
}

func PrettyPrint(node minsk.SyntaxNode, indent string, isLast bool) {
    if node == nil {
        return
    }

    var marker string
    if isLast {
        marker = "└─"
    } else {
        marker = "├─"
    }

    fmt.Printf("%s%s%s", indent, marker, node.Kind())

    if node.Value() != nil {
        fmt.Print(" ")
        fmt.Print(node.Value())
    }
    fmt.Printf("\n")

    if isLast {
        indent = indent + "     "
    } else {
        indent = indent + "│    "
    }

    children := node.GetChildren()
    lenChildren := len(children) - 1
    for i, child := range children {
        PrettyPrint(child, indent, i == lenChildren)
    }
}
