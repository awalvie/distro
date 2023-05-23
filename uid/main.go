package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func GenerateUID() string {
	// Get the current time in UnixNano format
	now := time.Now().UnixNano()

	// Generate a random number
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1000)

	// Combine the current time and random number
	id := strconv.FormatInt(now, 10) + strconv.Itoa(randomNum)

	return id
}

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		// Unmarshal the message as a map
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Generate a new UID
		body["id"] = GenerateUID()

		// Update the message type to return back
		body["type"] = "generate_ok"

		// Echo the original message back with updated type
		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
