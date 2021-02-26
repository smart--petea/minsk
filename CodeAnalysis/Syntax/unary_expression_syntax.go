package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
)

type UnaryExpressionSyntax struct {
    *syntaxNodeChildren
    OperatorNode SyntaxNode
    Operand ExpressionSyntax
}

func (u *UnaryExpressionSyntax) Value() interface{} {
    return nil
}

func NewUnaryExpressionSyntax(operatorNode SyntaxNode, operand ExpressionSyntax) *UnaryExpressionSyntax {
    return &UnaryExpressionSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(operatorNode, operand),
        OperatorNode: operatorNode,
        Operand: operand,
    }
}

func (u *UnaryExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.UnaryExpression
}

func (u *UnaryExpressionSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(u)
}
