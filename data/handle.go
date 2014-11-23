// Parses and analyzes collected data before being stored in db

package data

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"github.com/sorcix/irc"
)

// Main datastore
var DB *neoism.Database

// Map with usernames as key and their associated rooms as the values
var users map[string][]string

func init() {
	DB, err := neoism.Connect("http://127.0.0.1:7474/db/data")
	if err != nil {
		fmt.Println("Could not find Neo4j instance")
		fmt.Println(DB, err)
	}
}

func Handle(input string) {
	msg := irc.ParseMessage(input)
	if msg == nil {
		fmt.Println("Could not parse message")
	}
	store(msg)
}

func store(msg *irc.Message) {
	switch msg.Command {
	case "JOIN":
		joined(msg)
	case "PRIVMSG":
		messaged(msg)
	case "ACTION":
		messaged(msg)
	// List of nicks in Channel before start.
	case "353":
		inchan(msg)
	}
}
