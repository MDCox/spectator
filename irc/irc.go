// Handles irc connections and logs to be sent to the parser
package irc

import (
	"fmt"
	"github.com/sorcix/irc"
	"log"
)

func Connect(hostname string) {
	//channels := []string{"#bottesting"}

	fmt.Printf("connection to %s\n", hostname)
	c, err := irc.Dial(hostname)
	if err != nil {
		log.Fatal(err)
	}
	msg, _ := c.Decode()
	fmt.Println(msg)
}
