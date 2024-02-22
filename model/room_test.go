package model

import "testing"

func TestDivide(t *testing.T) {
	room := Room{}

	resetRoom := func() {
		room = Room{
			PosX:   50,
			PosY:   100,
			Width:  350,
			Height: 400,
		}
	}

	resetRoom()

	tests := []struct {
		PosX, PosY   int32
		Dir          Direction
		ExpectedRoom Room
	}{
		{
			PosX: 49,
			PosY: 150,
			Dir:  DirHorizontal,
			// x out of bound
			ExpectedRoom: room,
		},
		{
			PosX: 401,
			PosY: 150,
			Dir:  DirHorizontal,
			// x out of bound
			ExpectedRoom: room,
		},
		{
			PosX: 100,
			PosY: 99,
			Dir:  DirHorizontal,
			// y out of bound
			ExpectedRoom: room,
		},
		{
			PosX: 100,
			PosY: 501,
			Dir:  DirHorizontal,
			// y out of bound
			ExpectedRoom: room,
		},
        {
			PosX: 51,
			PosY: 200,
			Dir:  DirHorizontal,
            // lower room 
            ExpectedRoom: Room{
                PosX: 50,
                PosY: 200,
                Width: 350,
                Height: 300,
            },
        },
        {
			PosX: 51,
			PosY: 350,
			Dir:  DirHorizontal,
            // upper room 
            ExpectedRoom: Room{
                PosX: 50,
                PosY: 100,
                Width: 350,
                Height: 250,
            },
        },
        {
			PosX: 150,
			PosY: 101,
			Dir:  DirVertical,
            // right room 
            ExpectedRoom: Room{
                PosX: 150,
                PosY: 100,
                Width: 250,
                Height: 400,
            },
        },
        {
			PosX: 250,
			PosY: 101,
			Dir:  DirVertical,
            // left room 
            ExpectedRoom: Room{
                PosX: 50,
                PosY: 100,
                Width: 200,
                Height: 400,
            },
        },
	}

	for _, test := range tests {
		room.Divide(test.PosX, test.PosY, test.Dir)
        if room != test.ExpectedRoom {
            t.Errorf("Expected %v, got %v", test.ExpectedRoom, room)
        }
        resetRoom()
	}
}
