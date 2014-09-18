package data

// Main datastructure to be marshalled into GraphJSON.
// Nodes and Edges are stored in a map with their key
// set to the node/edge ID.
type db struct {
	Nodes map[string]*Node `json:"nodes"`
	Edges map[string]*Edge `json:"edges"`
}

// Node representing an irc user or channel
//
// The ID is the user's nick or the channel name.
// The NodeType is a string that determines if the node is
// a room or a user.  Depending on NodeType, Count tallies
// the number of either users in a room, or the number of
// references made to that person.
type Node struct {
	ID       string `json:"id"`
	NodeType string `json:"nodeType"`
	Count    int    `json:"count"`
}

// Edge representing a relationship or action between nodes.
//
// The Source and Target fields are strings that contain the
// source and target's node IDs.
// The EdgeType can be set to either "IS_IN" in the case of a
// user being a member of a room, or "REFERENCED" in the case
// of a user referencing another user.
type Edge struct {
	Source   string `json:"source"`
	Target   string `json:"target"`
	EdgeType string `json:"edgeType"`
}
