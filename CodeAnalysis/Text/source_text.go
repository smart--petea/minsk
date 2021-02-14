package Text

type SourceText struct{
    Lines []*TextLine
    runes []rune
}

func newSourceText(text string) *SourceText {
    sourceText := &SourceText{
        runes: []rune(text),
    }
    sourceText.Lines = ParseLines(sourceText)

    return sourceText
}

func (st *SourceText) GetRune(index int) rune {
    return st.runes[index]
}

func (st *SourceText) GetRunes(left, right int) []rune {
    return st.runes[left:right]
}

func (st *SourceText) Length() int {
    return len(st.runes)
}

func SourceTextFrom(text string) *SourceText {
    return newSourceText(text)
}

func ParseLines(sourceText *SourceText) []*TextLine {
    var result []*TextLine

    var position, lineStart int
    for position < len(sourceText.runes) {
        lineBreakWidth := GetLineBreakWidth(sourceText.runes, position)
        if lineBreakWidth == 0 {
            position += 1
        } else {
            lineLength := position - lineStart
            lineLengthIncludingLineBreak := lineLength + lineBreakWidth
            line := NewTextLine(sourceText, lineStart, lineLength, lineLengthIncludingLineBreak)
            result = append(result, line)

            position += lineBreakWidth
            lineStart = position
        }
    }

    if position >= lineStart {
        lineLength := position - lineStart
        line := NewTextLine(sourceText, lineStart, lineLength, 0)
        result = append(result, line)
    }

    return result
}

func (st *SourceText) GetLineIndex(position int) int {
    lower := 0;
    upper := len(st.Lines) - 1

    for lower <= upper {
        index := lower + (upper - lower) / 2

        if position >= st.Lines[index].Start && position < st.Lines[index].End() {
            return index
        }

        if st.Lines[index].Start > position {
            upper = index - 1
        } else {
            lower = index + 1
        }
    }

    return lower - 1
}

func (st *SourceText) String() string {
    return string(st.runes)
}

func GetLineBreakWidth(runes []rune, i int) int {
    c := runes[i]

    l := '\x00'
    if (i + 1) > len(runes) {
        l = runes[i+1]
    }

    if c == '\r' && l == '\n' {
        return 2
    }

    if c == '\r' || c == '\n' {
        return 1
    }

    return 0
}

func (st *SourceText) StringBySpan(span *TextSpan) string {
    return st.String()[span.Start:span.End()]
}
