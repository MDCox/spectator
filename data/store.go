package data

import (
	"github.com/sorcix/irc"
)

// Create user node and "IS_IN" edge to room if non-existent.
func joined(msg *irc.Message) {

}

// Check if reference to a person associated with that room.
// If it is, create a "REFERENCED" edge between speaker and
// the reference.  If that edge already exists, increment
// the "times" property by 1.
func action(msg *irc.Message) {

}

// Check if reference to a person associated with that room.
// If it is, create a "REFERENCED" edge between speaker and
// the reference.  If that edge already exists, increment
// the "times" property by 1.
func pmsged(msg *irc.Message) {

}
