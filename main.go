package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/t-hg/divide/model"
)

const (
    ScreenWidth = 800
    ScreenHeight = 640
)

func main() {
    room := model.Room{
        PosX: 300,
        PosY: 300,
        Width: 150,
        Height: 150,
    }
	balls := model.RandomBalls(room, 25)
	rl.InitWindow(ScreenWidth, ScreenHeight, "divide")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
        rl.DrawRectangle(room.PosX, room.PosY, room.Width, room.Height, rl.RayWhite)
		for i := range balls {
			rl.DrawCircle(balls[i].PosX, balls[i].PosY, 4, rl.Black)
			balls[i].NextPosition(room)
		}
		rl.EndDrawing()
	}
}
