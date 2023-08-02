package game

import (
	"os"
	"path"

	ray "github.com/gen2brain/raylib-go/raylib"
)

type Game struct{}

func (g *Game) Init() {
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
