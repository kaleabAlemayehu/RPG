package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Image *ebiten.Image
	X, Y  float64
}

type Player struct {
	*Sprite
	Health float64
}

type Enemy struct {
	*Sprite
	FollowPlayer bool
}

type Posion struct {
	*Sprite
	AmountOfHeal float64
}

type Game struct {
	player      *Player
	enemies     []*Enemy
	posions     []*Posion
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyK) {
		g.player.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.player.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyL) {
		g.player.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyH) {
		g.player.X -= 2
	}

	for _, enemy := range g.enemies {
		if enemy.FollowPlayer {
			if enemy.X < g.player.X {
				enemy.X += 1
			}
			if enemy.X > g.player.X {
				enemy.X -= 1
			}
			if enemy.Y < g.player.Y {
				enemy.Y += 1
			}
			if enemy.Y > g.player.Y {
				enemy.Y -= 1
			}
		}
	}

	for _, posion := range g.posions {
		if ((g.player.X >= posion.X-4.5) && (g.player.X <= (posion.X + 4.5))) && ((g.player.Y >= posion.Y-5.5) && (g.player.Y <= (posion.Y + 5.5))) {
			g.player.Health += posion.AmountOfHeal
			fmt.Printf("curent Health: %v\n", g.player.Health)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y)
	screen.DrawImage(g.player.Image.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), &opts)
	opts.GeoM.Reset()
	// INFO: Drawing enemies
	for _, enemy := range g.enemies {
		opts.GeoM.Translate(enemy.X, enemy.Y)
		screen.DrawImage(enemy.Image.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), &opts)
		opts.GeoM.Reset()
	}
	// INFO: Drawing position
	for _, posion := range g.posions {
		opts.GeoM.Translate(posion.X, posion.Y)
		screen.DrawImage(posion.Image.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), &opts)
		opts.GeoM.Reset()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 350, 360
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	PlayerImage, _, err := ebitenutil.NewImageFromFile("./assets/ninjaDark.png")
	if err != nil {
		log.Fatal(err)
	}
	EnemyImage, _, err := ebitenutil.NewImageFromFile("./assets/skeleton.png")
	if err != nil {
		log.Fatal(err)
	}

	PosionImage, _, err := ebitenutil.NewImageFromFile("./assets/LifePot.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(&Game{
		player: &Player{
			Sprite: &Sprite{Image: PlayerImage, X: 100.0, Y: 200.0},
			Health: 100.0,
		},
		enemies: []*Enemy{
			{
				Sprite: &Sprite{
					Image: EnemyImage,
					X:     200.0,
					Y:     300.0,
				},
				FollowPlayer: true,
			},
			{
				Sprite: &Sprite{
					Image: EnemyImage,
					X:     100.0,
					Y:     200.0,
				},
				FollowPlayer: false,
			},
			{
				Sprite: &Sprite{
					Image: EnemyImage,
					X:     400.0,
					Y:     600.0,
				},
				FollowPlayer: true,
			},
		},

		posions: []*Posion{
			{
				Sprite: &Sprite{
					Image: PosionImage,
					X:     20.0,
					Y:     30.0,
				},
				AmountOfHeal: 100,
			},
			{
				Sprite: &Sprite{
					Image: PosionImage,
					X:     190.0,
					Y:     280.0,
				},
				AmountOfHeal: 100,
			},
			{
				Sprite: &Sprite{
					Image: PosionImage,
					X:     400.0,
					Y:     600.0,
				},
				AmountOfHeal: 100,
			},
		},
	}); err != nil {
		log.Fatal(err)
	}
}
