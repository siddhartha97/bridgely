package servers

import "testing"

func TestBoundaryCheck1(t *testing.T) {
	param := 1
	err := checkBoundaries(int64(param))
	expected := "number of servers should be more than 0"
	if err != nil {
		t.Errorf("checkBoundaries(%d) failed, expected %v, got %v", param, expected, err)
	} else {
		t.Logf("checkBoundaries(%d) sucess, expected %v, got %v", param, expected, err)
	}
}

func TestBoundaryCheck2(t *testing.T) {
	param := 0
	err := checkBoundaries(int64(param))
	expected := "number of servers should be more than 0"
	if err == nil {
		t.Errorf("checkBoundaries(%d) failed, expected %v, got %v", param, expected, err)
	} else {
		t.Logf("checkBoundaries(%d) sucess, expected %v, got %v", param, expected, err)
	}
}

func TestBoundaryCheck3(t *testing.T) {
	param := -1000
	err := checkBoundaries(int64(param))
	expected := "number of servers should be more than 0"
	if err == nil {
		t.Errorf("checkBoundaries(%d) failed, expected %v, got %v", param, expected, err)
	} else {
		t.Logf("checkBoundaries(%d) sucess, expected %v, got %v", param, expected, err)
	}
}

func TestBoundaryCheck4(t *testing.T) {
	param := -1
	err := checkBoundaries(int64(param))
	expected := "number of servers should be more than 0"
	if err == nil {
		t.Errorf("checkBoundaries(%d) failed, expected %v, got %v", param, expected, err)
	} else {
		t.Logf("checkBoundaries(%d) sucess, expected %v, got %v", param, expected, err)
	}
}
