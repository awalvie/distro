package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var n *maelstrom.Node
var messages []float64

// TopologyHandler stores the topology of the network
func TopologyHandler(msg maelstrom.Message) error {
	// Unmarshal the message as a map
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	// Read and store node topology
	body["type"] = "topology_ok"
	delete(body, "topology")

	// Reply to the message
	return n.Reply(msg, body)
}

// ReadHandler prints all the messages that have been
// read by this node until now
func ReadHandler(msg maelstrom.Message) error {
	// Unmarshal the message as a map
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	// Return message history
	body["type"] = "read_ok"

	// Add all messages to the body
	body["messages"] = messages

	// Reply to the message
	return n.Reply(msg, body)
}

// BroadcastHandler broadcasts a message to all nodes
func BroadcastHandler(msg maelstrom.Message) error {
	// Unmarshal the message as a map
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	// Store message
	messages = append(messages, body["message"].(float64))

	// Broadcast the message to all nodes
	for _, node := range n.NodeIDs() {
		if err := n.Send(node, body); err != nil {
			return err
		}
	}

	// Set the message type to broadcast_ok
	body["type"] = "broadcast_ok"

	// Remove the message key from the body
	delete(body, "message")

	return n.Reply(msg, body)
}

func main() {
	n = maelstrom.NewNode()

	// Declare handlers for each message type
	n.Handle("broadcast", BroadcastHandler)
	n.Handle("read", ReadHandler)
	n.Handle("topology", TopologyHandler)

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
