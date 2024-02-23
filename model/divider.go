package model

const DividerExpandSpeed = 4

type Divider struct {
	StartX, StartY int32
	EndX, EndY     int32
	Dir            Direction
	FullyExpanded  bool
}

func (divider *Divider) Expand(room Room) {
	switch divider.Dir {
	case DirHorizontal:
		expanding := false
		if divider.StartX > room.PosX {
			divider.StartX -= DividerExpandSpeed
			expanding = true
		}
		if divider.EndX < room.PosX+room.Width {
			divider.EndX += DividerExpandSpeed
			expanding = true
		}
		if !expanding {
			divider.FullyExpanded = true
		}
	case DirVertical:
		expanding := false
		if divider.StartY > room.PosY {
			divider.StartY -= DividerExpandSpeed
			expanding = true
		}
		if divider.EndY < room.PosY+room.Height {
			divider.EndY += DividerExpandSpeed
			expanding = true
		}
		if !expanding {
			divider.FullyExpanded = true
		}
	}
}
