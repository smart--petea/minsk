package Binding

import (
    "minsk/Util"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
)

type BoundGotoStatement struct {
    *Util.ChildrenProvider

    Label *BoundLabel
}

func NewBoundGotoStatement(label *BoundLabel) *BoundGotoStatement {
    return &BoundGotoStatement{
        ChildrenProvider: Util.NewChildrenProvider(),

        Label: label,
    }
}

func (b *BoundGotoStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.GotoStatement
}

func (b *BoundGotoStatement) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{
        {
            Name: "label",
            Value: b.Label.Name,
        },
    }
}
