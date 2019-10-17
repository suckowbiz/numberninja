package main

import (
	"github.com/suckowbiz/1x1Pass/internal/interaction"
	"github.com/suckowbiz/1x1Pass/internal/school"
)

const (
	// repeat a mistake every n-th task
	repeatModulus = 5
)

func main() {
	io := interaction.NewCmdLine()
	duration := io.ReadDuration()
	observer := school.NewObserver(duration, repeatModulus)

	school.NewLesson().Run(io, observer)
}
