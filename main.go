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
        
        if model.AllDestroyed(balls) {
            rl.ClearBackground(rl.Green)
            text := "YOU WIN!"
            textSize := int32(32)
            textWidth := rl.MeasureText(text, textSize)
            rl.DrawText(text, ScreenWidth/2 - textWidth/2, ScreenHeight/2-textSize/2, textSize, rl.Black)
            rl.EndDrawing()
            continue
        }
		
        rl.ClearBackground(rl.Black)
        rl.DrawRectangle(room.PosX, room.PosY, room.Width, room.Height, rl.RayWhite)

		for i := range balls {
            if !balls[i].Destroyed {
                rl.DrawCircle(balls[i].PosX, balls[i].PosY, model.BallRadius, rl.Black)
                balls[i].NextPosition(room)
            }
		}

        if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
            pos := rl.GetMousePosition()
            room.Divide(int32(pos.X), int32(pos.Y), model.DirVertical)
        }

        if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
            pos := rl.GetMousePosition()
            room.Divide(int32(pos.X), int32(pos.Y), model.DirHorizontal)
        }

		rl.EndDrawing()
	}
}
