package SyntaxTest

import (
    "testing"
    "minsk/CodeAnalysis/Syntax"
    SyntaxKind "minsk/CodeAnalysis/Syntax/Kind"
)

type AssertingEnumerator struct {
    enumerator <-chan Syntax.CoreSyntaxNode
    t *testing.T
}

func NewAssertingEnumerator(node Syntax.CoreSyntaxNode, t *testing.T) *AssertingEnumerator {
    return &AssertingEnumerator{
        enumerator: flatten(node),
        t: t,
    }
}

func (ae *AssertingEnumerator) AssertToken(kind SyntaxKind.SyntaxKind, text string) {
    if current, isOpen := <- ae.enumerator; isOpen {
        var token *Syntax.SyntaxToken
        var ok bool
        if token, ok = current.(*Syntax.SyntaxToken); !ok {
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
        if _, ok := current.(*Syntax.SyntaxToken); ok {
            ae.t.Errorf("current should not be a SyntaxToken")
        }

        if current.Kind() != kind {
            ae.t.Errorf("current.Kind=%s, expected=%s", string(current.Kind()), string(kind))
        }

        return
    }

    ae.t.Errorf("no next token")
}
