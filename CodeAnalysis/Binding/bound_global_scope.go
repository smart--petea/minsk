package Binding

import (
    "minsk/Util"
    "minsk/CodeAnalysis/Syntax"
    "minsk/CodeAnalysis/Symbols"
    "minsk/CodeAnalysis/Symbols/BuiltinFunctions"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type BoundGlobalScope struct {
    *Util.DiagnosticBag

    Previous *BoundGlobalScope
    Functions []*Symbols.FunctionSymbol
    Variables []Symbols.IVariableSymbol
    Statement BoundStatement
}

func NewBoundGlobalScope(previous *BoundGlobalScope, functions []*Symbols.FunctionSymbol, variables []Symbols.IVariableSymbol, statement BoundStatement) *BoundGlobalScope {
    return &BoundGlobalScope{
        DiagnosticBag: Util.NewDiagnosticBag(),

        Previous: previous,
        Functions: functions,
        Variables: variables,
        Statement: statement,
    }
}

func BoundGlobalScopeBindGlobalScope(previous *BoundGlobalScope, syntax *Syntax.CompilationUnitSyntax) *BoundGlobalScope {
    parentScope := CreateParentScopes(previous)
    binder := NewBinder(parentScope, nil)

    for _, memberSyntax := range syntax.Members.OfType(SyntaxKind.FunctionDeclaration) {
        function := memberSyntax.(*Syntax.FunctionDeclarationSyntax)
        binder.BindFunctionDeclaration(function)
    }

    var statementBuilder []BoundStatement

    for _, memberSyntax := range syntax.Members.OfType(SyntaxKind.GlobalStatement) {
        globalStatement := memberSyntax.(*Syntax.GlobalStatementSyntax)
        s := binder.BindStatement(globalStatement.Statement)
        statementBuilder = append(statementBuilder, s)
    }

    statement := NewBoundBlockStatement(statementBuilder)
    functions := binder.scope.GetDeclaredFunctions()
    variables := binder.scope.GetDeclaredVariables()

    boundGlobalScope := NewBoundGlobalScope(previous, functions, variables, statement)
    boundGlobalScope.DiagnosticBag.AddRange(binder.DiagnosticBag)
    if previous != nil {
        boundGlobalScope.DiagnosticBag.AddRange(previous.DiagnosticBag)
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

        for _, f := range previous.Functions {
            scope.TryDeclareFunction(f)
        }

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

