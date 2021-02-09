package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type BinaryExpressionSyntax struct {
    *syntaxNodeChildren
    Left ExpressionSyntax
    Right ExpressionSyntax
    OperatorNode SyntaxNode
}

func (b *BinaryExpressionSyntax) Value() interface{} {
    return nil
}

func NewBinaryExpressionSyntax(left ExpressionSyntax, operatorNode SyntaxNode, right ExpressionSyntax) *BinaryExpressionSyntax {
    return &BinaryExpressionSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(left, operatorNode, right),
        Left: left,
        Right: right,
        OperatorNode: operatorNode,
    }
}

func (b *BinaryExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.BinaryExpression
}
