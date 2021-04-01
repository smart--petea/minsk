package Binding

import (
    "minsk/CodeAnalysis/Symbols"
)

type BoundExpression interface {
    BoundNode

    GetType() *Symbols.TypeSymbol
}
