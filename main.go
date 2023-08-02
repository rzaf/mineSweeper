package main

import (
	"github.com/rzaf/mineSweeper/game"
)

func main() {
	mineSweeper := game.Game{}
	mineSweeper.Init()
	mineSweeper.Run()
	mineSweeper.Destroy()
}
