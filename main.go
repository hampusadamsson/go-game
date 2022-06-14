package main

import (
	"image"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hampusadamsson/go-game/game"
)

const (
	screenWidth  = 120
	screenHeight = 120
	padding      = 25
)

const (
	tileSize = 15
)

var (
	tilesImage *ebiten.Image
)

func init() {
	img, _, err := ebitenutil.NewImageFromFile("props/advancewars.png")
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

type GameEbiten struct {
	game         *game.Game
	x            int
	y            int
	hasSelection bool
	selection    game.Coord
	playerAction chan game.Action
}

func (g *GameEbiten) Update() error {
	// ebitenutil.DebugPrint(g, "Our first game in Ebiten!")
	return nil
}

func (g *GameEbiten) Draw(screen *ebiten.Image) {
	for i, l := range g.game.Board.Tiles {
		for j, tile := range l {
			op := &ebiten.DrawImageOptions{}
			// Where to draw it
			op.GeoM.Translate(float64((j*tileSize)+padding), float64((i*tileSize))+padding)
			// What to draw
			x, y, size, _ := tile.Img.GetImage()
			screen.DrawImage(tilesImage.SubImage(image.Rect(x, y, x+size, y+size)).(*ebiten.Image), op) // What part of a larger image to draw
			if u, _ := tile.GetUnit(); u != nil {
				x, y, size, animation := u.Img.GetImage()
				if animation {
					op.GeoM.Scale(1, 0.98)
				}
				if u.ExhaustedMove { // Change hue when exhausted?
					op.ColorM.ChangeHSV(1, 1, 0.2)
				}
				screen.DrawImage(tilesImage.SubImage(image.Rect(x, y, x+size, y+size)).(*ebiten.Image), op)
			}
		}
	}
	g.drawSelection(screen)
	g.handleInput(screen)
}

// Draws the selection where the cursor is at
func (g *GameEbiten) drawSelection(screen *ebiten.Image) {
	x := g.x
	y := g.y
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64((x*tileSize)-5+padding), float64((y*tileSize)-5+padding))
	xImg, yImg, size, _ := game.Picker.GetImage()
	screen.DrawImage(tilesImage.SubImage(image.Rect(xImg, yImg, xImg+size, yImg+size)).(*ebiten.Image), op)
}

func (g *GameEbiten) handleInput(screen *ebiten.Image) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.y >= 1 {
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			g.y--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && g.y <= 1 {
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			g.y++
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.x >= 1 {
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			g.x--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && g.x <= 1 {
		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			g.x++
		}
	}
	// End turn
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		if inpututil.IsKeyJustPressed(ebiten.KeyE) {
			g.hasSelection = false
			g.playerAction <- game.Action{ActionType: game.ActionEnd}
		}
	}
	// Action key
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if g.hasSelection == false {
				g.selection = game.Coord{g.y, g.x}
				g.hasSelection = true
			} else {
				g.playerAction <- game.Action{ActionType: game.ActionAttack, From: g.selection, To: game.Coord{g.y, g.x}}
				g.playerAction <- game.Action{ActionType: game.ActionMove, From: g.selection, To: game.Coord{g.y, g.x}}
				g.hasSelection = false
			}
		}
	}
}

func (g *GameEbiten) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	gf := game.GameFactory{}

	p1Action := make(chan game.Action)
	p1 := &game.Player{"A", p1Action}

	p2Action := make(chan game.Action)
	p2 := &game.Player{"B", p2Action}
	go endTurnAi(p2Action)

	g := gf.OneVsOne(p1, p2)
	g.Run()

	ge := &GameEbiten{
		game:         g,
		playerAction: p1Action,
	}

	ebiten.SetWindowSize(screenWidth*3, screenHeight*3)
	ebiten.SetWindowTitle("Go-Game")
	//go g.warp()

	if err := ebiten.RunGame(ge); err != nil {
		log.Fatal(err)
	}
}

// warp adds a 'bouncy' animation to units
func (g *GameEbiten) warp() {
	var cur = g.game.Board.Tiles
	for {
		time.Sleep(time.Millisecond * 1000)
		for i := 0; i < len(cur); i++ {
			for j := 0; j < len(cur[i]); j++ {
				var tile = &cur[i][j]
				if unit, err := tile.GetUnit(); err == nil {
					unit.Img.ToggleAnimation()
				}
			}
		}
	}
}

func endTurnAi(ac chan game.Action) {
	for {
		time.Sleep(time.Second * 1)
		ac <- game.Action{ActionType: game.ActionEnd}
	}
}
