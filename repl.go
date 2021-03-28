package minsk

import (
    "fmt"
    "bufio"
    "os"
    "os/exec"
    "log"
    "strings"
    "golang.org/x/sys/unix"
    "syscall"
    "strconv"

    "minsk/Util/Console"
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
    ConsoleInit()
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
    document := NewObservableCollection("")
    view := NewSubmissionView(ir.RenderLine, document)

    for r.done == false {
        key := ConsoleReadKey()
        r.handleKey(key, document, view)
    }
    fmt.Println()

    return strings.Join(document.Collection, "\n")
}

func ConsoleInit() {
    //clean screen
    fmt.Print("\033[2J")
    ConsoleSetCursorPos(1,1)

    //disable input buffering
    exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
    //do not display entered characters on the screen
    exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func ConsoleReadKey() *ConsoleKeyInfo {
    b := make([]byte, 8)
    size, _ := os.Stdin.Read(b)
    log.Printf("ConsoleReadKey %+v %s", b, string(b))

    return NewConsoleKeyInfo(b[:size])
}

type ConsoleKeyInfo struct {
    Bytes []byte
}

const (
    NoModifiers int = 0
    Alt int = 1
    Control int = 2
    Shift int = 4
)

func NewConsoleKeyInfo(bytes []byte) *ConsoleKeyInfo {
    return &ConsoleKeyInfo{
        Bytes: bytes,
    }
}

func (c *ConsoleKeyInfo) Kind() ConsoleKey {
    if len(c.Bytes) == 4 {
        if c.Bytes[0] == 27 && c.Bytes[1] == 91 && c.Bytes[3] == 126 {
            switch  c.Bytes[2] {
            case 49:
                return Home
            case 51:
                return Delete
            case 52:
                return End
            case 53:
                return PageUp
            case 54:
                return PageDown
            }
        }
    }

    if len(c.Bytes) == 3 {
        if c.Bytes[0] == 27 && c.Bytes[1] == 91 {
            switch c.Bytes[2] {
            case 67:
                return RightArrow
            case 68:
                return LeftArrow
            case 65:
                return UpArrow
            case 66:
                return DownArrow
            default:
                panic(fmt.Sprintf("Unknown console command %+v", c.Bytes[:3]))
            }
        }
    }

    if len(c.Bytes) == 2 {
        if c.Bytes[0] == 27 && c.Bytes[1] == 10 {
            return AltEnter
        }
    }

    switch c.Bytes[0] {
    case 27:
        return Escape
    case 10:
        return Enter
    case 127:
        return Backspace
    case 9:
        return Tab
    }

    return Symbol
}

type ConsoleKey string

const (
        Backspace ConsoleKey = "Backspace"
        Enter ConsoleKey = "Enter"
        Escape ConsoleKey = "Escape"
        Tab ConsoleKey = "Tab"
        LeftArrow ConsoleKey = "LeftArrow"
        RightArrow ConsoleKey = "RightArrow"
        UpArrow ConsoleKey = "UpArrow"
        DownArrow ConsoleKey = "DownArrow"
        Symbol ConsoleKey = "Symbol"
        Delete ConsoleKey = "Delete"
        Home ConsoleKey = "Home"
        End ConsoleKey = "End"
        PageUp ConsoleKey = "PageUp"
        PageDown ConsoleKey = "PageDown"
        AltEnter ConsoleKey = "AltEnter"
)

type NotifyCollectionChangedEventArgs struct {
}

func NewNotifyCollectionChangedEventArgs() *NotifyCollectionChangedEventArgs {
    return &NotifyCollectionChangedEventArgs{
    }
}

type ObservableCollection struct {
    Collection []string
    Listeners []func(interface{}, *NotifyCollectionChangedEventArgs)
}

func (o *ObservableCollection) CollectionChanged(listener func(interface{}, *NotifyCollectionChangedEventArgs) ) {
    o.Listeners = append(o.Listeners, listener)
}

func (o *ObservableCollection) fireCollectionChanged() {
    //log.Printf("fireCollectionChanged")
    for _, l := range o.Listeners {
        args := NewNotifyCollectionChangedEventArgs()
        l(nil, args)
    }
}

func (o *ObservableCollection) Clear() {
    //log.Printf("ObservableCollection.Clear")
    o.Collection = []string{}
    o.fireCollectionChanged()
}


func (o *ObservableCollection) Add(e string) {
    //log.Printf("ObservableCollection.Add")
    o.Collection = append(o.Collection, e)
    o.fireCollectionChanged()
}

func (o *ObservableCollection) Get(index int) string {
    return o.Collection[index]
}

func (o *ObservableCollection) Set(index int, val string) {
    //log.Printf("ObservableCollection.Set")
    o.Collection[index] = val
    o.fireCollectionChanged()
}

func (o *ObservableCollection) Count() int {
    return len(o.Collection)
}

func NewObservableCollection(collection... string) *ObservableCollection {
    return &ObservableCollection{
        Collection: collection,
    }
}

func (r *Repl) handleKey(key *ConsoleKeyInfo, document *ObservableCollection, view *SubmissionView) {
    log.Printf("handleKey kind=%s", key.Kind())
    switch key.Kind() {
    case Backspace:
        r.HandleBackspace(document, view)
    case Enter:
        r.HandleEnter(document, view)
    case Escape:
        r.HandleEscape(document, view)
    case LeftArrow:
        r.HandleLeftArrow(document, view)
    case RightArrow:
        r.HandleRightArrow(document, view)
    case UpArrow:
        r.HandleUpArrow(document, view)
    case DownArrow:
        r.HandleDownArrow(document, view)
    case Delete:
        r.HandleDelete(document, view)
    case Home:
        r.HandleHome(document, view)
    case End:
        r.HandleEnd(document, view)
    case PageUp:
        r.HandlePageUp(document, view)
    case PageDown:
        r.HandlePageDown(document, view)
    case AltEnter:
        r.HandleAltEnter(document, view)
    case Tab:
        r.HandleTab(document, view)
    }

    if key.Kind() == Symbol {
        r.HandleTyping(document, view, string(key.Bytes))
    }
}

const TabWidth int = 4

func (r *Repl) HandleTab(document *ObservableCollection, view *SubmissionView) {
    start := view.GetCurrentCharacter()
    remainingSpaces := TabWidth - start % TabWidth;

    lineIndex := view.GetCurrentLine()
    line := document.Get(lineIndex)
    before := line[:start]

    after := strings.Repeat(" ", remainingSpaces) + line[start:]
    view.Print(after)

    line = before + after 
    document.Set(lineIndex, line)
    view.SetCurrentCharacter(start + remainingSpaces)
}

func (r *Repl) HandleAltEnter(document *ObservableCollection, view *SubmissionView) {
    r.done = true

    view.SetCurrentCharacter(0)
    view.SetCurrentLine(document.Count() - 1)
}

func (r *Repl) HandleHome(document *ObservableCollection, view *SubmissionView) {
    view.SetCurrentCharacter(0)
}

func (r *Repl) HandlePageUp(document *ObservableCollection, view *SubmissionView) {
    r.submissionHistoryIndex = r.submissionHistoryIndex - 1
    if r.submissionHistoryIndex < 0 {
        r.submissionHistoryIndex = len(r.submissionHistory) - 1
    }

    r.UpdateDocumentFromHistory(document, view)
}

func (r *Repl) UpdateDocumentFromHistory(document *ObservableCollection, view *SubmissionView) {
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

func (r *Repl) HandlePageDown(document *ObservableCollection, view *SubmissionView) {
    r.submissionHistoryIndex = r.submissionHistoryIndex + 1
    if r.submissionHistoryIndex > len(r.submissionHistory)-1 {
        r.submissionHistoryIndex = 0
    }

    r.UpdateDocumentFromHistory(document, view)
}

func (r *Repl) HandleEnd(document *ObservableCollection, view *SubmissionView) {
    lineIndex := view.GetCurrentLine()
    line := document.Get(lineIndex)
    view.SetCurrentCharacter(len(line))
}

func (r *Repl) HandleDelete(document *ObservableCollection, view *SubmissionView) {
    lineIndex := view.GetCurrentLine()
    line := document.Get(lineIndex)
    start := view.GetCurrentCharacter()

    if len(line) == 0 || start >= len(line) {
        return
    }

    line = line[:start] + line[start + 1:]
    document.Set(lineIndex, line)
    view.Print(line[start:] + " ")
}

func (r *Repl) HandleBackspace(document *ObservableCollection, view *SubmissionView) {
    lineIndex := view.GetCurrentLine()
    line := document.Get(lineIndex)
    start := view.GetCurrentCharacter()
    //log.Printf("HandleBackspace len(line)=%d start=%d", len(line), start)
    if len(line) == 0 {
        return
    }

    if start == 0 {
        return
    }


    before := line[:start-1]
    after := line[start:]
    line = before + after
    document.Set(lineIndex, line)

    view.SetCurrentCharacter(start - 1)
    view.Print(after + " ")
}

func (r *Repl) HandleEscape(document *ObservableCollection, view *SubmissionView) {
    view.SetCurrentCharacter(0)

    currentLineIndex := view.GetCurrentLine()
    line := document.Get(currentLineIndex)
    document.Set(currentLineIndex, "")
    view.Print(strings.Repeat(" ", len(line)))
}

func (r *Repl) HandleEnter(document *ObservableCollection, view *SubmissionView) {
    lines := document.Collection
    submissionText :=  strings.Join(lines, "\n")
    if strings.HasPrefix(submissionText, "#") || r.IsCompleteSubmission(submissionText) {
        r.done = true

        view.SetCurrentCharacter(0)
        view.SetCurrentLine(document.Count() - 1)

        return
    }

    document.Add("")
    view.SetCurrentCharacter(0)
    view.SetCurrentLine(document.Count() - 1)
}

func (r *Repl) HandleLeftArrow(document *ObservableCollection, view *SubmissionView) {
    currentCharacter := view.GetCurrentCharacter()
    if currentCharacter > 0 {
        view.SetCurrentCharacter(currentCharacter - 1)
    }
}

func (r *Repl) HandleRightArrow(document *ObservableCollection, view *SubmissionView){
    line := document.Get(view.GetCurrentLine())

    currentCharacter := view.GetCurrentCharacter()
    if currentCharacter < len(line) {
        view.SetCurrentCharacter(currentCharacter + 1)
    }
}

func (r *Repl) HandleUpArrow(document *ObservableCollection, view *SubmissionView) {
    currentLineIndex := view.GetCurrentLine()
    if currentLineIndex > 0 {
        view.SetCurrentLine(currentLineIndex - 1)
    }
}

func (r *Repl) HandleDownArrow(document *ObservableCollection, view *SubmissionView) {
    currentLine := view.GetCurrentLine()
    if currentLine < document.Count() - 1 {
        view.SetCurrentLine(currentLine + 1)
    }
}

func (r *Repl) HandleTyping(document *ObservableCollection, view *SubmissionView, text string) {
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

type SubmissionView struct {
    _currentLine int
    _currentCharacter int
    SubmissionDocument *ObservableCollection
    lineRenderer func(string)

    cursorTop int
    renderedLineCount int
}

func (s *SubmissionView) Print(text string) {
    log.Printf("SubmissionView.Print %s", text)
    left, top := ConsoleGetCursorPos()

    fmt.Printf(text)

    ConsoleSetCursorPos(top, left)
}

func ConsoleGetCursorPos() (left int, top int) {
    fd := os.Stdin.Fd()
    term, err := unix.IoctlGetTermios(int(fd), unix.TCGETS)
    if err != nil {
        panic(err)
    }

    restore := *term
    term.Lflag &^= (syscall.ICANON | syscall.ECHO)
    err = unix.IoctlSetTermios(int(fd), unix.TCSETS, term)
    if err != nil {
        panic(err)
    }

    _, err = os.Stdout.Write([]byte("\033[6n"))//todo save []byte("\033[6n") in a constant
    if err != nil {
        panic(err)
    }

    var x, y int
    var zeroes int = 1
    var firstNumber bool = true
    b := make([]byte, 1)
    for {
        _, err = os.Stdin.Read(b)
        if err != nil {
            panic(err)
        }

        if b[0] == 'R' {
            break
        }

        if b[0] == ';' {
            firstNumber = false
            zeroes = 1
            continue
        }

        if b[0] >= '0' && b[0] <= '9' {
            if firstNumber {
                x = x * zeroes + int(b[0]) - int('0') //todo constants
            } else {
                y = y * zeroes + int(b[0]) - int('0') //todo constants
            }
            zeroes *= 10
            continue
        }
    }

    err = unix.IoctlSetTermios(int(fd), unix.TCSETS, &restore)
    if err != nil {
        panic(err)
    }

    return x, y
}

type ActionString struct {
}

func NewSubmissionView(lineRenderer func(string), submissionDocument *ObservableCollection) *SubmissionView {
    top, _ := ConsoleGetCursorPos()
    s := &SubmissionView{
        SubmissionDocument: submissionDocument,
        cursorTop: top,
        lineRenderer: lineRenderer,
    }

    submissionDocument.CollectionChanged(s.SubmissionDocumentChanged)
    s.Render()

    return s
}

func (s *SubmissionView) SubmissionDocumentChanged(sender interface{}, e *NotifyCollectionChangedEventArgs) {
    s.Render()
}

func ConsoleSetCursorPos(left, top int) {
    fmt.Printf("\033[%d;%dH", top, left)
}


func ConsoleSetCursorVisibile(v bool) {
    if v {
        fmt.Printf("\033]?25h")
    } else {
        fmt.Printf("\033]?25l")
    }
}

func (s *SubmissionView) Render() {
    left := 1
    top := s.cursorTop
    log.Printf("SubmissionView.Render left=%d top=%d ", left, top)

    ConsoleSetCursorVisibile(false)

    var lineCount int
    for _, line := range s.SubmissionDocument.Collection {
        ConsoleSetCursorPos(left, s.cursorTop + lineCount)
        Console.ForegroundColour(Console.COLOUR_GREEN)
        if lineCount == 0 {
            fmt.Print("»")
        } else {
            fmt.Print("·")
        }

        Console.ResetColour()
        s.lineRenderer(line)

        lineCount = lineCount + 1
    }

    numberOfBlankLines := s.renderedLineCount - lineCount
    if numberOfBlankLines > 0 {
        blankLine := strings.Repeat(" ", ConsoleWindowWidth())
        for i := 0; i < numberOfBlankLines; i = i+1 {
            ConsoleSetCursorPos(0, s.cursorTop + lineCount + i)
            fmt.Println(blankLine)
        }
    }

    s.renderedLineCount = lineCount
    ConsoleSetCursorVisibile(true)
    s.UpdateCursorPosition() 
}

func ConsoleWindowWidth() int {
    cmd := exec.Command("stty", "size")
    cmd.Stdin = os.Stdin
    out, err := cmd.Output()
    if err != nil {
        panic(err)
    }

    fields := strings.Fields(string(out))
    i, err := strconv.Atoi(fields[1])
    if err != nil {
        panic(err)
    }

    return i
}

func (s *SubmissionView) UpdateCursorPosition() {
    top := s.cursorTop + s.GetCurrentLine()
    left := 2 + s._currentCharacter

    //log.Printf("UpdateCursorPosition cursorTop=%d lineIndex=%d currentCharacter=%d left=%d top=%d", s.cursorTop, s.GetCurrentLine(), s._currentCharacter, left, top)

    ConsoleSetCursorPos(left, top)
}

func (s *SubmissionView) GetCurrentLine() int {
    return s._currentLine
}

func (s *SubmissionView) SetCurrentLine(value int) {
    //log.Printf("SetCurrentLine old=%d new=%d", s._currentLine, value)
    if value != s._currentLine {
        s._currentLine  = value

        lineLen := len(s.SubmissionDocument.Collection[s._currentLine])
        if s._currentCharacter > lineLen {
            s._currentCharacter  = lineLen
        }

        s.UpdateCursorPosition()
    }
}

func (s *SubmissionView) GetCurrentCharacter() int {
    return s._currentCharacter 
}

func (s *SubmissionView) SetCurrentCharacter(value int) {
    //log.Printf("SetCurrentCharacter old=%d new=%d", s._currentCharacter, value)
    if value != s._currentCharacter {
        s._currentCharacter  = value
        s.UpdateCursorPosition()
    }
}
