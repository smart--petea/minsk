package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/CodeAnalysis/Symbols"
    "minsk/Util"
)

type BoundConversionExpression struct {
    *Util.ChildrenProvider

    Type *Symbols.TypeSymbol
    Expression BoundExpression
}

func NewBoundConversionExpression(ttype *Symbols.TypeSymbol, expression BoundExpression) *BoundConversionExpression {
    return &BoundConversionExpression{
        ChildrenProvider: Util.NewChildrenProvider(expression),

        Type: ttype,
        Expression: expression,
    }
}

func (b *BoundConversionExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.ConversionExpression
}

func (b *BoundConversionExpression) GetType() *Symbols.TypeSymbol {
    return b.Type
}

func (b *BoundConversionExpression) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{
        {
            Name: "type",
            Value: b.GetType(),
        },
    }
}
