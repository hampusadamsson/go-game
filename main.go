package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hampusadamsson/go-game/game"
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
	game            *game.Game
	x               int
	y               int
	hasSelection    bool
	selection       game.Coord
	playerAction    chan game.Action
	cursor          *game.ImageMeta
	highlightAttack map[game.Coord]bool
	highlightPaths  map[game.Coord]int
	message         string
}

func (g *GameEbiten) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *GameEbiten) Update() error {
	// ebitenutil.DebugPrint(g, "Our first game in Ebiten!")
	return nil
}

func (g *GameEbiten) drawTerrain(screen *ebiten.Image) {

	for i, l := range g.game.Board.Tiles {
		for j, tile := range l {
			op := &ebiten.DrawImageOptions{}
			// Where to draw it
			xPos := float64((j * tileSize) + padding)
			yPos := float64((i * tileSize)) + padding
			op.GeoM.Translate(xPos, yPos)
			// What to draw
			screen.DrawImage(tilesImage.SubImage(image.Rect(tile.Img.X, tile.Img.Y, tile.Img.X+tile.Img.Size, tile.Img.Y+tile.Img.Size)).(*ebiten.Image), op) // What part of a larger image to draw

			// Draw movement
			if g.highlightPaths != nil {
				if _, ok := g.highlightPaths[game.Coord{i, j}]; ok {
					screen.DrawImage(tilesImage.SubImage(image.Rect(game.GreenZone.X, game.GreenZone.Y, game.GreenZone.X+game.GreenZone.Size, game.GreenZone.Y+game.GreenZone.Size)).(*ebiten.Image), op)
				}
			}

			// Draw attacks
			if g.highlightAttack != nil {
				if _, ok := g.highlightAttack[game.Coord{i, j}]; ok {
					screen.DrawImage(tilesImage.SubImage(image.Rect(game.RedZone.X, game.RedZone.Y, game.RedZone.X+game.RedZone.Size, game.RedZone.Y+game.RedZone.Size)).(*ebiten.Image), op)
				}
			}

			// Draw units
			if u, _ := tile.GetUnit(); u != nil {

				// Animate all unit
				if u.Img.Animation {
					op.GeoM.Scale(1, 0.99)
				}

				// Change hue if exhausted
				if u.ExhaustedMove || u.ExhaustedAttack {
					op.ColorM.ChangeHSV(1, 1, 0.25)
				}

				// Different colors for team
				var team int
				if u.Owner == g.game.Players[0] {
					team = 0
				} else {
					team = 18
				}
				screen.DrawImage(tilesImage.SubImage(image.Rect(u.Img.X, u.Img.Y+team, u.Img.X+u.Img.Size, u.Img.Y+u.Img.Size+team)).(*ebiten.Image), op)

				// Draw HP
				// if u.HP > 9 {
				// 	op := &ebiten.DrawImageOptions{}
				// 	xPos := float64((j * tileSize) + padding)
				// 	yPos := float64((i * tileSize)) + padding
				// 	op.GeoM.Translate(xPos, yPos)
				// 	imgNr := game.NumberImage(u.HP / 10)
				// 	screen.DrawImage(tilesImage.SubImage(image.Rect(imgNr.X, imgNr.Y, imgNr.X+imgNr.Size, imgNr.Y+imgNr.Size*2)).(*ebiten.Image), op)
				// 	imgNr = game.NumberImage(u.HP % 10)
				// 	op = &ebiten.DrawImageOptions{}
				// 	xPos = float64((j * tileSize) + padding)
				// 	yPos = float64((i * tileSize)) + padding
				// 	op.GeoM.Translate(xPos+9, yPos)
				// 	screen.DrawImage(tilesImage.SubImage(image.Rect(imgNr.X, imgNr.Y, imgNr.X+imgNr.Size, imgNr.Y+imgNr.Size*2)).(*ebiten.Image), op)
				// } else {
				imgNr := game.NumberImage(u.HP)
				screen.DrawImage(tilesImage.SubImage(image.Rect(imgNr.X, imgNr.Y, imgNr.X+imgNr.Size, imgNr.Y+imgNr.Size*2)).(*ebiten.Image), op)
				// }
			}
		}
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Main loop
func (g *GameEbiten) Draw(screen *ebiten.Image) {
	g.handleInput()

	//PrintMemUsage()

	clearing := fmt.Sprintf("FPS: %f\nTPS: %f\ngameOver: %t", ebiten.CurrentFPS(), ebiten.CurrentTPS(), g.game.GameOver)
	ebitenutil.DebugPrint(screen, clearing)

	g.drawTerrain(screen)
	g.drawSelection(screen)
	// g.drawStats(screen, g.message, 12, 12, color.White)
}

// Draws the selection where the cursor is at
// func (g *GameEbiten) drawStats(screen *ebiten.Image, msg string, x int, y int, c color.Color) {
// 	f, _ := opentype.Parse(gomedium.TTF)
// 	fontFace, _ := opentype.NewFace(f, &opentype.FaceOptions{
// 		Size: 8,
// 		DPI:  100,
// 	})
// 	text.Draw(screen, msg, fontFace, x, y, c)
// }

// Draws the selection where the cursor is at
func (g *GameEbiten) drawSelection(screen *ebiten.Image) {
	x := g.x
	y := g.y
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64((x*tileSize)-5+padding), float64((y*tileSize)-5+padding))
	screen.DrawImage(tilesImage.SubImage(image.Rect(g.cursor.X, g.cursor.Y, g.cursor.X+g.cursor.Size, g.cursor.Y+g.cursor.Size)).(*ebiten.Image), op)
}

// The cursor is picker, or attacker based on status
func (g *GameEbiten) updateCursorImage() {
	if g.hasSelection {
		if defender, err := g.game.Board.GetUnit(g.y, g.x); err == nil { // Make selection easier
			if attacker, err := g.game.Board.GetUnit(g.selection.X, g.selection.Y); err == nil { // Make selection easier
				if defender.Owner != attacker.Owner {
					g.cursor = game.PickerAttack
					return
				}
			}
		}
	}
	g.cursor = game.Picker
}

// Handle basic game input
// left, right, up, down
// q - quit game
// e - end turn
// space - execute action
// esc - cancel
func (g *GameEbiten) handleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) && g.y >= 1 {
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			g.y--
			g.updateCursorImage()
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && g.y < len(g.game.Board.Tiles[0])-2 {
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			g.y++
			g.updateCursorImage()
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && g.x >= 1 {
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			g.x--
			g.updateCursorImage()
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && g.x < len(g.game.Board.Tiles) {
		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			g.x++
			g.updateCursorImage()
		}
	}
	// Cancel
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.hasSelection = false
			g.message = "Cancel"
			g.updateCursorImage()
			g.highlightAttack = nil
			g.highlightPaths = nil
		}
	}
	// Exit game
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			os.Exit(0)
		}
	}
	// End turn
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		if inpututil.IsKeyJustPressed(ebiten.KeyE) {
			g.hasSelection = false
			g.playerAction <- game.Action{ActionType: game.ActionEnd}
			g.message = "End turn"
			g.updateCursorImage()
			g.highlightAttack = nil
			g.highlightPaths = nil
		}
	}
	// Action key
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if g.hasSelection == false { // 1st selection
				if u, err := g.game.Board.GetUnit(g.y, g.x); err == nil {
					g.selection = game.Coord{g.y, g.x}
					if u.ExhaustedAttack == false {
						g.highlightAttack = u.GetAllAttackCoords()
						u, _ := g.game.Board.GetUnit(g.selection.X, g.selection.Y)
						if u.ExhaustedMove == false {
							path := g.game.Board.GetAllPaths(u)
							g.highlightPaths = path
						}
					}
					g.hasSelection = true
					g.message = fmt.Sprintf("%d:%d", g.y, g.x)
					// Update path to

				}
			} else { // 2d selection
				if _, err := g.game.Board.GetUnit(g.y, g.x); err == nil {
					g.playerAction <- game.Action{ActionType: game.ActionAttack, From: g.selection, To: game.Coord{g.y, g.x}}
					g.message = "Attacking"
				} else {
					g.playerAction <- game.Action{ActionType: game.ActionMove, From: g.selection, To: game.Coord{g.y, g.x}}
					g.message = "Moving"
				}
				g.hasSelection = false
				g.highlightAttack = nil
				g.highlightPaths = nil
				g.updateCursorImage()
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

	//g := gf.OneVsOne(p1, p2)
	g := gf.OneVsOneFirstGame(p1, p2)

	//go endTurnAi(p2Action)
	// go DumbAi(g, p1, p1Action)
	go DumbAi(g, p2, p2Action)

	g.Run()

	ge := &GameEbiten{
		game:         g,
		playerAction: p1Action,
		cursor:       game.Picker,
	}

	ebiten.SetWindowSize(screenWidth*3, screenHeight*3)
	ebiten.SetWindowTitle("Go-Game")

	go ge.warp()

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

// End turn AI
func endTurnAi(ac chan game.Action) {
	for {
		time.Sleep(time.Millisecond * 250)
		ac <- game.Action{ActionType: game.ActionEnd}
	}
}

// Very stupid AI
func DumbAi(g *game.Game, p *game.Player, ac chan game.Action) {
	for {
		time.Sleep(time.Millisecond * 10)
		if g.Turn == p { // Naive 2 player AI
			for _, u := range g.Board.GetUnits(p) {
				time.Sleep(time.Millisecond * 10)
				from := game.Coord{u.X, u.Y}
				// Try attacking
				for a, _ := range u.GetAllAttackCoords() {
					ac <- game.Action{ActionType: game.ActionAttack, From: from, To: a}
				}

				// Try moving
				for _, m := range g.Board.GetShortestPathToNearestEnemy(u) {
					ac <- game.Action{ActionType: game.ActionMove, From: from, To: m}
				}

			// Try attacking again
				from = game.Coord{u.X, u.Y}
				for a, _ := range u.GetAllAttackCoords() {
					ac <- game.Action{ActionType: game.ActionAttack, From: from, To: a}
				}

			}
			// Last action - end turn
			time.Sleep(time.Millisecond * 220)
			ac <- game.Action{ActionType: game.ActionEnd}
		}
	}
}
