package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type BinaryExpressionSyntax struct {
    *Util.ChildrenProvider

    Left ExpressionSyntax
    Right ExpressionSyntax
    OperatorNode SyntaxNode
}

func (b *BinaryExpressionSyntax) Value() interface{} {
    return nil
}

func NewBinaryExpressionSyntax(left ExpressionSyntax, operatorNode SyntaxNode, right ExpressionSyntax) *BinaryExpressionSyntax {
    return &BinaryExpressionSyntax{
        ChildrenProvider: Util.NewChildrenProvider(left, operatorNode, right),

        Left: left,
        Right: right,
        OperatorNode: operatorNode,
    }
}

func (b *BinaryExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.BinaryExpression
}

func (b *BinaryExpressionSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(b)
}
