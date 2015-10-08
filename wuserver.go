package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"math/rand"
	"time"
)

func main() {
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Connect("tcp://127.0.0.1:5556")

	var applications []string

	applications = append(applications, "rfgateway")
	applications = append(applications, "gsmdaemon")
	applications = append(applications, "gwagent")

	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// loop for a while aparently
	for {
		// Get values that will fool the boss
		zipcode := rand.Intn(100000)
		temperature := rand.Intn(215) - 80
		relhumidity := rand.Intn(50) + 10

		application := applications[rand.Intn(len(applications))]

		// Send message to all subscribers
		//msg := fmt.Sprintf("%05d %d %d", zipcode, temperature, relhumidity)
		msg := fmt.Sprintf("%s{zipcode=%d&temperature=%d&relhumidity=%d}", application, zipcode, temperature, relhumidity)
		fmt.Printf("Sending %s\n", msg)
		publisher.Send(msg, 0)
		time.Sleep(5 * time.Duration(relhumidity) * time.Millisecond)
	}
}
