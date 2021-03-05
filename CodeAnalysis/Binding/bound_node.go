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
    GetChildren() <-chan interface{} 
    GetProperties() []*BoundNodeProperty
}

func prettyPrint(writer io.StringWriter, nodeI interface{}, indent string, isLast bool) {
    isToConsole := (writer == os.Stdout)
    if fmt.Sprintf("%v", nodeI) == "<nil>" {
        return
    }
    node := nodeI.(BoundNode)

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

    if isToConsole {
        Console.ForegroundColour(getColor(node))
    }

    text := getText(node)
    writer.WriteString(text)

    isFirstProperty := true
    for _, p := range node.GetProperties() {
        if isFirstProperty {
            isFirstProperty = false
        } else {
            if isToConsole {
                Console.ForegroundColour(Console.COLOUR_GRAY)
            }

            writer.WriteString(", ")
        }
        writer.WriteString(" ")

        if isToConsole {
            Console.ForegroundColour(Console.COLOUR_YELLOW)
        }
        writer.WriteString(p.Name)

        if isToConsole {
            Console.ForegroundColour(Console.COLOUR_GRAY)
        }
        writer.WriteString(" = ")

        if isToConsole {
            Console.ForegroundColour(Console.COLOUR_DARK_YELLOW)
        }
        writer.WriteString(fmt.Sprintf("%v", p.Value))
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

func getText(node BoundNode) string {
    if b, ok := node.(*BoundBinaryExpression); ok {
        return fmt.Sprintf("%s Expression", b.Op.Kind)
    }

    if u, ok := node.(*BoundUnaryExpression); ok {
        return fmt.Sprintf("%s Expression", u.Op.Kind)
    }

    return fmt.Sprintf("%s", node.Kind())
}

