package Binding

import (
    "minsk/Util"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundGotoStatement struct {
    *Util.ChildrenProvider

    Label *Util.LabelSymbol
}

func NewBoundGotoStatement(label *Util.LabelSymbol) *BoundGotoStatement {
    return &BoundGotoStatement{
        ChildrenProvider: Util.NewChildrenProvider(label),

        Label: label,
    }
}

func (b *BoundGotoStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.GotoStatement
}

func (b *BoundBlockStatement) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{}
}
