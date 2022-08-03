package main

import (
	"fmt"
	"golang-chat/websockets"
	"net/http"
)

// Lets user connect and starts the running of the client
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websockets.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	// Run the connection with the client
	go websockets.RunClient(conn)
}

// Sets up the route and the default channel
func setupRoutes() {
	// Create a default channel that all users join when they connect
	pool := websockets.NewPool()
	pool.Name = "#default"
	websockets.Pools[pool.Name] = pool
	go pool.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r) 
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}






















// var addr = flag.String("addr", "localhost:8080", "http service address")

// var upgrader = websocket.Upgrader{} // use default options

// func home(w http.ResponseWriter, r *http.Request) {
// 	c, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Print("upgrade:", err)
// 		return
// 	}
// 	defer c.Close()
// 	for {
// 		mt, message, err := c.ReadMessage()
// 		if err != nil {
// 			log.Println("read:", err)
// 			break
// 		}
// 		log.Printf("recv: %s", message)
// 		err = c.WriteMessage(mt, message)
// 		if err != nil {
// 			log.Println("write:", err)
// 			break
// 		}
// 	}

// }

// func main() {
// 	flag.Parse()
// 	log.SetFlags(0)
// 	http.HandleFunc("/", home)
// 	log.Fatal(http.ListenAndServe(*addr, nil))
// }


// func serveWs(pool *websockets.Pool, w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("WebSocket Endpoint Hit")
// 	conn, err := websockets.Upgrade(w, r)
// 	if err != nil {
// 		fmt.Fprintf(w, "%+v\n", err)
// 	}

// 	client := &websockets.Client{
// 		Conn: conn,
// 	}
// 	client.Pool = append(client.Pool, pool)
// 	// client.Username = 
	
// 	pool.Register <- client
// 	client.Read(pool)
// }


// func setupRoutes() {
// 	pool := websockets.NewPool()
// 	go pool.Start() // Server is in listening mode

// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		serveWs(pool, w, r) // Read from client
// 	})
// }

// func main() {
// 	fmt.Println("Distributed Chat App v0.01")
// 	setupRoutes()
// 	http.ListenAndServe(":8080", nil)
// }

