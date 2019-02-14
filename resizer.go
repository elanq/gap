package gap

import (
	"bytes"
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

//Resize resizes current application based on its properties
func Resize(app *Application) error {
	if err := app.Validate(); err != nil {
		return err
	}

	return resize(app)
}

func script(app *Application) string {
	template := `
		tell application "System Events" to tell process "%s"
			tell window 1
				set position to %s
				set size to %s
			end tell
		end tell
	`
	return fmt.Sprintf(template, app.name, app.Position(), app.Size())
}

func resize(app *Application) error {
	s := script(app)
	cmd := exec.Command("osascript")
	cmd.Stdin = bytes.NewBuffer([]byte(s))
	out, err := cmd.CombinedOutput()

	log.Println(string(out))
	if err != nil {
		return err
	}
	return nil
}
