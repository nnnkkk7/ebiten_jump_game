package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image

const (
	screenWidth  = 640
	screenHeight = 480
)

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("ebi.png")
	if err != nil {
		log.Fatal(err)
	}

}

type Game struct {
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
