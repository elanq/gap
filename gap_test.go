package gap_test

import (
	"testing"

	"github.com/elanq/gap"
)

type MockScreen struct {
	height float64
	width  float64
}

func NewMockScreen(h float64, w float64) gap.Screen {
	return &MockScreen{height: h, width: w}
}

func (m *MockScreen) Height() float64 {
	return m.height
}

func (m *MockScreen) Width() float64 {
	return m.width
}

var cases = []struct {
	screen                gap.Screen
	expectedLeftSize      string
	expectedRightSize     string
	expectedPositionLeft  string
	expectedPositionRight string
}{
	{NewMockScreen(3840.0, 2160.0), "{1862, 2116}", "{3801, 2116}", "{38, 64}", "{1939, 64}"},
	{NewMockScreen(2880.0, 1880.0), "{1396, 1764}", "{2851, 1764}", "{28, 54}", "{1454, 54}"},
	{NewMockScreen(1920.0, 1200.0), "{931, 1176}", "{1900, 1176}", "{19, 36}", "{969, 36"},
	{NewMockScreen(1076.0, 768.0), "{496, 752}", "{1013, 752}", "{10, 23}", "{517, 23}"},
}

func TestCalculation(t *testing.T) {
	for _, c := range cases {
		left := leftApp(c.screen)
		if left.IsLeft() == false {
			t.Error("App should be on the left")
		}

		if left.Position() != c.expectedPositionLeft {
			t.Errorf("Invalid position. expected %s got %s", c.expectedPositionLeft, left.Position())
		}

		if left.Size() != c.expectedLeftSize {
			t.Errorf("Invalid size. expected %s got %s", c.expectedLeftSize, left.Size())
		}

		right := rightApp(c.screen)
		if right.IsLeft() == true {
			t.Error("App should be on the right")
		}

		if right.Position() != c.expectedPositionRight {
			t.Errorf("Invalid position. expected %s got %s", c.expectedPositionRight, right.Position())
		}

		if right.Size() != c.expectedRightSize {
			t.Errorf("Invalid size. expected %s got %s", c.expectedRightSize, right.Size())
		}
	}
}

func leftApp(screen gap.Screen) *gap.Application {
	return gap.NewApplication("left-app").Left(screen)
}

func rightApp(screen gap.Screen) *gap.Application {
	return gap.NewApplication("right-app").Right(screen)
}
