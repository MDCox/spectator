// Handles irc connections and logs to be sent to the parser
package irc

import (
	"fmt"
	"github.com/sorcix/irc"
	"log"
	"time"
)

// Connect() takes a string with url:port of the irc server, and a channel
// that accepts strings.  It connects to the irc server specified and
// identifies with the nick "spectator."  It connects to every room specified
// in the rooms variable.
func Connect(hostname string, c chan string) {
	rooms := []string{
		"#Node.js", "##javascript",
		"#python", "#haskell", "#vim", "#go-nuts", "#ruby",
		"#clojure", "#perl", "##php",
		"#erlang", "#scheme", "#lisp", "#R", "#swift-lang",
	}

	fmt.Printf("connecting to %s\n", hostname)
	conn, err := irc.Dial(hostname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	go func() {
		for {
			msg, _ := conn.Decode()
			if msg.Command == "PING" {
				conn.Encode(&irc.Message{
					Command:  "PONG",
					Params:   msg.Params,
					Trailing: msg.Trailing,
				})
			}
			c <- msg.String()
		}
	}()

	identify(conn)
	fmt.Println("Identified")

	joinRooms(conn, rooms)
	fmt.Printf("Join: %s\n", rooms)
}

// identify() sends the USER and NICK commands to identify upon conenction.
func identify(conn *irc.Conn) {
	var messages []*irc.Message
	messages = append(messages, &irc.Message{
		Command:  irc.USER,
		Params:   []string{"testin", "0", "*"},
		Trailing: "testin",
	})
	messages = append(messages, &irc.Message{
		Command: irc.NICK,
		Params:  []string{"testin"},
	})
	for _, msg := range messages {
		err := conn.Encode(msg)
		if err != nil {
			fmt.Printf("Err: %s \n%s\n", err, msg)
		}
	}
}

// joinRooms() joins any irc channel that is included in the `rooms` slice.
func joinRooms(conn *irc.Conn, rooms []string) {
	for _, room := range rooms {
		time.Sleep(time.Second * 2)
		conn.Encode(&irc.Message{
			Command: irc.JOIN,
			Params:  []string{room},
		})
	}
}
