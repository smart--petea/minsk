package SyntaxFacts

import (
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

func GetUnaryOperatorPrecedence(kind SyntaxKind) int {
    switch kind {
        case PlusToken: 
            return 3
        case MinusToken: 
            return 3

        default:
            return 0
    }
}

func GetBinaryOperatorPrecedence(kind SyntaxKind) int {
    switch kind {
        case StarToken:
            return 2
        case SlashToken:
            return 2

        case PlusToken: 
            return 1
        case MinusToken: 
            return 1

        default:
            return 0
    }
}

func GetKeywordKind(text string) SyntaxKind {
    switch text {
    case "true":
        return SyntaxKind.TrueKeyword
    case "false":
        return SyntaxKind.TrueKeyword
    default:
        return SyntaxKind.IdentifierToken
    }
}
