package gap

import (
	"os/exec"
	"strconv"
	"strings"
)

const cmdGetScreen = `
system_profiler SPDisplaysDataType | grep Resolution | awk '/Resolution/{print $2}{print $4}'
`

//Screen is interface for screen size
type Screen interface {
	Height() float64
	Width() float64
}

//ScreenSize represents current display screen size
type ScreenSize struct {
	height int64
	width  int64
}

//GetScreenSize return screen size of current display
func GetScreenSize() (*ScreenSize, error) {
	out, err := execCommand()
	if err != nil {
		return nil, err
	}

	return parseOutput(out)

}

//Height returns screen height
func (s *ScreenSize) Height() float64 {
	return float64(s.height)
}

//Width returns screen width
func (s *ScreenSize) Width() float64 {
	return float64(s.width)
}

//Separator is line to separate between left and right window
func (s *ScreenSize) Separator() int64 {
	return s.width / 2
}

func execCommand() ([]byte, error) {
	cmd := exec.Command("bash", "-c", cmdGetScreen)
	out, err := cmd.Output()
	return out, err
}

func parseOutput(out []byte) (*ScreenSize, error) {
	var screenSize *ScreenSize
	var err error

	strForm := string(out)
	sizes := strings.Split(strForm, "\n")
	width, err := strconv.ParseInt(sizes[0], 10, 64)
	if err != nil {
		return screenSize, err
	}
	height, err := strconv.ParseInt(sizes[1], 10, 64)
	if err != nil {
		return screenSize, err
	}

	screenSize = &ScreenSize{
		width:  width,
		height: height,
	}

	return screenSize, err
}
