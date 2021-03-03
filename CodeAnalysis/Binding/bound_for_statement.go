package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
)

type BoundForStatement struct {
    *boundNodeChildren

    Variable *Util.VariableSymbol
    LowerBound BoundExpression
    UpperBound BoundExpression
    Body BoundStatement
}

func NewBoundForStatement(variable *Util.VariableSymbol, lowerBound BoundExpression, upperBound BoundExpression, body BoundStatement) *BoundForStatement {
    return &BoundForStatement{
        boundNodeChildren: newBoundNodeChildren(lowerBound, upperBound, body),

        Variable: variable,
        LowerBound: lowerBound,
        UpperBound: upperBound,
        Body: body,
    }
}

func (b *BoundForStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.ForStatement
}
