package Binding

import (
    "fmt"
)

type BoundTreeRewriter struct {
}

func (b *BoundTreeRewriter) RewriteStatement(node BoundStatement) BoundStatement {
    switch n := node.(type) {
        case *BoundBlockStatement:
            return b.RewriteBlockStatement(n)
        case *BoundExpressionStatement:
            return b.RewriteExpressionStatement(n)
        case *BoundIfStatement:
            return b.RewriteIfStatement(n)
        case *BoundWhileStatement:
            return b.RewriteWhileStatement(n)
        case *BoundForStatement:
            return b.RewriteForStatement(n)
        case *BoundVariableDeclaration:
            return b.RewriteVariableDeclaration(n)
        default:
            panic(fmt.Sprintf("Unexpected node: %s", node.Kind()))
    }
}

func (b *BoundTreeRewriter) RewriteExpression(node BoundExpression) BoundExpression {
    switch n := node.(type) {
        case *BoundUnaryExpression:
            return b.RewriteUnaryExpression(n)
        case *BoundLiteralExpression:
            return b.RewriteLiteralExpression(n)
        case *BoundBinaryExpression:
            return b.RewriteBinaryExpression(n)
        case *BoundVariableExpression:
            return b.RewriteVariableExpression(n)
        case *BoundAssignmentExpression:
            return b.RewriteAssignmentExpression(n)
        default:
            panic(fmt.Sprintf("Unexpected node: %s", node.Kind()))
    }
}

func (b *BoundTreeRewriter) RewriteBlockStatement(node *BoundBlockStatement) BoundStatement {
    var builder []BoundStatement
    var isBuilderNew bool

    for _, oldStatement := range node.Statements {
        newStatement := b.RewriteStatement(oldStatement)
        if newStatement != oldStatement {
            builder = append(builder, newStatement)
            isBuilderNew = true
        } else {
            builder = append(builder, oldStatement)
        }
    }

    if isBuilderNew == false {
        return node
    }

    return NewBoundBlockStatement(builder)
}

func (b *BoundTreeRewriter) RewriteExpressionStatement(node *BoundExpressionStatement) BoundStatement {
    expression := b.RewriteExpression(node.Expression)
    if expression == node.Expression {
        return node
    }

    return NewBoundExpressionStatement(expression)
}

func (b *BoundTreeRewriter) RewriteIfStatement(node *BoundIfStatement) BoundStatement {
    condition := b.RewriteExpression(node.Condition)
    thenStatement := b.RewriteStatement(node.ThenStatement)
    var elseStatement BoundStatement
    if node.ElseStatement != nil {
        elseStatement = b.RewriteStatement(node.ElseStatement)
    }

    if condition == node.Condition && thenStatement == node.ThenStatement && elseStatement == node.ElseStatement {
        return node
    }

    return NewBoundIfStatement(condition, thenStatement, elseStatement)
}

func (b *BoundTreeRewriter) RewriteWhileStatement(node *BoundWhileStatement) BoundStatement {
    condition := b.RewriteExpression(node.Condition)
    body := b.RewriteStatement(node.Body)
    if condition == node.Condition && body == node.Body {
        return node
    }

    return NewBoundWhileStatement(condition, body)
}

func (b *BoundTreeRewriter) RewriteForStatement(node *BoundForStatement) BoundStatement {
    lowerBound := b.RewriteExpression(node.LowerBound)
    upperBound := b.RewriteExpression(node.UpperBound)
    body := b.RewriteStatement(node.Body)
    if lowerBound == node.LowerBound && upperBound == node.UpperBound && body == node.Body {
        return node
    }

    return NewBoundForStatement(node.Variable, lowerBound, upperBound, body)
}

func (b *BoundTreeRewriter) RewriteVariableDeclaration(node *BoundVariableDeclaration) BoundStatement {
    initializer := b.RewriteExpression(node.Initializer)
    if initializer == node.Initializer {
        return node
    }

    return NewBoundVariableDeclaration(node.Variable, initializer)
}

func (b *BoundTreeRewriter) RewriteUnaryExpression(node *BoundUnaryExpression) BoundExpression {
    operand := b.RewriteExpression(node.Operand)
    if operand == node.Operand {
        return node
    }

    return NewBoundUnaryExpression(node.Op, operand)
}

func (b *BoundTreeRewriter) RewriteLiteralExpression(node *BoundLiteralExpression) BoundExpression {
    return node
}

func (b *BoundTreeRewriter) RewriteBinaryExpression(node *BoundBinaryExpression) BoundExpression {
    left := b.RewriteExpression(node.Left)
    right := b.RewriteExpression(node.Right)
    if left == node.Left && right == node.Right {
        return node
    }

    return NewBoundBinaryExpression(left, node.Op, right)
}

func (b *BoundTreeRewriter) RewriteVariableExpression(node *BoundVariableExpression) BoundExpression {
    return node
}

func (b *BoundTreeRewriter) RewriteAssignmentExpression(node *BoundAssignmentExpression) BoundExpression {
    expression := b.RewriteExpression(node.Expression)
    if expression == node.Expression {
        return node
    }

    return NewBoundAssignmentExpression(node.Variable, expression)
}
