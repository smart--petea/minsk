package SyntaxTest

import (
    "testing"
    "minsk/CodeAnalysis/Syntax/SyntaxFacts"
    "minsk/CodeAnalysis/Syntax"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
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
            expression := Syntax.ParseSyntaxTree(text).Root

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


func flatten(node Syntax.SyntaxNode) <-chan Syntax.SyntaxNode {
    out := make(chan Syntax.SyntaxNode)

    var stack syntaxNodeStack
    stack.Push(node)

    go func() {
        for stack.Count() > 0 {
            n := stack.Pop()
            out<-n

            children := n.GetChildren()
            for i := len(children)-1; i >= 0; i = i - 1 {
                stack.Push(children[i])
            }
        }

        close(out)
    }()

    return out
}
