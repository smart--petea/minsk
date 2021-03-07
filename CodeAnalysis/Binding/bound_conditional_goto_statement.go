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
        ChildrenProvider: Util.NewChildrenProvider(label, condition, jumpIfFalse),

        Label: label,
        Condition: condition,
        JumpIfFalse: jumpIfFalse,
    }
}

func (b *BoundConditionalGotoStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.ConditionalGotoStatement
}

func (b *BoundConditionalGotoStatement) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{}
}
