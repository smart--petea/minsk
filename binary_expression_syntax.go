package minsk

type BinaryExpressionSyntax struct {
    Left ExpressionSyntax
    Right ExpressionSyntax
    OperatorNode SyntaxNode
}

func (b *BinaryExpressionSyntax) Value() interface{} {
    return nil
}

func NewBinaryExpressionSyntax(left ExpressionSyntax, operatorNode SyntaxNode, right ExpressionSyntax) *BinaryExpressionSyntax {
    return &BinaryExpressionSyntax{
        Left: left,
        Right: right,
        OperatorNode: operatorNode,
    }
}

func (b *BinaryExpressionSyntax) Kind() SyntaxKind {
    return BinaryExpression
}

func (b *BinaryExpressionSyntax) GetChildren() []SyntaxNode {
    return []SyntaxNode{
        b.Left,
        b.OperatorNode,
        b.Right,
    }
}
