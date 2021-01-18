package SyntaxFacts

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

func GetUnaryOperatorPrecedence(kind SyntaxKind.SyntaxKind) int {
    switch kind {
        case SyntaxKind.PlusToken: 
            return 3
        case SyntaxKind.MinusToken: 
            return 3

        default:
            return 0
    }
}

func GetBinaryOperatorPrecedence(kind SyntaxKind.SyntaxKind) int {
    switch kind {
        case SyntaxKind.StarToken:
            return 2
        case SyntaxKind.SlashToken:
            return 2

        case SyntaxKind.PlusToken: 
            return 1
        case SyntaxKind.MinusToken: 
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
        return SyntaxKind.TrueKeyword
    default:
        return SyntaxKind.IdentifierToken
    }
}
