package gap

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
)

const (
	horizontalGap, verticalGap = 10.0, 10.0
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
		x: horizontalGap * screenSize.Width(),
		y: verticalGap * screenSize.Height(),
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

	temp := int64(2 * horizontalGap)
	app.size = calculateWindow(screenSize)
	app.location = point{
		x: (screenSize.Width() * temp) + app.size.width,
		y: verticalGap * screenSize.Height(),
	}
	r.app = app

	return r.resize()
}

func (r *Resizer) printWindowFormat() string {
	format := fmt.Sprintf("{%d, %d, %d, %d}",
		r.app.location.x/100,
		r.app.location.y/100,
		r.app.size.width/100,
		r.app.size.height/100)
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

func calculatePoint(size ScreenSize) point {
	point := point{
		x: 0,
		y: 0,
	}

	return point
}

//calculate width and height of the window
func calculateWindow(size *ScreenSize) windowSize {
	window := windowSize{
		width:  windowHeight(size.Height()),
		height: windowWidth(size.Width()),
	}

	return window
}

func windowHeight(height int64) int64 {
	temp := int64(100 - (2 * horizontalGap))
	return height * temp
}

func windowWidth(width int64) int64 {
	temp := int64(100 - (3 * verticalGap))
	return (width * temp) / 2
}
