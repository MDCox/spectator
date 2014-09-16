// Parses and analyzes collected data before being stored in db

package data

import (
	"fmt"
	"github.com/sorcix/irc"
)

// Main datastore
var DB db

func init() {
	DB.Nodes = make(map[string]*Node)
	DB.Edges = make(map[string]*Edge)
}

func Handle(input string) {
	msg := irc.ParseMessage(input)
	if msg == nil {
		fmt.Println("Could not parse message")
	}
	store(msg)
	viewDB(input)
}

func store(msg *irc.Message) {
	switch msg.Command {
	case "JOIN":
		joined(msg)
	case "PRIVMSG":
		messaged(msg)
	case "ACTION":
		messaged(msg)
	}
}

// Print DB to console in human readable format.
// For debugging purposes only.
func viewDB(msg string) {
	fmt.Println("MSG: ", msg)
	fmt.Println("DB:")
	fmt.Println("    Nodes:")
	for k, v := range DB.Nodes {
		fmt.Printf("        %s, %s\n", k, v)
	}
	fmt.Println("    Edges:")
	for k, v := range DB.Edges {
		fmt.Printf("        %s, %s\n", k, v)
	}
	fmt.Println("--------\n")
}
