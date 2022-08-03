package websockets

import "fmt"

type Pool struct {
	Name       string `json:"poolname"`
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool // all clients in channel
	Broadcast  chan string
}

// returns a new instance of a channel
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan string),
	}
}

var Pools = make(map[string]*Pool)

// Starts the pool and allows for new clients to join, leave, and send messages
func (pool *Pool) Start() {
	for {
		select {
		case newClient := <-pool.Register:
			pool.Clients[newClient] = true
			fmt.Println("Size of", pool.Name, ":", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)

				client.Conn.WriteJSON("From " + pool.Name + ": " + newClient.Username + " has joined the channel...")
			}
		case departingClient := <-pool.Unregister:
			delete(pool.Clients, departingClient)
			fmt.Println("Size of", pool.Name, ":", len(pool.Clients), " ", departingClient.Username, "disconnected...")
			for client := range pool.Clients {
				client.Conn.WriteJSON("From " + pool.Name + ": " + departingClient.Username + " has left the channel...")
			}
			if pool.Name != "#default" && len(pool.Clients) == 0{
				delete(Pools, pool.Name)
				fmt.Println("All users have left", pool.Name)
				fmt.Println("Removing channel:", pool.Name)
				return
			}
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in pool")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}