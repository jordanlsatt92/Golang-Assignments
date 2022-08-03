package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"

	"github.com/gorilla/websocket"
)

// var addr = flag.String("addr", "localhost:8080", "http service address")

func main(){
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Available server ports:\n8080\n8081\n8082")
	fmt.Print("Please enter the desired port number from above: ")
	input, err := reader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}
	input = input[:len(input)-2]
	fmt.Println([]byte(input))
	u := url.URL{Scheme: "ws", Host: ":" + input, Path: "/"}
	log.Printf("Connecting to %s\n", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil{
		log.Fatal("dial:", err)
	}
	defer c.Close()
	fmt.Println("Connected...")

	done := make(chan struct{})

	go func() {
		defer close(done)
		var message string
		for {
			// _, message, err :=  c.ReadMessage()
			err := c.ReadJSON(&message)
			if err != nil {
				log.Println("read:", err)
				os.Exit(0)
			}
			log.Printf("%s", message)
		}
	}()
	
	fmt.Println("Enter your user name:")
	input, err = reader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}
	input = strings.TrimSuffix(input, "\n")
	inputChan := make(chan string, 1)
	inputChan <- string(input)

	for {
		select {
		case <-done:
			close(inputChan)
			return
		case t := <-inputChan:
			err := c.WriteJSON(t) //WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
				close(inputChan)
			}
			return
		}
		input, err = reader.ReadString('\n')
		if err != nil{
			log.Fatal(err)
		}
		inputChan <- string(input)
	}
	
}