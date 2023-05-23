# Think

### Problem Statement

In this challenge, you'll need to implement a globally-unique ID
generation system that runs against Maelstrom's unique-ids workload.
Your service should be totally available, meaning that it can continue
to operate even in the face of network partitions.

- Input

```
{
  "type": "generate"
}
```

- Output

```
{
  "type": "generate_ok"
  "id": 123
}
```

### Questions

- How would I generate a unique ID
    - What is an ID?
        - ID can be: string, boolean, integer, float, array
    - How do we know if they're unique?
        - An ID is unique if there exist no duplicates. The same ID
          wasn't generated before and won't be generated after.
    - How do you generate a unique ID?
        - Track all generated IDs. The database approach. Start from a
          base value and increment the value everytime you want to
          generate an ID
        - Use a unique value to generate IDs. Time is constantly
          changing and every instance of it is unique. You don't need to
          store anything, and can be relatively sure that the value
          generated is unique.

- Network Partitions
  - What is meant by network partitions
    - Network failure, hardware issue etc of some sort which means nodes
      can't talk to each other. Data inconsistencies, split brain,
      conflicts etc.
  - How does this impact the choice we make with generating unique ids?
      - The current time approach seems the easiest. Even if there are
        network partitions, the node's system clock will ensure that the
        generated ID is unique.

- How do I test maelstrom routes
  - The binaries read from STDIN. To test a route once, remvoe all `\n`
    from the json and pipe it as input to the binary


### Attempts

- Generate a unique id using:

```go
	// Get the current time in UnixNano format
	now := time.Now().UnixNano()

	// Generate a random number
	randomNum := rand.Intn(1000)

	// Combine the current time and random number
	id := strconv.FormatInt(now, 10) + strconv.Itoa(randomNum)
```
