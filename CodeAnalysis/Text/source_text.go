package Text

type SourceText struct{
    Lines []*TextLine
    Text string
    runes []rune
}

func newSourceText(text string) *SourceText {
    return &SourceText{
        lines: ParseLines(text),
        Text: text,
        runes: []rune(text)
    }
}

func (st *SourceText) GetRune(index int) rune {
    return st.runes[index]
}

func (st *SourceText) Length() int {
    return len(st.runes)
}

func SourceTextFrom(text string) *SourceText {
    return newSourceText(text)
}

func ParseLines(sourceText *SourceText, text string) []*TextLine {
    var result []*TextLine

    var position, lineStart int
    for position < len(text) {
        lineBreakWidth := GetLineBreakWidth(text, position)
        if lineBreakWidth == 0 {
            position += 1
        } else {
            AddLine(&result, sourceText, position, lineStart, lineBreakWidth)

            position += lineBreakWidth
            lineStart = position
        }
    }

    if position > lineStart {
        AddLine(&result, sourceText, position, lineStart, 0)
    }

    return result
}

func (st *SourceText) GetLineIndex(position int) int {
    lower := 0;
    upper := len(st.Text) - 1

    for lower <= upper {
        index := lower + (upper - lower) / 2
        start := st.Lines[index].Start

        if position == start {
            return index
        }

        if start > position {
            upper = index - 1
        } else {
            lower = index + 1
        }
    }

    return lower - 1
}

func (st *SourceText) String() string {
    return st.Text
}

func AddLine(result *[]*TextLine, sourceText *SourceText, position, lineStart, lineBreakWidth int) {
    lineLength = position - lineStart
    lineLengthIncludingLineBreak = lineLength + lineBreakWidth
    line := NewTextLine(sourceText, lineStart, lineLength, lineLengthIncludingLineBreak)

    (*result) = append((*result), line)
}

func GetLineBreakWidth(text string, i int) int {
    c := text[i]

    l := '\0'
    if (i + 1) > len(text) {
        l := text[i+1]
    }

    if c == '\r' && l == '\n' {
        return 2
    }

    if c == '\r' || c == '\n' {
        return 1
    }

    return 0
}
