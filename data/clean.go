package data

import "strings"

// Remove IRC op symbols from nick (eg: ~.@.+)
// A common IRC convention is to change your nick with a variation on "Away"
// appended when you are away from the computer, which we also want to try
// to remove intelligenly so we don't mess up people's nicks but also don't
// get duplicate nodes for people who talk as both "nick" and "nick_away."
func cleanName(nick string) string {
	var cleaned string

	// IRC op codes
	cleaned = strings.Replace(nick, "+", "", 1)
	cleaned = strings.Replace(nick, "@", "", 1)
	cleaned = strings.Replace(nick, "~", "", 1)
	cleaned = strings.Replace(nick, "&", "", 1)

	// Possible "away" tokens
	cleaned = strings.Replace(nick, "_away", "", 1)
	cleaned = strings.Replace(nick, "_AWAY", "", 1)
	cleaned = strings.Replace(nick, "away", "", 1)
	cleaned = strings.Replace(nick, "AWAY", "", 1)
	cleaned = strings.Replace(nick, "[AWAY]", "", 1)
	cleaned = strings.Replace(nick, "|away", "", 1)
	cleaned = strings.Replace(nick, "|AWAY", "", 1)
	cleaned = strings.Replace(nick, "[GONE]", "", 1)
	cleaned = strings.Replace(nick, "GONE", "", 1)

	// other
	cleaned = strings.Replace(nick, "/", "", 1)
	cleaned = strings.Replace(nick, "_", "", 1)
	cleaned = strings.Replace(nick, "\\", "", 1)
	cleaned = strings.Replace(nick, "\"", "", 1)
	return cleaned
}
