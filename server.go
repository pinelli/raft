package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func CreateServer(id int, name string, host string) *Server {
	server := &Server{id, name, host, nil}
	return server
}

func (server *Server) Run(service Service) {
	server.Service = service
	service.SetServer(server)
	service.Launch()
	go server.Listen()
}

func (server *Server) Send(req Request) {
	conn, err := net.Dial("tcp", req.Server)
	if err != nil {
		//fmt.Println("Cannot send message: ", err)
		return
	}

	encoder := json.NewEncoder(conn)
	err = encoder.Encode(*req.Message)
	if err != nil {
		fmt.Println("Error encoding message")
		conn.Close()
		return
	}

	conn.Close()
}

func (server *Server) Listen() {
	ln, err := net.Listen("tcp", server.Host)
	if err != nil {
		fmt.Println("Cannot create server")
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go server.handleConnection(conn)
	}
}

func (server *Server) handleConnection(c net.Conn) {

	d := json.NewDecoder(c)
	var msg Message
	err := d.Decode(&msg)

	if err != nil {
		fmt.Println("Error decoding message")
		c.Close()
		return
	}

	var remote string = c.RemoteAddr().String()

	req := Request{&msg, remote}
	server.Service.Receive(req)

	c.Close()
}
