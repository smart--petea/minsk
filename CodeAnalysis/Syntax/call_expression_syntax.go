package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

//print("hello")
type CallExpressionSyntax struct {
    *Util.ChildrenProvider
}

func (a *CallExpressionSyntax) Value() interface{} {
    return nil
}

func NewCallExpressionSyntax() *CallExpressionSyntax {
    return &CallExpressionSyntax{
        ChildrenProvider: Util.NewChildrenProvider(),

    }
}

func (c *CallExpressionSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.CallExpression 
}
