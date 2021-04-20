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
