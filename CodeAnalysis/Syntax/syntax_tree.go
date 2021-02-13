package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type SyntaxTree struct {
    Util.DiagnosticBag

    Root ExpressionSyntax
    EndOfFileToken *SyntaxToken
    Text *Text.SourceText
}

func ParseSyntaxTree(text string) *SyntaxTree {
    sourceText := Text.SourceTextFrom(text)

    parser := NewParser(sourceText)
    rootExpression, endOfFileToken := parser.Parse()

    syntaxTree := &SyntaxTree {
        Root: rootExpression,
        EndOfFileToken: endOfFileToken,
        Text: sourceText,
    }

    syntaxTree.AddDiagnosticsRange(parser.GetDiagnostics())

    return syntaxTree
}

func ParseTokens(text string) (tokens []*SyntaxToken) {
    sourceText := Text.SourceTextFrom(text)
    lexer := NewLexer(sourceText)

    for {
        token := lexer.Lex()
        if token.Kind() == SyntaxKind.EndOfFileToken {
            break
        }

        tokens = append(tokens, token)
    }

    return tokens
}
