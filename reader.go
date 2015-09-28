package main

import stomp "github.com/go-stomp/stomp"
import "fmt"

//Connect to ActiveMQ and listen for messages
func main() {
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		fmt.Println(err)
	}

	sub, err := conn.Subscribe("/queue/test-1", stomp.AckAuto)
	if err != nil {
		fmt.Println(err)
	}
	for {
		msg := <-sub.C
		fmt.Println(msg)
	}

	err = sub.Unsubscribe()
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Disconnect()
}
