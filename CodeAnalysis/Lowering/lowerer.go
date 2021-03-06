package Lowering

import (
    "minsk/CodeAnalysis/Binding"
)

type Lowerer struct {
    Binding.BoundTreeRewriter
}

func newLowerer() *Lowerer {
    return &Lowerer{}
}

func LowererLower(statement Binding.BoundStatement) Binding.BoundStatement {
    lowerer := newLowerer()

    return lowerer.RewriteStatement(statement)
}
