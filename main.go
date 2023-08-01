package main

import (
	"minesweeper/game"
	"os"
	"path"

	ray "github.com/gen2brain/raylib-go/raylib"
)

const (
	WINDOW_WIDTH  = 1000
	WINDOW_HEIGHT = 600
)

func main() {
	currentPath, _ := os.Executable()
	currentPath = path.Dir(currentPath)
	os.Chdir(currentPath)
	ray.SetTraceLog(ray.LogError)
	ray.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Mine Sweeper")
	defer ray.CloseWindow()
	ray.SetWindowState(ray.FlagWindowResizable)
	ray.SetWindowPosition(100, 100)
	ray.SetTargetFPS(60)
	game.Load()
	defer game.Unload()
	for !ray.WindowShouldClose() {
		game.UpdateGame()

		ray.BeginDrawing()
		ray.ClearBackground(ray.Gray)
		game.DrawGame()
		ray.EndDrawing()
	}
}
