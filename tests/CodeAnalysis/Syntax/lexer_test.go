package SyntaxTest

import (
    "testing"
    "strings"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Syntax"

    "fmt"
)

type testToken struct {
    kind SyntaxKind.SyntaxKind
    text string
}

func getSeparators() []testToken {
    return []testToken {
        {kind: SyntaxKind.WhitespaceToken, text: " "},
        {kind: SyntaxKind.WhitespaceToken, text: "  "},
        {kind: SyntaxKind.WhitespaceToken, text: "\r"},
        {kind: SyntaxKind.WhitespaceToken, text: "\n\r"},
    }
}

func getTokens() []testToken {
    return []testToken{
        {kind: SyntaxKind.PlusToken, text: "+"},
        {kind: SyntaxKind.MinusToken, text: "-"},
        {kind: SyntaxKind.StarToken, text: "*"},
        {kind: SyntaxKind.SlashToken, text: "/"},
        {kind: SyntaxKind.BangToken, text: "!"},
        {kind: SyntaxKind.EqualsToken, text: "="},
        {kind: SyntaxKind.LessToken, text: "<"},
        {kind: SyntaxKind.LessOrEqualsToken, text: "<="},
        {kind: SyntaxKind.GreaterToken, text: ">"},
        {kind: SyntaxKind.GreaterOrEqualsToken, text: ">="},
        {kind: SyntaxKind.AmpersandAmpersandToken, text: "&&"},
        {kind: SyntaxKind.AmpersandToken, text: "&"},
        {kind: SyntaxKind.PipeToken, text: "|"},
        {kind: SyntaxKind.TildeToken, text: "~"},
        {kind: SyntaxKind.HatToken, text: "^"},
        {kind: SyntaxKind.PipePipeToken, text: "||"},
        {kind: SyntaxKind.EqualsEqualsToken, text: "=="},
        {kind: SyntaxKind.BangEqualsToken, text: "!="},
        {kind: SyntaxKind.OpenParenthesisToken, text: "("},
        {kind: SyntaxKind.CloseParenthesisToken, text: ")"},
        {kind: SyntaxKind.FalseKeyword, text: "false"},
        {kind: SyntaxKind.TrueKeyword, text: "true"},
        {kind: SyntaxKind.ToKeyword, text: "to"},
        {kind: SyntaxKind.WhileKeyword, text: "while"},
        {kind: SyntaxKind.ForKeyword, text: "for"},
        {kind: SyntaxKind.IfKeyword, text: "if"},

        {kind: SyntaxKind.NumberToken, text: "1"},
        {kind: SyntaxKind.NumberToken, text: "123"},
        {kind: SyntaxKind.IdentifierToken, text: "a"},
        {kind: SyntaxKind.IdentifierToken, text: "abc"},
        {kind: SyntaxKind.StringToken, text: "\"Test\""},
        {kind: SyntaxKind.StringToken, text: "\"Te\"\"st\""},
    }
}

func TestLexerLexesUnterminatedString(t *testing.T) {
    text := "\"text"
    tokens, diagnostics := Syntax.ParseTokensWithDiagnostics(text)

    if len(tokens) != 1 {
        t.Errorf("len(%+v) expected=1", tokens)
    }

    token := tokens[0]
    if token.Kind() != SyntaxKind.StringToken {
        t.Errorf("kind=%s, expected=%s", token.Kind(), SyntaxKind.StringToken)
    }
    if string(token.Runes) != text {
        t.Errorf("runes=%s, expected=%s", string(token.Runes), text)
    }

    if len(diagnostics) != 1 {
        t.Errorf("diagnostics len(%+v)=%d should be 1", diagnostics, len(diagnostics))
    }
    diagnostic := diagnostics[0]

    if diagnostic.Span.Start != 0 || diagnostic.Span.Length != 1 {
        t.Errorf("diagnostic.Span is (%d, %d), should be (%d, %d)", diagnostic.Span.Start, diagnostic.Span.Length, 0, 1)
    }

    if diagnostic.Message != "Unterminated string literal." {
        t.Errorf("diagnostic.Message = %s, should be '%s'", diagnostic.Message, "Unterminated string literal.")
    }
}

func TestLexerCoversAllTokens(t *testing.T) {
    var tokenKinds SyntaxKind.SyntaxKindSlice
    for _, kind := range SyntaxKind.GetValues() {
        if strings.HasSuffix(string(kind), "Keyword") || strings.HasSuffix(string(kind), "Token") {
            tokenKinds = append(tokenKinds, kind)
        }
    }

    var testedTokenKinds SyntaxKind.SyntaxKindSlice
    for _, testtoken := range append(getTokens(), getSeparators()...) {
        testedTokenKinds = append(testedTokenKinds, testtoken.kind)
    }

    testedTokenKinds = append(testedTokenKinds, SyntaxKind.BadToken, SyntaxKind.EndOfFileToken)
    untestedTokenKinds := tokenKinds.Unique().ExceptWith(testedTokenKinds)

    if len(untestedTokenKinds) != 0 {
        t.Errorf("(%+v) - untested token kinds", untestedTokenKinds)
    }
}

func TestLexerLexesToken(t *testing.T) {
    tokensData := append(getTokens(), getSeparators()...)

    for _, tokenData := range tokensData {
        tokens := Syntax.ParseTokens(tokenData.text)

        if len(tokens) != 1 {
            t.Errorf("len(%+v) expected=1", tokens)
        }

        token := tokens[0]
        if token.Kind() != tokenData.kind {
            t.Errorf("kind=%s, expected=%s", token.Kind(), tokenData.kind)
        }
        if string(token.Runes) != tokenData.text {
            t.Errorf("runes=%s, expected=%s", string(token.Runes), tokenData.text)
        }
    }
}

func TestLexerTokenPairs(t *testing.T) {
    tokensData := getTokens()

    for _, tokenData0 := range tokensData {
        for _, tokenData1 := range tokensData {
            if requiresSeparator(tokenData0.kind, tokenData1.kind) {
                //fmt.Printf("109. %+v %+v \n", tokenData0.kind, tokenData1.kind)
                continue
            }
            fmt.Printf("112. %+v %+v \n", tokenData0.kind, tokenData1.kind)
            text := tokenData0.text + tokenData1.text
            fmt.Printf("113. %+v\n", text)
            tokens := Syntax.ParseTokens(text)

            if len(tokens) != 2 {
                t.Errorf("len(%+v) expected=2", tokens)
            }

            if tokens[0].Kind() != tokenData0.kind {
                t.Errorf("0. kind=%s, expected=%s", tokens[0].Kind(), tokenData0.kind)
            }
            if string(tokens[0].Runes) != tokenData0.text {
                t.Errorf("0. runes=%s, expected=%s", string(tokens[0].Runes), tokenData0.text)
            }

            if tokens[1].Kind() != tokenData1.kind {
                t.Errorf("1. kind=%s, expected=%s", tokens[1].Kind(), tokenData1.kind)
            }
            if string(tokens[1].Runes) != tokenData1.text {
                t.Errorf("1. runes=%s, expected=%s", string(tokens[1].Runes), tokenData1.text)
            }
        }
    }
}

func TestLexerTokenPairsWithSeparators(t *testing.T) {
    for _, tokenData0 := range getTokens() {
        for _, tokenData2 := range getTokens() {
            if requiresSeparator(tokenData0.kind, tokenData2.kind) {
                continue
            }

            for _, separatorData := range getSeparators() {
                text := tokenData0.text + separatorData.text + tokenData2.text
                tokens := Syntax.ParseTokens(text)

                if len(tokens) != 3 {
                    t.Errorf("len(%+v) expected=3", tokens)
                }

                if tokens[0].Kind() != tokenData0.kind {
                    t.Errorf("0. kind=%s, expected=%s", tokens[0].Kind(), tokenData0.kind)
                }
                if string(tokens[0].Runes) != tokenData0.text {
                    t.Errorf("0. runes=%s, expected=%s", string(tokens[0].Runes), tokenData0.text)
                }

                if tokens[1].Kind() != separatorData.kind {
                    t.Errorf("1. kind=%s, expected=%s", tokens[1].Kind(), separatorData.kind)
                }
                if string(tokens[1].Runes) != separatorData.text {
                    t.Errorf("1. runes=%s, expected=%s", string(tokens[1].Runes), separatorData.text)
                }

                if tokens[2].Kind() != tokenData2.kind {
                    t.Errorf("2. kind=%s, expected=%s", tokens[2].Kind(), tokenData2.kind)
                }
                if string(tokens[2].Runes) != tokenData2.text {
                    t.Errorf("2. runes=%s, expected=%s", string(tokens[2].Runes), tokenData2.text)
                }
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

    if t1kind == SyntaxKind.StringToken && t2kind == SyntaxKind.StringToken {
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

    if t1kind == SyntaxKind.LessToken && t2kind == SyntaxKind.EqualsToken {
        return true
    }

    if t1kind == SyntaxKind.LessToken && t2kind == SyntaxKind.EqualsEqualsToken {
        return true
    }

    if t1kind == SyntaxKind.GreaterToken && t2kind == SyntaxKind.EqualsToken {
        return true
    }

    if t1kind == SyntaxKind.GreaterToken && t2kind == SyntaxKind.EqualsEqualsToken {
        return true
    }

    if t1kind == SyntaxKind.AmpersandToken && t2kind == SyntaxKind.AmpersandAmpersandToken {
        //fmt.Printf("%+v%+v\n", t1kind, t2kind)
        return true
    }

    if t1kind == SyntaxKind.AmpersandAmpersandToken && t2kind == SyntaxKind.AmpersandToken {
        return true
    }

    if t1kind == SyntaxKind.AmpersandToken && t2kind == SyntaxKind.AmpersandToken {
        return true
    }

    if t1kind == SyntaxKind.PipeToken && t2kind == SyntaxKind.PipePipeToken {
        //fmt.Printf("%+v%+v\n", t1kind, t2kind)
        return true
    }

    if t1kind == SyntaxKind.PipePipeToken && t2kind == SyntaxKind.PipeToken {
        return true
    }

    if t1kind == SyntaxKind.PipeToken && t2kind == SyntaxKind.PipeToken {
        return true
    }

    return false
}
