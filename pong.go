package main

import (
	"flag"
	"fmt"
    "net/http"
    "log"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8080", "Ping service address")

var wsupgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("received a message: %v+", msg)
		conn.WriteMessage(t, msg)
	}
}

func main() {
	fmt.Println("Pong has started")
	flag.Parse()
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
