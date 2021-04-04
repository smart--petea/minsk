package Syntax

type SeparatedSyntaxList struct {
    separatorsAndNodes []SyntaxNode
}

func NewSeparatedSyntaxList(separatorsAndNodes []SyntaxNode) *SeparatedSyntaxList {
    return &SeparatedSyntaxList{
        separatorsAndNodes: separatorsAndNodes,
    }
}

func (s *SeparatedSyntaxList) Count() int {
    return (len(s.separatorsAndNodes) + 1) / 2
}

func (s *SeparatedSyntaxList) Get(index int) SyntaxNode {
    return s.separatorsAndNodes[index * 2]
}

func (s *SeparatedSyntaxList) GetSeparator(index int) *SyntaxToken {
    return s.separatorsAndNodes[index * 2 + 1]
}

func (s *SeparatedSyntaxList) GetWithSeparators(index int) []SyntaxNode {
    return s.separatorsAndNodes
}

func (s *SeparatedSyntaxList) GetEnumerator(index int) <-chan SyntaxNode {
    c := make(chan SyntaxNode)

    go func() {
        for _, s := range s.separatorsAndNodes {
            c<-s
        }
    }()

    return c
}
