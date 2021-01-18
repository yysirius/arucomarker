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
		/*{
			SlideSize:  6,
			MarkerSize: 7,
			BorderSize: 1,
			Name:       "7x7_8_10",
			Data:       []uint8{137, 7, 31, 86, 150, 158, 0, 112, 74, 104, 49, 231, 125, 0},
		},*/
		{
			SlideSize:  150,
			MarkerSize: 6,
			BorderSize: 1,
			Name:       "6x6_0_150",
			Data:       []uint8{30, 61, 216, 42, 6, 227, 186, 70, 49, 9, 101, 65, 187, 199, 8, 152, 198, 37, 220, 7},
		},
		{
			SlideSize:  150,
			MarkerSize: 6,
			BorderSize: 1,
			Name:       "6x6_446_150",
			Data:       []uint8{33, 194, 87, 244, 1, 60, 69, 186, 88, 2, 130, 254, 164, 56, 4, 65, 165, 218, 35, 12},
		},
		{
			SlideSize:  150,
			MarkerSize: 6,
			BorderSize: 1,
			Name:       "6x6_999_150",
			Data:       []uint8{254, 214, 62, 31, 15, 206, 125, 253, 183, 5, 255, 135, 198, 183, 15, 174, 219, 251, 231, 3},
		},
	}

	for _, c := range cases {
		err := DrawMarker(c)
		if err != nil {
			t.Fatal("draw failed", c.Name, err)
		}
	}

}
