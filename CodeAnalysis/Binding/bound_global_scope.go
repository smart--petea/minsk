package Binding

import (
    "minsk/Util"
    "minsk/CodeAnalysis/Syntax"
)

type BoundGlobalScope struct {
    Util.DiagnosticBag
    Previous *BoundScope
    Variables []*Util.VariableSymbol
    Expression BoundExpression
}

func NewBoundGlobalScope(previous *BoundScope, variables []*Util.VariableSymbol, expression BoundExpression) *BoundGlobalScope {
    return &BoundGlobalScope{
        Previous: previous,
        Variables: variables,
        Expression: expression,
    }
}

func BoundGlobalScopeFromCompilationUnitSyntax(syntax *Syntax.CompilationUnitSyntax) *BoundGlobalScope {
    binder := NewBinder(nil)
    expression := binder.BindExpression(syntax.Expression)
    variables := binder.scope.GetDeclaredVariables()

    boundGlobalScope := NewBoundGlobalScope(nil, variables, expression)
    boundGlobalScope.AddDiagnosticsRange(binder.GetDiagnostics())

    return boundGlobalScope
}
