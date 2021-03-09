package Lowering

import (
    "minsk/CodeAnalysis/Binding"
    "minsk/CodeAnalysis/Binding/BoundBinaryOperator"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/Util"

    "reflect"
    "fmt"
)

type Lowerer struct {
    Binding.BoundTreeRewriter

    labelCount int
}

func newLowerer() *Lowerer {
    return &Lowerer{}
}

func LowererLower(statement Binding.BoundStatement) *Binding.BoundBlockStatement {
    lowerer := newLowerer()
    result := lowerer.RewriteStatement(lowerer, statement)

    return lowererFlatten(result)
}

func lowererFlatten(statement Binding.BoundStatement) *Binding.BoundBlockStatement {
    var builder []Binding.BoundStatement

    stack := Util.NewStack()
    stack.Push(statement)

    for stack.Count() > 0 {
        current := stack.Pop()

        if boundBlockStatement, ok := current.(*Binding.BoundBlockStatement); ok {
            statements := boundBlockStatement.Statements
            for i := len(statements) - 1; i >= 0; i = i - 1 {
                stack.Push(statements[i])
            }
        } else {
            builder = append(builder, current.(Binding.BoundStatement))
        }
    }

    return Binding.NewBoundBlockStatement(builder)
}

func (l *Lowerer) GenerateLabel() *Util.LabelSymbol {
    name := fmt.Sprintf("Label%d", l.labelCount)
    l.labelCount = l.labelCount + 1

    return Util.NewLabelSymbol(name)
}

func (l *Lowerer) RewriteIfStatement(b Binding.BoundITreeRewriter, node *Binding.BoundIfStatement) Binding.BoundStatement {
    if node.ElseStatement == nil {
        //if <condition>
        //      <then>
        //
        // ---->
        //gotoFalse <condition> end
        //<then>
        //end:

        endLabel := l.GenerateLabel()
        gotoFalse := Binding.NewBoundConditionalGotoStatement(endLabel, node.Condition, true)
        endLabelStatement := Binding.NewBoundLabelStatement(endLabel)
        result := Binding.NewBoundBlockStatement([]Binding.BoundStatement{gotoFalse, node.ThenStatement, endLabelStatement}) 
        return b.RewriteStatement(b, result)
    } else {
        //if <condition>
        //          <then>
        //else
        //          <else
        //----->
        //
        //gotoFalse <condition> else
        //<then>
        //goto end
        //else:
        //<else>
        //end:
        endLabel := l.GenerateLabel()
        elseLabel := l.GenerateLabel()

        gotoFalse := Binding.NewBoundConditionalGotoStatement(endLabel, node.Condition, true)
        gotoEndStatement := Binding.NewBoundGotoStatement(endLabel)
        elseLabelStatement := Binding.NewBoundLabelStatement(elseLabel)
        endLabelStatement := Binding.NewBoundLabelStatement(endLabel)
        result := Binding.NewBoundBlockStatement([]Binding.BoundStatement{
            gotoFalse,
            node.ThenStatement,
            gotoEndStatement,
            elseLabelStatement,
            node.ElseStatement,
            endLabelStatement,
        }) 
        return b.RewriteStatement(b, result)
    }
}

func (l *Lowerer) RewriteWhileStatement(b Binding.BoundITreeRewriter, node *Binding.BoundWhileStatement) Binding.BoundStatement {
    //while <condition>
    //      <body>
    //
    //---->
    //
    //goto check
    //continue:
    //<body> 
    //check:
    // gotoTrue <condition> continue
    //end:

    continueLabel := l.GenerateLabel()
    checkLabel := l.GenerateLabel()
    endLabel := l.GenerateLabel()

    gotoCheck := Binding.NewBoundGotoStatement(checkLabel)
    continueLabelStatement := Binding.NewBoundLabelStatement(continueLabel)
    checkLabelStatement := Binding.NewBoundLabelStatement(checkLabel)
    gotoTrue := Binding.NewBoundConditionalGotoStatement(continueLabel, node.Condition, false)
    endLabelStatement := Binding.NewBoundLabelStatement(endLabel)
    result := Binding.NewBoundBlockStatement([]Binding.BoundStatement{
        gotoCheck,
        continueLabelStatement,
        node.Body,
        checkLabelStatement,
        gotoTrue,
        endLabelStatement,
    })

    return b.RewriteStatement(b, result)
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
