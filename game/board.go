package game

import "github.com/hajimehoshi/ebiten/v2"

type Board struct {
	Squares [8][8]*Piece
}

func NewBoard(pieceImages map[string]*ebiten.Image) *Board {
	board := &Board{}
	board.initialize(pieceImages)
	return board
}

func (b *Board) initialize(pieceImages map[string]*ebiten.Image) {
	// Place pawns
	for i := 0; i < 8; i++ {
		b.Squares[1][i] = NewPiece(Pawn, White, pieceImages["white_pawn"])
		b.Squares[6][i] = NewPiece(Pawn, Black, pieceImages["black_pawn"])
	}

	// Place rooks
	b.Squares[0][0] = NewPiece(Rook, White, pieceImages["white_rook"])
	b.Squares[0][7] = NewPiece(Rook, White, pieceImages["white_rook"])
	b.Squares[7][0] = NewPiece(Rook, Black, pieceImages["black_rook"])
	b.Squares[7][7] = NewPiece(Rook, Black, pieceImages["black_rook"])

	// Place knights
	b.Squares[0][1] = NewPiece(Knight, White, pieceImages["white_knight"])
	b.Squares[0][6] = NewPiece(Knight, White, pieceImages["white_knight"])
	b.Squares[7][1] = NewPiece(Knight, Black, pieceImages["black_knight"])
	b.Squares[7][6] = NewPiece(Knight, Black, pieceImages["black_knight"])

	// Place bishops
	b.Squares[0][2] = NewPiece(Bishop, White, pieceImages["white_bishop"])
	b.Squares[0][5] = NewPiece(Bishop, White, pieceImages["white_bishop"])
	b.Squares[7][2] = NewPiece(Bishop, Black, pieceImages["black_bishop"])
	b.Squares[7][5] = NewPiece(Bishop, Black, pieceImages["black_bishop"])

	// Place queens
	b.Squares[0][3] = NewPiece(Queen, White, pieceImages["white_queen"])
	b.Squares[7][3] = NewPiece(Queen, Black, pieceImages["black_queen"])

	// Place kings
	b.Squares[0][4] = NewPiece(King, White, pieceImages["white_king"])
	b.Squares[7][4] = NewPiece(King, Black, pieceImages["black_king"])
}
