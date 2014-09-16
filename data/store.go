package data

import (
	"fmt"
	"github.com/sorcix/irc"
	"strings"
)

// Create user node and "IS_IN" edge to room if non-existent.
func joined(msg *irc.Message) {
	// Create new room node if non-existent
	if DB.Nodes[msg.Params[0]] == nil {
		fmt.Printf("New Room: %s\n", msg.Params[0])
		DB.Nodes[msg.Params[0]] = &Node{
			ID:       msg.Params[0],
			NodeType: "room",
		}
	}

	// Create new user node if non-existent
	if DB.Nodes[msg.Prefix.Name] == nil {
		fmt.Printf("New User: %s\n", msg.Prefix.Name)
		DB.Nodes[msg.Prefix.Name] = &Node{
			ID:       msg.Prefix.Name,
			NodeType: "user",
		}
	}

	// Associate user with channel
	edgeID := fmt.Sprintf("%s-%s", msg.Prefix.Name, msg.Params[0])
	if DB.Edges[edgeID] == nil {
		DB.Edges[edgeID] = &Edge{
			Source:   DB.Nodes[msg.Prefix.Name].ID,
			Target:   DB.Nodes[msg.Params[0]].ID,
			EdgeType: "IS_IN",
		}
	}
}

// Check if reference to a person associated with that room.
// If it is, create a "REFERENCED" edge between speaker and
// the reference.  If that edge already exists, increment
// the "times" property by 1.
func messaged(msg *irc.Message) {
	message := msg.Trailing
	for name, _ := range DB.Nodes {
		if strings.Contains(message, name) {
			edgeID := fmt.Sprintf("%s-%s", msg.Prefix.Name, name)
			DB.Edges[edgeID] = &Edge{
				Source:   msg.Prefix.Name,
				Target:   name,
				EdgeType: "REFERENCED",
			}
			fmt.Printf("\n%s referenced %s\n", msg.Prefix.Name, name)
		}
	}
}
