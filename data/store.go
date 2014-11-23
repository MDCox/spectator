package data

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"github.com/sorcix/irc"
	"strings"
)

// Create user node and "IS_IN" edge to room if non-existent.
func joined(msg *irc.Message) {
	room := msg.Params[0]
	user := cleanName(msg.Prefix.Name)

	contains := false
	if users[user] == nil {
		users[user] = []string{room}
	} else {
		for _, r := range users[user] {
			if r == room {
				contains = true
				break
			}
		}

	}

	if contains != true {
		users[user] = append(users[user], room)

		// Create new room node if non-existent
		query := neoism.CypherQuery{
			Statement:  "MERGE (n:Room {name: {room}})<-[:IS_IN]-(u:User {name: {user}})",
			Parameters: neoism.Props{"room": room, "user": user},
		}
		DB.Cypher(&query)
	}
}

// Check if reference to a person associated with that room.
// If it is, create a "REFERENCED" edge between speaker and
// the reference.  If that edge already exists, increment
// the "times" property by 1.
func messaged(msg *irc.Message) {
	message := msg.Trailing
	for name, _ := range users {
		if strings.Contains(message, name) {
			speaker := cleanName(msg.Prefix.Name)
			reference := name

			query := neoism.CypherQuery{
				Statement: `MERGE (s:User {name: {speaker}})-[r:REFERENCED]-(u:User {name: {reference}})
					    ON MATCH r.times = coalesce(r.times, 0) + 1`,
				Parameters: neoism.Props{"speaker": speaker, "reference": reference},
			}
			DB.Cypher(&query)
		}
	}
}

// Adds nodes for users who were already in the room before
// spectator was started.  If they were added from another
// room, they are not added.  An edge is drawn to show that
// they are in the room.
func inchan(msg *irc.Message) {
	nicks := strings.Split(msg.Trailing, " ")
	room := msg.Params[2]
	rawQuery := []string{"MERGE"}

	for _, u := range nicks {
		if users[u] == nil {
			pattern := fmt.Sprintf(`(:User {name: "%v"})-[:IS_IN]->(:Room {name: "%v"}),`, u, room)
			rawQuery = append(rawQuery, pattern)
		}
	}

	query := neoism.CypherQuery{
		Statement: strings.Join(rawQuery, " "),
	}
	DB.Cypher(&query)
}
