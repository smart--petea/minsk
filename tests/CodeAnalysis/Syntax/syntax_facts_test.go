package SyntaxTest

import (
    "testing"
    "minsk/CodeAnalysis/Syntax/SyntaxFacts"
    "minsk/CodeAnalysis/Syntax"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

func TestSyntaxFactGetTextRoundTrips(t *testing.T) {
    kinds := []SyntaxKind.SyntaxKind{
        SyntaxKind.NumberToken,
        SyntaxKind.WhitespaceToken,
        SyntaxKind.PlusToken,
        SyntaxKind.EndOfFileToken,
        SyntaxKind.MinusToken,
        SyntaxKind.StarToken,
        SyntaxKind.SlashToken,
        SyntaxKind.OpenParenthesisToken,
        SyntaxKind.CloseParenthesisToken,
        SyntaxKind.BangToken,
        SyntaxKind.BadToken,
        SyntaxKind.IdentifierToken,
        SyntaxKind.AmpersandAmpersandToken,
        SyntaxKind.PipePipeToken,
        SyntaxKind.EqualsEqualsToken,
        SyntaxKind.EqualsToken,
        SyntaxKind.BangEqualsToken,
        SyntaxKind.TrueKeyword,
        SyntaxKind.FalseKeyword,
        SyntaxKind.AssignmentExpression, 
        SyntaxKind.BinaryExpression, 
        SyntaxKind.LiteralExpression, 
        SyntaxKind.NameExpression, 
        SyntaxKind.ParenthesizedExpression,
        SyntaxKind.UnaryExpression, 
    }

    for _, kind := range kinds {
        text := SyntaxFacts.GetText(kind)
        if text == "" {
            continue
        }

        tokens := Syntax.ParseTokens(text)

        if len(tokens) != 1 {
            t.Errorf("len(%+v)=%d expected=1", tokens, len(tokens))
        }

        token := tokens[0]
        if token.Kind() != kind {
            t.Errorf("kind=%+v, expected=%+v", token.Kind(), kind)
        }

        if text != string(token.Runes) {
            t.Errorf("text=%s, expected=%s", text, string(token.Runes))
        }
    }
}
