package main

import (
	"github.com/suckowbiz/numberninja/internal/interaction"
	"github.com/suckowbiz/numberninja/internal/school"
)

const (
	// repeat a mistake every n-th task
	repeatModulus = 5
)

func main() {
	io := interaction.NewCmdLine()
	duration := io.ReadDuration()
	mode := io.ReadArithmetic()
	observer := school.NewObserver(duration, repeatModulus)

	school.NewLesson().Run(io, observer, mode)
}
