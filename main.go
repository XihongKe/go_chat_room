package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
var upgrader = websocket.Upgrader{}

func main(){
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}

func echo( w http.ResponseWriter, r *http.Request){
	fmt.Println(upgrader)
	c, _ := upgrader.Upgrade(w, r, nil)
	fmt.Println(1111)
	defer c.Close()
	fmt.Println(2222)
	for {
		mt, message, _ := c.ReadMessage()
		fmt.Println(message)
		c.WriteMessage(mt, append([]byte("hello "),message[:]...))
	}
}