package data

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"github.com/sorcix/irc"
	"strings"
)

// Create user node and "IS_IN" edge to room if non-existent.
func joined(msg *irc.Message, DB *neoism.Database) {
	room := msg.Params[0]
	user := cleanName(msg.Prefix.Name)

	if users[user] == nil {
		users[user] = []string{room}
	}

	contains := false
	for _, r := range users[user] {
		if r == room {
			contains = true
			break
		}
	}

	if contains != true {
		users[user] = append(users[user], room)

		statement := fmt.Sprintf(
			`MERGE (n:Room {name: "%v"})
			 MERGE (u:User {name: "%v"})
			 MERGE (n)<-[:IS_IN]-(u)`,
			room, user)

		// Create new room node if non-existent
		query := neoism.CypherQuery{
			Statement: statement,
		}
		err := DB.Cypher(&query)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Check if reference to a person associated with that room.
// If it is, create a "REFERENCED" edge between speaker and
// the reference.  If that edge already exists, increment
// the "times" property by 1.
func messaged(msg *irc.Message, DB *neoism.Database) {
	message := msg.Trailing
	for name, _ := range users {
		if strings.Contains(message, name) {
			speaker := cleanName(msg.Prefix.Name)
			reference := name

			fmt.Printf("%v was referenced by %v", speaker, reference)

			statement := fmt.Sprintf(
				`MATCH (s:User {name: "%v"}), (u:User {name: "%v"})
				 MERGE (s)-[r:REFERENCED]->(u)
				 ON MATCH SET r.times = coalesce(r.times, 0) + 1`, speaker, reference)

			query := neoism.CypherQuery{
				Statement: statement,
			}
			err := DB.Cypher(&query)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// Adds nodes for users who were already in the room before
// spectator was started.  If they were added from another
// room, they are not added.  An edge is drawn to show that
// they are in the room.
func inchan(msg *irc.Message, DB *neoism.Database) {
	nicks := strings.Split(msg.Trailing, " ")
	room := msg.Params[2]

	queryStart := fmt.Sprintf(`MERGE (n:Room {name: "%v"}) `, room)
	rawQuery := []string{queryStart}

	i := 0
	for _, u := range nicks {
		cu := cleanName(u)
		if users[u] == nil {
			pattern := fmt.Sprintf(`MERGE (u%v:User {name: "%v"}) MERGE (u%v)-[:IS_IN]->(n)`, i, cu, i)
			rawQuery = append(rawQuery, pattern)
			i++
		}
	}

	query := neoism.CypherQuery{
		Statement: strings.Join(rawQuery, " "),
	}
	err := DB.Cypher(&query)
	if err != nil {
		fmt.Println(err)
	}
}
