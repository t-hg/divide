package model

import (
	"math/rand"

	"github.com/t-hg/divide/screen"
)

type Ball struct {
    PosX, PosY int32
    DirX, DirY int32
    Force float64
}

func RandomBall() Ball {
    for {
        ball := Ball {
            PosX: int32(rand.Intn(int(screen.Width))),
            PosY: int32(rand.Intn(int(screen.Height))),
            DirX: int32(rand.Intn(3) - 1),
            DirY: int32(rand.Intn(3) - 1),
            Force: rand.Float64() * 2 + 1,
        }
        if ball.DirX == 0 && ball.DirY == 0 {
            continue
        }
        return ball
    }
}

func RandomBalls(n int) []Ball {
    balls := make([]Ball, n)
    for i := range n {
        balls[i] = RandomBall()
    }
    return balls
}

func (ball *Ball) NextPosition() {
    ball.PosX += int32(float64(ball.DirX) * ball.Force)
    ball.PosY += int32(float64(ball.DirY) * ball.Force)
    if ball.PosX < 0 || ball.PosX > screen.Width {
        ball.DirX *= -1        
    }
    if ball.PosY < 0 || ball.PosY > screen.Height {
        ball.DirY *= -1        
    }
}
