package main

type Service interface {
	Launch()
	Receive(req Request)
	Send(req Request)
	SetServer(server *Server)
}
