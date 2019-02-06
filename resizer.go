package gap

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
)

const (
	horizontalGap, verticalGap = 0.1, 0.1
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

//Resizer represents osascript command to resize window
type Resizer struct {
	app *Application
}

//Left resizes current Application to left display
func (r *Resizer) Left(app *Application) error {
	if err := app.Validate(); err != nil {
		return err
	}

	screenSize, err := GetScreenSize()
	if err != nil {
		return err
	}

	app.size = calculateWindow(screenSize)
	app.location = point{
		x: int64((float64(horizontalGap) / 100) * float64(screenSize.Width())),
		y: int64((float64(verticalGap) / 100) * float64(screenSize.Height())),
	}
	r.app = app

	return r.resize()
}

//Right resizes current Application to right display
func (r *Resizer) Right(app *Application) error {
	if err := app.Validate(); err != nil {
		return err
	}

	screenSize, err := GetScreenSize()
	if err != nil {
		return err
	}

	temp := float64(2*horizontalGap) / 100
	app.size = calculateWindow(screenSize)
	app.location = point{
		x: int64((float64(screenSize.Width()) * temp) + float64(app.size.width)),
		y: int64((float64(verticalGap) / 100) * float64(screenSize.Height())),
	}
	r.app = app

	return r.resize()
}

func (r *Resizer) printWindowFormat() string {
	format := fmt.Sprintf("{%d, %d, %d, %d}",
		r.app.location.x,
		r.app.location.y,
		r.app.size.width,
		r.app.size.height,
	)
	return format
}

func (r *Resizer) resize() error {
	cmds := []string{
		"-e",
		`'tell application "` + r.app.name + `"'`,
		"-e",
		`'set bounds of front window to ` + r.printWindowFormat() + `'`,
		"-e",
		"'end tell'",
	}
	cmd, err := exec.Command("osascript", cmds...).Output()
	log.Println(string(cmd))
	if err != nil {
		return err
	}
	return nil

}

//calculate width and height of the window
func calculateWindow(size *ScreenSize) windowSize {
	window := windowSize{
		width:  windowHeight(float64(size.Height())),
		height: windowWidth(float64(size.Width())),
	}

	return window
}

func windowHeight(height float64) int64 {
	temp := float64(100 - (2 * horizontalGap))
	return int64(height * (temp / 100))
}

func windowWidth(width float64) int64 {
	temp := float64(100 - (3 * verticalGap))
	return int64((width * (temp / 100)) / 2)
}
