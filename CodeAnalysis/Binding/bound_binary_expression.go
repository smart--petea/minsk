package Binding

import (
    "minsk/CodeAnalysis/Binding/BoundBinaryOperator"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/CodeAnalysis/Symbols"
    "minsk/Util"
)

type BoundBinaryExpression struct {
    *Util.ChildrenProvider

    Left BoundExpression
    Op *BoundBinaryOperator.BoundBinaryOperator
    Right BoundExpression
}

func NewBoundBinaryExpression(left BoundExpression, op *BoundBinaryOperator.BoundBinaryOperator, right BoundExpression) *BoundBinaryExpression {
    return &BoundBinaryExpression{
        ChildrenProvider: Util.NewChildrenProvider(left, right),

        Left: left,
        Op: op,
        Right: right,
    }
}

func (b *BoundBinaryExpression) Kind() BoundNodeKind.BoundNodeKind {
   return BoundNodeKind.BinaryExpression 
}

func (b *BoundBinaryExpression) GetType() *Symbols.TypeSymbol {
    return b.Op.ResultType
}

func (b *BoundBinaryExpression) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{
        {
            Name: "type",
            Value: b.GetType(),
        },
    }
}
