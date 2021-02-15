package Syntax

import (
    "io"
    "os"
    "fmt"
    "strings"

    "minsk/Util/Console"
    "minsk/CodeAnalysis/Text"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type SyntaxNode interface {
    Kind() SyntaxKind.SyntaxKind
    Value()  interface{}
    GetChildren() <-chan SyntaxNode 
}

func SyntaxNodeToTextSpan(sn SyntaxNode) *Text.TextSpan {
    if syntaxToken, ok := sn.(*SyntaxToken); ok {
        return Text.NewTextSpan(syntaxToken.Position, len(syntaxToken.Runes)) 
    }

    children := sn.GetChildren()
    first := <-children
    last := first
    for last = range children {}

    start := SyntaxNodeToTextSpan(first).Start
    end := SyntaxNodeToTextSpan(last).End()
    return Text.NewTextSpan(start, end) 
}

func prettyPrint(writer io.StringWriter, node SyntaxNode, indent string, isLast bool) {
    isToConsole := (writer == os.Stdout)
    if fmt.Sprintf("%v", node) == "<nil>" {
        return
    }

    var marker string
    if isLast {
        marker = "└─"
    } else {
        marker = "├─"
    }

    writer.WriteString(indent)

    if isToConsole {
        Console.ForegroundColour(Console.COLOUR_GRAY)
        writer.WriteString(marker)
        Console.ResetColour()
    }

    _, isSyntaxToken := node.(*SyntaxToken) 
    if isToConsole {
        if isSyntaxToken {
            Console.ForegroundColour(Console.COLOUR_BLUE)
        } else {
            Console.ForegroundColour(Console.COLOUR_CYAN)
        }
    }

    kindS := fmt.Sprintf("%s", node.Kind())
    writer.WriteString(kindS)

    if isSyntaxToken && node.Value() != nil {
        var s string
        switch val := node.Value().(type) {
        case int:
            s = fmt.Sprintf(" %d", val)
        default:
            s = fmt.Sprintf(" %s", val)
        }

        writer.WriteString(s)
    }
    if isToConsole {
        Console.ResetColour()
    }
    writer.WriteString("\n")

    if isLast {
        indent = indent + "    "
    } else {
        indent = indent + "│   "
    }

    var nextChild, prevChild SyntaxNode
    var ok bool
    children := node.GetChildren()
    prevChild, ok = <-children 
    for ok {
        nextChild, ok = <-children
        if ok {
            prettyPrint(writer, prevChild, indent, false)
            prevChild = nextChild
        }
    }
    prettyPrint(writer, prevChild, indent, true)
}

func WriteTo(writer io.StringWriter, node SyntaxNode) {
    indent := ""
    isLast := true
    prettyPrint(writer, node, indent, isLast)
}

func ToString(node SyntaxNode) string {
    writer := &strings.Builder{}
    WriteTo(writer, node)

    return writer.String()
}
