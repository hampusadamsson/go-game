package main

// import (
// 	"bufio"
// 	"fmt"
// 	_ "image/png"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/hampusadamsson/go-game/game"
// )

// func main() {
// 	gf := game.GameFactory{}

// 	p1Action := make(chan game.Action)
// 	p1 := &game.Player{"A", p1Action}

// 	p2Action := make(chan game.Action)
// 	p2 := &game.Player{"B", p2Action}

// 	g := gf.OneVsOne(p1, p2)
// 	g.Run()

// 	// Start AI loop
// 	go endTurnAi(p2Action)

// 	// Start player loop
// 	for {
// 		printBoard(g)
// 		fmt.Println()
// 		fmt.Println("(m)ove - ex. m 0 0 3 3")
// 		fmt.Println("(a)ttack - ex. m 0 0 3 3")
// 		fmt.Println("(e)nd")
// 		fmt.Println("(q)uit")
// 		fmt.Print("Action: ")

// 		reader := bufio.NewReader(os.Stdin)
// 		cliInput, _ := reader.ReadString('\n')
// 		cliInput = strings.Replace(cliInput, "\n", " ", 1)
// 		switch cliInput[:1] {
// 		case "m":
// 			l := strings.Split(cliInput, " ")
// 			x1, _ := strconv.Atoi(l[1])
// 			y1, _ := strconv.Atoi(l[2])
// 			x2, _ := strconv.Atoi(l[3])
// 			y2, _ := strconv.Atoi(l[4])
// 			p1Action <- game.Action{ActionType: game.ActionMove, From: game.Coord{x1, y1}, To: game.Coord{x2, y2}}
// 		case "a":
// 			l := strings.Split(cliInput, " ")
// 			x1, _ := strconv.Atoi(l[1])
// 			y1, _ := strconv.Atoi(l[2])
// 			x2, _ := strconv.Atoi(l[3])
// 			y2, _ := strconv.Atoi(l[4])
// 			p1Action <- game.Action{ActionType: game.ActionAttack, From: game.Coord{x1, y1}, To: game.Coord{x2, y2}}
// 		case "e":
// 			p1Action <- game.Action{ActionType: game.ActionEnd}
// 		case "q":
// 			break
// 		}

// 	}
// }

// func printBoard(g *game.Game) {
// 	for _, v := range g.Board.Tiles {
// 		for _, t := range v {
// 			if u, err := t.GetUnit(); err == nil {
// 				fmt.Printf("[%s%d ]", u.Owner.Name[:1], t.Cost)
// 			} else {
// 				fmt.Printf("[ %d ]", t.Cost)
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// func endTurnAi(ac chan game.Action) {
// 	for {
// 		time.Sleep(time.Second * 1)
// 		ac <- game.Action{ActionType: game.ActionEnd}
// 	}
// }
