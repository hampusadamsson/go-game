package main

import (
	"fmt"
	_ "image/png"

	"github.com/hampusadamsson/go-game/game"
)

func main() {
	gf := game.GameFactory{}
	p := game.Player{}
	g := gf.Tutorial(p)
	fmt.Println(g)
}
