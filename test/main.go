package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "strings"

    CA "minsk/CodeAnalysis"
    Syntax "minsk/CodeAnalysis/Syntax"
    Console "minsk/Util/Console"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    var showTree bool

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
        result := compilation.Evaluate()

        if showTree {
            Console.ForegroundColour(Console.COLOUR_GRAY)
            PrettyPrint(syntaxTree.Root, "", true)
            Console.ResetColour()
        } 

        if len(result.Diagnostics) == 0  {
            fmt.Println(result.Value)
        } else {
            Console.ForegroundColour(Console.COLOUR_RED)
            for _, diagnostic := range result.Diagnostics {
                fmt.Println(diagnostic)
            }

            Console.ResetColour()
        } 
    }
}

func PrettyPrint(node Syntax.SyntaxNode, indent string, isLast bool) {
    if fmt.Sprintf("%v", node) == "<nil>" {
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
        indent = indent + "    "
    } else {
        indent = indent + "│   "
    }

    children := node.GetChildren()
    lenChildren := len(children) - 1
    for i, child := range children {
        PrettyPrint(child, indent, i == lenChildren)
    }
}
