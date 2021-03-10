package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
)

type BoundConditionalGotoStatement struct {
    *Util.ChildrenProvider

    Label *Util.LabelSymbol
    Condition BoundExpression
    JumpIfFalse bool
}

func NewBoundConditionalGotoStatement(label *Util.LabelSymbol, condition BoundExpression, jumpIfFalse bool) *BoundConditionalGotoStatement {
    return &BoundConditionalGotoStatement{
        ChildrenProvider: Util.NewChildrenProvider(condition),

        Label: label,
        Condition: condition,
        JumpIfFalse: jumpIfFalse,
    }
}

func (b *BoundConditionalGotoStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.ConditionalGotoStatement
}

func (b *BoundConditionalGotoStatement) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{
        {
            Name: "Label",
            Value: b.Label.Name,
        },
        {
            Name: "JumpIfFalse",
            Value: b.JumpIfFalse,
        },
    }
}
