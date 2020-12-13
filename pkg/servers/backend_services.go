package servers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const basePort = 9000
const serverName = "server-"

func printPort(res http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprint(res, "Hello: "+req.Host)
}
func checkBoundaries(numberOfServer int) error {
	if numberOfServer < 1 {
		return errors.New("number of servers should be more than 0")
	}
	return nil
}

func createWaitGroup(numberOfServer int) *sync.WaitGroup {
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(numberOfServer)
	return waitGroup
}

func createServer(serverName string, serverPort string) *http.Server {
	mux := http.NewServeMux()
	fmt.Printf("Creating server %s: on port: %s\n", serverName, serverPort)
	mux.HandleFunc("/", printPort)
	serverObj := http.Server{
		Addr:    serverPort,
		Handler: mux,
	}
	return &serverObj
}

//StartServer spins the numberOfServers concurrently
func StartServer(numberOfServer int) (int, error) {
	if err := checkBoundaries(numberOfServer); err != nil {
		return numberOfServer, fmt.Errorf("Error: \nservers.StartServer(numberOfServer): Cannot deploy %d servers, %w\n", numberOfServer, err)
	}
	waitGroup := createWaitGroup(numberOfServer)
	for i := 0; i < numberOfServer; i++ {
		currPort := basePort + i
		currAddr := ":" + strconv.Itoa(currPort)
		currServerName := serverName + strconv.Itoa(i+1)
		go func() {
			serverObj := createServer(currServerName, currAddr)
			err := serverObj.ListenAndServe()
			if err != nil {
				log.Fatalf("Error: \nservers.StartServer(numberOfServer): Cannot deploy %d servers, %s\n", numberOfServer, err.Error())
			}
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()

	return numberOfServer, nil
}
