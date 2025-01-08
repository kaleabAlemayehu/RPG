package main

import (
	"fmt"
	"game/entities"
	"game/tiles"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player      *entities.Player
	enemies     []*entities.Enemy
	posions     []*entities.Posion
	tileMapJSON *tiles.TilesetMapJSON
	tileMapImg  *ebiten.Image
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
	// INFO: Drawing A Map
	for _, layer := range g.tileMapJSON.Layers {
		for index, id := range layer.Data {
			x := index % layer.Width
			y := index / layer.Width
			x *= 16
			y *= 16

			srcX := (id - 1) % 21
			srcY := (id - 1) / 21
			srcX *= 16
			srcY *= 16

			opts.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(g.tileMapImg.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image), &opts)
			opts.GeoM.Reset()
		}
	}
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
	return ebiten.WindowSize()
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
	tileMapJSON, err := tiles.NewTileMapJSON("./assets/spawn.json")
	if err != nil {
		log.Fatal(err)
	}
	tileMapImg, _, err := ebitenutil.NewImageFromFile("./assets/TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.RunGame(&Game{
		player: &entities.Player{
			Sprite: &entities.Sprite{Image: PlayerImage, X: 100.0, Y: 200.0},
			Health: 100.0,
		},
		tileMapJSON: tileMapJSON,
		tileMapImg:  tileMapImg,
		enemies: []*entities.Enemy{
			{
				Sprite: &entities.Sprite{
					Image: EnemyImage,
					X:     200.0,
					Y:     300.0,
				},
				FollowPlayer: true,
			},
			{
				Sprite: &entities.Sprite{
					Image: EnemyImage,
					X:     100.0,
					Y:     200.0,
				},
				FollowPlayer: false,
			},
			{
				Sprite: &entities.Sprite{
					Image: EnemyImage,
					X:     400.0,
					Y:     600.0,
				},
				FollowPlayer: true,
			},
		},

		posions: []*entities.Posion{
			{
				Sprite: &entities.Sprite{
					Image: PosionImage,
					X:     20.0,
					Y:     30.0,
				},
				AmountOfHeal: 100,
			},
			{
				Sprite: &entities.Sprite{
					Image: PosionImage,
					X:     190.0,
					Y:     280.0,
				},
				AmountOfHeal: 100,
			},
			{
				Sprite: &entities.Sprite{
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
