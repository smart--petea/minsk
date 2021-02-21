package Syntax

type VariableDeclarationSyntax struct {
    *syntaxNodeChildren

    Keyword SyntaxToken
    Identifier SyntaxToken
    EqualsToken SyntaxToken
    Initializer ExpressionSyntax
}

func (v *VariableDeclarationSyntax) Kind() SyntaxKind.SyntaxKind {
    return SyntaxKind.VariableDeclaration
}

func NewVariableDeclarationSyntax(keyword, identifier, equalsToken  SyntaxToken, initializer ExpressionSyntax) *VariableDeclarationSyntax {
    return &VariableDeclarationSyntax{
        syntaxNodeChildren: newSyntaxNodeChildren(keyword, identifier, equalsToken, initializer),

        Keyword: keyword,
        Identifier: identifier,
        EqualsToken: equalsToken,
        Initializer: initializer,
    }
}

func (v *VariableDeclarationSyntax) Value() interface{} {
    return nil
}
