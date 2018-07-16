package bingo

import "math/rand"

type NumberBox interface {
	Draw() int
}

type RandomNumberBox struct {
	numbers       []int
	spinnedNumber map[int]bool
}

func (rn *RandomNumberBox) Draw() int {
	var value int
	value, rn.numbers = rn.numbers[0], rn.numbers[1:]
	return value
}

func (rn *RandomNumberBox) Shuffle() {
	for index := len(rn.numbers) - 1; index > 0; index-- {
		swapIndex := rand.Intn(index + 1)
		rn.numbers[index], rn.numbers[swapIndex] = rn.numbers[swapIndex], rn.numbers[index]
	}
}

type MockNumberBox struct {
	numbers []int
}

func (mn *MockNumberBox) Set(numbers []int) {
	mn.numbers = numbers
}

func (mn *MockNumberBox) Draw() int {
	var value int
	value, mn.numbers = mn.numbers[0], mn.numbers[1:]
	return value
}

type Game struct {
	numberBox NumberBox
}

func (g Game) PickNumber() int {
	return g.numberBox.Draw()
}
