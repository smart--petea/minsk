package Lowering

import (
    "minsk/CodeAnalysis/Binding"
    "minsk/CodeAnalysis/Binding/BoundBinaryOperator"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "reflect"
)

type Lowerer struct {
    Binding.BoundTreeRewriter
}

func newLowerer() *Lowerer {
    return &Lowerer{}
}

func LowererLower(statement Binding.BoundStatement) Binding.BoundStatement {
    lowerer := newLowerer()

    return lowerer.RewriteStatement(lowerer, statement)
}

func (*Lowerer) RewriteForStatement(b Binding.BoundITreeRewriter, node *Binding.BoundForStatement) Binding.BoundStatement {
    //for i = <lower> to <upper>
    //      <body>
    //
    // ---->
    //{
    //  var <var> = <lower>
    //  while (<var> <= <upper>)
    //  {
    //     <body>
    //     <var> = <var> + 1
    //  }
    //}

    variableDeclaration := Binding.NewBoundVariableDeclaration(node.Variable, node.LowerBound)
    variableExpression := Binding.NewBoundVariableExpression(node.Variable)
    condition := Binding.NewBoundBinaryExpression(
        variableExpression,
        BoundBinaryOperator.Bind(SyntaxKind.LessOrEqualsToken, reflect.Int, reflect.Int),
        node.UpperBound,
    )

    increment := Binding.NewBoundExpressionStatement(
        Binding.NewBoundAssignmentExpression(
            node.Variable,
            Binding.NewBoundBinaryExpression(
                variableExpression,
                BoundBinaryOperator.Bind(SyntaxKind.PlusToken, reflect.Int, reflect.Int),
                Binding.NewBoundLiteralExpression(1),
            ),
        ),
    )

    whileBody := Binding.NewBoundBlockStatement([]Binding.BoundStatement{node.Body, increment})
    whileStatement := Binding.NewBoundWhileStatement(condition, whileBody)
    result := Binding.NewBoundBlockStatement([]Binding.BoundStatement{variableDeclaration, whileStatement})

    return b.RewriteStatement(b, result)
}
