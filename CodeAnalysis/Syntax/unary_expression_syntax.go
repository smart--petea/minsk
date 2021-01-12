package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type UnaryExpressionSyntax struct {
    OperatorNode SyntaxNode
    Operand ExpressionSyntax
}

func (u *UnaryExpressionSyntax) Value() interface{} {
    return nil
}

func NewUnaryExpressionSyntax(operatorNode SyntaxNode, operand ExpressionSyntax) *UnaryExpressionSyntax {
    return &UnaryExpressionSyntax{
        OperatorNode: operatorNode,
        Operand: operand,
    }
}

func (u *UnaryExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.UnaryExpression
}

func (u *UnaryExpressionSyntax) GetChildren() []SyntaxNode {
    return []SyntaxNode{
        u.OperatorNode,
        u.Operand,
    }
}
