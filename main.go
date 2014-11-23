// Collects, parses, and displays Portland IRC community data.

package main

import (
	"./data"
	"./irc"
	"bufio"
	"fmt"
	"os"
	"time"
)

var startTime time.Time = time.Now()

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
	go irc.Connect("irc.freenode.net:6665", c)
	go cli()
	for {
		select {
		case msg = <-c:
			data.Handle(msg)
		}
	}
}

// simple cli that allows checking of db stats and runtime
func cli() {
	var inp string
	reader := bufio.NewReader(os.Stdin)
	for {
		inp, _ = reader.ReadString('\n')
		switch inp {
		case "runtime\n":
			fmt.Println(time.Now().Sub(startTime))
		case "quit\n", "q\n", "exit\n":
			fmt.Println(time.Now().Sub(startTime))
			os.Exit(1)
		}
	}
}
