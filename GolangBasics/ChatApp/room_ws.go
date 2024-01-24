package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type room struct{
	
	//client holds all current clients in this room
	clients map[*client]bool
	
	//joining channel for client
	join chan *client

	//leaving channel for client
	leave chan *client

	//a channel to frwrd msg to other clients
	forward chan []byte
	

}

func newRoom() *room{
	return &room{
		forward: make(chan []byte),
		join: make(chan *client),
		leave: make(chan *leave)
		clients: make(map[ *client]bool)
	}
}

func (r *room) run(){

	for{
		select{
		case client := <- r.join:
			r.clients(client)= true
		
		case client := <- r.leave:
				delete(r.clients , client)
				close(client.recieve)
		
		case msg := <- r.forward:
			for client:= range r.clients{
				client.receive <- msg
			}
		}
	}
}	

const(
	socketBufferSize =1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrade(ReadBufferSize: socketBufferSize , WriteBufferSize: messageBufferSize )

func (r *room) ServerHTTP(w http.ResponseWriter , req *http.Request)
{
	socket , err:= upgrader.Upgrade(w ,req, nil)
	if err !=nil {
		log.Fatal("ServerHTTP: ",err)
		return
	}

	client := &client{
		socket: socket,
		recieve: make(chan []byte , messageBufferSize),
		room: r,
	}
	r.join <- client
	defer func(){r.leave <- client} {}
	go client.write(){}
	client.read()
}