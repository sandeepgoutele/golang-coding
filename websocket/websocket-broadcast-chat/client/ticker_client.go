package main

/*
func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	serverUrl := "ws://localhost:8081/ws"
	url, err := url.Parse(serverUrl)
	if err != nil {
		log.Fatal("Error parsing server URL: ", err)
	}

	log.Printf("Connecting to %s", url.String())
	c, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Fatal("Dial error: ", err)
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Read error: ", err)
				return
			}
			log.Printf("Received message: %s", message)
		}
	}()

	// ticker := time.NewTicker(2 * time.Second)
	// defer ticker.Stop()
	msg := make(chan string)
	defer close(msg)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-time.After(2 * time.Second):
				msg <- "Hello from client"
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case t := <-msg:
			err := c.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println("Write error: ", err)
				return
			}
		case <-interrupt:
			log.Println("Interrupt signal received")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Write close error: ", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
*/
