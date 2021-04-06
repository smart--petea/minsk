package Syntax

type SeparatedSyntaxList struct {
    nodesAndSeparators []SyntaxNode
}

func NewSeparatedSyntaxList(nodesAndSeparators []SyntaxNode) *SeparatedSyntaxList {
    return &SeparatedSyntaxList{

        nodesAndSeparators: nodesAndSeparators,
    }
}

func (s *SeparatedSyntaxList) Count() int {
    return (len(s.nodesAndSeparators) + 1) / 2
}

//class index operator
func (s *SeparatedSyntaxList) Get(index int) SyntaxNode {
    return s.nodesAndSeparators[index * 2]
}

func (s *SeparatedSyntaxList) GetSeparator(index int) *SyntaxToken {
    if index == len(s.nodesAndSeparators) {
        return nil
    }

    syntaxToken, ok := s.nodesAndSeparators[index * 2 + 1].(*SyntaxToken)
    if !ok {
        panic("Can't transform syntaxNode to syntaxToken")
    }
    return syntaxToken
}

func (s *SeparatedSyntaxList) GetWithSeparators() []SyntaxNode {
    return s.nodesAndSeparators
}

//get arguments without comma
func (s *SeparatedSyntaxList) GetEnumerator() <-chan SyntaxNode {
    c := make(chan SyntaxNode)

    go func() {
        defer close(c)

        for i, s := range s.nodesAndSeparators {
            if i % 2 == 1 {
                continue
            }

            c<-s
        }
    }()

    return c
}
