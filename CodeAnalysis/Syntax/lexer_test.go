package Syntax

import (
    "testing"
    "strings"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

var tokensData = []struct{
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

    //{kind: SyntaxKind.WhitespaceToken, text: " "},
    //{kind: SyntaxKind.WhitespaceToken, text: "  "},
    //{kind: SyntaxKind.WhitespaceToken, text: "\r"},
    //{kind: SyntaxKind.WhitespaceToken, text: "\n\r"},
    {kind: SyntaxKind.NumberToken, text: "1"},
    {kind: SyntaxKind.NumberToken, text: "123"},
    {kind: SyntaxKind.IdentifierToken, text: "a"},
    {kind: SyntaxKind.IdentifierToken, text: "abc"},
}

func TestLexerLexesToken(t *testing.T) {
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

func TestLexerTokenPairs(t *testing.T) {
    for _, tokenData0 := range tokensData {
        for _, tokenData1 := range tokensData {
            if requiresSeparator(tokenData0.kind, tokenData1.kind) {
                continue
            }
            text := tokenData0.text + tokenData1.text
            tokens := ParseTokens(text)

            if len(tokens) != 2 {
                t.Errorf("len(%+v) expected=2", tokens)
            }

            if tokens[0].kind != tokenData0.kind {
                t.Errorf("0. kind=%s, expected=%s", tokens[0].kind, tokenData0.kind)
            }
            if string(tokens[0].Runes) != tokenData0.text {
                t.Errorf("0. runes=%s, expected=%s", string(tokens[0].Runes), tokenData0.text)
            }

            if tokens[1].kind != tokenData1.kind {
                t.Errorf("1. kind=%s, expected=%s", tokens[1].kind, tokenData1.kind)
            }
            if string(tokens[1].Runes) != tokenData1.text {
                t.Errorf("1. runes=%s, expected=%s", string(tokens[1].Runes), tokenData1.text)
            }
        }
    }
}

func requiresSeparator(t1kind, t2kind SyntaxKind.SyntaxKind) bool {
    t1kindIsKeyword := strings.HasSuffix(string(t1kind), "Keyword")
    t2kindIsKeyword := strings.HasSuffix(string(t2kind), "Keyword")

    if t1kind == SyntaxKind.IdentifierToken && t2kind == SyntaxKind.IdentifierToken {
        return true
    }

    if t1kindIsKeyword && t2kindIsKeyword {
        return true
    }

    if t1kindIsKeyword && t2kind == SyntaxKind.IdentifierToken {
        return true
    }

    if t1kind == SyntaxKind.IdentifierToken && t2kindIsKeyword {
        return true
    }

    if t1kind == SyntaxKind.NumberToken && t2kind == SyntaxKind.NumberToken {
        return true
    }

    if t1kind == SyntaxKind.BangToken && t2kind == SyntaxKind.EqualsToken {
        return true
    }

    if t1kind == SyntaxKind.BangToken && t2kind == SyntaxKind.EqualsEqualsToken {
        return true
    }

    if t1kind == SyntaxKind.EqualsToken && t2kind == SyntaxKind.EqualsToken {
        return true
    }

    if t1kind == SyntaxKind.EqualsToken && t2kind == SyntaxKind.EqualsEqualsToken {
        return true
    }

    return false
}
