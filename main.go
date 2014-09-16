// Collects, parses, and displays Portland IRC community data.

package main

import (
	"./data"
	"./irc"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Handles all the different packages to make sure that the data-collecting bot
// can pass its output to the data pkg, and that the http server has no problem
// querying that data.
func main() {
	// Channel to be used to pass information back from irc server
	c := make(chan string)

	// Dump JSON on ^C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		if sig != nil {
			data.Dump()
			os.Exit(1)
		}
	}()

	// Holds current irc message
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
