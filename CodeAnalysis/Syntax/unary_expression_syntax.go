package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type UnaryExpressionSyntax struct {
    *Util.ChildrenProvider

    OperatorNode SyntaxNode
    Operand ExpressionSyntax
}

func (u *UnaryExpressionSyntax) Value() interface{} {
    return nil
}

func NewUnaryExpressionSyntax(operatorNode SyntaxNode, operand ExpressionSyntax) *UnaryExpressionSyntax {
    return &UnaryExpressionSyntax{
        ChildrenProvider: Util.NewChildrenProvider(operatorNode, operand),

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
