package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundIfStatement struct {
    Condition BoundExpression 
    ThenStatement BoundStatement
    ElseStatement BoundStatement
}

func NewBoundIfStatement(condition BoundExpression, thenStatement, elseStatement BoundStatement) *BoundIfStatement {
    return &BoundIfStatement{
        Condition: condition,
        ThenStatement: thenStatement,
        ElseStatement: elseStatement,
    }
}

func (b *BoundIfStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.IfStatement
}
