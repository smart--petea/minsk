package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
)

type BoundLabelStatement struct {
    *Util.ChildrenProvider

    Label *BoundLabel
}

func NewBoundLabelStatement(label *BoundLabel) *BoundLabelStatement {
    return &BoundLabelStatement{
        ChildrenProvider: Util.NewChildrenProvider(),

        Label: label,
    }
}

func (b *BoundLabelStatement) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.LabelStatement
}

func (b *BoundLabelStatement) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{
        {
            Name: "label",
            Value: b.Label.Name,
        },
    }
}
