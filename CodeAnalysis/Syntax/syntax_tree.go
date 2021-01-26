package Syntax

import (
    "minsk/Util"
)

type SyntaxTree struct {
    Util.DiagnosticBag

    Root ExpressionSyntax
    EndOfFileToken *SyntaxToken
}

func ParseSyntaxTree(text string) *SyntaxTree {
    parser := NewParser(text)
    rootExpression, endOfFileToken := parser.Parse()

    syntaxTree := &SyntaxTree {
        Root: rootExpression,
        EndOfFileToken: endOfFileToken,
    }

    syntaxTree.AddDiagnosticsRange(parser.GetDiagnostics())

    return syntaxTree
}
