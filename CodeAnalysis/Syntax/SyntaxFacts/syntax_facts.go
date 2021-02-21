package SyntaxFacts

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

func GetUnaryOperatorPrecedence(kind SyntaxKind.SyntaxKind) int {
    switch kind {
        case SyntaxKind.PlusToken, SyntaxKind.MinusToken, SyntaxKind.BangToken: 
            return 6

        default:
            return 0
    }
}

func GetBinaryOperatorPrecedence(kind SyntaxKind.SyntaxKind) int {
    switch kind {
        case SyntaxKind.StarToken, SyntaxKind.SlashToken:
            return 5

        case SyntaxKind.PlusToken, SyntaxKind.MinusToken: 
            return 4

        case SyntaxKind.EqualsEqualsToken, SyntaxKind.BangEqualsToken: 
            return 3

        case SyntaxKind.AmpersandAmpersandToken:
            return 2

        case SyntaxKind.PipePipeToken: 
            return 1

        default:
            return 0
    }
}

func GetKeywordKind(text string) SyntaxKind.SyntaxKind {
    switch text {
    case "let":
        return SyntaxKind.LetKeyword
    case "false":
        return SyntaxKind.FalseKeyword
    case "true":
        return SyntaxKind.TrueKeyword
    case "var":
        return SyntaxKind.VarKeyword
    default:
        return SyntaxKind.IdentifierToken
    }
}

func GetText(kind SyntaxKind.SyntaxKind) string {
    switch kind {
    case SyntaxKind.PlusToken:
        return "+"
    case SyntaxKind.MinusToken:
        return "-"
    case SyntaxKind.StarToken:
        return "*"
    case SyntaxKind.SlashToken:
        return "/"
    case SyntaxKind.BangToken:
        return "!"
    case SyntaxKind.EqualsToken:
        return "="
    case SyntaxKind.AmpersandAmpersandToken:
        return "&&"
    case SyntaxKind.PipePipeToken:
        return "||"
    case SyntaxKind.EqualsEqualsToken: 
        return "=="
    case SyntaxKind.BangEqualsToken:
        return "!="
    case SyntaxKind.OpenParenthesisToken:
        return "("
    case SyntaxKind.CloseParenthesisToken:
        return ")"
    case SyntaxKind.OpenBraceToken:
        return "{"
    case SyntaxKind.CloseBraceToken:
        return "}"
    case SyntaxKind.FalseKeyword:
        return "false"
    case SyntaxKind.TrueKeyword:
        return "true"
    case SyntaxKind.LetKeyword:
        return "let"
    case SyntaxKind.VarKeyword:
        return "var"
    default:
        return ""
    }
}

func GetUnaryOperatorKinds() (kinds []SyntaxKind.SyntaxKind) {
    for _, kind := range SyntaxKind.GetValues() {
        if GetUnaryOperatorPrecedence(kind) > 0 {
            kinds = append(kinds, kind)
        }
    }

    return
}

func GetBinaryOperatorKinds() (kinds []SyntaxKind.SyntaxKind) {
    for _, kind := range SyntaxKind.GetValues() {
        if GetBinaryOperatorPrecedence(kind) > 0 {
            kinds = append(kinds, kind)
        }
    }

    return
}
