package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
    "minsk/CodeAnalysis/Symbols"
)

type BoundAssignmentExpression struct {
    *Util.ChildrenProvider

    Variable *Symbols.VariableSymbol
    Expression BoundExpression
}

func NewBoundAssignmentExpression(variable *Symbols.VariableSymbol, expression BoundExpression) *BoundAssignmentExpression {
    return &BoundAssignmentExpression{
        ChildrenProvider: Util.NewChildrenProvider(expression),

        Variable: variable,
        Expression: expression,
    }
}

func (b *BoundAssignmentExpression) GetType() *Symbols.TypeSymbol {
    return b.Expression.GetType()
}

func (b *BoundAssignmentExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.AssignmentExpression
}

func (b *BoundAssignmentExpression) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{
        {
            Name: "type",
            Value: b.GetType(),
        },
        {
            Name: "variable",
            Value: b.Variable.Name,
        },
    }
}
