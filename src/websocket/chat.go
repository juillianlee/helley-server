package websocket

// var upgrader = websocket.Upgrader{}
// var messages []string

// func getCmd(input string) string {
// 	inputArr := strings.Split(input, " ")
// 	return inputArr[0]
// }

// func getMessage(input string) string {
// 	inputArr := strings.Split(input, " ")
// 	var result string
// 	for i := 1; i < len(inputArr); i++ {
// 		result += inputArr[i]
// 	}
// 	return result
// }

// func updateTodoList(input string) {
// 	tmpList := messages
// 	messages = []string{}
// 	for _, val := range tmpList {
// 		if val == input {
// 			continue
// 		}
// 		messages = append(messages, val)
// 	}
// }

// type ChatWebsocket interface {
// 	Chat(http.ResponseWriter, *http.Request)
// 	ChatView(http.ResponseWriter, *http.Request)
// }

// type chatWebsocket struct {
// }

// func NewChatWebsocket() ChatWebsocket {
// 	return &chatWebsocket{}
// }

// func (c *chatWebsocket) Chat(w http.ResponseWriter, r *http.Request) {

// 	conn, err := upgrader.Upgrade(w, r, nil)

// 	if err != nil {
// 		log.Print("upgrade fail: ", err)
// 		return
// 	}

// 	defer conn.Close()

// 	for {
// 		mt, message, err := conn.ReadMessage()

// 		fmt.Println(mt)

// 		if err != nil {
// 			log.Println("read failed: ", err)
// 		}

// 		input := string(message)
// 		cmd := getCmd(input)
// 		msg := getMessage(input)

// 		if cmd == "add" {
// 			messages = append(messages, msg)
// 		} else if cmd == "done" {
// 			updateTodoList(msg)
// 		}

// 		output := "Current Messages: \n"

// 		for _, message := range messages {
// 			output += "\n - " + message + "\n"
// 		}

// 		output += "\n----------------------------------------"

// 		message = []byte(output)

// 		err = conn.WriteMessage(mt, message)

// 		if err != nil {
// 			log.Println("write failed: ", err)
// 			break
// 		}
// 	}
// }

// func (c *chatWebsocket) ChatView(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "websockets.html")
// }
