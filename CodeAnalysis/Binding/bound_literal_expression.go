package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/CodeAnalysis/Symbols"
    "minsk/Util"

    "fmt"
)

type BoundLiteralExpression struct {
    *Util.ChildrenProvider

    Type *Symbols.TypeSymbol
    Value interface{}
}

func NewBoundLiteralExpression(value interface{}) *BoundLiteralExpression {
    var t *Symbols.TypeSymbol
    switch value.(type) {
    case int:
        t = Symbols.TypeSymbolInt
    case bool:
        t = Symbols.TypeSymbolBool
    case string:
        t = Symbols.TypeSymbolString
    default:
        panic(fmt.Sprintf("Unexpected literal %+v of type %t", value, value))
    }

    return &BoundLiteralExpression{
        ChildrenProvider: Util.NewChildrenProvider(),

        Type: t,
        Value: value,
    }
}

func (b *BoundLiteralExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.LiteralExpression
}

func (b *BoundLiteralExpression) GetType() *Symbols.TypeSymbol {
    return b.Type
}

func (b *BoundLiteralExpression) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{
        {
            Name: "type",
            Value: b.GetType(),
        },
        {
            Name: "value",
            Value: b.Value,
        },
    }
}
