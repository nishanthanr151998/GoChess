package game

type Board struct {
	Squares [8][8]*Piece
}

func NewBoard() *Board {
	board := &Board{}
	board.initialize()
	return board
}

func (b *Board) initialize() {
	// Initialize the board with pieces in their starting positions
	// For now, leave it blank
}
