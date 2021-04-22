package BoundBinaryOperator

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundBinaryOperatorKind"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Symbols"
)

type BoundBinaryOperator struct {
        SyntaxKind SyntaxKind.SyntaxKind
        Kind BoundBinaryOperatorKind.BoundBinaryOperatorKind
        LeftType *Symbols.TypeSymbol
        RightType *Symbols.TypeSymbol
        ResultType *Symbols.TypeSymbol
}

func NewBoundBinaryOperator(
        syntaxKind SyntaxKind.SyntaxKind,
        kind BoundBinaryOperatorKind.BoundBinaryOperatorKind,
        leftType *Symbols.TypeSymbol,
        rightType *Symbols.TypeSymbol,
        resultType *Symbols.TypeSymbol,
    ) *BoundBinaryOperator {
        return &BoundBinaryOperator{
            SyntaxKind: syntaxKind,
            Kind: kind,
            LeftType: leftType,
            RightType: rightType,
            ResultType: resultType,
        }
}

var _operators = []*BoundBinaryOperator{
    NewBoundBinaryOperator(SyntaxKind.PlusToken, BoundBinaryOperatorKind.Addition, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt),
    NewBoundBinaryOperator(SyntaxKind.MinusToken, BoundBinaryOperatorKind.Subtraction, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt),
    NewBoundBinaryOperator(SyntaxKind.StarToken, BoundBinaryOperatorKind.Multiplication, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt),
    NewBoundBinaryOperator(SyntaxKind.SlashToken, BoundBinaryOperatorKind.Division, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt),
    NewBoundBinaryOperator(SyntaxKind.EqualsEqualsToken, BoundBinaryOperatorKind.Equals, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.BangEqualsToken, BoundBinaryOperatorKind.NotEquals, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.LessToken, BoundBinaryOperatorKind.Less, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.LessOrEqualsToken, BoundBinaryOperatorKind.LessOrEquals, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.GreaterToken, BoundBinaryOperatorKind.Greater, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.GreaterOrEqualsToken, BoundBinaryOperatorKind.GreaterOrEquals, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.AmpersandToken, BoundBinaryOperatorKind.BitwiseAnd, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt),
    NewBoundBinaryOperator(SyntaxKind.PipeToken, BoundBinaryOperatorKind.BitwiseOr, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt),
    NewBoundBinaryOperator(SyntaxKind.HatToken, BoundBinaryOperatorKind.BitwiseXor, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt),

    NewBoundBinaryOperator(SyntaxKind.AmpersandToken, BoundBinaryOperatorKind.BitwiseAnd, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.AmpersandAmpersandToken, BoundBinaryOperatorKind.LogicalAnd, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.PipeToken, BoundBinaryOperatorKind.BitwiseOr, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.PipePipeToken, BoundBinaryOperatorKind.LogicalOr, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.HatToken, BoundBinaryOperatorKind.BitwiseXor, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.EqualsEqualsToken, BoundBinaryOperatorKind.Equals, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.BangEqualsToken, BoundBinaryOperatorKind.NotEquals, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool),

    NewBoundBinaryOperator(SyntaxKind.PlusToken, BoundBinaryOperatorKind.Addition, Symbols.TypeSymbolString, Symbols.TypeSymbolString, Symbols.TypeSymbolString),
    NewBoundBinaryOperator(SyntaxKind.EqualsEqualsToken, BoundBinaryOperatorKind.Equals, Symbols.TypeSymbolString, Symbols.TypeSymbolString, Symbols.TypeSymbolBool),
    NewBoundBinaryOperator(SyntaxKind.BangEqualsToken, BoundBinaryOperatorKind.NotEquals, Symbols.TypeSymbolString, Symbols.TypeSymbolString, Symbols.TypeSymbolBool),
}

func Bind(syntaxKind SyntaxKind.SyntaxKind, leftType *Symbols.TypeSymbol, rightType *Symbols.TypeSymbol) *BoundBinaryOperator {
    for _, op := range _operators {
        if op.SyntaxKind == syntaxKind && leftType == op.LeftType && rightType == op.RightType {
            return op
        }
    }

    return nil
}
