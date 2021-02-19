package Syntax

type ExpressionStatementSyntax struct {
    *syntaxNodeChildren

    Expression ExpressionSyntax
}

func NewExpressionStatementSyntax(expression ExpressionSyntax) *ExpressionStatementSyntax {
    return &ExpressionStatementSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(expression),

        Expression: expression,
    }
}

func (e *ExpressionStatementSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.ExpressionStatement
}

func (e *ExpressionStatementSyntax) Value() interface{} {
    return e.Expression.Value()
}
