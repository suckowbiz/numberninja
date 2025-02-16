package interaction

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Mode int

const (
	Multiplication Mode = iota
	Division
	Addition
	Subtraction
	MultiplicationAndDivision
	AdditionAndSubtraction
	All
)

func IntToMode(n int) (Mode, error) {
	m := Mode(n)
	if m >= Multiplication && m <= All {
		return m, nil
	}
	return 0, errors.New("invalid mode")
}

// CmdLiner interacts via command line.
type CmdLiner interface {
	AskProblem(loop int, operator rune, left, right int) int
	PrintDone(count, mistakes int)
	PrintFailure()
	PrintSuccess()
	ReadInt() int
	ReadDuration() time.Duration
	ReadArithmetic() Mode
	PrintStart()
}

// CmdLine is a representation of the commandline as I/O interface.
type CmdLine struct {
	reader *bufio.Reader
}

// NewCmdLine constructs a new CmdLiner
func NewCmdLine() CmdLiner {
	return &CmdLine{reader: bufio.NewReader(os.Stdin)}
}

// AskProblem outputs a problem to solve.
func (i *CmdLine) AskProblem(loop int, operator rune, left, right int) int {
	fmt.Printf("\n\n%d) %d %c %d = ", loop, left, operator, right)
	result := i.ReadInt()
	return result
}

// PrintFailure outputs a failure message.
func (i *CmdLine) PrintFailure() {
	fmt.Printf("Oh, das war nicht richtig!")
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
	fmt.Printf("\nOk, es geht los!")
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
		fmt.Printf("Das kann ich nicht lesen!\nAuswahl: ")
	}
	return result
}

// ReadDuration asks for input of duration and returns it.
func (i *CmdLine) ReadDuration() time.Duration {
	fmt.Printf("Wie viele Minuten möchtest du Rechnen? Gib die Anzahl ein: ")
	input := i.ReadInt()
	duration := fmt.Sprintf("%dm", input)
	res, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
	}
	return res
}

func (i *CmdLine) ReadArithmetic() Mode {
	fmt.Printf("Welche Rechenarten möchtest du rechnen - "+
		"Multiplikation (%d), Division(%d), Addition (%d), Subtraktion (%d), "+
		"Multiplikation und Division (%d), Subtraktion und Addition (%d), Alle (%d)?",
		Multiplication, Division, Addition, Subtraction, MultiplicationAndDivision, AdditionAndSubtraction, All)
	fmt.Printf("\nGib die Zahl ein: ")

	var mode Mode
	var err error
	for {
		result := i.ReadInt()
		mode, err = IntToMode(result)
		if err == nil {
			break
		}
		fmt.Printf("Das kann ich nicht lesen!\nZahl: ")
	}
	return mode
}
