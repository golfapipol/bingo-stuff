package bingo

import "testing"

func Test_ATDD_Vertical_Bingo(t *testing.T) {
	ticketA := NewTicket()
	ticketB := NewTicket()

	playerA := NewPlayer("A", ticketA)
	playerB := NewPlayer("A", ticketB)

	numberBox := NewNumberBox(75)

	game := NewGame([]Player{playerA, playerB}, numberBox)

	number := game.PickUpNumber()
	//Round 1
	positionOfTicketA := playerA.CheckNumber(number)
	playerA.MarkGrid(positionOfTicketA)
	positionOfTicketB := playerB.CheckNumber(number)
	playerB.MarkGrid(positionOfTicketB)
	//Round 2
	positionOfTicketA = playerA.CheckNumber(number)
	playerA.MarkGrid(positionOfTicketA)
	playerB.CheckNumber(number)
	//Round 3
	positionOfTicketA = playerA.CheckNumber(number)
	playerA.MarkGrid(positionOfTicketA)
	playerB.MarkGrid(positionOfTicketB)
	//Round 4
	positionOfTicketA = playerA.CheckNumber(number)
	playerA.MarkGrid(positionOfTicketA)
	positionOfTicketB = playerB.CheckNumber(number)
	playerB.MarkGrid(positionOfTicketB)
	//Round 5
	positionOfTicketA = playerA.CheckNumber(number)
	playerA.MarkGrid(positionOfTicketA)
	positionOfTicketB = playerB.CheckNumber(number)
	playerB.MarkGrid(positionOfTicketB)
	//Round 6
	positionOfTicketA = playerA.CheckNumber(number)
	playerA.MarkGrid(positionOfTicketA)
	positionOfTicketB = playerB.CheckNumber(number)
	playerB.MarkGrid(positionOfTicketB)
	//Round 7
	positionOfTicketA = playerA.CheckNumber(number)
	playerA.MarkGrid(positionOfTicketA)
	positionOfTicketB = playerB.CheckNumber(number)
	playerB.MarkGrid(positionOfTicketB)
	//Round 8
	positionOfTicketA = playerA.CheckNumber(number)
	playerA.MarkGrid(positionOfTicketA)
	playerA.CheckBingo(positionOfTicketA)

}
