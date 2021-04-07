package Binding

import (
    "minsk/Util"
    "minsk/CodeAnalysis/Syntax"
    "minsk/CodeAnalysis/Symbols"
    "minsk/CodeAnalysis/Symbols/BuiltinFunctions"
)

type BoundGlobalScope struct {
    Util.DiagnosticBag
    Previous *BoundGlobalScope
    Variables []*Symbols.VariableSymbol
    Statement BoundStatement
}

func NewBoundGlobalScope(previous *BoundGlobalScope, variables []*Symbols.VariableSymbol, statement BoundStatement) *BoundGlobalScope {
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

    parent := CreateRootScope()

    for stack.Count() > 0 {
        previous = stack.Pop().(*BoundGlobalScope)

        scope := NewBoundScope(parent)
        for _, v := range previous.Variables {
            scope.TryDeclareVariable(v)
        }

        parent = scope
    }

    return parent
}

func CreateRootScope() *BoundScope {
    result :=  NewBoundScope(nil)

    for f := range BuiltinFunctions.GetAll() {
        result.TryDeclareFunction(f)
    }

    return result
}
