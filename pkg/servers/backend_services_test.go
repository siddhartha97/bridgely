package servers

import (
	"testing"
)

func TestBoundaryCheck1(t *testing.T) {
	params := [5]int{1, 1000, 45, 999, 322092}
	for i := 0; i < 5; i++ {
		numberOfServers := params[i]
		err := checkBoundaries(numberOfServers)
		expected := "nil"
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

func TestGenerateServerParams(t *testing.T) {
	params1 := [4]int{4, 0, 1, 5}
	for i := 0; i < 4; i++ {
		numberOfServers := params1[i]
		_, _, err := generateServerParams(numberOfServers)
		expected := "nil"
		if err != nil {
			t.Errorf("generateServerParams(%d) failed, expected %v, got %v", numberOfServers, expected, err)
		} else {
			t.Logf("generateServerParams(%d) success, expected %v, got %v", numberOfServers, expected, err)
		}
	}
	params2 := [4]int{-4, -2, -1, -5000}
	for i := 0; i < 4; i++ {
		numberOfServers := params2[i]
		_, _, err := generateServerParams(numberOfServers)
		errorExpected := "server number should be non-negative"
		if err == nil {
			t.Errorf("generateServerParams(%d) failed, expected %v, got %v", numberOfServers, errorExpected, err)
		} else {
			t.Logf("generateServerParams(%d) success, expected %v, got %v", numberOfServers, errorExpected, err)
		}
	}
}
