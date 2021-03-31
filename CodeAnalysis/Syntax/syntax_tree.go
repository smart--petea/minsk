package Syntax

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
)

type SyntaxTree struct {
    Util.DiagnosticBag

    Root *CompilationUnitSyntax
    Text *Text.SourceText
}

func newSyntaxTree(sourceText *Text.SourceText) *SyntaxTree {
    parser := NewParser(sourceText)
    root := parser.ParseCompilationUnit()

    syntaxTree := &SyntaxTree {
        Root: root,
        Text: sourceText,
    }
    syntaxTree.AddDiagnosticsRange(parser.GetDiagnostics())

    return syntaxTree
}

func SyntaxTreeParse(text string) *SyntaxTree {
    sourceText := Text.SourceTextFrom(text)
    return newSyntaxTree(sourceText)
}

func ParseTokens(text string) (tokens []*SyntaxToken) {
    tokens, _ = ParseTokensWithDiagnostics(text)

    return tokens
}

func ParseTokensWithDiagnostics(text string) (tokens []*SyntaxToken, diagnostics []*Util.Diagnostic) {
    sourceText := Text.SourceTextFrom(text)
    lexer := NewLexer(sourceText)

    for {
        token := lexer.Lex()
        if token.Kind() == SyntaxKind.EndOfFileToken {
            break
        }

        tokens = append(tokens, token)
    }

    return tokens,  lexer.GetDiagnostics()
}
