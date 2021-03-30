package Binding

import (
    "minsk/CodeAnalysis/Binding/BoundUnaryOperator"
    "minsk/CodeAnalysis/Binding/BoundBinaryOperator"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Syntax"
    "minsk/Util"

    "fmt"
    "log"
    "reflect"
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

func (b *Binder) BindExpressionWithType(syntax Syntax.ExpressionSyntax, expectedType reflect.Kind) BoundExpression {
    result := b.BindExpression(syntax)
    if result.GetType() != expectedType {
        span := syntax.GetSpan()
        b.ReportCannotConvert(span, result.GetType(), expectedType)
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

    default:
        panic(fmt.Sprintf("Unexpected syntax %s", syntax.Kind()))
    }
}

func (b *Binder) BindAssignmentExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    assignmentExpressionSyntax := syntax.(*Syntax.AssignmentExpressionSyntax)

    name := string(assignmentExpressionSyntax.IdentifierToken.Runes)
    boundExpression := b.BindExpression(assignmentExpressionSyntax.Expression)

    var variable *Util.VariableSymbol
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

func (b *Binder) BindNameExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    log.Printf("BindNameExpression")
    nameExpressionSyntax := syntax.(*Syntax.NameExpressionSyntax)
    name := string(nameExpressionSyntax.IdentifierToken.Runes)

    if len(name) == 0 {
        //This means the token was inserted by the parser. We already reported
        //error so we can just return an error expression.
        return NewBoundLiteralExpression(0)
    }

    var variable *Util.VariableSymbol
    if b.scope.TryLookup(name, &variable) {
        return NewBoundVariableExpression(variable)
    }

    span := nameExpressionSyntax.IdentifierToken.GetSpan()
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
    boundOperator := BoundUnaryOperator.Bind(unarySyntax.OperatorNode.Kind(), boundOperand.GetType()) 

    if boundOperator == nil {
        syntaxToken := unarySyntax.OperatorNode.(*Syntax.SyntaxToken)
        span := syntaxToken.GetSpan()
        b.ReportUndefinedUnaryOperator(span, syntaxToken.Runes, boundOperand.GetType())
        return boundOperand;
    }

    return NewBoundUnaryExpression(boundOperator, boundOperand)
}

func (b *Binder) BindBinaryExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    binarySyntax := (syntax).(*Syntax.BinaryExpressionSyntax)

    boundLeft := b.BindExpression(binarySyntax.Left)
    boundRight := b.BindExpression(binarySyntax.Right)
    boundOperator := BoundBinaryOperator.Bind(binarySyntax.OperatorNode.Kind(), boundLeft.GetType(), boundRight.GetType()) 

    if boundOperator == nil {
        syntaxToken := binarySyntax.OperatorNode.(*Syntax.SyntaxToken)
        span := syntaxToken.GetSpan()
        b.ReportUndefinedBinaryOperator(span, syntaxToken.Runes, boundLeft.GetType(), boundRight.GetType())

        return boundLeft;
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

    lowerBound := b.BindExpressionWithType(forStatementSyntax.LowerBound, reflect.Int)
    upperBound := b.BindExpressionWithType(forStatementSyntax.UpperBound, reflect.Int)

    b.scope = NewBoundScope(b.scope)

    name := string(forStatementSyntax.Identifier.Runes)
    variable := Util.NewVariableSymbol(name, true, reflect.Int)
    if b.scope.TryDeclare(variable) == false {
        b.ReportVariableAlreadyDeclared(forStatementSyntax.Identifier.GetSpan(), name)
    }

    body := b.BindStatement(forStatementSyntax.Body)

    b.scope = b.scope.Parent
    return NewBoundForStatement(variable, lowerBound, upperBound, body)
}

func (b *Binder) BindWhileStatement(syntax Syntax.StatementSyntax) BoundStatement {
    whileStatementSyntax := (syntax).(*Syntax.WhileStatementSyntax)

    condition := b.BindExpressionWithType(whileStatementSyntax.Condition, reflect.Bool) 
    body := b.BindStatement(whileStatementSyntax.Body)

    return NewBoundWhileStatement(condition, body)
}

func (b *Binder) BindIfStatement(syntax Syntax.StatementSyntax) BoundStatement {
    ifStatementSyntax := (syntax).(*Syntax.IfStatementSyntax)

    condition := b.BindExpressionWithType(ifStatementSyntax.Condition, reflect.Bool) 
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

func (b *Binder) BindVariableDeclaration(syntax Syntax.StatementSyntax) BoundStatement {
    variableDeclarationSyntax := (syntax).(*Syntax.VariableDeclarationSyntax)

    name := string(variableDeclarationSyntax.Identifier.Runes)
    isReadOnly := variableDeclarationSyntax.Keyword.Kind() == SyntaxKind.LetKeyword
    initializer := b.BindExpression(variableDeclarationSyntax.Initializer)
    variable := Util.NewVariableSymbol(name, isReadOnly, initializer.GetType())

    if !b.scope.TryDeclare(variable) {
        span := variableDeclarationSyntax.Identifier.GetSpan()
        b.ReportVariableAlreadyDeclared(span, name)
    }

    return NewBoundVariableDeclaration(variable, initializer)
}
