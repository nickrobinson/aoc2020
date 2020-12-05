package boarding

import (
	"testing"
)

func TestGetSeatNumber(t *testing.T) {
	var tests = []struct {
		name    string
		bp      BoardingPass
		seatNum int
	}{
		{"Initial Example", BoardingPass{SpacePartition: "FBFBBFFRLR"}, 357},
		{"Basic Boarding Pass 1", BoardingPass{SpacePartition: "BFFFBBFRRR"}, 567},
		{"Basic Boarding Pass 2", BoardingPass{SpacePartition: "FFFBBBFRRR"}, 119},
		{"Basic Boarding Pass 2", BoardingPass{SpacePartition: "BBFFBBFRLL"}, 820},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seatNum := tt.bp.GetSeatNumber()
			if seatNum != tt.seatNum {
				t.Errorf("GetSeatNumber(%s) got %d, want %d", tt.name, seatNum, tt.seatNum)
			}
		})
	}
}
