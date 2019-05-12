# Taint tracking in Golang
Project done in the Language-based Security course at Chalmers University of Technology.  
This is a dummy project for trying to implement Taint-tracking for Microservices written in Golang. We use the GRPC-toolchain to try to simulate an autonomous car. 

Remote-procedure calls will be issued to the car for increasing and decreasing the velocity of the moving car. There exists an service used for logging different types of actions that the car has been made, and thus if you want to filter what is being logged, an additional flow is inserted. 

Every action that coming from the log-service should be marked as tainted, and should not be able to reach the sinks. Test-cases is to try to inject rpc-calls or something that will make the car issue stuff like interpret "filter-input" as "increase velocity", and then try to prevent it by using the principles of taint tracking.
### Path
Setting the path correctly is quite important, do the following:

    $GOPATH/src/taint-tracking-in-golang
### Updating the protocol buffers 
    protoc -I taint-tracking/ taint-tracking-in-golang/taint-tracking.proto --go_out=plugins=grpc:taint-tracking-in-golang
### Compiling Go files
    go build

### Architecture of the system
![alt text](https://i.imgur.com/BE5K0M4.jpg)
# Useful links
### Proposal / Project Idea
- Written by Per Hallgren @ Einride
https://docs.google.com/document/d/1FiiFnHUEg-CaLzTe8mgLnbbiA1u6iNjulv9qPpXE500/edit
### Papers
- All You Ever Wanted to Know About Dynamic Taint Analysis and Forward Symbolic Execution
https://users.ece.cmu.edu/~aavgerin/papers/Oakland10.pdf
- Explicit Secrecy: A Policy for Taint Tracking
https://www.cse.chalmers.se/~andrei/eurosp16.pdf
- Information Flow Analysis for Go
https://link.springer.com/chapter/10.1007/978-3-319-47166-2_30
### Tools & Resources
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

### Concepts
- Taint checking
https://en.wikipedia.org/wiki/Taint_checking
- Non-interference (security)
https://en.wikipedia.org/wiki/Non-interference_(security)
- Information flow (information theory)
https://en.wikipedia.org/wiki/Information_flow_(information_theory)
