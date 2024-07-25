package main

import (
	"GoChess/game"
	"image"
	"image/color"
	_ "image/png" // Needed to decode PNG images
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 640
	gridSize     = 8
	cellSize     = screenWidth / gridSize
)

type Game struct {
	board *game.Board
}

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the chessboard
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			rect := image.Rect(x*cellSize, y*cellSize, (x+1)*cellSize, (y+1)*cellSize)
			col := color.RGBA{0xff, 0xff, 0xff, 0xff} // White color
			if (x+y)%2 == 1 {
				col = color.RGBA{0x00, 0x00, 0x00, 0xff} // Black color
			}
			ebitenutil.DrawRect(screen, float64(rect.Min.X), float64(rect.Min.Y), float64(cellSize), float64(cellSize), col)
		}
	}

	// Draw the pieces
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			piece := g.board.Squares[y][x]
			if piece != nil {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(float64(cellSize)/float64(piece.Image.Bounds().Dx()), float64(cellSize)/float64(piece.Image.Bounds().Dy()))
				op.GeoM.Translate(float64(x*cellSize), float64(y*cellSize))
				screen.DrawImage(piece.Image, op)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	// Load piece images
	pieceImages := map[string]*ebiten.Image{
		"white_pawn":   loadImage("assets/wp.png"),
		"black_pawn":   loadImage("assets/bp.png"),
		"white_rook":   loadImage("assets/wr.png"),
		"black_rook":   loadImage("assets/br.png"),
		"white_knight": loadImage("assets/wn.png"),
		"black_knight": loadImage("assets/bn.png"),
		"white_bishop": loadImage("assets/wb.png"),
		"black_bishop": loadImage("assets/bb.png"),
		"white_queen":  loadImage("assets/wq.png"),
		"black_queen":  loadImage("assets/bq.png"),
		"white_king":   loadImage("assets/wk.png"),
		"black_king":   loadImage("assets/bk.png"),
	}

	board := game.NewBoard(pieceImages)
	g := &Game{board: board}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Chess Application")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
