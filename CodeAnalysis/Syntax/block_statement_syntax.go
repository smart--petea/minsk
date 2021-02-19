package Syntax

type BlockStatementSyntax struct {
    *syntaxNodeChildren

    OpenBraceToken *SyntaxToken
    Statements []*StatementSyntax
    CloseBraceToken *SyntaxToken
}

func NewBlockStatementSyntax(openBraceToken *SyntaxToken, statements []*StatementSyntax, closeBraceToken *SyntaxToken) *BlockStatementSyntax {
    return &BlockStatementSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(openBraceToken, statements..., closeBraceToken),

        OpenBraceToken: openBraceToken,
        Statements: statements,
        CloseBraceToken: closeBraceToken,
    }
}

func (b *BlockStatementSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.BlockStatement
}

func (b *BlockStatementSyntax) Value() interface{} {
    return nil
}
