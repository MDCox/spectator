// Parses and analyzes collected data before being stored in Neo4j

package data

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"github.com/sorcix/irc"
	"log"
)

func ConnectToNeo() *neoism.Database {
	db, err := neoism.Connect("http://localhost:7474/db/data")
	if err != nil {
		log.Fatal("COULD NOT CONNECT TO NEO4J: ", err)
	}
	return db
}

func Handle(input string, db *neoism.Database) {
	msg := irc.ParseMessage(input)
	if msg == nil {
		fmt.Println("Could not parse message")
	}

	format(msg)
}

func format(msg *irc.Message) {

}

func store() {

}
