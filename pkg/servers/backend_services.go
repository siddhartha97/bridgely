package servers

import (
	"errors"
	"fmt"
)

func checkBoundaries(numberOfServer int64) error {
	if numberOfServer < 1 {
		return errors.New("number of servers should be more than 0")
	}
	return nil
}

//Start spins the numberOfServers
func Start(numberOfServer int64) (int64, error) {
	if err := checkBoundaries(numberOfServer); err != nil {
		return numberOfServer, fmt.Errorf("Error: \nservers.Start(numberOfServer): Cannot deploy %d servers, %w\n", numberOfServer, err)
	}
	return numberOfServer, nil
}
