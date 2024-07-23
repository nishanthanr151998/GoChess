package main

import (
	"GoChess/game"
	"image"
	"image/color"
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

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
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
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	board := game.NewBoard()
	g := &Game{board: board}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Chess Application")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
