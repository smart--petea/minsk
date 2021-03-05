package Binding

import (
    "minsk/CodeAnalysis/Binding/BoundUnaryOperator"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"

    "reflect"
)

type BoundUnaryExpression struct {
    *Util.ChildrenProvider

    Operand BoundExpression
    Op *BoundUnaryOperator.BoundUnaryOperator
}

func NewBoundUnaryExpression(op *BoundUnaryOperator.BoundUnaryOperator, operand BoundExpression) *BoundUnaryExpression {
    return &BoundUnaryExpression{
        ChildrenProvider: Util.NewChildrenProvider(operand),

        Operand: operand,
        Op: op,
    }
}

func (b *BoundUnaryExpression) GetType() reflect.Kind {
    return b.Op.ResultType
}

func (b *BoundUnaryExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.UnaryExpression
}

func (b *BoundUnaryExpression) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{
        {
            Name: "type",
            Value: b.GetType(),
        },
        {
            Name: "operand",
            Value: reflect.TypeOf(b.Operand),
        },
        {
            Name: "op",
            Value: reflect.TypeOf(b.Op),
        },
    }
}
