package Binding

import (
    "minsk/CodeAnalysis/Symbols"
    "minsk/Util"
)

type BoundProgram struct {
    *Util.DiagnosticBag

    GlobalScope *BoundGlobalScope
    FunctionBodies map[*Symbols.FunctionSymbol]*BoundBlockStatement
}

func NewBoundProgram(globalScope *BoundGlobalScope, diagnostics *Util.DiagnosticBag, functionBodies map[*Symbols.FunctionSymbol]*BoundBlockStatement) *BoundProgram {
    if diagnostics == nil {
        diagnostics = Util.NewDiagnosticBag()
    }

    return &BoundProgram{
        DiagnosticBag: diagnostics,

        GlobalScope: globalScope,
        FunctionBodies: functionBodies,
    }
}
