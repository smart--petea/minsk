package Syntax

import (
    "testing"
    "minsk/CodeAnalysis/Syntax/SyntaxFacts"
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
            expression := ParseSyntaxTree(text).Root

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
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentifierToken, "b")
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
                    e.AssertNode(SyntaxKind.BinaryExpression)
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentifierToken, "b")
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentifierToken, "c")
            }
        }
    }
}

type AssertingEnumerator struct {
    //todo C# IDisposable. pay some heed to io.Close
    enumerator <-chan SyntaxNode
    t *testing.T
}

/* todo
public void Dispose()
{
    select {
    case current <- ae.enumerator:
        t.Errorf("it should be no next token")
    default:
    }

    enumerator.Dispose()
}
*/

func NewAssertingEnumerator(node SyntaxNode, t *testing.T) *AssertingEnumerator {
    return &AssertingEnumerator{
        enumerator: flatten(node),
        t: t,
    }
}

func (ae *AssertingEnumerator) AssertToken(kind SyntaxKind.SyntaxKind, text string) {
    if current, isOpen := <- ae.enumerator; isOpen {
        fmt.Printf("AssertToken %+v %+v\n", current, current.Kind())
        var token *SyntaxToken
        var ok bool
        if token, ok = current.(*SyntaxToken); !ok {
            ae.t.Errorf("current should be a SyntaxToken")
        }

        if token.Kind() != kind {
            ae.t.Errorf("current.Kind=%s, expected=%s", string(token.Kind()), string(kind))
        }

        if string(token.Runes) != text {
            ae.t.Errorf("current.Text=%s, expected=%s", string(token.Runes), text)
        }

        return
    }

    ae.t.Errorf("no next token")
}

func (ae *AssertingEnumerator) AssertNode(kind SyntaxKind.SyntaxKind) {
    if current, isChanOpen := <- ae.enumerator; isChanOpen {
        fmt.Printf("AssertNode %+v %+v\n", current, current.Kind())
        if _, ok := current.(*SyntaxToken); ok {
            ae.t.Errorf("current should not be a SyntaxToken")
        }

        if current.Kind() != kind {
            ae.t.Errorf("current.Kind=%s, expected=%s", string(current.Kind()), string(kind))
        }

        return
    }

    ae.t.Errorf("no next token")
}

func flatten(node SyntaxNode) <-chan SyntaxNode {
    out := make(chan SyntaxNode)

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

type syntaxNodeStack struct {
    stack []SyntaxNode
}

func (stack *syntaxNodeStack) Count() int {
    return len(stack.stack)
}

func (stack *syntaxNodeStack) Push(node SyntaxNode) {
    stack.stack = append(stack.stack, node)
}

func (stack *syntaxNodeStack) Pop() SyntaxNode {
    node := stack.stack[len(stack.stack) - 1]
    stack.stack = stack.stack[:len(stack.stack) - 1]

    return node
}
