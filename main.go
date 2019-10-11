package main

import (
	"github.com/suckowbiz/mentalbreak/internal/interaction"
	"github.com/suckowbiz/mentalbreak/internal/school"
)

const (
	// repeat a mistake every n-th task
	repeatModulus = 5
)

func main() {
	io := interaction.NewCmdLine()
	duration := io.ReadDuration()
	observer := school.NewObserver(duration, repeatModulus)

	io.PrintStart()
	school.NewLesson().Run(io, observer)
}
