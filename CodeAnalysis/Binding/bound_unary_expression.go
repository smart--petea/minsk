package Binding

import (
    "minsk/CodeAnalysis/Binding/BoundUnaryOperator"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"

    "reflect"
)

type BoundUnaryExpression struct {
    *boundNodeChildren

    Operand BoundExpression
    Op *BoundUnaryOperator.BoundUnaryOperator
}

func NewBoundUnaryExpression(op *BoundUnaryOperator.BoundUnaryOperator, operand BoundExpression) *BoundUnaryExpression {
    return &BoundUnaryExpression{
        boundNodeChildren: newBoundNodeChildren(operand),

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
