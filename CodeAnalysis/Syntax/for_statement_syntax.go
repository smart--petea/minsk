package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type ForStatementSyntax struct {
    *Util.ChildrenProvider

    Keyword *SyntaxToken
    Identifier *SyntaxToken
    EqualsToken *SyntaxToken
    LowerBound ExpressionSyntax
    ToKeyword *SyntaxToken
    UpperBound ExpressionSyntax
    Body StatementSyntax
}

func NewForStatementSyntax(keyword, identifier, equalsToken *SyntaxToken, lowerBound ExpressionSyntax, toKeyword *SyntaxToken, upperBound ExpressionSyntax, body StatementSyntax) *ForStatementSyntax {
    return &ForStatementSyntax{
        ChildrenProvider: Util.NewChildrenProvider(keyword, identifier, equalsToken, lowerBound, toKeyword, upperBound, body),

        Keyword: keyword,
        Identifier: identifier,
        EqualsToken: equalsToken,
        LowerBound: lowerBound,
        ToKeyword: toKeyword,
        UpperBound: upperBound,
        Body: body,
    }
}

func (f *ForStatementSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.ForStatement
}

func (f *ForStatementSyntax) Value() interface{} {
    return nil
}

func (f *ForStatementSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(f)
}
