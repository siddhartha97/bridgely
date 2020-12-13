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

//checkBoundaries checks if number of servers are less than 1
func checkBoundaries(numberOfServer int) error {
	if numberOfServer < 1 {
		return errors.New("number of servers should be more than 0")
	}
	return nil
}

//createWaitGroup creates a waiting group of number of servers
func createWaitGroup(numberOfServer int) *sync.WaitGroup {
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(numberOfServer)
	return waitGroup
}

//createServer creates the server
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

//createServerInfo creates a server name and assigns a port
func createServerInfo(param int) (string, string, error) {
	if param < 0 {
		return "", "", errors.New("server number should be non-negative")
	}
	currPort := basePort + param
	currPortAddr := ":" + strconv.Itoa(currPort)
	currServerName := serverName + strconv.Itoa(param+1)
	return currServerName, currPortAddr, nil
}

//StartServer spins the numberOfServers concurrently
func StartServer(numberOfServer int) (int, error) {
	if err := checkBoundaries(numberOfServer); err != nil {
		return numberOfServer, fmt.Errorf("Error: \ncheckBoundaries(numberOfServer): Cannot deploy %d servers, %w\n", numberOfServer, err)
	}
	waitGroup := createWaitGroup(numberOfServer)
	for i := 0; i < numberOfServer; i++ {
		currServerName, currPortAddr, err := createServerInfo(i)
		if err != nil {
			log.Fatalf("Error: \ncreateServerInfo(param): Cannot deploy %d servers, %s\n", numberOfServer, err.Error())
		}
		go func() {
			serverObj := createServer(currServerName, currPortAddr)
			err := serverObj.ListenAndServe()
			if err != nil {
				log.Fatalf("Error: \nListenAndServe(): Cannot deploy %d servers, %s\n", numberOfServer, err.Error())
			}
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	return numberOfServer, nil
}
