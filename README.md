# Taint tracking in Golang
The path should be $GOPATH/src/taint-tracking-in-golang for this to work, the grpc librarie is included aswell but should perhaps not be ¯\_(ツ)_/¯

### Updating the protocol buffers 
    protoc -I taint-tracking/ taint-tracking-in-golang/taint-tracking.proto --go_out=plugins=grpc:taint-tracking-in-golang
### Compiling Go files
    go build
    
### USEFUL STUFF
# Proposal / Project Idea
- Written by Per Hallgren @ Einride
https://docs.google.com/document/d/1FiiFnHUEg-CaLzTe8mgLnbbiA1u6iNjulv9qPpXE500/edit
# Papers
- All You Ever Wanted to Know About Dynamic Taint Analysis and Forward Symbolic Execution
https://users.ece.cmu.edu/~aavgerin/papers/Oakland10.pdf
- Explicit Secrecy: A Policy for Taint Tracking
https://www.cse.chalmers.se/~andrei/eurosp16.pdf
- Information Flow Analysis for Go
https://link.springer.com/chapter/10.1007/978-3-319-47166-2_30
# Tools & Resources
- Protocol Buffers
https://developers.google.com/protocol-buffers/docs/overview
- GRPC
https://grpc.io/docs/guides/
- Golang
https://golang.org/doc/
- Learning Go in 12 minute
https://www.youtube.com/watch?v=C8LgvuEBraI&t=19s
- Concurrency tutorial in Go
https://www.youtube.com/watch?v=LvgVSSpwND8&t=4s
- Static Analysis in Go
https://www.slideshare.net/takuyaueda967/static-analysis-in-go

# Concepts
- Taint checking
https://en.wikipedia.org/wiki/Taint_checking
- Non-interference (security)
https://en.wikipedia.org/wiki/Non-interference_(security)
- Information flow (information theory)
https://en.wikipedia.org/wiki/Information_flow_(information_theory)

# Video that is relevant when 1 may hits
https://www.youtube.com/watch?v=K8M7fFwdM9w
