// Handles irc connections and logs to be sent to the parser
package irc

import (
	"fmt"
	"github.com/sorcix/irc"
	"log"
)

// Connect() takes a string with url:port of the irc server, and a channel
// that accepts strings.  It connects to the irc server specified and
// identifies with the nick "spectator."  It connects to every room specified
// in the rooms variable.
func Connect(hostname string, c chan string) {
	//rooms := []string{"#bottesting"}

	fmt.Printf("connecting to %s\n", hostname)
	conn, err := irc.Dial(hostname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	go func() {
		for {
			msg, _ := conn.Decode()
			c <- msg.String()
		}
	}()

	identify(conn)
}

// identify() sends the USER and NICK commands to identify upon conenction.
func identify(conn *irc.Conn) {
	var messages []*irc.Message
	messages = append(messages, &irc.Message{
		Command:  irc.USER,
		Params:   []string{"spectator", "0", "*"},
		Trailing: "spectator",
	})
	messages = append(messages, &irc.Message{
		Command: irc.NICK,
		Params:  []string{"spectator"},
	})
	for _, msg := range messages {
		err := conn.Encode(msg)
		if err != nil {
			fmt.Printf("Err: %s \n%s\n", err, msg)
		}
	}
}
