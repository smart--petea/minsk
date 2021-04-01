package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
    "minsk/CodeAnalysis/Symbols"
)

type BoundErrorExpression struct {
    *Util.ChildrenProvider
}

func NewBoundErrorExpression() *BoundErrorExpression {
    return &BoundErrorExpression{
        ChildrenProvider: Util.NewChildrenProvider(),
    }
}

func (b *BoundErrorExpression) GetType() *Symbols.TypeSymbol {
    return Symbols.TypeSymbolError
}

func (b *BoundErrorExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.ErrorExpression
}

func (b *BoundErrorExpression) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{
        {
            Name: "type",
            Value: b.GetType(),
        },
    }
}
