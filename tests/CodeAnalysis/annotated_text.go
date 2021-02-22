package CodeAnalysisTest

import (
    "minsk/CodeAnalysis/Text"
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
}

func AnnotatedTextUnindent(text string) string {
    stringReader := strings.NewReader(text)
    reader := bufio.NewScanner(stringReader)

    var lines []string
    for reader.Scan() {
        line := reader.Text()
        lines = append(lines, line)
    }

    minIndentation := math.MaxInt32
    for _, line := range lines {
        if strings.TrimSpace(line) == "" {
            lines[i] = ""
            continue
        }

        indentation := getIndentation(line) 
        if indentation < minIndetation {
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
    for i < len(lines) && len(lines[i]) == 0; i = i + 1 {}
    lines = lines[i:]

    i = len(lines) - 1
    for i >= 0 && lines[i] == 0; i = i - 1 {}
    if i + 1 < len(lines) {
        lines = lines[:i+1]
    }

    newline := fmt.Sprintln()
    return strings.Join(lines, newline) 
}

func getIndentation(text string) int {
    trimmed := strings.TrimSpace(text)
    return strings.Index(text, trimmed)
}
