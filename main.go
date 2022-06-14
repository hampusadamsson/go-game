package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hampusadamsson/go-game/game"
	"golang.org/x/image/font/gofont/gomedium"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 220
	screenHeight = 220
	padding      = screenWidth / 3 //25
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
	message      string
}

func (g *GameEbiten) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *GameEbiten) Update() error {
	// ebitenutil.DebugPrint(g, "Our first game in Ebiten!")
	return nil
}

// Main loop
func (g *GameEbiten) Draw(screen *ebiten.Image) {
	for i, l := range g.game.Board.Tiles {
		for j, tile := range l {
			op := &ebiten.DrawImageOptions{}
			// Where to draw it
			xPos := float64((j * tileSize) + padding)
			yPos := float64((i * tileSize)) + padding
			op.GeoM.Translate(xPos, yPos)
			// What to draw
			x, y, size, _ := tile.Img.GetImage()
			screen.DrawImage(tilesImage.SubImage(image.Rect(x, y, x+size, y+size)).(*ebiten.Image), op) // What part of a larger image to draw
			if u, _ := tile.GetUnit(); u != nil {
				x, y, size, animation := u.Img.GetImage()
				if animation {
					op.GeoM.Scale(1, 0.98)
				}
				if u.ExhaustedMove { // Change hue if exhausted
					op.ColorM.ChangeHSV(1, 1, 0.25)
				}
				screen.DrawImage(tilesImage.SubImage(image.Rect(x, y, x+size, y+size)).(*ebiten.Image), op)
				hp := fmt.Sprintf("%d", u.HP)
				fmt.Println(x, y)
				g.drawStats(screen, hp, int(xPos), int(yPos), color.Black)
			}
		}
	}
	g.drawSelection(screen)
	g.drawStats(screen, g.message, 12, 12, color.White)
	g.handleInput(screen)
}

// Draws the selection where the cursor is at
func (g *GameEbiten) drawStats(screen *ebiten.Image, msg string, x int, y int, c color.Color) {
	f, _ := opentype.Parse(gomedium.TTF)
	fontFace, _ := opentype.NewFace(f, &opentype.FaceOptions{
		Size: 8,
		DPI:  100,
	})
	text.Draw(screen, msg, fontFace, x, y, c)
}

// Draws the selection where the cursor is at
func (g *GameEbiten) drawSelection(screen *ebiten.Image) {
	x := g.x
	y := g.y
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64((x*tileSize)-5+padding), float64((y*tileSize)-5+padding))
	xImg, yImg, size, _ := game.Picker.GetImage()

	if g.hasSelection {
		if defender, err := g.game.Board.GetUnit(g.y, g.x); err == nil { // Make selection easier
			if attacker, err := g.game.Board.GetUnit(g.selection.X, g.selection.Y); err == nil { // Make selection easier
				if defender.Owner != attacker.Owner {
					xImg, yImg, size, _ = game.PickerAttack.GetImage()
				}
			}

		}
	}
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
			g.message = "End turn"
		}
	}
	// Action key
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if g.hasSelection == false {
				if _, err := g.game.Board.GetUnit(g.y, g.x); err == nil { // Make selection easier
					g.selection = game.Coord{g.y, g.x}
					g.hasSelection = true
					g.message = fmt.Sprintf("%d:%d", g.y, g.x)
				}
			} else {
				if _, err := g.game.Board.GetUnit(g.y, g.x); err == nil { // Make selection easier
					g.playerAction <- game.Action{ActionType: game.ActionAttack, From: g.selection, To: game.Coord{g.y, g.x}}
					g.message = "Attacking"
				} else {
					g.playerAction <- game.Action{ActionType: game.ActionMove, From: g.selection, To: game.Coord{g.y, g.x}}
					g.message = "Moving"
				}
				g.hasSelection = false
			}
		}
	}
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
