package gap

import (
	"errors"
	"log"
	"os/exec"
)

const (
	horizontalGap, verticalGap = 0.01, 0.03
)

var (
	//ErrEmptyAppName raised when application name is not specified
	ErrEmptyAppName = errors.New("Error: Application name not specified")
)

type point struct {
	x int64
	y int64
}

type windowSize struct {
	height int64
	width  int64
}

//Resize resizes current application based on its properties
func Resize(app *Application) error {
	if err := app.Validate(); err != nil {
		return err
	}

	return resize(app)
}

func resize(app *Application) error {
	cmds := []string{
		"-e",
		`'tell application "System Events" to tell process "` + app.name + `"'`,
		"-e",
		`'tell window 1'`,
		"-e",
		`'set position to ` + app.Position() + "'",
		"-e",
		`'set size to ` + app.Size() + "'",
		"-e",
		"'end tell'",
		"-e",
		"'end tell'",
	}
	log.Println(cmds)
	cmd, err := exec.Command("osascript", cmds...).Output()
	log.Println(string(cmd))
	if err != nil {
		return err
	}
	return nil
}

func windowHeight(height float64) int64 {
	wHeight := height * (1 - (2 * verticalGap))
	return int64(wHeight)
}

func leftWindowWidth(width float64) int64 {
	lwWidth := (width - (3*horizontalGap)*width) / 2
	return int64(lwWidth)
}

func rightWindowWidth(rightPointX int64, leftWindowWidth int64) int64 {
	return rightPointX + leftWindowWidth
}

func rightPoint(screen Screen, leftWindowWidth int64, pY int64) (int64, int64) {
	pX := int64(((2 * horizontalGap) * screen.Width())) + leftWindowWidth
	return pX, pY
}

func leftPoint(screen Screen) (int64, int64) {
	pX := horizontalGap * screen.Width()
	pY := verticalGap * screen.Height()

	return int64(pX), int64(pY)
}
