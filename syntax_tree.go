package minsk

type SyntaxTree struct {
    Root ExpressionSyntax
    EndOfFileToken *SyntaxToken
    Diagnostics []string
}

func NewSyntaxTree(diagnostics []string, root ExpressionSyntax, endOfFileToken *SyntaxToken) *SyntaxTree {
    return &SyntaxTree {
        Root: root,
        EndOfFileToken: endOfFileToken,
        Diagnostics: diagnostics,
    }
}
