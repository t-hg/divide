package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/t-hg/divide/model"
	"github.com/t-hg/divide/screen"
)

func main() {
	balls := model.RandomBalls(25)
	rl.InitWindow(screen.Width, screen.Height, "divide")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for i := range balls {
			rl.DrawCircle(balls[i].PosX, balls[i].PosY, 4, rl.Black)
			balls[i].NextPosition()
		}
		rl.EndDrawing()
	}
}
