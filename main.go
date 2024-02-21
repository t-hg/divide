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
        PosX: 0,
        PosY: 0,
        Width: ScreenWidth,
        Height: ScreenHeight,
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
			rl.DrawCircle(balls[i].PosX, balls[i].PosY, 4, rl.Red)
			balls[i].NextPosition(room)
		}

        if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
            pos := rl.GetMousePosition()
            room.Divide(int32(pos.X), int32(pos.Y), model.DirHorizontal)
        }

        if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
            pos := rl.GetMousePosition()
            room.Divide(int32(pos.X), int32(pos.Y), model.DirVertical)
        }

		rl.EndDrawing()
	}
}
