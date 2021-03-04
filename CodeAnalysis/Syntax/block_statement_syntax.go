package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type BlockStatementSyntax struct {
    *Util.ChildrenProvider

    OpenBraceToken *SyntaxToken
    Statements []StatementSyntax
    CloseBraceToken *SyntaxToken
}

func NewBlockStatementSyntax(openBraceToken *SyntaxToken, statements []StatementSyntax, closeBraceToken *SyntaxToken) *BlockStatementSyntax {
    var children []interface{}
    children = append(children, openBraceToken)
    for _, statement := range statements {
        children = append(children, statement)
    }
    children = append(children, closeBraceToken)

    return &BlockStatementSyntax{
        ChildrenProvider: Util.NewChildrenProvider(children...),

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

func (b *BlockStatementSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(b)
}
