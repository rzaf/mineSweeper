package game

import (
	"flag"
	"fmt"
	"os"
	"path"

	ray "github.com/gen2brain/raylib-go/raylib"
)

type Game struct{}

func handleFlags() {
	flag.IntVar(&boardWidth, "w", boardWidth, "number of vertical cells (minimum=5,maximum=100)")
	flag.IntVar(&boardHeight, "h", boardHeight, "number of horizontal cells (minimum=5,maximum=100)")
	flag.IntVar(&bombChance, "b", bombChance, "chance of a cell being bomb (minimum=0,maximum=100)")
	flag.Parse()
	w := boardWidth < 5 || boardWidth > 100
	h := boardHeight < 5 || boardHeight > 100
	b := bombChance < 0 || bombChance > 100
	if w || h || b {
		if w {
			fmt.Printf("`w=%d` not in range \n", boardWidth)
		}
		if h {
			fmt.Printf("`h=%d` not in range \n", boardHeight)
		}
		if b {
			fmt.Printf("`b=%d` not in range \n", bombChance)
		}
		flag.Usage()
		os.Exit(1)
	}
	if len(os.Args) == 1 {
		fmt.Println("run with --help to see options")
	}
}

func (g *Game) Init() {
	handleFlags()
	currentPath, _ := os.Executable()
	currentPath = path.Dir(currentPath)
	os.Chdir(currentPath)
	ray.SetTraceLog(ray.LogError)
	ray.InitWindow(1000, 600, "Mine Sweeper")
	ray.SetWindowState(ray.FlagWindowResizable)
	ray.SetWindowPosition(100, 100)
	ray.SetTargetFPS(60)
	Load()
}

func (g *Game) Run() {
	for !ray.WindowShouldClose() {
		UpdateGame()

		ray.BeginDrawing()
		ray.ClearBackground(ray.Gray)
		DrawGame()
		ray.EndDrawing()
	}
}

func (g *Game) Destroy() {
	Unload()
	ray.CloseWindow()
}
