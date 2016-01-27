package main

// go get golang.org/x/net/websocket
import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"os"
)

func main() {

	conn, err := websocket.Dial("ws://192.168.2.10:8080", "gabbo", "http://localhost/")
	checkError(err)
	var msg string
	for {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				fmt.Println(err)
				// graceful shutdown by server
				break
			}
			fmt.Println("Couldn't receive msg " + err.Error())
			break
		}
		fmt.Println("----------------------------\n")
		fmt.Println("Received from server: " + msg)
		fmt.Println("----------------------------\n")
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
