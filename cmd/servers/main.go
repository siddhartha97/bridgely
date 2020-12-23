package main

import (
	"flag"
	"github.com/siddhartha97/bridgely/pkg/servers"
	"log"
)

func main() {
	numberOfServer := flag.Int("s", 1, "a flag for number of servers")
	flag.Parse()
	_, err := servers.StartServer(*numberOfServer)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
