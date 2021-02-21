package Binding

import (
    "minsk/Util"
    "minsk/CodeAnalysis/Syntax"
)

type BoundGlobalScope struct {
    Util.DiagnosticBag
    Previous *BoundGlobalScope
    Variables []*Util.VariableSymbol
    Statement BoundStatement
}

func NewBoundGlobalScope(previous *BoundGlobalScope, variables []*Util.VariableSymbol, statement BoundStatement) *BoundGlobalScope {
    return &BoundGlobalScope{
        Previous: previous,
        Variables: variables,
        Statement: statement,
    }
}

func BoundGlobalScopeFromCompilationUnitSyntax(previous *BoundGlobalScope, syntax *Syntax.CompilationUnitSyntax) *BoundGlobalScope {
    parentScope := CreateParentScopes(previous)
    binder := NewBinder(parentScope)
    statement := binder.BindStatement(syntax.Statement)
    variables := binder.scope.GetDeclaredVariables()

    boundGlobalScope := NewBoundGlobalScope(previous, variables, statement)
    boundGlobalScope.AddDiagnosticsRange(binder.GetDiagnostics())
    if previous != nil {
        boundGlobalScope.AddDiagnosticsRange(previous.GetDiagnostics())
    }

    return boundGlobalScope
}

func CreateParentScopes(previous *BoundGlobalScope) *BoundScope {
    stack := Util.NewStack()
    for previous != nil {
        stack.Push(interface{}(previous))
        previous = previous.Previous
    }

    var parent *BoundScope
    for stack.Count() > 0 {
        previous = stack.Pop().(*BoundGlobalScope)

        scope := NewBoundScope(parent)
        for _, v := range previous.Variables {
            scope.TryDeclare(v)
        }

        parent = scope
    }

    return parent
}
