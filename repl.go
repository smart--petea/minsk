package minsk

import (
    "fmt"
    "bufio"
    "os"
    //"log"
    "strings"

    "minsk/Util"
    "minsk/Util/Console"
    "minsk/Util/Console/KeyInfo"
)

type Repl struct {
    submissionHistory []string
    submissionHistoryIndex int

    done bool
    reader *bufio.Reader
    submissionText string

    EvaluateSubmission func(text string) 
    IsCompleteSubmission func(text string) bool
}

func (r *Repl) clearHistory() {
    r.submissionHistory = []string{}
}

func (r *Repl) RenderLine(line string) {
    //log.Printf("Repl.RenderLine %s", line)
    fmt.Println(line)
}

func (r *Repl) getReader() *bufio.Reader {
    if r.reader == nil {
        r.reader = bufio.NewReader(os.Stdin)
    }

    return r.reader
}

func (r *Repl) GetInput() (string, error) {
    return r.getReader().ReadString('\n')
}

func (r *Repl) Run(ir IRepl) {
    Console.Init()
    for {
        text := r.editSubmission(ir)
        if text == "" {
            return
        }

        if strings.Contains(text, "\n") == false && strings.HasPrefix(text, "#") {
            ir.EvaluateMetaCommand(text)
        } else {
            ir.EvaluateSubmission(text)
        }

        r.submissionHistory = append(r.submissionHistory, text)
        r.submissionHistoryIndex = 0

        //log.Printf("Run len(submissionHIstory)=%d history=%+v", len(r.submissionHistory), r.submissionHistory)
    }
}

func (r *Repl) editSubmission(ir IRepl) string {
    r.done = false

    r.submissionText = ""
    document := Util.NewObservableCollection("")
    view := NewSubmissionView(ir.RenderLine, document)

    for r.done == false {
        key := Console.ReadKey()
        r.handleKey(key, document, view)
    }
    fmt.Println()

    return strings.Join(document.Collection, "\n")
}

func (r *Repl) handleKey(key *KeyInfo.KeyInfo, document *Util.ObservableCollection, view *SubmissionView) {
    //log.Printf("handleKey kind=%s", key.Kind())
    switch key.Kind() {
    case KeyInfo.Backspace:
        r.HandleBackspace(document, view)
    case KeyInfo.Enter:
        r.HandleEnter(document, view)
    case KeyInfo.Escape:
        r.HandleEscape(document, view)
    case KeyInfo.LeftArrow:
        r.HandleLeftArrow(document, view)
    case KeyInfo.RightArrow:
        r.HandleRightArrow(document, view)
    case KeyInfo.UpArrow:
        r.HandleUpArrow(document, view)
    case KeyInfo.DownArrow:
        r.HandleDownArrow(document, view)
    case KeyInfo.Delete:
        r.HandleDelete(document, view)
    case KeyInfo.Home:
        r.HandleHome(document, view)
    case KeyInfo.End:
        r.HandleEnd(document, view)
    case KeyInfo.PageUp:
        r.HandlePageUp(document, view)
    case KeyInfo.PageDown:
        r.HandlePageDown(document, view)
    case KeyInfo.AltEnter:
        r.HandleAltEnter(document, view)
    case KeyInfo.Tab:
        r.HandleTab(document, view)
    }

    if key.Kind() == KeyInfo.Symbol {
        r.HandleTyping(document, view, string(key.Bytes))
    }
}

const TabWidth int = 4

func (r *Repl) HandleTab(document *Util.ObservableCollection, view *SubmissionView) {
    start := view.GetCurrentCharacter()
    remainingSpaces := TabWidth - start % TabWidth;

    lineIndex := view.GetCurrentLine()
    line := document.Get(lineIndex)
    before := line[:start]

    after := strings.Repeat(" ", remainingSpaces) + line[start:]
    Console.Print(after)

    line = before + after 
    document.Set(lineIndex, line)
    view.SetCurrentCharacter(start + remainingSpaces)
}

func (r *Repl) HandleAltEnter(document *Util.ObservableCollection, view *SubmissionView) {

    view.SetCurrentCharacter(0)
    view.SetCurrentLine(document.Count() - 1)
    r.done = true
}

func (r *Repl) HandleHome(document *Util.ObservableCollection, view *SubmissionView) {
    view.SetCurrentCharacter(0)
}

func (r *Repl) HandlePageUp(document *Util.ObservableCollection, view *SubmissionView) {
    r.submissionHistoryIndex = r.submissionHistoryIndex - 1
    if r.submissionHistoryIndex < 0 {
        r.submissionHistoryIndex = len(r.submissionHistory) - 1
    }

    r.UpdateDocumentFromHistory(document, view)
}

func (r *Repl) UpdateDocumentFromHistory(document *Util.ObservableCollection, view *SubmissionView) {
    if len(r.submissionHistory) == 0 {
        return 
    }

    document.Clear()

    historyItem := r.submissionHistory[r.submissionHistoryIndex]
    lines := strings.Split(historyItem, "\n")

    for i, l := range lines {
        document.Add(l)
        view.SetCurrentCharacter(0)
        view.SetCurrentLine(i)
        view.lineRenderer(l)
    }

    line := lines[len(lines) - 1]
    view.SetCurrentCharacter(len(line))
}

func (r *Repl) HandlePageDown(document *Util.ObservableCollection, view *SubmissionView) {
    r.submissionHistoryIndex = r.submissionHistoryIndex + 1
    if r.submissionHistoryIndex > len(r.submissionHistory)-1 {
        r.submissionHistoryIndex = 0
    }

    r.UpdateDocumentFromHistory(document, view)
}

func (r *Repl) HandleEnd(document *Util.ObservableCollection, view *SubmissionView) {
    lineIndex := view.GetCurrentLine()
    line := document.Get(lineIndex)
    view.SetCurrentCharacter(len(line))
}

func (r *Repl) HandleDelete(document *Util.ObservableCollection, view *SubmissionView) {
    lineIndex := view.GetCurrentLine()
    line := document.Get(lineIndex)
    start := view.GetCurrentCharacter()

    if len(line) == 0 || start >= len(line) {
        if document.Count() - 1 == lineIndex {
            return
        }

        nextLine := document.Get(lineIndex + 1)
        line = line + nextLine
        document.Set(lineIndex, line)
        document.RemoveAt(lineIndex + 1)
        return
    }

    line = line[:start] + line[start + 1:]
    document.Set(lineIndex, line)
    Console.Print(line[start:] + " ")
}

func (r *Repl) HandleBackspace(document *Util.ObservableCollection, view *SubmissionView) {
    currentLineIndex := view.GetCurrentLine()
    currentLine := document.Get(currentLineIndex)
    start := view.GetCurrentCharacter()
    //log.Printf("HandleBackspace len(line)=%d start=%d", len(line), start)
    if len(currentLine) == 0 || start == 0 {
        if currentLineIndex <= 0 {
            return
        }

        previousLine := document.Get(currentLineIndex - 1)
        document.RemoveAt(currentLineIndex)
        view.SetCurrentLine(currentLineIndex - 1)
        document.Set(currentLineIndex - 1, previousLine + currentLine)
        view.SetCurrentCharacter(len(previousLine))

        return
    } else {

        before := currentLine[:start-1]
        after := currentLine[start:]
        currentLine = before + after
        document.Set(currentLineIndex, currentLine)

        view.SetCurrentCharacter(start - 1)
        Console.Print(after + " ")
    }
}

func (r *Repl) HandleEscape(document *Util.ObservableCollection, view *SubmissionView) {
    view.SetCurrentCharacter(0)

    document.Clear()
    document.Add("")
    view.SetCurrentLine(0)

    currentLineIndex := view.GetCurrentLine()
    line := document.Get(currentLineIndex)
    Console.Print(strings.Repeat(" ", len(line)))
}

func (r *Repl) HandleEnter(document *Util.ObservableCollection, view *SubmissionView) {
    lines := document.Collection
    submissionText :=  strings.Join(lines, "\n")
    if strings.HasPrefix(submissionText, "#") || r.IsCompleteSubmission(submissionText) {
        r.done = true

        view.SetCurrentCharacter(0)
        view.SetCurrentLine(document.Count() - 1)

        return
    }

    ReplInsertLine(document, view)
}

func ReplInsertLine(document *Util.ObservableCollection, view *SubmissionView) {
    lineIndex := view.GetCurrentLine()
    currentCharacter := view.GetCurrentCharacter()

    line := document.Get(lineIndex)

    document.Set(lineIndex, line[:currentCharacter])
    document.Insert(lineIndex + 1, line[currentCharacter:])
    view.SetCurrentCharacter(0)
    view.SetCurrentLine(lineIndex + 1)
}

func (r *Repl) HandleLeftArrow(document *Util.ObservableCollection, view *SubmissionView) {
    currentCharacter := view.GetCurrentCharacter()
    if currentCharacter > 0 {
        view.SetCurrentCharacter(currentCharacter - 1)
    }
}

func (r *Repl) HandleRightArrow(document *Util.ObservableCollection, view *SubmissionView){
    line := document.Get(view.GetCurrentLine())

    currentCharacter := view.GetCurrentCharacter()
    if currentCharacter < len(line) {
        view.SetCurrentCharacter(currentCharacter + 1)
    }
}

func (r *Repl) HandleUpArrow(document *Util.ObservableCollection, view *SubmissionView) {
    currentLineIndex := view.GetCurrentLine()
    if currentLineIndex > 0 {
        view.SetCurrentLine(currentLineIndex - 1)
    }
}

func (r *Repl) HandleDownArrow(document *Util.ObservableCollection, view *SubmissionView) {
    currentLine := view.GetCurrentLine()
    if currentLine < document.Count() - 1 {
        view.SetCurrentLine(currentLine + 1)
    }
}

func (r *Repl) HandleTyping(document *Util.ObservableCollection, view *SubmissionView, text string) {
    lineIndex := view.GetCurrentLine()
    start := view.GetCurrentCharacter()
    //log.Printf("handleTyping start=%d lineIndex=%d", start, lineIndex)

    line := document.Get(lineIndex)
    line =  line[:start] + text + line[start:]
    fmt.Printf("%s", line[start:])

    //log.Printf("handleTyping line=%s", string(line))

    document.Set(lineIndex, line)
    currentCharacter := view.GetCurrentCharacter() 
    currentCharacter = currentCharacter + len(text)
    view.SetCurrentCharacter(currentCharacter)
}

func (r *Repl) EvaluateMetaCommand(input string) {
    Console.ForegroundColour(Console.COLOUR_RED)
    fmt.Printf("Invalid command %s.", input)
    Console.ResetColour()
}
