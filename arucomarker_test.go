package arucomarker

import (
	"testing"
)

func TestDrawMarker(t *testing.T) {
	err := DrawMarker()
	if err != nil {
		t.Log(err)
	} else {
		t.Log("test draw marker")
	}
}
