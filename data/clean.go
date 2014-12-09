package data

import "strings"

// Remove IRC op symbols from nick (eg: ~.@.+)
// A common IRC convention is to change your nick with a variation on "Away"
// appended when you are away from the computer, which we also want to try
// to remove intelligenly so we don't mess up people's nicks but also don't
// get duplicate nodes for people who talk as both "nick" and "nick_away."
func cleanName(nick string) string {
	var cleaned string

	// Make alphanumeric only
	cleaned = alphanum.ReplaceAllString(nick, "")

	// Possible "away" tokens
	cleaned = strings.Replace(cleaned, "away", "", 1)
	cleaned = strings.Replace(cleaned, "AWAY", "", 1)
	cleaned = strings.Replace(cleaned, "gone", "", 1)
	cleaned = strings.Replace(cleaned, "GONE", "", 1)

	return cleaned
}
