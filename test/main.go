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
        syntaxTree := parser.Parse()

        fmt.Print("\033[90m")
        PrettyPrint(syntaxTree.Root, "", true)
        fmt.Print("\033[37m")

        if len(syntaxTree.Diagnostics) > 0  {
            fmt.Print("\033[31m")

            for _, diagnostic := range parser.Diagnostics {
                fmt.Println(diagnostic)
            }

            fmt.Print("\033[37m")
        } else {
            e := minsk.NewEvaluator(syntaxTree.Root)
            result := e.Evaluate()
            fmt.Println(result)
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

    fmt.Printf("acesta r.62 %+v %+v\n", node, node==nil)
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
    fmt.Printf("acesta r.78 %+v\n", node)
    fmt.Printf("acesta r.79 %+v\n", node.GetChildren())
    fmt.Printf("acesta r.80 %+v\n", len(node.GetChildren()))
    for i, child := range children {
        fmt.Printf("acesta r.80 %+v\n", node)
        fmt.Printf("acesta r.81 %+v %+v\n", child, child==nil)
        PrettyPrint(child, indent, i == lenChildren)
    }
}
