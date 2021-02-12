package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/Util"
)

type SyntaxTree struct {
    Util.DiagnosticBag

    Root ExpressionSyntax
    EndOfFileToken *SyntaxToken
}

func ParseSyntaxTree(text string) *SyntaxTree {
    sourceText := SourceTextFrom(text)

    parser := NewParser(sourceText)
    rootExpression, endOfFileToken := parser.Parse()

    syntaxTree := &SyntaxTree {
        Root: rootExpression,
        EndOfFileToken: endOfFileToken,
    }

    syntaxTree.AddDiagnosticsRange(parser.GetDiagnostics())

    return syntaxTree
}

func ParseTokens(text string) (tokens []*SyntaxToken) {
    lexer := NewLexer([]rune(text))

    for {
        token := lexer.Lex()
        if token.Kind() == SyntaxKind.EndOfFileToken {
            break
        }

        tokens = append(tokens, token)
    }

    return tokens
}
