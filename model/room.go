package model

import "math/rand"

type Room struct {
	PosX, PosY    int32
	Width, Height int32
}

func (room *Room) Divide(posX, posY int32, dir Direction) {
	if posX < room.PosX || posX > room.PosX+room.Width {
		return
	}
	if posY < room.PosY || posY > room.PosY+room.Height {
		return
	}
	switch dir {
	case DirHorizontal:
		upperRoom := Room{
			PosX:   room.PosX,
			PosY:   room.PosY,
			Width:  room.Width,
			Height: posY - room.PosY,
		}
		lowerRoom := Room{
			PosX:   room.PosX,
			PosY:   posY,
			Width:  room.Width,
			Height: room.PosY + room.Height - posY,
		}
        biggerRoom := biggerRoom(upperRoom, lowerRoom)
        room.PosX = biggerRoom.PosX
        room.PosY = biggerRoom.PosY
        room.Height = biggerRoom.Height
	case DirVertical:
        leftRoom := Room {
            PosX: room.PosX,
            PosY: room.PosY,
            Width: posX - room.PosX,
            Height: room.Height, 
        }
        rightRoom := Room {
            PosX: posX,
            PosY: room.PosY,
            Width: room.PosX + room.Width - posX,
            Height: room.Height, 
        }
        biggerRoom := biggerRoom(leftRoom, rightRoom)
        room.PosX = biggerRoom.PosX
        room.PosY = biggerRoom.PosY
        room.Width = biggerRoom.Width
	default:
		return
	}
}

func biggerRoom(roomA, roomB Room) Room {
	if area(roomA) == area(roomB) {
		return randomRoom(roomA, roomB)
	} else if area(roomA) > area(roomB) {
		return roomA
	} else {
		return roomB
	}
}

func randomRoom(roomA, roomB Room) Room {
	n := rand.Intn(2)
	if n == 0 {
		return roomA
	} else {
		return roomB
	}
}

func area(room Room) int32 {
	return room.Width * room.Height
}
