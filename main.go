// Collects, parses, and displays Portland IRC community data.

package main

import (
	"./data"
	"./irc"
	"fmt"
)

// Handles all the different packages to make sure that the data-collecting bot
// can pass its output to the data pkg, and that the http server has no problem
// querying that data.
func main() {
	// Channel to be used to pass information back from irc server
	c := make(chan string)
	var msg string

	fmt.Println("Server starting...")
	go irc.Connect("irc.freenode.net:6665", c)
	for {
		select {
		case msg = <-c:
			data.Handle(msg)
		}
	}
}
