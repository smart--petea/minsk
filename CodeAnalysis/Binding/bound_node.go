package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util/Console"
    "io"
    "os"
    "fmt"
)

type BoundNode interface {
    Kind() BoundNodeKind.BoundNodeKind
    GetChildren() <-chan BoundNode 
}

func prettyPrint(writer io.StringWriter, node BoundNode, indent string, isLast bool) {
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
    }
    writer.WriteString(marker)
    Console.ResetColour()

    /*
    _, isSyntaxToken := node.(*SyntaxToken) 
    if isToConsole {
        if isSyntaxToken {
            Console.ForegroundColour(Console.COLOUR_BLUE)
        } else {
            Console.ForegroundColour(Console.COLOUR_CYAN)
        }
    }
    */

    writeNode(writer, node)
    writeProperties(writer, node)

    /*
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
    */

    if isToConsole {
        Console.ResetColour()
    }
    writer.WriteString("\n")

    if isLast {
        indent = indent + "    "
    } else {
        indent = indent + "│   "
    }

    var nextChild, prevChild BoundNode
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

func WriteTo(writer io.StringWriter, node BoundNode) {
    indent := ""
    isLast := true
    prettyPrint(writer, node, indent, isLast)
}

func getColor(node BoundNode) Console.Colour {
    if _, ok := node.(BoundExpression); ok {
        return Console.COLOUR_BLUE
    }

    if _, ok := node.(BoundStatement); ok {
        return Console.COLOUR_CYAN
    }

    return Console.COLOUR_YELLOW
}

func writeProperties(writer io.StringWriter, node BoundNode) {
}

func writeNode(writer io.StringWriter, node BoundNode) {
    Console.ForegroundColour(getColor(node))

    text := getText(node)
    writer.WriteString(text)

    Console.ResetColour()
}

func getText(node BoundNode) string {
    if b, ok := node.(*BoundBinaryExpression); ok {
        return fmt.Sprintf("%s Expression", b.Op.Kind)
    }

    if u, ok := node.(*BoundUnaryExpression); ok {
        return fmt.Sprintf("%s Expression", u.Op.Kind)
    }

    return fmt.Sprintf("%s", node.Kind())
}

