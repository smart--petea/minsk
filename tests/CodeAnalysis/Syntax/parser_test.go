package SyntaxTest

import (
    "testing"
    "minsk/CodeAnalysis/Syntax/SyntaxFacts"
    "minsk/CodeAnalysis/Syntax"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
    "minsk/Util"
    "fmt"
)

func TestParserBinaryExpressionHonorsPrecedences(t *testing.T) {
    for _, op1 := range SyntaxFacts.GetBinaryOperatorKinds() {
        for _, op2 := range SyntaxFacts.GetBinaryOperatorKinds() {
            op1Precedence := SyntaxFacts.GetBinaryOperatorPrecedence(op1)
            op2Precedence := SyntaxFacts.GetBinaryOperatorPrecedence(op2)
            op1Text := SyntaxFacts.GetText(op1)
            op2Text := SyntaxFacts.GetText(op2)
            text := fmt.Sprintf("a %s b %s c", op1Text, op2Text) 
            expression := parseExpression(text, t).(*Syntax.ExpressionStatementSyntax).Expression

            if op1Precedence >= op2Precedence {
                //     op2
                //     / \
                //   op1  c
                //  /  \   
                // a    b
                e := NewAssertingEnumerator(expression, t)
                e.AssertNode(SyntaxKind.BinaryExpression)
                    e.AssertNode(SyntaxKind.BinaryExpression)
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentifierToken, "a")
                        e.AssertToken(op1, op1Text)
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentifierToken, "b")
                    e.AssertToken(op2, op2Text)
                    e.AssertNode(SyntaxKind.NameExpression)
                        e.AssertToken(SyntaxKind.IdentifierToken, "c")
            } else {
                //   op1
                //  /  \
                // a   op2
                //    /  \
                //   b    c
                e := NewAssertingEnumerator(expression, t)
                e.AssertNode(SyntaxKind.BinaryExpression)
                    e.AssertNode(SyntaxKind.NameExpression)
                        e.AssertToken(SyntaxKind.IdentifierToken, "a")
                    e.AssertToken(op1, op1Text)
                    e.AssertNode(SyntaxKind.BinaryExpression)
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentifierToken, "b")
                        e.AssertToken(op2, op2Text)
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentifierToken, "c")
            }
        }
    }
}

func TestParserUnaryExpressionHonorsPrecedences(t *testing.T) {
    for _, unary := range SyntaxFacts.GetUnaryOperatorKinds() {
        for _, binary := range SyntaxFacts.GetBinaryOperatorKinds() {
            unaryPrecedence := SyntaxFacts.GetUnaryOperatorPrecedence(unary)
            binaryPrecedence := SyntaxFacts.GetBinaryOperatorPrecedence(binary)
            unaryText := SyntaxFacts.GetText(unary)
            binaryText := SyntaxFacts.GetText(binary)
            text := fmt.Sprintf("%s a %s b", unaryText, binaryText) 
            expression := parseExpression(text, t).(*Syntax.ExpressionStatementSyntax).Expression

            if unaryPrecedence >= binaryPrecedence {
                //   binary  
                //   /    \   
                // unary   b
                //  |
                //  a
                e := NewAssertingEnumerator(expression, t)
                e.AssertNode(SyntaxKind.BinaryExpression)
                    e.AssertNode(SyntaxKind.UnaryExpression)
                        e.AssertToken(unary, unaryText)
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentifierToken, "a")
                    e.AssertToken(binary, binaryText)
                    e.AssertNode(SyntaxKind.NameExpression)
                        e.AssertToken(SyntaxKind.IdentifierToken, "b")
            } else {
                //   unary
                //     |
                //  binary
                //  /    \
                // a      b

                e := NewAssertingEnumerator(expression, t)
                e.AssertNode(SyntaxKind.UnaryExpression)
                    e.AssertToken(unary, unaryText)
                    e.AssertNode(SyntaxKind.BinaryExpression)
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentifierToken, "a")
                        e.AssertToken(binary, binaryText)
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentifierToken, "b")
            }
        }
    }
}


func parseExpression(text string, t *testing.T) Syntax.ExpressionSyntax {
    var syntaxTree *Syntax.SyntaxTree = Syntax.SyntaxTreeParse(text)
    var root *Syntax.CompilationUnitSyntax = syntaxTree.Root
    var statement Syntax.StatementSyntax = root.Statement

    var expression Syntax.ExpressionSyntax
    var ok bool
    if expression, ok = statement.(Syntax.ExpressionSyntax); !ok {
        t.Errorf("statement %+v can't be converted to ExpressionSyntax interface", statement)
    }

    return expression
}

func flatten(node Syntax.CoreSyntaxNode) <-chan Syntax.CoreSyntaxNode {
    out := make(chan Syntax.CoreSyntaxNode)

    stack := Util.NewStack()
    stack.Push(interface{}(node))

    go func() {
        for stack.Count() > 0 {
            n := stack.Pop().(Syntax.CoreSyntaxNode)
            out<-n

            var children []Syntax.CoreSyntaxNode
            for child := range n.GetChildren() {
                children = append(children, child)
            }

            for i := len(children)-1; i >= 0; i = i - 1 {
                stack.Push(interface{}(children[i]))
            }
        }

        close(out)
    }()

    return out
}
