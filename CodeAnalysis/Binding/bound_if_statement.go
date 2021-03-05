package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
)

type BoundIfStatement struct {
    *Util.ChildrenProvider

    Condition BoundExpression 
    ThenStatement BoundStatement
    ElseStatement BoundStatement
}

func NewBoundIfStatement(condition BoundExpression, thenStatement, elseStatement BoundStatement) *BoundIfStatement {
    return &BoundIfStatement{
        ChildrenProvider: Util.NewChildrenProvider(condition, thenStatement, elseStatement),

        Condition: condition,
        ThenStatement: thenStatement,
        ElseStatement: elseStatement,
    }
}

func (b *BoundIfStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.IfStatement
}

//todo
func (b *BoundIfStatement) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{}
}
