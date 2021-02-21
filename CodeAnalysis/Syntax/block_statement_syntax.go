package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type BlockStatementSyntax struct {
    *syntaxNodeChildren

    OpenBraceToken *SyntaxToken
    Statements []StatementSyntax
    CloseBraceToken *SyntaxToken
}

func NewBlockStatementSyntax(openBraceToken *SyntaxToken, statements []StatementSyntax, closeBraceToken *SyntaxToken) *BlockStatementSyntax {
    var children []SyntaxNode
    children = append(children, openBraceToken)
    for _, statement := range statements {
        children = append(children, statement.(SyntaxNode))
    }
    children = append(children, closeBraceToken)

    return &BlockStatementSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(children...),

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
