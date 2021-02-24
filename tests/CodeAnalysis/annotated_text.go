package CodeAnalysisTest

import (
    "minsk/CodeAnalysis/Text"
    "minsk/Util"
    "strings"
    "bufio"
    "math"
    "fmt"
)


type AnnotatedText struct {
    Text string
    Spans []*Text.TextSpan
}

func NewAnnotatedText(text string, spans []*Text.TextSpan) *AnnotatedText {
    return &AnnotatedText{
        Text: text,
        Spans: spans,
    }
}

func AnnotatedTextParse(text string) *AnnotatedText {
    text = annotatedTextUnindent(text)

    var textBuilder []rune
    var spanBuilder []*Text.TextSpan
    startStack := Util.NewStack()

    var position int
    for _, c := range text {
        if c == '[' {
            startStack.Push(position)
        } else if c == ']' {
            if startStack.Count() == 0 {
                message := fmt.Sprintf("Too many ']' in text %s", text)
                panic(message)
            }

            start := startStack.Pop().(int)
            end := position
            span := Text.NewTextSpan(start, end - start)
            spanBuilder = append(spanBuilder, span)
        } else {
            position = position + 1
            textBuilder = append(textBuilder, c)
        }
    }

    if startStack.Count() != 0 {
        message := fmt.Sprintf("Missing ']' in text %s", text)
        panic(message)
    }

    return NewAnnotatedText(string(textBuilder), spanBuilder)
}

func annotatedTextUnindent(text string) string {
    lines := AnnotatedTextUnindentLines(text)

    newline := fmt.Sprintln()
    return strings.Join(lines, newline) 
}

func AnnotatedTextUnindentLines(text string) []string {
    stringReader := strings.NewReader(text)
    reader := bufio.NewScanner(stringReader)

    var lines []string
    for reader.Scan() {
        line := reader.Text()
        lines = append(lines, line)
    }

    minIndentation := math.MaxInt32
    for i, line := range lines {
        if strings.TrimSpace(line) == "" {
            lines[i] = ""
            continue
        }

        indentation := getIndentation(line) 
        if indentation < minIndentation {
            minIndentation = indentation
        }
    }

    for i, _ := range lines {
        if len(lines[i]) <= minIndentation {
            continue
        }

        lines[i] = lines[i][minIndentation:]
    }

    var i int
    for i = 0; i < len(lines) && len(lines[i]) == 0; i = i + 1 {}
    lines = lines[i:]

    for i = len(lines) - 1; i >= 0 && len(lines[i]) == 0; i = i - 1 {}
    if i + 1 < len(lines) {
        lines = lines[:i+1]
    }

    return lines
}

func getIndentation(text string) int {
    trimmed := strings.TrimSpace(text)
    return strings.Index(text, trimmed)
}
