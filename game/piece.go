package game

import "github.com/hajimehoshi/ebiten/v2"

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
	Image *ebiten.Image
}

func NewPiece(pieceType PieceType, color Color, image *ebiten.Image) *Piece {
	return &Piece{
		Type:  pieceType,
		Color: color,
		Image: image,
	}
}
