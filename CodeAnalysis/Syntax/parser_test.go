package Syntax

import (
    "testing"
     "minsk/CodeAnalysis/Syntax/SyntaxFacts"
     "fmt"
)

func TestParserBinaryExpressionHonorsPrecedences(t *testing.T) {
    for _, op1 := SyntaxFacts.GetBinaryOperatorKinds() {
        for _, op2 := SyntaxFacts.GetBinaryOperatorKinds() {
            op1Precedence := SyntaxFacts.GetBinaryOperatorPrecedence(op1)
            op2Precedence := SyntaxFacts.GetBinaryOperatorPrecedence(op2)
            op1Text := SyntaxFacts.GetText(op1)
            op2Text := SyntaxFacts.GetText(op2)
            text := fmt.Sprintf("a %s b %s c", op1Text, op2Text) 
            expression := Syntax.ParseSyntaxTree(line).Root

            if op1Precedence >= op2Precedence {
                //     op2
                //     / \
                //   op1  c
                //  /  \   
                // a    b
                e := NewAssertingEnumerator(expression)
                e.AssertNode(SyntaxKind.BinaryExpression)
                    e.AssertNode(SyntaxKind.BinaryExpression)
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentiferToken, "a")
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentiferToken, "b")
                    e.AssertNode(SyntaxKind.NameExpression)
                        e.AssertToken(SyntaxKind.IdentiferToken, "c")
            } else {
                //   op1
                //  /  \
                // a   op2
                //    /  \
                //   b    c
                e := NewAssertingEnumerator(expression)
                e.AssertNode(SyntaxKind.BinaryExpression)
                    e.AssertNode(SyntaxKind.NameExpression)
                        e.AssertToken(SyntaxKind.IdentiferToken, "a")
                    e.AssertNode(SyntaxKind.BinaryExpression)
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentiferToken, "b")
                        e.AssertNode(SyntaxKind.NameExpression)
                            e.AssertToken(SyntaxKind.IdentiferToken, "c")
            }
        }
    }
}

type AssertingEnumerator struct {
    //todo C# IDisposable. pay some heed to io.Close
    enumerator: <-chan SyntaxNode
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

func NewAssertingEnumerator(node SyntaxNode) *AssertingEnumerator {
    return &AssertingEnumerator{
        enumerator: flatten(node),
    }
}

func (ae *AssertingEnumerator) AssertToken(kind SyntaxKind.SyntaxKind, text string) {
    var current SyntaxNode
    select {
    case current <- ae.enumerator:
        //
    default:
        t.Errorf("no next token")
    }

    var token SyntaxToken
    if token, ok := current.(SyntaxToken); !ok {
        t.Errorf("current should be a SyntaxToken")
    }

    if token.Kind() != kind {
        t.Errorf("current.Kind=%s, expected=%s", string(token.Kind()), string(kind))
    }
    
    if string(token.Runes) != text {
        t.Errorf("current.Text=%s, expected=%s", string(token.Runes), text)
    }
}

func (ae *AssertingEnumerator) AssertNode(kind SyntaxKind.SyntaxKind) {
    var current SyntaxNode
    select {
    case current <- ae.enumerator:
        //
    default:
        t.Errorf("no next token")
    }

    if _, ok := current.(SyntaxToken); ok {
        t.Errorf("current should not be a SyntaxToken")
    }

    if current.Kind() != kind {
        t.Errorf("current.Kind=%s, expected=%s", string(current.Kind()), string(kind))
    }
}

flatten(node SyntaxNode) out <-chan SyntaxNode {
    var stack syntaxNodeStack
    stack.Push(node)

    for stack.Count() > 0 {
        n := stack.Pop()
        out<-n

        children := n.GetChildren()
        for _, child := range Util.Reverse(children) { //todo Reverse
            stack.Push(child)
        }
    }
}

type syntaxNodeStack {
    stack []SyntaxNode
}


