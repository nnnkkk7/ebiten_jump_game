package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var img *ebiten.Image

const (
	screenWidth  = 640
	screenHeight = 480
	baseX        = 100
	basey        = 100
	jumpingPower = 15
	gravity      = 1
)

const (
	modeStart = iota
	modeGame
	modeGameover
)

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("ebi.png")
	if err != nil {
		log.Fatal(err)
	}

}

type Game struct {
	mode  int
	count int
	score int
	ebix  int
	ebiy  int
	gy    int
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}
func (g *Game) init() {
	g.count = 0
	g.score = 0
	g.ebix = baseX
	g.ebiy = basey
	g.gy = 0
}

func (g *Game) Update() error {
	switch g.mode {
	case modeStart:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.mode = modeGame
		}
	case modeGame:
		g.count++
		g.score = g.count / 5
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.gy -= jumpingPower
			g.ebiy += g.gy
			g.gy += gravity
		}

	case modeGameover:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.init()
			g.mode = modeGame
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(baseX, float64(g.ebiy))
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
