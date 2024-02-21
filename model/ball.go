package model

import (
	"math/rand"
)

type Ball struct {
	PosX, PosY int32
	DirX, DirY int32
	Force      float64
}

func RandomBall(room Room) Ball {
	for {
		ball := Ball{
			PosX:  rand.Int31n(room.Width) + room.PosX,
			PosY:  rand.Int31n(room.Height) + room.PosY,
			DirX:  int32(rand.Intn(3) - 1),
			DirY:  int32(rand.Intn(3) - 1),
			Force: rand.Float64()*2 + 1,
		}
		if ball.DirX == 0 && ball.DirY == 0 {
			continue
		}
		return ball
	}
}

func RandomBalls(room Room, n int) []Ball {
	balls := make([]Ball, n)
	for i := range n {
		balls[i] = RandomBall(room)
	}
	return balls
}

func (ball *Ball) NextPosition(room Room) {
	ball.PosX += int32(float64(ball.DirX) * ball.Force)
	ball.PosY += int32(float64(ball.DirY) * ball.Force)
	if ball.PosX < room.PosX || ball.PosX > room.PosX+room.Width {
		ball.DirX *= -1
	}
	if ball.PosY < room.PosY || ball.PosY > room.PosY+room.Height {
		ball.DirY *= -1
	}
}
