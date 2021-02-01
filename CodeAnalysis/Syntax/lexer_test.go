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

        {kind: SyntaxKind.PlusToken, text: "+"},
        {kind: SyntaxKind.MinusToken, text: "-"},
        {kind: SyntaxKind.StarToken, text: "*"},
        {kind: SyntaxKind.SlashToken, text: "/"},
        {kind: SyntaxKind.BangToken, text: "!"},
        {kind: SyntaxKind.EqualsToken, text: "="},
        {kind: SyntaxKind.AmpersandAmpersandToken, text: "&&"},
        {kind: SyntaxKind.PipePipeToken, text: "||"},
        {kind: SyntaxKind.EqualsEqualsToken, text: "=="},
        {kind: SyntaxKind.BangEqualsToken, text: "!="},
        {kind: SyntaxKind.OpenParenthesisToken, text: "("},
        {kind: SyntaxKind.CloseParenthesisToken, text: ")"},
        {kind: SyntaxKind.FalseKeyword, text: "false"},
        {kind: SyntaxKind.TrueKeyword, text: "true"},

        {kind: SyntaxKind.WhitespaceToken, text: " "},
        {kind: SyntaxKind.WhitespaceToken, text: "  "},
        {kind: SyntaxKind.WhitespaceToken, text: "\r"},
        {kind: SyntaxKind.WhitespaceToken, text: "\n\r"},
        {kind: SyntaxKind.NumberToken, text: "1"},
        {kind: SyntaxKind.NumberToken, text: "123"},
        {kind: SyntaxKind.IdentifierToken, text: "a"},
        {kind: SyntaxKind.IdentifierToken, text: "abc"},
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
