package Binding

import (
    "minsk/CodeAnalysis/Binding/BoundUnaryOperator"
    "minsk/CodeAnalysis/Binding/BoundBinaryOperator"
    "minsk/CodeAnalysis/Binding/Conversion"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Symbols"
    "minsk/CodeAnalysis/Syntax"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"

    "fmt"
//    "log"
)

type Binder struct {
    Util.DiagnosticBag 

    scope *BoundScope
    function *Symbols.FunctionSymbol
}

func NewBinder(parent *BoundScope, function *Symbols.FunctionSymbol) *Binder {
    return &Binder{
        scope: NewBoundScope(parent),
        function: function,
    }
}

func (b *Binder) BindStatement(syntax Syntax.StatementSyntax) BoundStatement {
    switch syntax.Kind() {
    case SyntaxKind.BlockStatement:
        return b.BindBlockStatement(syntax)

    case SyntaxKind.VariableDeclaration:
        return b.BindVariableDeclaration(syntax)

    case SyntaxKind.IfStatement:
        return b.BindIfStatement(syntax)

    case SyntaxKind.WhileStatement:
        return b.BindWhileStatement(syntax)

    case SyntaxKind.ForStatement:
        return b.BindForStatement(syntax)

    case SyntaxKind.ExpressionStatement:
        return b.BindExpressionStatement(syntax)

    default:
        panic(fmt.Sprintf("Unexpected syntax %s", syntax.Kind()))
    }
}

func (b *Binder) BindExpressionWithType(syntax Syntax.ExpressionSyntax, targetType *Symbols.TypeSymbol) BoundExpression {
    return b.BindConversion(syntax, targetType, false)
}

func (b *Binder) BindExpression(syntax Syntax.ExpressionSyntax, canBeVoid bool) BoundExpression {
    result := b.BindExpressionInternal(syntax)
    if !canBeVoid && result.GetType() == Symbols.TypeSymbolVoid {
        b.ReportExpressionMustHaveValue(syntax.GetSpan())
        return NewBoundErrorExpression()
    }

    return result
}

func (b *Binder) BindExpressionInternal(syntax Syntax.ExpressionSyntax) BoundExpression {
    switch syntax.Kind() {
    case SyntaxKind.ParenthesizedExpression:
        return b.BindParenthesizedExpression(syntax)

    case SyntaxKind.LiteralExpression:
        return b.BindLiteralExpression(syntax)

    case SyntaxKind.NameExpression:
        return b.BindNameExpression(syntax)

    case SyntaxKind.AssignmentExpression:
        return b.BindAssignmentExpression(syntax)

    case SyntaxKind.UnaryExpression:
        return b.BindUnaryExpression(syntax)

    case SyntaxKind.BinaryExpression:
        return b.BindBinaryExpression(syntax)

    case SyntaxKind.CallExpression:
        return b.BindCallExpression(syntax)

    default:
        panic(fmt.Sprintf("Unexpected syntax %s", syntax.Kind()))
    }
}

func (b *Binder) BindAssignmentExpression(assignmentExpressionSyntax Syntax.ExpressionSyntax) BoundExpression {
    syntax := assignmentExpressionSyntax.(*Syntax.AssignmentExpressionSyntax)

    name := string(syntax.IdentifierToken.Runes)
    boundExpression := b.BindExpression(syntax.Expression, false)

    var variable *Symbols.VariableSymbol
    if b.scope.TryLookupVariable(name, &variable) == false {
        span := syntax.IdentifierToken.GetSpan()
        b.ReportUndefinedName(span, name)
        return boundExpression
    }

    if variable.IsReadOnly {
        span := syntax.EqualsToken.GetSpan()
        b.ReportCannotAssign(span, name)
    }

    convertedExpression := b.BindConversionInner(syntax.Expression.GetSpan(), boundExpression, variable.Type, false)
    return NewBoundAssignmentExpression(variable, convertedExpression)
}

func (b *Binder) BindNameExpression(expressionSyntax Syntax.ExpressionSyntax) BoundExpression {
    syntax := expressionSyntax.(*Syntax.NameExpressionSyntax)
    if syntax.IdentifierToken.IsMissing() {
        //This means the token was inserted by the parser. We already reported
        //error so we can just return an error expression.
        return NewBoundErrorExpression()
    }

    name := string(syntax.IdentifierToken.Runes)
    var variable *Symbols.VariableSymbol
    if b.scope.TryLookupVariable(name, &variable) {
        return NewBoundVariableExpression(variable)
    }

    span := syntax.IdentifierToken.GetSpan()
    b.ReportUndefinedName(span, name)
    return NewBoundLiteralExpression(0)
}

func (b *Binder) BindParenthesizedExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    pS := syntax.(*Syntax.ParenthesizedExpressionSyntax)
    return b.BindExpression(pS.Expression, false)
}

func (b *Binder) BindLiteralExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    //log.Printf("BindLiteralExpression %+v", syntax)
    literalSyntax := syntax.(*Syntax.LiteralExpressionSyntax)

    var value interface{}

    switch literalSyntax.LiteralToken.Kind() {
    case SyntaxKind.TrueKeyword:
        value = true
    case SyntaxKind.FalseKeyword:
        value = false
    case SyntaxKind.IdentifierToken:
        value = string(literalSyntax.LiteralToken.Runes)
    default:
        //if val, ok := literalSyntax.Value().(int); ok {
        //    value = val
        //}
        value = literalSyntax.Value()
    }

    return NewBoundLiteralExpression(value)
}

func (b *Binder) BindUnaryExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    unarySyntax := syntax.(*Syntax.UnaryExpressionSyntax)
    boundOperand := b.BindExpression(unarySyntax.Operand, false)
    if boundOperand.GetType() == Symbols.TypeSymbolError {
        return NewBoundErrorExpression()
    }

    boundOperator := BoundUnaryOperator.Bind(unarySyntax.OperatorNode.Kind(), boundOperand.GetType()) 

    if boundOperator == nil {
        syntaxToken := unarySyntax.OperatorNode.(*Syntax.SyntaxToken)
        span := syntaxToken.GetSpan()
        b.ReportUndefinedUnaryOperator(span, syntaxToken.Runes, boundOperand.GetType())

        return NewBoundErrorExpression()
    }

    return NewBoundUnaryExpression(boundOperator, boundOperand)
}

func (b *Binder)  LookupType(name string) *Symbols.TypeSymbol {
    switch name {
    case "bool":
        return Symbols.TypeSymbolBool
    case "int":
        return Symbols.TypeSymbolInt
    case "string":
        return Symbols.TypeSymbolString
    default:
        return nil
    }
}

func (b *Binder) BindConversion(syntax Syntax.ExpressionSyntax, ttype *Symbols.TypeSymbol, allowExplicit bool) BoundExpression {
    expression := b.BindExpression(syntax, false)
    return b.BindConversionInner(syntax.GetSpan(), expression, ttype, allowExplicit)
}

func (b *Binder) BindConversionInner(diagnosticSpan *Text.TextSpan, expression BoundExpression, ttype *Symbols.TypeSymbol, allowExplicit bool) BoundExpression {
    conversion := Conversion.ConversionClassify(expression.GetType(), ttype)

    if !conversion.Exists {
        if expression.GetType() != Symbols.TypeSymbolError && ttype != Symbols.TypeSymbolError {
            b.ReportCannotConvert(diagnosticSpan, expression.GetType(), ttype)
        }
        return NewBoundErrorExpression()
    }

    if !allowExplicit && conversion.IsExplicit {
        b.ReportCannotConvertImplicitly(diagnosticSpan, expression.GetType(), ttype)
    }

    if conversion.IsIdentity {
        return expression
    }


    return NewBoundConversionExpression(ttype, expression)
}

func (b *Binder) BindCallExpression(expressionSyntax Syntax.ExpressionSyntax) BoundExpression {
    syntax := expressionSyntax.(*Syntax.CallExpressionSyntax)

    if syntax.Arguments.Count() == 1 {
        ttype := b.LookupType(string(syntax.Identifier.Runes)) 
        if ttype != nil {
            allowExplicit := true
            return b.BindConversion(syntax.Arguments.Get(0), ttype, allowExplicit)
        }
    }

    var boundArguments []BoundExpression
    for argument := range syntax.Arguments.GetEnumerator() {
        boundArgument := b.BindExpression(argument, false)
        boundArguments = append(boundArguments, boundArgument)
    }

    var function *Symbols.FunctionSymbol
    functionName := string(syntax.Identifier.Runes)
    if b.scope.TryLookupFunction(functionName, &function) == false {
        b.ReportUndefinedFunction(syntax.Identifier.GetSpan(), functionName)
        return NewBoundErrorExpression()
    }

    if syntax.Arguments.Count() != len(function.Parameter) {
        b.ReportWrongArgumentCount(syntax.Identifier.GetSpan(), functionName, len(function.Parameter), syntax.Arguments.Count())
        return NewBoundErrorExpression()
    }

    for i := 0; i < syntax.Arguments.Count(); i = i + 1 {
        argument := boundArguments[i]
        parameter := function.Parameter[i]

        if argument.GetType() != parameter.Type {
            b.ReportWrongArgumentType(syntax.Identifier.GetSpan(), parameter.Name, parameter.Type, argument.GetType())
            return NewBoundErrorExpression()
        }
    }

    return NewBoundCallExpression(function, boundArguments)
}

func (b *Binder) BindBinaryExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    binarySyntax := (syntax).(*Syntax.BinaryExpressionSyntax)

    boundLeft := b.BindExpression(binarySyntax.Left, false)
    boundRight := b.BindExpression(binarySyntax.Right, false)

    if boundLeft.GetType() == Symbols.TypeSymbolError || boundRight.GetType() == Symbols.TypeSymbolError {
        return NewBoundErrorExpression()
    }

    boundOperator := BoundBinaryOperator.Bind(binarySyntax.OperatorNode.Kind(), boundLeft.GetType(), boundRight.GetType()) 

    if boundOperator == nil {
        syntaxToken := binarySyntax.OperatorNode.(*Syntax.SyntaxToken)
        span := syntaxToken.GetSpan()
        b.ReportUndefinedBinaryOperator(span, syntaxToken.Runes, boundLeft.GetType(), boundRight.GetType())

        return NewBoundErrorExpression()
    }

    return NewBoundBinaryExpression(boundLeft, boundOperator, boundRight)
}

func (b *Binder) BindBlockStatement(syntax Syntax.StatementSyntax) *BoundBlockStatement {
    blockStatementSyntax := (syntax).(*Syntax.BlockStatementSyntax)
    b.scope = NewBoundScope(b.scope)

    var statements []BoundStatement
    for _, statementSyntax := range blockStatementSyntax.Statements {
        statement := b.BindStatement(statementSyntax)
        statements = append(statements, statement)
    }
    
    b.scope = b.scope.Parent

    return NewBoundBlockStatement(statements)
}

func (b *Binder) BindForStatement(syntax Syntax.StatementSyntax) BoundStatement {
    forStatementSyntax := (syntax).(*Syntax.ForStatementSyntax)

    lowerBound := b.BindExpressionWithType(forStatementSyntax.LowerBound, Symbols.TypeSymbolInt)
    upperBound := b.BindExpressionWithType(forStatementSyntax.UpperBound, Symbols.TypeSymbolInt)

    b.scope = NewBoundScope(b.scope)

    variable := b.BindVariable(forStatementSyntax.Identifier, true, Symbols.TypeSymbolInt)

    body := b.BindStatement(forStatementSyntax.Body)

    b.scope = b.scope.Parent
    return NewBoundForStatement(variable, lowerBound, upperBound, body)
}

func (b *Binder) BindWhileStatement(syntax Syntax.StatementSyntax) BoundStatement {
    whileStatementSyntax := (syntax).(*Syntax.WhileStatementSyntax)

    condition := b.BindExpressionWithType(whileStatementSyntax.Condition, Symbols.TypeSymbolBool) 
    body := b.BindStatement(whileStatementSyntax.Body)

    return NewBoundWhileStatement(condition, body)
}

func (b *Binder) BindIfStatement(syntax Syntax.StatementSyntax) BoundStatement {
    ifStatementSyntax := (syntax).(*Syntax.IfStatementSyntax)

    condition := b.BindExpressionWithType(ifStatementSyntax.Condition, Symbols.TypeSymbolBool) 
    thenStatement := b.BindStatement(ifStatementSyntax.ThenStatement)
    var elseStatement BoundStatement
    if ifStatementSyntax.ElseClause != nil {
        elseStatement = b.BindStatement(ifStatementSyntax.ElseClause.ElseStatement)
    }

    return NewBoundIfStatement(condition, thenStatement, elseStatement)
}

func (b *Binder) BindExpressionStatement(syntax Syntax.StatementSyntax) BoundStatement {
    expressionStatementSyntax := (syntax).(*Syntax.ExpressionStatementSyntax)

    canBeVoid := true
    expression := b.BindExpression(expressionStatementSyntax.Expression, canBeVoid)

    return NewBoundExpressionStatement(expression)
}

func (b *Binder) BindVariableDeclaration(statementSyntax Syntax.StatementSyntax) BoundStatement {
    syntax := (statementSyntax).(*Syntax.VariableDeclarationSyntax)

    isReadOnly := syntax.Keyword.Kind() == SyntaxKind.LetKeyword
    ttype := b.BindTypeClause(syntax.TypeClause)
    initializer := b.BindExpression(syntax.Initializer, false)

    variableType := ttype
    if variableType == nil {
        variableType = initializer.GetType()
    }

    variable := b.BindVariable(syntax.Identifier, isReadOnly, variableType)
    convertedInitializer := b.BindConversionInner(syntax.Initializer.GetSpan(), initializer, variableType, false)
    return NewBoundVariableDeclaration(variable, convertedInitializer)
}

func (b *Binder) BindTypeClause(syntax *Syntax.TypeClauseSyntax) *Symbols.TypeSymbol {
    if syntax == nil {
        return nil
    }

    ttype := b.LookupType(string(syntax.Identifier.Runes))
    if ttype == nil {
        b.ReportUndefinedType(syntax.Identifier.GetSpan(), string(syntax.Identifier.Runes))
    }

    return ttype
}

func (b *Binder) BindVariable(identifier *Syntax.SyntaxToken, isReadOnly bool, ttype *Symbols.TypeSymbol) *Symbols.VariableSymbol {
    name := string(identifier.Runes)
    if name == "" {
        name = "?"
    }
    variable := Symbols.NewVariableSymbol(name, isReadOnly, ttype)
    declare := !identifier.IsMissing()

    if declare && b.scope.TryDeclareVariable(variable) == false {
        b.ReportSymbolAlreadyDeclared(identifier.GetSpan(), name)
    }

    return variable
}

func (b *Binder) BindFunctionDeclaration(syntax *Syntax.FunctionDeclarationSyntax) {
    var parameters []*Symbols.ParameterSymbol
    seenParameterNames := Util.NewHashSetString()

    for  syntaxNode := range syntax.Parameters.GetEnumerator() {
        parameterSyntax := syntaxNode.(*Syntax.ParameterSyntax)
        parameterName := string(parameterSyntax.Identifier.Runes)
        parameterType := b.BindTypeClause(parameterSyntax.Type)

        if !seenParameterNames.Add(parameterName) {
            b.ReportParameterAlreadyDeclared(parameterSyntax.GetSpan(), parameterName) 
        } else {
            parameter := Symbols.NewParameterSymbol(parameterName, parameterType)
            parameters = append(parameters, parameter)
        }
    }

    ttype := b.BindTypeClause(syntax.Type)
    if ttype == nil {
        ttype = Symbols.TypeSymbolVoid
    }

    if ttype != Symbols.TypeSymbolVoid {
        b.XXX_ReportFunctionsAreUnsupported(syntax.Type.GetSpan())
    }

    function := Symbols.NewFunctionSymbol(string(syntax.Identifier.Runes), parameters, ttype, syntax)
    if !b.scope.TryDeclareFunction(function) {
        b.ReportSymbolAlreadyDeclared(syntax.Identifier.GetSpan(), function.Name)
    }
}

func BinderBindProgram(globalScope *BoundGlobalScope) *BoundProgram {
    parentScope := CreateParentScopes(globalScope)
    functionBodies := make(map[*Symbols.FunctionSymbol]*BoundBlockStatement)
    diagnostics := Util.NewDiagnosticBag()

    for _, function := range globalScope.Functions {
        binder := NewBinder(parentScope, function)
        body := binder.BindStatement(function.Declaration.Body)
        loweredBody := Lowering.LowererLower(body)
        functionBodies[function] = loweredBody

        diagnostics.AddRange(binder.DiagnosticBag)
    }

    boundProgram := NewBoundProgram(global, diagnostics, functionBodies)
    return boundProgram
}
