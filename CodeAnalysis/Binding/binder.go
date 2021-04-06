package Binding

import (
    "minsk/CodeAnalysis/Binding/BoundUnaryOperator"
    "minsk/CodeAnalysis/Binding/BoundBinaryOperator"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Symbols"
    "minsk/CodeAnalysis/Symbols/BuiltinFunctions"
    "minsk/CodeAnalysis/Syntax"
    "minsk/Util"

    "fmt"
//    "log"
)

type Binder struct {
    Util.DiagnosticBag 

    scope *BoundScope
}

func NewBinder(parent *BoundScope) *Binder {
    return &Binder{
        scope: NewBoundScope(parent),
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
    result := b.BindExpression(syntax)
    if targetType != Symbols.TypeSymbolError &&
        result.GetType() != Symbols.TypeSymbolError &&
        result.GetType() != targetType {
        span := syntax.GetSpan()
        b.ReportCannotConvert(span, result.GetType(), targetType)
    }

    return result
}

func (b *Binder) BindExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
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

func (b *Binder) BindAssignmentExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    assignmentExpressionSyntax := syntax.(*Syntax.AssignmentExpressionSyntax)

    name := string(assignmentExpressionSyntax.IdentifierToken.Runes)
    boundExpression := b.BindExpression(assignmentExpressionSyntax.Expression)

    var variable *Symbols.VariableSymbol
    if b.scope.TryLookup(name, &variable) == false {
        span := assignmentExpressionSyntax.IdentifierToken.GetSpan()
        b.ReportUndefinedName(span, name)
        return boundExpression
    }

    if variable.IsReadOnly {
        span := assignmentExpressionSyntax.EqualsToken.GetSpan()
        b.ReportCannotAssign(span, name)
    }

    if boundExpression.GetType() != variable.Type {
        span := assignmentExpressionSyntax.IdentifierToken.GetSpan()
        b.ReportCannotConvert(span, boundExpression.GetType(), variable.Type)
        return boundExpression
    }

    return NewBoundAssignmentExpression(variable, boundExpression)
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
    if b.scope.TryLookup(name, &variable) {
        return NewBoundVariableExpression(variable)
    }

    span := syntax.IdentifierToken.GetSpan()
    b.ReportUndefinedName(span, name)
    return NewBoundLiteralExpression(0)
}

func (b *Binder) BindParenthesizedExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    pS := syntax.(*Syntax.ParenthesizedExpressionSyntax)
    return b.BindExpression(pS.Expression)
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
    boundOperand := b.BindExpression(unarySyntax.Operand)
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

func (b *Binder) BindCallExpression(expressionSyntax Syntax.ExpressionSyntax) BoundExpression {
    syntax := expressionSyntax.(*Syntax.CallExpressionSyntax)

    var boundArguments []BoundExpression
    for argument := range syntax.Arguments.GetEnumerator() {
        boundArgument := b.BindExpression(argument)
        boundArguments = append(boundArguments, boundArgument)
    }

    functionName := string(syntax.Identifier.Runes)
    var function *Symbols.FunctionSymbol
    for f := range BuiltinFunctions.GetAll() {
        if f.Name == functionName {
            function = f
            break
        }
    }

    if function == nil {
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

    boundLeft := b.BindExpression(binarySyntax.Left)
    boundRight := b.BindExpression(binarySyntax.Right)

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
    expression := b.BindExpression(expressionStatementSyntax.Expression)

    return NewBoundExpressionStatement(expression)
}

func (b *Binder) BindVariableDeclaration(statementSyntax Syntax.StatementSyntax) BoundStatement {
    syntax := (statementSyntax).(*Syntax.VariableDeclarationSyntax)

    isReadOnly := syntax.Keyword.Kind() == SyntaxKind.LetKeyword
    initializer := b.BindExpression(syntax.Initializer)
    variable := b.BindVariable(syntax.Identifier, isReadOnly, initializer.GetType())

    return NewBoundVariableDeclaration(variable, initializer)
}

func (b *Binder) BindVariable(identifier *Syntax.SyntaxToken, isReadOnly bool, ttype *Symbols.TypeSymbol) *Symbols.VariableSymbol {
    name := string(identifier.Runes)
    if name == "" {
        name = "?"
    }
    variable := Symbols.NewVariableSymbol(name, isReadOnly, ttype)
    declare := !identifier.IsMissing()

    if declare && b.scope.TryDeclare(variable) == false {
        b.ReportVariableAlreadyDeclared(identifier.GetSpan(), name)
    }

    return variable
}

