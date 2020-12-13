package servers

import (
	"testing"
)

func TestBoundaryCheck1(t *testing.T) {
	params := [5]int{1, 1000, 45, 999, 322092}
	for i := 0; i < 5; i++ {
		numberOfServers := params[i]
		err := checkBoundaries(numberOfServers)
		expected := "number of servers should be more than 0"
		if err != nil {
			t.Errorf("checkBoundaries(%d) failed, expected %v, got %v", numberOfServers, expected, err)
		} else {
			t.Logf("checkBoundaries(%d) success, expected %v, got %v", numberOfServers, expected, err)
		}
	}
}

func TestBoundaryCheck2(t *testing.T) {
	params := [5]int{-1, -1000, 0, -45, -999}
	for i := 0; i < 5; i++ {
		numberOfServers := params[i]
		err := checkBoundaries(numberOfServers)
		expected := "number of servers should be more than 0"
		if err == nil {
			t.Errorf("checkBoundaries(%d) failed, expected %v, got %v", numberOfServers, expected, err)
		} else {
			t.Logf("checkBoundaries(%d) success, expected %v, got %v", numberOfServers, expected, err)
		}
	}
}
