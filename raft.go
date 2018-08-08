package main

import (
	"fmt"
	"time"
)

const TIME = 5

func (node *Node) Raft() {
	go node.follower()
}

func (node *Node) follower() {
	fmt.Println("I'm a FOLLOWER")
	timer := time.NewTimer(TIME * time.Second).C
	for {
		select {
		case <-timer:
			go node.candidate()
			return
		default:
			cmd := processFollower(node)
			switch cmd {
			case "kill":
				return
			case "updateTimer":
				timer = time.NewTimer(TIME * time.Second).C
			case "nop":
			}
		}
	}
}

func processFollower(node *Node) string {
	select {
	case req := <-node.InputReq:
		if req.Message.Mcode == "elect" {
			msg := Message{Mcode: "vote"}
			node.Reply(req, &msg)
			go node.voted()
			return "kill"
		} else if req.Message.Mcode == "hb" {
			processHB(node, req.Message)
			return "updateTimer"
		}
		return "nop"

	default:
		return "nop"
	}
}

func processHB(node *Node, msg *Message) {
	fmt.Println("I got heartbeat from", msg.Sender)
}

func (node *Node) candidate() {
	fmt.Println("I'm a CANDIDATE")
	timer := time.NewTimer(TIME * time.Second).C
	for {
		select {
		case <-timer:
			go node.follower()
			return
		default:
			cmd := processCandidate(node)
			switch cmd {
			case "kill":
				return
			case "updateTimer":
				timer = time.NewTimer(TIME * time.Second).C
			case "nop":
			}
		}
	}
}

func processCandidate(node *Node) string {
	select {
	case req := <-node.InputReq:
		if req.Message.Mcode == "vote" {
			go node.master()
			return "kill"
		} else if req.Message.Mcode == "hb" {
			processHB(node, req.Message)
			go node.follower()
			return "kill"
		}
		return "nop"

	default:
		msg := Message{Mcode: "elect"}
		node.SendAll(&msg)
		return "nop"
	}
}
func (node *Node) voted() {
	fmt.Println("I'm a VOTED")
	timer := time.NewTimer(TIME * time.Second).C
	for {
		select {
		case <-timer:
			go node.follower()
			return
		default:
			cmd := processVoted(node)
			switch cmd {
			case "kill":
				return
			case "updateTimer":
				timer = time.NewTimer(TIME * time.Second).C
			case "nop":
			}
		}
	}
}

func processVoted(node *Node) string {
	select {
	case req := <-node.InputReq:
		if req.Message.Mcode == "hb" {
			processHB(node, req.Message)
			return "updateTimer"
		}
		return "nop"

	default:
		return "nop"
	}
}

func (node *Node) master() {
	fmt.Println("I'm a MASTER")
	timer := time.NewTimer(1 * time.Second).C
	for {
		select {
		case <-timer:
			sendHeartbeat(node)
			timer = time.NewTimer(1 * time.Second).C
		default:
			cmd := processMaster(node)
			switch cmd {
			case "kill":
				return
			case "nop":
			}
		}
	}
}

func processMaster(node *Node) string {
	select {
	case req := <-node.InputReq:
		if req.Message.Mcode == "hbResponse" {
			return "nop"
		} else if req.Message.Mcode == "hb" {
			msg := Message{Mcode: "elect"}
			node.Reply(req, &msg)
			go node.voted()
			return "kill"

		} else if req.Message.Mcode == "elect" {
			msg := Message{Mcode: "vote"}
			node.Reply(req, &msg)
			go node.voted()
			return "kill"
		}
		return "nop"

	default:
		return "nop"
	}
}

func sendHeartbeat(node *Node) {
	msg := Message{Mcode: "hb"}
	node.SendAll(&msg)
}
