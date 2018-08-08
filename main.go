package main

import "os"

func main() {
	validateArgs()

	myHostPort := os.Args[1]
	networkServersFile := os.Args[2]

	var myNode *Node = CreateNode()
	myNode.LoadNetworkServers(networkServersFile)

	var server *Server = CreateServer(1, "serv1", myHostPort)
	server.Run(myNode)

	select {}
}

func validateArgs(){
}