package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/Util"
)

type FunctionDeclarationSyntax struct {
    *Util.ChildrenProvider

    FunctionKeyword *SyntaxToken
    Identifier *SyntaxToken
    OpenParenthesisToken *SyntaxToken
    Parameters *SeparatedSyntaxList
    CloseParenthesisToken *SyntaxToken
    Type *TypeClauseSyntax
}

func NewFunctionDeclarationSyntax(functionKeyword *SyntaxToken, identifier *SyntaxToken, openParenthesisToken *SyntaxToken, parameters *SeparatedSyntaxList, closeParenthesisToken *SyntaxToken, ttype *TypeClauseSyntax) *FunctionDeclarationSyntax {
    var children []interface{}
    children = append(children, interface{}(functionKeyword))
    children = append(children, interface{}(identifier))
    children = append(children, interface{}(openParenthesisToken))
    for _, parameter := range parameters.GetWithSeparators() {
        children = append(children, interface{}(paramter))
    }
    children = append(children, interface{}(closeParenthesisToken))
    children = append(children, interface{}(ttype))
    return &FunctionDeclarationSyntax{
        ChildrenProvider: Util.NewChildrenProvider(children...),

        FunctionKeyword: functionKeyword,
        Identifier: identifier,
        OpenParenthesisToken: openParenthesisToken,
        Parameters: parameters,
        CloseParenthesisToken: closeParenthesisToken,
        Type: ttype,
    }
}

func (g *FunctionDeclarationSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.FunctionDeclaration
}

func (g *FunctionDeclarationSyntax) Value() interface{} {
    return nil
}
