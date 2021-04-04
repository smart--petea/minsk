package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

//print("hello")
type CallExpressionSyntax struct {
    *Util.ChildrenProvider

    Identifier: *SyntaxToken,
    Arguments: *SeparatedSyntaxList,
}

func (a *CallExpressionSyntax) Value() interface{} {
    return nil
}

func NewCallExpressionSyntax(identifier *SyntaxToken, arguments  *SeparatedSyntaxList) *CallExpressionSyntax {
    return &CallExpressionSyntax{
        ChildrenProvider: Util.NewChildrenProvider(identifier, arguments.GetWithSeparators()...),

        Identifier: identifier,
        Arguments: arguments,
    }
}

func (c *CallExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.CallExpression 
}
