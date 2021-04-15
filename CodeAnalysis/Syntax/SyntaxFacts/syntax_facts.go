package SyntaxFacts

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

func GetUnaryOperatorPrecedence(kind SyntaxKind.SyntaxKind) int {
    switch kind {
        case SyntaxKind.PlusToken, SyntaxKind.MinusToken, SyntaxKind.BangToken, SyntaxKind.TildeToken: 
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

        case SyntaxKind.EqualsEqualsToken, SyntaxKind.BangEqualsToken, SyntaxKind.LessToken, SyntaxKind.LessOrEqualsToken, SyntaxKind.GreaterToken, SyntaxKind.GreaterOrEqualsToken: 
            return 3

        case SyntaxKind.AmpersandAmpersandToken, SyntaxKind.AmpersandToken:
            return 2

        case SyntaxKind.PipePipeToken, SyntaxKind.PipeToken, SyntaxKind.HatToken: 
            return 1

        default:
            return 0
    }
}

func GetKeywordKind(text string) SyntaxKind.SyntaxKind {
    switch text {
    case "false":
        return SyntaxKind.FalseKeyword
    case "if":
        return SyntaxKind.IfKeyword
    case "else":
        return SyntaxKind.ElseKeyword
    case "let":
        return SyntaxKind.LetKeyword
    case "to":
        return SyntaxKind.ToKeyword
    case "true":
        return SyntaxKind.TrueKeyword
    case "var":
        return SyntaxKind.VarKeyword
    case "while":
        return SyntaxKind.WhileKeyword
    case "for":
        return SyntaxKind.ForKeyword
    case "function":
        return SyntaxKind.FunctionKeyword
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
    case SyntaxKind.CommaToken:
        return ","
    case SyntaxKind.ColonToken:
        return ":"
    case SyntaxKind.EqualsToken:
        return "="
    case SyntaxKind.LessToken:
        return "<"
    case SyntaxKind.LessOrEqualsToken:
        return "<="
    case SyntaxKind.GreaterToken:
        return ">"
    case SyntaxKind.GreaterOrEqualsToken:
        return ">="
    case SyntaxKind.TildeToken:
        return "~"
    case SyntaxKind.HatToken:
        return "^"
    case SyntaxKind.PipeToken:
        return "|"
    case SyntaxKind.AmpersandToken:
        return "&"
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
    case SyntaxKind.IfKeyword:
        return "if"
    case SyntaxKind.ElseKeyword:
        return "else"
    case SyntaxKind.LetKeyword:
        return "let"
    case SyntaxKind.ToKeyword:
        return "to"
    case SyntaxKind.VarKeyword:
        return "var"
    case SyntaxKind.WhileKeyword:
        return "while"
    case SyntaxKind.ForKeyword:
        return "for"
    case SyntaxKind.FunctionKeyword:
        return "function"
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
