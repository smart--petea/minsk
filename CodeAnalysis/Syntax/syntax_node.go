package Syntax

import (
    "io"
    "os"
    "fmt"
    "strings"

    "minsk/Util/Console"
    "minsk/CodeAnalysis/Text"
)

type SyntaxNode interface {
    CoreSyntaxNode

    GetSpan() *Text.TextSpan
}

func SyntaxNodeToTextSpan(coreSyntaxNode CoreSyntaxNode) *Text.TextSpan {
    if syntaxToken, ok := coreSyntaxNode.(*SyntaxToken); ok {
        return Text.NewTextSpan(syntaxToken.Position, len(syntaxToken.Runes)) 
    }

    children := coreSyntaxNode.GetChildren()
    first := <-children
    last := first
    for last = range children {}

    start := SyntaxNodeToTextSpan(first.(CoreSyntaxNode)).Start
    end := SyntaxNodeToTextSpan(last.(CoreSyntaxNode)).End()
    return Text.NewTextSpan(start, end - start) 
}

func prettyPrint(writer io.StringWriter, nodeI interface{}, indent string, isLast bool) {
    isToConsole := (writer == os.Stdout)
    if fmt.Sprintf("%v", nodeI) == "<nil>" {
        return
    }
    node := nodeI.(CoreSyntaxNode)

    var marker string
    if isLast {
        marker = "└─"
    } else {
        marker = "├─"
    }

    writer.WriteString(indent)

    if isToConsole {
        Console.ForegroundColour(Console.COLOUR_GRAY)
    }
    writer.WriteString(marker)
    Console.ResetColour()

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

    var nextChild, prevChild interface{}
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

func WriteTo(writer io.StringWriter, node CoreSyntaxNode) {
    indent := ""
    isLast := true
    prettyPrint(writer, node, indent, isLast)
}

func ToString(node CoreSyntaxNode) string {
    writer := &strings.Builder{}
    WriteTo(writer, node)

    return writer.String()
}

func SyntaxNodeGetLastToken(node SyntaxNode) *SyntaxToken {
    if syntaxToken, ok := node.(*SyntaxToken); ok {
        return syntaxToken
    }

    var lastChild, tmp interface{}
    var ok bool
    children := node.GetChildren()
    lastChild, ok = <-children 
    for ok {
        tmp, ok = <-children
        if ok {
            lastChild = tmp
        }
    }

    lastSyntaxToken := SyntaxNodeGetLastToken(lastChild.(SyntaxNode))
    return lastSyntaxToken
}
