## raft-blockchain
My own implementation of the Raft consensus algorithm.
For now nodes can elect a leader and receive empty heartbeats from it.

### Description of the raft algorithm:

https://en.wikipedia.org/wiki/Raft_(computer_science)

https://raft.github.io/

http://thesecretlivesofdata.com/raft/


### What is not implemented yet
By the Raft specification, a node becomes a leader when it gets votes from majority.
For now, after the node announces an election, it does not count the number of votes it receives and thinks that it's a leader when it receives at least one vote.

### Usage
#### -Build a node
go build
#### -Run a node
./raft-blockchain \<host:port\> <file_with_other_nodes_in_the_network>

### To quickly run 3 nodes:

cd config

./runNode1.sh

Open 2 more terminals and run ./runNode2.sh and ./runNode3.sh respectively
