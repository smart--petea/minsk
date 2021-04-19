package Binding

import (
    "minsk/CodeAnalysis/Symbols"
    "minsk/Util"
)

type BoundProgram struct {
    GlobalScope *BoundGlobalScope
    Diagnostics *Util.DiagnosticBag
    FunctionBodies map[*Symbols.FunctionSymbol]*BoundBlockStatement
}

func NewBoundProgram(globalScope *BoundGlobalScope, diagnostics *Util.DiagnosticBag, functionBodies map[*Symbols.FunctionSymbol]*BoundBlockStatement) *BoundProgram {
    return &BoundProgram{
        //ChildrenProvider: Util.NewChildrenProvider(),

        GlobalScope: globalScope,
        Diagnostics: diagnostics,
        FunctionBodies: functionBodies,
    }
}
