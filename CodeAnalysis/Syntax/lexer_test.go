package Syntax

import (
    "testing"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

func TestLexerLexesToken(t *testing.T) {
    tokensData := []struct{
        kind SyntaxKind.SyntaxKind
        text string
    } {
        {
            kind: SyntaxKind.IdentifierToken, 
            text: "a",
        },
        {
            kind: SyntaxKind.IdentifierToken, 
            text: "abc",
        },
    }

    for _, tokenData := range tokensData {
        tokens := ParseTokens(tokenData.text)

        if len(tokens) != 1 {
            t.Errorf("len(%+v) expected=1", tokens)
        }

        token := tokens[0]
        if token.kind != tokenData.kind {
            t.Errorf("kind=%s, expected=%s", token.kind, tokenData.kind)
        }
        if string(token.Runes) != tokenData.text {
            t.Errorf("runes=%s, expected=%s", string(token.Runes), tokenData.text)
        }
    }
}
