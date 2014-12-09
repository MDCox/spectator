// Parses and analyzes collected data before being stored in db

package data

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"github.com/sorcix/irc"
	"regexp"
)

// Map with usernames as key and their associated rooms as the values
var users map[string][]string = make(map[string][]string)

// Regex used for name cleaning
var alphanum *regexp.Regexp

func init() {
	alphanum, _ = regexp.Compile("[^a-zA-Z0-9]")
}

func Handle(input string, DB *neoism.Database) {
	msg := irc.ParseMessage(input)
	if msg == nil {
		fmt.Println("Could not parse message")
	}
	store(msg, DB)
}

func store(msg *irc.Message, DB *neoism.Database) {
	switch msg.Command {
	case "JOIN":
		joined(msg, DB)
	case "PRIVMSG":
		messaged(msg, DB)
	case "ACTION":
		messaged(msg, DB)
	// List of nicks in Channel before start.
	case "353":
		inchan(msg, DB)
	}
}
