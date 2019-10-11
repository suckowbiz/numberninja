package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Cmdliner interacts via command line.
type Cmdliner interface {
	AskProblem(loop int, operator rune, left, right int) int
	PrintDone(count, mistakes int)
	PrintFailure(correctAnswer int)
	PrintSuccess()
	ReadInt() int
	ReadDuration() time.Duration
	PrintStart()
}

// CmdLine is a representation of the commandline as I/O interface.
type CmdLine struct {
	reader *bufio.Reader
}

// NewCmdLine constructs a new Cmdliner
func NewCmdLine() Cmdliner {
	return &CmdLine{reader: bufio.NewReader(os.Stdin)}
}

// AskProblem outputs a problem to solve.
func (i *CmdLine) AskProblem(loop int, operator rune, left, right int) int {
	fmt.Printf("\n\n%d.) %d %c %d = ", loop, left, operator, right)
	result := i.ReadInt()
	return result
}

// PrintFailure outputs a failure message.
func (i *CmdLine) PrintFailure(correctAnswer int) {
	fmt.Printf("Oh, das war nicht richtig! Richtig ist: %d", correctAnswer)
}

// PrintDone outputs a completion message.
func (i *CmdLine) PrintDone(count, mistakes int) {
	fmt.Printf("\n\nFertig, die Zeit ist abgelaufen. ")
	fmt.Printf("\nBei %d Aufgabe(n) hast du dich insgesamt %d mal verrechnet!\n\n", count, mistakes)
}

// PrintSuccess outputs a success message.
func (i *CmdLine) PrintSuccess() {
	fmt.Printf("Richtig!")
}

// PrintStart outputs a start message.
func (i *CmdLine) PrintStart() {
	fmt.Printf("Ok, es geht los!")
}

// ReadInt asks for integer input and returns it.
func (i *CmdLine) ReadInt() int {
	var result int
	var err error
	for {
		s, _ := i.reader.ReadString('\n')
		s = strings.TrimSpace(s)
		result, err = strconv.Atoi(s)
		if err == nil {
			break
		}
		fmt.Printf("Das kann ich nicht lesen!\nBitte gib eine Zahl ein: ")
	}
	return result
}

// ReadDuration asks for input of duration and returns it.
func (i *CmdLine) ReadDuration() time.Duration {
	fmt.Printf("Wie viele Minuten möchtest du Rechnen? Gib die Anzahl ein: ")
	min := i.ReadInt()
	duration := fmt.Sprintf("%dm", min)
	res, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}
	return res
}
