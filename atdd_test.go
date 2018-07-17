package bingo

import "testing"

func MockTicket(mockNumbers []int, ticket Ticket) Ticket {
	for rowIndex := range ticket.Grid {
		for colIndex := range ticket.Grid[rowIndex] {
			value, mockNumbers := mockNumbers[0], mockNumbers[1:]
			ticket.Grid[rowIndex][colIndex] = value
		}
	}
	return ticket
}
func Test_ATDD_Vertical_Bingo(t *testing.T) {
	// สร้าง ticket
	ticketA := NewTicket()
	ticketB := NewTicket()
	// ถ้าตอนสร้างมันเป็น สุ่มเลขจริงๆ ต้องทำ mock ตรงนี้ก่อน
	mockNumbersTicketA := []int{1, 17, 33, 51, 74, 9, 21, 41, 38, 79, 2, 23, 0, 45, 68, 14, 29, 32, 49, 66, 11, 30, 39, 36, 70}
	mockNumbersTicketB := []int{3, 21, 39, 53, 55, 12, 29, 32, 34, 67, 11, 30, 0, 49, 70, 9, 16, 41, 45, 68, 7, 19, 44, 52, 72}
	ticketA = MockTicket(mockNumbersTicketA, ticketA)
	ticketB = MockTicket(mockNumbersTicketB, ticketB)
	// สร้าง player
	playerA := NewPlayer("A", ticketA)
	playerB := NewPlayer("A", ticketB)
	// สร้าง numberbox
	numberBox := NewNumberBox(75)
	// ถ้าสร้างเป็น array เก็บเลข และอยากให้ตอน pickup หยิบแบบเรียงต้องจำลองค่าให้มันเป็นตาม test case
	mockNumberBox := []int{9, 51, 47, 29, 56, 49, 39, 58}
	numberBox = mockNumberBox
	// สร้าง game (ถ้ามองว่าแค่ pickup เฉยๆ อาจจะไม่ต้องมีก็ได้)
	game := NewGame([]Player{playerA, playerB}, numberBox)
	// ถ้าไม่มี game จะต้องเปลี่ยนเป็นของ numberBox ไม่ก็ function
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
