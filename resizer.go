package gap

import (
	"errors"
	"fmt"
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

//Application represents window application that should be resized
type Application struct {
	isLeft   bool
	name     string
	location point
	size     windowSize
}

//NewApplication creates instance of Application
func NewApplication(name string) *Application {
	return &Application{
		name: name,
	}
}

//Validate validates application
func (a *Application) Validate() error {
	if a.name == "" {
		return ErrEmptyAppName
	}

	return nil
}

//Left sets current Application properties to left display
func (a *Application) Left(screen Screen) *Application {
	a.isLeft = true
	a.calculate(screen)

	return a
}

//Right sets current Application properties to right display
func (a *Application) Right(screen Screen) *Application {
	a.isLeft = false
	a.calculate(screen)

	return a
}

//IsLeft checks whether Application is on the left side of the screen. If false the screen is on the right
func (a *Application) IsLeft() bool {
	return a.isLeft
}

//Position prints formatted window position represented by {x, y}
func (a *Application) Position() string {
	return fmt.Sprintf("{%d, %d}", a.location.x, a.location.y)
}

//Size prints formatted window size represented by {width, height}
func (a *Application) Size() string {
	return fmt.Sprintf("{%d, %d}", a.size.width, a.size.height)
}

//calculate calculates window size and position
func (a *Application) calculate(screen Screen) {
	height := windowHeight(screen.Height())
	width := leftWindowWidth(screen.Width())
	pointX, pointY := leftPoint(screen)

	if !a.IsLeft() {
		pointX, pointY = rightPoint(screen, width, pointY)
		width = rightWindowWidth(pointX, width)
	}

	a.size = windowSize{
		height: height,
		width:  width,
	}

	a.location = point{
		x: pointX,
		y: pointY,
	}
}

//Resizer represents osascript command to resize window
type Resizer struct {
	app *Application
}

//Do resizes current application based on its properties
func (r *Resizer) Do(app *Application) error {
	if err := app.Validate(); err != nil {
		return err
	}
	r.app = app

	return r.resize()
}

func (r *Resizer) resize() error {
	cmds := []string{
		"-e",
		`'tell application "System Events" to tell process "` + r.app.name + `"'`,
		"-e",
		`'tell window 1'`,
		"-e",
		`'set position to ` + r.app.Position() + "'",
		"-e",
		`'set size to ` + r.app.Size() + "'",
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
