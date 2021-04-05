package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type CallExpressionSyntax struct {
    *Util.ChildrenProvider

    Identifier *SyntaxToken
    OpenParenthesisToken *SyntaxToken
    CloseParenthesisToken *SyntaxToken
    Arguments *SeparatedSyntaxList
}

func (a *CallExpressionSyntax) Value() interface{} {
    return nil
}

func NewCallExpressionSyntax(identifier *SyntaxToken, openParenthesisToken *SyntaxToken, arguments  *SeparatedSyntaxList, closeParenthesisToken *SyntaxToken) *CallExpressionSyntax {
    var children []interface{}
    children = append(children, interface{}(identifier))
    children = append(children, interface{}(openParenthesisToken))
    for _, argument := range arguments.GetWithSeparators() {
        children = append(children, interface{}(argument))
    }
    children = append(children, interface{}(closeParenthesisToken))

    return &CallExpressionSyntax{
        ChildrenProvider: Util.NewChildrenProvider(children...),

        Identifier: identifier,
        Arguments: arguments,
        OpenParenthesisToken: openParenthesisToken,
        CloseParenthesisToken: closeParenthesisToken,
    }
}

func (c *CallExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.CallExpression 
}

func (c *CallExpressionSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(c)
}
