package arucomarker

import (
	"testing"
)

func TestDrawMarker(t *testing.T) {
	cases := []MarkerInfo{
		{
			SlideSize:  200,
			MarkerSize: 5,
			BorderSize: 1,
			Name:       "5x5_0",
			Data:       []uint8{162, 217, 94, 0, 82, 46, 217, 1},
		},
		{
			SlideSize:  150,
			MarkerSize: 7,
			BorderSize: 1,
			Name:       "7x7_8_150",
			Data:       []uint8{137, 7, 31, 86, 150, 158, 0, 112, 74, 104, 49, 231, 125, 0},
		},
		{
			SlideSize:  6,
			MarkerSize: 7,
			BorderSize: 1,
			Name:       "7x7_8_10",
			Data:       []uint8{137, 7, 31, 86, 150, 158, 0, 112, 74, 104, 49, 231, 125, 0},
		},
	}

	for _, c := range cases {
		err := DrawMarker(c)
		if err != nil {
			t.Fatal("draw failed", c.Name, err)
		}
	}

}
