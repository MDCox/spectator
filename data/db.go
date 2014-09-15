package data

// Main datastructure to be marshalled into GraphJSON.
// Nodes and Edges are stored in a map with their key
// set to the node/edge ID.
type db struct {
	nodes map[string]node
	edges map[string]edge
}

// Node representing an irc user or channel
//
// The ID is the user's nick or the channel name.
// The NodeType is a string that determines if the node is
// a room or a user.  Depending on NodeType, Count tallies
// the number of either users in a room, or the number of
// references made to that person.
type node struct {
	ID       string
	NodeType string
	Count    int
}

// Edge representing a relationship or action between nodes.
//
// The Source and Target fields are strings that contain the
// source and target's node IDs.  The ID field is a string
// that combines the two seperated by a dash (eg: "ID1-ID2".)
// The EdgeType can be set to either "IS_IN" in the case of a
// user being a member of a room, or "REFERENCED" in the case
// of a user referencing another user.
type edge struct {
	ID       string
	Source   string
	Target   string
	EdgeType string
}
