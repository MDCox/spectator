// Collects, parses, and displays Portland IRC community data.

package main

import (
	"./data"
	"./irc"
	"bufio"
	"fmt"
	"github.com/jmcvetta/neoism"
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

	DB, err := neoism.Connect("http://127.0.0.1:7474/db/data")
	if err != nil {
		fmt.Println("Could not find Neo4j instance")
		fmt.Println(DB, err)
	}

	go irc.Connect("irc.freenode.net:6665", c)
	go cli()
	for {
		select {
		case msg = <-c:
			data.Handle(msg, DB)
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
