package websockets

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var Users = make(map[string]*Client)

type Client struct {
	Username  string `json:"username"`
	Conn      *websocket.Conn
	Pools     map[string]*Pool
	MutexLock sync.Mutex
}

// Runs the client. Takes input from the user and routes it to the correct functions
func RunClient(conn *websocket.Conn){
	client := &Client{
		Conn: conn,
		Pools: make(map[string]*Pool),
	}

	defer func(){
		client.Conn.Close()
	}()
	// Get desired Username
	err := conn.ReadJSON(&client.Username)
	if err != nil{
		log.Fatal(err)
	}
	
	client.Username = "@" + client.Username[:len(client.Username)-1]

	//If the name is not unique, add a tag to it
	validName := false
	for !validName{
		if Users[client.Username] != nil{
			tag := (strconv.Itoa(rand.Intn(9999-1000) + 1000))
			client.Username += tag
		}else{
			validName = true
		}
	}

	// Add client to Users
	Users[client.Username] = client
	// Send the client their username
	conn.WriteJSON("From server: Your user name: " + client.Username)
	fmt.Println(client.Username, "has connected...")

	// Add client to default pool
	defaultPool := Pools["#default"]
	defaultPool.Register <- client
	client.Pools[defaultPool.Name] = defaultPool
	conn.WriteJSON("From server: You have joined the channel " + defaultPool.Name)
	// client.Read(defaultPool)
	// pool.Register <- client
	// client.Read()

	keepReading := true
	var message string

	// Read input form client
	for {
		if !keepReading{
			break
		}
		err = conn.ReadJSON(&message)
		if err != nil{
			fmt.Println(err)
			return
		}
		message = message[:len(message)-2]
		ParseCommand(message, client.Conn, client, &keepReading)
	}
}

func ParseCommand(msg string, conn *websocket.Conn, client *Client, keepReading *bool){
	str := strings.Split(msg, " ")
	commandType := str[0]
	switch commandType {
	// Send message to specific user or channel
	case "/msg":
		Message(str, client)
	// Join a channel
	case "/join":
		Join(str, client)
	// Leave a channel
	case "/leave":
		Leave(str, client)
	// Display users in channel
	case "/who":
		Who(str, client)
	// Create channel
	case "/create":
		Create(str, client)
	// Quit the chat program and disconnect from server
	case "/quit":
		Quit(str, client, keepReading)
	// Display all channels
	case "/channels":
		GetChannels(str, client)
	// Display all users
	case "/users":
		GetUsers(str, client)
	// Display list of possible commands
	case "/commands":
		GetCommands(str, client)
	default:
		conn.WriteJSON("From server: Command '" + commandType + "' unrecognized. Try another command.")
	}
}

// Message a specific client or channel
func Message(str []string, client *Client){
	if len(str) < 2{
		client.Conn.WriteJSON("You must have a destination for the message. Try '/msg #channel_name' or '/msg @user_name.'")
		return
	}
	if string(str[1][0]) == "@"{
		destination := Users[str[1]]
		if destination == nil{
			client.Conn.WriteJSON("From server: " + str[1] + " is not a current user. Use the command '/users' to see all users.")
			return
		}
		destination.Conn.WriteJSON("From " + client.Username + ": " + strings.Join(str[2:], " "))
	}else if string(str[1][0]) == "#"{
		destination := Pools[str[1]]
		if destination == nil{
			client.Conn.WriteJSON("From server: " + str[1] + " is not a current channel. Use the command '/channels' to see all available channels.")
			return
		}
		destination.Broadcast <- "From " + destination.Name + ": " + client.Username + ": " + strings.Join(str[2:], " ")
	}else{
		client.Conn.WriteJSON("From server: Unable to send message. User names begin with '@' and channel names begin with '#'.")
		return
	}
}

// Join a specific channel
func Join(str []string, client *Client){
	if len(str) < 2{
		client.Conn.WriteJSON("You must have a channel to join. Try '/join #channel_name'.")
		return
	}
	if string(str[1][0]) != "#"{
		client.Conn.WriteJSON("From server: Channel names begin with '#'. Please retry with '/join #desired_channel_name'")
		return
	}
	channel := Pools[str[1]]
	if channel == nil{
		client.Conn.WriteJSON("From server: " +str[1] + " is not a current channel. Use the command '/channels' to see all available channels.")
		return
	}
	if channel.Clients[client]{
		client.Conn.WriteJSON("You have already joined this channel.")
		return
	}
	
	channel.Register <- client
	client.Pools[channel.Name] = channel
}

// Leave a specific channel
func Leave(str []string, client *Client){
	if len(str) < 2{
		client.Conn.WriteJSON("You must have a channel to leave. Try '/leave #channel_name'.")
		return
	}
	if string(str[1][0]) != "#"{
		client.Conn.WriteJSON("From server: Channel names begin with '#'. Please retry with '/join #desired_channel_name.'")
		return
	}
	channel := Pools[str[1]]
	if channel == nil{
		client.Conn.WriteJSON("From server: " + str[1] + " is not a current channel. Use the command '/channels' to see all available channels.")
		return
	}
	if !channel.Clients[client]{
		client.Conn.WriteJSON("You are not currently in this channel. You must be in this channel to leave it.")
		return
	}
	channel.Unregister <- client
	delete(client.Pools, channel.Name)
}

// Display all users in a channel
func Who(str []string, client *Client){
	if len(str) < 2{
		client.Conn.WriteJSON("You must enter a channel to find who is in it. Try '/who #channel_name'.")
		return
	}
	channel := Pools[str[1]]
	if channel == nil {
		client.Conn.WriteJSON("From server: " + str[1] + " is not a current channel. Use the command '/channels' to see all available channels.")
		return
	}
	var users string = `From server: ` + "\n" + `Users in ` + channel.Name + `:`
	for user := range channel.Clients{
		users += "\n" + user.Username
	}
	client.Conn.WriteJSON(users)
}

// Create a channel
func Create(str []string, client *Client){
	if len(str) < 2{
		client.Conn.WriteJSON("You must enter a channel name to create it. Try '/create #channel_name'.")
		return
	}
	channel := Pools[str[1]]
	if channel != nil{
		client.Conn.WriteJSON("A channel with this name already exists. Try a different name")
		return
	}
	newPool := NewPool()
	if string(str[1][0]) != "#"{
		newPool.Name = "#" + str[1]
	}else{
		newPool.Name = str[1]
	}
	Pools[newPool.Name] = newPool
	go newPool.Start()
	newPool.Register <- client
	client.Pools[newPool.Name] = newPool
}

// Quit the chat program
func Quit(str []string, client *Client, keepReading *bool){
	for _, channel := range client.Pools{
		channel.Unregister <- client
	}
	delete(Users, client.Username)
	client.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Your connection is closing")) 
	time.Sleep(3 * time.Second)
	*keepReading = false
	client.Conn.Close()
}

// Display all channels
func GetChannels(str []string, client *Client){
	var channels string = `From server: ` + "\n" + `Channels:`
	for _, channel := range Pools{
		channels += "\n" + channel.Name  
	}
	client.Conn.WriteJSON(channels)
}

// Display all users
func GetUsers(str []string, client *Client){
	var users string = `From server: ` + "\n" + `Users:`
	for _, user := range Users{
		users += "\n" + user.Username  
	}
	client.Conn.WriteJSON(users)
}

// Display all possible commands
func GetCommands(str []string, client *Client){
	var commands string = `From Server: ` + "\n" + `Available Commands`
	commands += "\n/join <#channelname> : joins a channel"
	commands += "\n/leave <#channelname> : leaves a channel"
	commands += "\n/msg <#channelname> <any text you want to send>: sends message to specific channel"
	commands += "\n/msg <@username> <any text you want to send>: sends message to specific user"
	commands += "\n/who <#channelname> : displays the usernames of all in the channel"
	commands += "\n/create <#channelname> : creates a channel with the specified and name and joins you to the channel"
	commands += "\n/channels : displays a list of all channels"
	commands += "\n/users : displays a list of all users"
	commands += "\n/quit : removes you from all channels and closes the connection to the server."
	client.Conn.WriteJSON(commands)
}