package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	PlayerImage *ebiten.Image
	X, Y        float64
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyK) {
		g.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyL) {
		g.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyH) {
		g.X -= 2
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.X, g.Y)
	screen.DrawImage(g.PlayerImage.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), &opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	eImg, _, err := ebitenutil.NewImageFromFile("./assets/ninjaDark.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(&Game{PlayerImage: eImg, X: 100.0, Y: 200.0}); err != nil {
		log.Fatal(err)
	}
}
