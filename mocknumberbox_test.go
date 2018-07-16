package bingo

import (
	"testing"
)

func Test_Game_With_Mock_NumberBox_Should_Be_4_11_17_21_30_32(t *testing.T) {
	mockNumbers := []int{4, 11, 17, 21, 30, 32}
	numberBox := MockNumberBox{}
	numberBox.Set(mockNumbers)
	game := Game{numberBox: &numberBox}
	for index, expectedValue := range []int{4, 11, 17, 21, 30, 32} {
		actualValue := game.PickNumber()
		if expectedValue != actualValue {
			t.Errorf("Round: %d Expected %d but it got %d", index, expectedValue, actualValue)
		}
	}
}
