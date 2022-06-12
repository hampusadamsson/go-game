package main

// import (
// 	"fmt"
// 	"image"
// 	_ "image/png"
// 	"log"
// 	"time"

// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
// 	"github.com/hampusadamsson/go-game/game"
// )

// const (
// 	screenWidth  = 120
// 	screenHeight = 120
// )

// const (
// 	tileSize = 15
// )

// var (
// 	tilesImage *ebiten.Image
// )

// func init() {
// 	img, _, err := ebitenutil.NewImageFromFile("props/advancewars.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	tilesImage = ebiten.NewImageFromImage(img)
// }

// type Game struct {
// 	Board   game.Board
// 	message interface{}
// }

// func (g *Game) Update() error {
// 	return nil
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// 	for i, l := range g.Board.Tiles {
// 		for j, tile := range l {
// 			op := &ebiten.DrawImageOptions{}
// 			// Where to draw it
// 			op.GeoM.Translate(float64(j*tileSize), float64(i*tileSize))
// 			// What to draw
// 			x, y, size, _ := tile.Img.GetImage()
// 			screen.DrawImage(tilesImage.SubImage(image.Rect(x, y, x+size, y+size)).(*ebiten.Image), op) // What part of a larger image to draw
// 			if u, _ := tile.GetUnit(); u != nil {
// 				x, y, size, animation := u.Img.GetImage()
// 				if animation {
// 					op.GeoM.Scale(1, 0.98)
// 				}
// 				screen.DrawImage(tilesImage.SubImage(image.Rect(x, y, x+size, y+size)).(*ebiten.Image), op)
// 			}
// 		}
// 	}
// 	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%+v\n", g.message), 12, 12)
// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
// 	return screenWidth, screenHeight
// }

// func main() {
// 	bf := game.BoardFactory{}
// 	g := &Game{
// 		Board: bf.Example(),
// 	}

// 	ebiten.SetWindowSize(screenWidth*3, screenHeight*3)
// 	ebiten.SetWindowTitle("Go-Game")
// 	go g.warp()

// 	if err := ebiten.RunGame(g); err != nil {
// 		log.Fatal(err)
// 	}
// }

// // warp adds a 'bouncy' animation to units
// func (g *Game) warp() {
// 	var cur = g.Board.Tiles
// 	for {
// 		time.Sleep(time.Millisecond * 1000)
// 		for i := 0; i < len(cur); i++ {
// 			for j := 0; j < len(cur[i]); j++ {
// 				var tile = &cur[i][j]
// 				if unit, err := tile.GetUnit(); err == nil {
// 					unit.Img.ToggleAnimation()
// 				}
// 			}
// 		}
// 	}
// }
