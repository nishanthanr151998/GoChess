package game

type Color string

const (
	White Color = "white"
	Black Color = "black"
)

type PieceType string

const (
	Pawn   PieceType = "pawn"
	Rook   PieceType = "rook"
	Knight PieceType = "knight"
	Bishop PieceType = "bishop"
	Queen  PieceType = "queen"
	King   PieceType = "king"
)

type Piece struct {
	Type  PieceType
	Color Color
}

func NewPiece(pieceType PieceType, color Color) *Piece {
	return &Piece{
		Type:  pieceType,
		Color: color,
	}
}
