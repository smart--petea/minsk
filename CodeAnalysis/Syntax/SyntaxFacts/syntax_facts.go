package SyntaxFacts

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

func GetUnaryOperatorPrecedence(kind SyntaxKind.SyntaxKind) int {
    switch kind {
        case SyntaxKind.PlusToken, SyntaxKind.MinusToken, SyntaxKind.BangToken: 
            return 5

        default:
            return 0
    }
}

func GetBinaryOperatorPrecedence(kind SyntaxKind.SyntaxKind) int {
    switch kind {
        case SyntaxKind.StarToken, SyntaxKind.SlashToken:
            return 4

        case SyntaxKind.PlusToken, SyntaxKind.MinusToken: 
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
    case "true":
        return SyntaxKind.TrueKeyword
    case "false":
        return SyntaxKind.FalseKeyword
    default:
        return SyntaxKind.IdentifierToken
    }
}
