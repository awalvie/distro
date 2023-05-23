package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var n *maelstrom.Node

// BroadcastHandler broadcasts a message to all nodes
func BroadcastHandler(msg maelstrom.Message) error {
	// Unmarshal the message as a map
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	// Update the message type to return back
	body["type"] = "broadcast_ok"

	// Echo the original message back with updated type
	return n.Reply(msg, body)
}

func main() {
	n = maelstrom.NewNode()

	// Handle the "echo" message type
	n.Handle("broadcast", BroadcastHandler)

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
