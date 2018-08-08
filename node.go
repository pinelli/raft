package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//implements Service
type Node struct {
	Server         *Server
	InputReq       chan Request
	NetworkServers []string
}

func CreateNode() *Node {
	return &Node{}
}

func (node *Node) LoadNetworkServers(filePath string) {
	file, err :=
		ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	json.Unmarshal(file, &node.NetworkServers)
	fmt.Printf("Loaded network nodes: %v\n", node.NetworkServers)
}

func (node *Node) Launch() {
	node.InputReq = make(chan Request)
	go node.Raft()

	fmt.Println("NODE LAUNCHED")
}

func (node *Node) Receive(req Request) {
	//	fmt.Println("RECEIVED from ", req.Server)
	//	fmt.Println(*req.Message)
	go func() {
		node.InputReq <- req //save request as input to node
	}()
}

func (node *Node) Send(req Request) {
	req.Message.Sender = node.Server.Host
	node.Server.Send(req)
}

func (node *Node) SendAll(msg *Message) {
	for _, serv := range node.NetworkServers {
		node.Send(Request{msg, serv})
	}
}

func (node *Node) Reply(req Request, msg *Message) {
	node.Send(Request{msg, req.Message.Sender})
}

func (node *Node) SetServer(server *Server) {
	node.Server = server
}
