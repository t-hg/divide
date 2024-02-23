package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/t-hg/divide/model"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 640
)

func main() {
	room := model.Room{
		PosX:   0,
		PosY:   0,
		Width:  ScreenWidth,
		Height: ScreenHeight,
	}

	balls := model.RandomBalls(room, 25)

	var divider *model.Divider

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
			rl.DrawText(text, ScreenWidth/2-textWidth/2, ScreenHeight/2-textSize/2, textSize, rl.Black)
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

		divideDirection := model.DirUndefined

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			divideDirection = model.DirVertical
		} else if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
			divideDirection = model.DirHorizontal
		}

		if divideDirection != model.DirUndefined {
			if divider == nil {
				pos := rl.GetMousePosition()
				divider = &model.Divider{
					StartX:        int32(pos.X),
					StartY:        int32(pos.Y),
					EndX:          int32(pos.X),
					EndY:          int32(pos.Y),
					Dir:           divideDirection,
					FullyExpanded: false,
				}
			}
		}

		if divider != nil {
			if divider.FullyExpanded {
				room.Divide(divider.StartX, divider.StartY, divider.Dir)
				divider = nil
			} else {
				rl.DrawLineEx(rl.Vector2{
					X: float32(divider.StartX),
					Y: float32(divider.StartY),
				}, rl.Vector2{
					X: float32(divider.EndX),
					Y: float32(divider.EndY),
				}, 4, rl.Black)
				divider.Expand(room)
			}
		}

		rl.EndDrawing()
	}
}
