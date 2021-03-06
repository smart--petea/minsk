package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
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
    Body *BlockStatementSyntax
}

func NewFunctionDeclarationSyntax(functionKeyword *SyntaxToken, identifier *SyntaxToken, openParenthesisToken *SyntaxToken, parameters *SeparatedSyntaxList, closeParenthesisToken *SyntaxToken, ttype *TypeClauseSyntax, body *BlockStatementSyntax) *FunctionDeclarationSyntax {
    var children []interface{}
    children = append(children, interface{}(functionKeyword))
    children = append(children, interface{}(identifier))
    children = append(children, interface{}(openParenthesisToken))
    for _, parameter := range parameters.GetWithSeparators() {
        children = append(children, interface{}(parameter))
    }
    children = append(children, interface{}(closeParenthesisToken))
    children = append(children, interface{}(ttype))
    children = append(children, interface{}(body))

    return &FunctionDeclarationSyntax{
        ChildrenProvider: Util.NewChildrenProvider(children...),

        FunctionKeyword: functionKeyword,
        Identifier: identifier,
        OpenParenthesisToken: openParenthesisToken,
        Parameters: parameters,
        CloseParenthesisToken: closeParenthesisToken,
        Type: ttype,
        Body: body,
    }
}

func (g *FunctionDeclarationSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.FunctionDeclaration
}

func (g *FunctionDeclarationSyntax) Value() interface{} {
    return g.Body.Value()
}

func (g *FunctionDeclarationSyntax) GetSpan() *Text.TextSpan {
    return SyntaxNodeToTextSpan(g)
}
