package CodeAnalysis

type SyntaxTree struct {
    Root ExpressionSyntax
    EndOfFileToken *SyntaxToken
    Diagnostics []string
}

func ParseSyntaxTree(text string) *SyntaxTree {
    parser := NewParser(text)
    rootExpression, endOfFileToken := parser.Parse()

    return &SyntaxTree {
        Root: rootExpression,
        EndOfFileToken: endOfFileToken,
        Diagnostics: parser.Diagnostics,
    }
}
