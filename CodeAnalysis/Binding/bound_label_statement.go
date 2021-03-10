package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"

    "reflect"
)

type BoundLabelStatement struct {
    *Util.ChildrenProvider

    Label *Util.LabelSymbol
}

func NewBoundLabelStatement(label *Util.LabelSymbol) *BoundLabelStatement {
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
            Value: reflect.TypeOf(*b.Label).Name(),
        },
    }
}
