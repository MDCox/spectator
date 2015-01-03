// Collects, parses, and displays Portland IRC community data.

package main

import (
	"./data"
	"./irc"
	"fmt"
	"github.com/jmcvetta/neoism"
	"time"
)

// Handles all the different packages to make sure that the data-collecting bot
// can pass its output to the data pkg, and that the http server has no problem
// querying that data.
func main() {
	// Channel to be used to pass information back from irc server
	c := make(chan string)

	// Holds current irc message
	var msg string

	fmt.Println(time.Now())
	fmt.Println("Server starting...")

	DB, err := neoism.Connect("http://127.0.0.1:7474/db/data")
	if err != nil {
		fmt.Println("Could not find Neo4j instance")
		fmt.Println(DB, err)
	}

	go irc.Connect("irc.freenode.net:6665", c)
	for {
		select {
		case msg = <-c:
			data.Handle(msg, DB)
		}
	}
}
