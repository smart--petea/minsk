package Binding

import (
    "fmt"
)

type BoundTreeRewriter struct {
}

func (*BoundTreeRewriter) RewriteStatement(b BoundITreeRewriter, node BoundStatement) BoundStatement {
    switch n := node.(type) {
        case *BoundBlockStatement:
            return b.RewriteBlockStatement(b, n)
        case *BoundExpressionStatement:
            return b.RewriteExpressionStatement(b, n)
        case *BoundIfStatement:
            return b.RewriteIfStatement(b, n)
        case *BoundWhileStatement:
            return b.RewriteWhileStatement(b, n)
        case *BoundForStatement:
            return b.RewriteForStatement(b, n)
        case *BoundLabelStatement:
            return b.RewriteLabelStatement(b, n)
        case *BoundGotoStatement:
            return b.RewriteGotoStatement(b, n)
        case *BoundConditionalGotoStatement:
            return b.RewriteConditionalGotoStatement(b, n)
        case *BoundVariableDeclaration:
            return b.RewriteVariableDeclaration(b, n)
        default:
            panic(fmt.Sprintf("Unexpected node: %s", node.Kind()))
    }
}

func (*BoundTreeRewriter) RewriteExpression(b BoundITreeRewriter, node BoundExpression) BoundExpression {
    switch n := node.(type) {
        case *BoundUnaryExpression:
            return b.RewriteUnaryExpression(b, n)
        case *BoundLiteralExpression:
            return b.RewriteLiteralExpression(b, n)
        case *BoundBinaryExpression:
            return b.RewriteBinaryExpression(b, n)
        case *BoundVariableExpression:
            return b.RewriteVariableExpression(b, n)
        case *BoundAssignmentExpression:
            return b.RewriteAssignmentExpression(b, n)
        default:
            panic(fmt.Sprintf("Unexpected node: %s", node.Kind()))
    }
}

func (*BoundTreeRewriter) RewriteBlockStatement(b BoundITreeRewriter, node *BoundBlockStatement) BoundStatement {
    var builder []BoundStatement
    var isBuilderNew bool

    for _, oldStatement := range node.Statements {
        newStatement := b.RewriteStatement(b, oldStatement)
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

func (*BoundTreeRewriter) RewriteExpressionStatement(b BoundITreeRewriter, node *BoundExpressionStatement) BoundStatement {
    expression := b.RewriteExpression(b, node.Expression)
    if expression == node.Expression {
        return node
    }

    return NewBoundExpressionStatement(expression)
}

func (*BoundTreeRewriter) RewriteIfStatement(b BoundITreeRewriter, node *BoundIfStatement) BoundStatement {
    condition := b.RewriteExpression(b, node.Condition)
    thenStatement := b.RewriteStatement(b, node.ThenStatement)
    var elseStatement BoundStatement
    if node.ElseStatement != nil {
        elseStatement = b.RewriteStatement(b, node.ElseStatement)
    }

    if condition == node.Condition && thenStatement == node.ThenStatement && elseStatement == node.ElseStatement {
        return node
    }

    return NewBoundIfStatement(condition, thenStatement, elseStatement)
}

func (*BoundTreeRewriter) RewriteWhileStatement(b BoundITreeRewriter, node *BoundWhileStatement) BoundStatement {
    condition := b.RewriteExpression(b, node.Condition)
    body := b.RewriteStatement(b, node.Body)
    if condition == node.Condition && body == node.Body {
        return node
    }

    return NewBoundWhileStatement(condition, body)
}

func (*BoundTreeRewriter) RewriteForStatement(b BoundITreeRewriter, node *BoundForStatement) BoundStatement {
    lowerBound := b.RewriteExpression(b, node.LowerBound)
    upperBound := b.RewriteExpression(b, node.UpperBound)
    body := b.RewriteStatement(b, node.Body)
    if lowerBound == node.LowerBound && upperBound == node.UpperBound && body == node.Body {
        return node
    }

    return NewBoundForStatement(node.Variable, lowerBound, upperBound, body)
}

func (*BoundTreeRewriter) RewriteVariableDeclaration(b BoundITreeRewriter, node *BoundVariableDeclaration) BoundStatement {
    initializer := b.RewriteExpression(b, node.Initializer)
    if initializer == node.Initializer {
        return node
    }

    return NewBoundVariableDeclaration(node.Variable, initializer)
}

func (*BoundTreeRewriter) RewriteUnaryExpression(b BoundITreeRewriter, node *BoundUnaryExpression) BoundExpression {
    operand := b.RewriteExpression(b, node.Operand)
    if operand == node.Operand {
        return node
    }

    return NewBoundUnaryExpression(node.Op, operand)
}

func (*BoundTreeRewriter) RewriteLiteralExpression(b BoundITreeRewriter, node *BoundLiteralExpression) BoundExpression {
    return node
}

func (*BoundTreeRewriter) RewriteBinaryExpression(b BoundITreeRewriter, node *BoundBinaryExpression) BoundExpression {
    left := b.RewriteExpression(b, node.Left)
    right := b.RewriteExpression(b, node.Right)
    if left == node.Left && right == node.Right {
        return node
    }

    return NewBoundBinaryExpression(left, node.Op, right)
}

func (*BoundTreeRewriter) RewriteVariableExpression(b BoundITreeRewriter, node *BoundVariableExpression) BoundExpression {
    return node
}

func (*BoundTreeRewriter) RewriteAssignmentExpression(b BoundITreeRewriter, node *BoundAssignmentExpression) BoundExpression {
    expression := b.RewriteExpression(b, node.Expression)
    if expression == node.Expression {
        return node
    }

    return NewBoundAssignmentExpression(node.Variable, expression)
}

func (*BoundTreeRewriter) RewriteLabelStatement(b BoundITreeRewriter, node *BoundLabelStatement) BoundStatement {
    return node
}

func (*BoundTreeRewriter) RewriteGotoStatement(b BoundITreeRewriter, node *BoundGotoStatement) BoundStatement {
    return node
}

func (*BoundTreeRewriter) RewriteConditionalGotoStatement(b BoundITreeRewriter, node *BoundConditionalGotoStatement) BoundStatement {
    condition := b.RewriteExpression(b, node.Condition)
    if condition == node.Condition {
        return node
    }

    return NewBoundConditionalGotoStatement(node.Label, condition, node.JumpIfFalse)
}
