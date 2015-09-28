package main

import stomp "github.com/go-stomp/stomp"
import "fmt"

func main() {
	conn, err := stomp.Dial("tcp", "localhost:61613")
	if err != nil {
		fmt.Println(err)
	}
	err = conn.Send(
		"/queue/test-1", // destination
		"text/plain", // content-type
		[]byte("Test message #1")) // body
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Golang + AMQ = Rocks")

	sub, err := conn.Subscribe("/queue/test-1", stomp.AckClient)
	if err != nil {
		fmt.Println(err)
	}
	for  {
		msg := <-sub.C
		fmt.Println(msg)

		// acknowledge the message
		err = conn.Ack(msg)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = sub.Unsubscribe()
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Disconnect()

}