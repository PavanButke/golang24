package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"os/signal"
	"time"
)

type client struct{
	//web socket
	socket *websocket.Conn
	//channel to recieve msg from others 
	receive chan []byte
	//chat room
	room *room
}

func (c* client) read(){
	//close if we have read
	defer c.socket.Close()

		for {
			_, msg, err := s.socket.ReadMessage()
			if err != nil{
				return
			}
			//msg recieving from channel
			c.room.forward <- msg
		}

}

func (c* client) write(){
	//close if we have read
	defer c.socket.Close()

		for {
			_, msg, err := s.socket.ReadMessage()
			if err != nil{
				return
			}
			//msg recieving from channel
			c.room.forward <- msg
		}

}

