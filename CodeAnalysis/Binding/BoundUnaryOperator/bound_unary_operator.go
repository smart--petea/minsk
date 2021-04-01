package BoundUnaryOperator

import (
    "minsk/CodeAnalysis/Binding/Kind/BoundUnaryOperatorKind"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/CodeAnalysis/Symbols"
)

type BoundUnaryOperator struct {
        SyntaxKind SyntaxKind.SyntaxKind
        Kind BoundUnaryOperatorKind.BoundUnaryOperatorKind
        OperandType *Symbols.TypeSymbol
        ResultType *Symbols.TypeSymbol
}

func NewBoundUnaryOperator(
        syntaxKind SyntaxKind.SyntaxKind,
        kind BoundUnaryOperatorKind.BoundUnaryOperatorKind,
        operandType *Symbols.TypeSymbol,
        resultType *Symbols.TypeSymbol,
    ) *BoundUnaryOperator {
        return &BoundUnaryOperator{
            SyntaxKind: syntaxKind,
            Kind: kind,
            OperandType: operandType,
            ResultType: resultType,
        }
}

var _operators = []*BoundUnaryOperator{
    NewBoundUnaryOperator(SyntaxKind.BangToken, BoundUnaryOperatorKind.LogicalNegation, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool),
    NewBoundUnaryOperator(SyntaxKind.PlusToken, BoundUnaryOperatorKind.Identity, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt),
    NewBoundUnaryOperator(SyntaxKind.MinusToken, BoundUnaryOperatorKind.Negation, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt),
    NewBoundUnaryOperator(SyntaxKind.TildeToken, BoundUnaryOperatorKind.OnesComplement, Symbols.TypeSymbolInt, Symbols.TypeSymbolInt),
    NewBoundUnaryOperator(SyntaxKind.TildeToken, BoundUnaryOperatorKind.OnesComplement, Symbols.TypeSymbolBool, Symbols.TypeSymbolBool),
}

func Bind(syntaxKind SyntaxKind.SyntaxKind, operandType *Symbols.TypeSymbol) *BoundUnaryOperator {
    for _, op := range _operators {
        if op.SyntaxKind == syntaxKind && op.OperandType == operandType {
            return op
        }
    }

    return nil
}
