package Syntax

import (
    "io"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/Util"
    "fmt"
    "strings"
)

type SyntaxNode interface {
    Kind() SyntaxKind.SyntaxKind
    Value()  interface{}
    GetChildren() <-chan SyntaxNode 
}

func SyntaxNodeToTextSpan(sn SyntaxNode) *Util.TextSpan {
    if syntaxToken, ok := sn.(*SyntaxToken); ok {
        return Util.NewTextSpan(syntaxToken.Position, len(syntaxToken.Runes)) 
    }

    children := sn.GetChildren()
    first := <-children
    last := first
    for last = range children {}

    start := SyntaxNodeToTextSpan(first).Start
    end := SyntaxNodeToTextSpan(last).End()
    return Util.NewTextSpan(start, end) 
}

func prettyPrint(writer io.StringWriter, node SyntaxNode, indent string, isLast bool) {
    if fmt.Sprintf("%v", node) == "<nil>" {
        return
    }

    var marker string
    if isLast {
        marker = "└─"
    } else {
        marker = "├─"
    }

    s := fmt.Sprintf("%s%s%s", indent, marker, node.Kind())
    writer.WriteString(s)

    if node.Value() != nil {
        var s string
        switch val := node.Value().(type) {
        case int:
            s = fmt.Sprintf(" %d", val)
        default:
            s = fmt.Sprintf(" %s", val)
        }

        writer.WriteString(s)
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
