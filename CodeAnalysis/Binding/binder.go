package Binding

import (
    "minsk/CodeAnalysis/Binding/BoundUnaryOperator"
    "minsk/CodeAnalysis/Binding/BoundBinaryOperator"

    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Syntax"

    "minsk/Util"

    "fmt"
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

    case SyntaxKind.ExpressionStatement:
        return b.BindExpressionStatement(syntax)

    default:
        panic(fmt.Sprintf("Unexpected syntax %s", syntax.Kind()))
    }
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
        variable = Util.NewVariableSymbol(name, boundExpression.GetType())
        b.scope.TryDeclare(variable)
    }

    if boundExpression.GetType() != variable.Type {
        span := Syntax.SyntaxNodeToTextSpan(assignmentExpressionSyntax.IdentifierToken)
        b.ReportCannotConvert(span, boundExpression.GetType(), variable.Type)
        return boundExpression
    }

    return NewBoundAssignmentExpression(variable, boundExpression)
}

func (b *Binder) BindNameExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    nameExpressionSyntax := syntax.(*Syntax.NameExpressionSyntax)
    name := string(nameExpressionSyntax.IdentifierToken.Runes)

    var variable *Util.VariableSymbol
    if b.scope.TryLookup(name, &variable) {
        return NewBoundVariableExpression(variable)
    }

    span := Syntax.SyntaxNodeToTextSpan(nameExpressionSyntax.IdentifierToken)
    b.ReportUndefinedName(span, name)
    return NewBoundLiteralExpression(0)
}

func (b *Binder) BindParenthesizedExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    pS := syntax.(*Syntax.ParenthesizedExpressionSyntax)
    return b.BindExpression(pS.Expression)
}

func (b *Binder) BindLiteralExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
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
        if val, ok := literalSyntax.Value().(int); ok {
            value = val
        }
    }

    return NewBoundLiteralExpression(value)
}

func (b *Binder) BindUnaryExpression(syntax Syntax.ExpressionSyntax) BoundExpression {
    unarySyntax := syntax.(*Syntax.UnaryExpressionSyntax)
    boundOperand := b.BindExpression(unarySyntax.Operand)
    boundOperator := BoundUnaryOperator.Bind(unarySyntax.OperatorNode.Kind(), boundOperand.GetType()) 

    if boundOperator == nil {
        syntaxToken := unarySyntax.OperatorNode.(*Syntax.SyntaxToken)
        span := Syntax.SyntaxNodeToTextSpan(syntaxToken)
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
        span := Syntax.SyntaxNodeToTextSpan(syntaxToken)
        b.ReportUndefinedBinaryOperator(span, syntaxToken.Runes, boundLeft.GetType(), boundRight.GetType())

        return boundLeft;
    }

    return NewBoundBinaryExpression(boundLeft, boundOperator, boundRight)
}

func (b *Binder) BindBlockStatement(syntax Syntax.StatementSyntax) *BoundBlockStatement {
    blockStatementSyntax := (syntax).(*Syntax.BlockStatementSyntax)

    var statements []BoundStatement
    for _, statementSyntax := range blockStatementSyntax.Statements {
        statement := b.BindStatement(statementSyntax)
        statements = append(statements, statement)
    }

    return NewBoundBlockStatement(statements)
}

func (b *Binder) BindExpressionStatement(syntax Syntax.StatementSyntax) BoundStatement {
    expressionStatementSyntax := (syntax).(*Syntax.ExpressionStatementSyntax)
    expression := b.BindExpression(expressionStatementSyntax.Expression)

    return NewBoundExpressionStatement(expression)
}

func (b *Binder) BindVariableDeclaration(syntax Syntax.StatementSyntax) BoundStatement {
    variableDeclarationSyntax := (syntax).(*Syntax.VariableDeclaration)

    name := variableDeclarationSyntax.Identifier.Runes
    isReadOnly := variableDeclarationSyntax.Keyword.Kind() == SyntaxKind.LetKeyword
    initializer := p.BindExpression(variableDeclarationSyntax.Initializer)
    variable := NewVariableSymbol( )
}
