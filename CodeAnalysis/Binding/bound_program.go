package Binding

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundNodeKind"
    "minsk/CodeAnalysis/Symbols"
    "minsk/Util"
)

type BoundProgram struct {
    GlobalScope *BoundGlobalScope
    Diagnostics *DiagnosticBag
    FunctionBodies map[*Symbols.FunctionSymbol]*BoundBlockStatement
}

func NewBoundProgram(globalScope *BoundGlobalScope, diagnostics *DiagnosticBag, functionBodies map[*Symbols.FunctionSymbol]*BoundBlockStatement) *BoundProgram {
    return &BoundProgram{
        ChildrenProvider: Util.NewChildrenProvider(),

        GlobalScope: globalScope
        Diagnostics: diagnostics,
        FunctionBodies: functionBodies,
    }
}
