package Symbols

type StackMapVariableSymbol struct {
    stack []MapVariableSymbol
}

func NewStackMapVariableSymbol() *StackMapVariableSymbol {
    return &StackMapVariableSymbol{
    }
}

func (s *StackMapVariableSymbol) Push(m MapVariableSymbol) {
    s.stack = append([]MapVariableSymbol{m}, s.stack...)
}

func (s *StackMapVariableSymbol) Peek() MapVariableSymbol {
    if len(s.stack) == 0 {
        return nil
    }


    return s.stack[0]
}

func (s *StackMapVariableSymbol) Pop() MapVariableSymbol {
    l := len(s.stack)
    if l == 0 {
        return nil
    }

    el := s.stack[0]

    if l == 1 {
        s.stack = []MapVariableSymbol{}
    } else {
        s.stack = s.stack[1:]
    }

    return el
}
