package main

type Block struct {
	PrevHash string
	CurHash  string
	Data     []string
}

type Message struct {
	Len    int
	Sender string
	Mcode  string
	Blocks []Block
}

/*Info contains data the nodes send to each other
 */
type Info struct {
	Msg      string
	ChainLen int
	Data     []Block
}

type Server struct {
	Id      int
	Name    string
	Host    string
	Service Service
}

type Request struct {
	Message *Message
	Server  string
}

type Raft struct {
}
