package Binding

import (
    "minsk/CodeAnalysis/Symbols"
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/Util"
)

type BoundCallExpression struct {
    *Util.ChildrenProvider

    Function *Symbols.FunctionSymbol
    Arguments []BoundExpression
}

func NewBoundCallExpression(function *Symbols.FunctionSymbol, arguments []BoundExpression) *BoundCallExpression {
    var children []interface{}

    for _, argument := range arguments {
        children = append(children, interface{}(argument))
    }

    return &BoundCallExpression{
        ChildrenProvider: Util.NewChildrenProvider(children...),

        Function: function,
        Arguments: arguments,
    }
}

func (b *BoundCallExpression) Kind() BoundNodeKind.BoundNodeKind {
    return BoundNodeKind.CallExpression
}

func (b *BoundCallExpression) GetProperties() []*BoundNodeProperty {
    return []*BoundNodeProperty{}
}

func (b *BoundCallExpression) GetType() *Symbols.TypeSymbol {
    return b.Function.Type
}
