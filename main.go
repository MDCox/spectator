// Collects, parses, and displays Portland IRC community data.

package main

import (
	"./irc"
	"fmt"
)

// Handles all the different packages to make sure the data-collecting bot
// can input it's data into Neo4j, and that the http server has no problem
// querying that data.
func main() {
	c := make(chan string)
	var msg string
	fmt.Println("Server starting...")
	go irc.Connect("irc.freenode.net:6665", c)
	for {
		select {
		case msg = <-c:
			fmt.Println(msg)
		}
	}
}
