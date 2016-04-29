# Altpoint Ventures Test App

Altpoint-Challenge is a test application for Altpoint company. 

The task:
> RPC client-server application
Create a server application, which can calculate simple math tasks. Also, implement a client which gets input via stdin and send it to the server. After the server response, show the result on stdout. Client must also support interactive input.
The server must support at least addition, subtraction, multiplication and division.
For the communication between Client und Server, you must implement the RPC Thrift Interface.
This problem should take no more than 8 hours.

and also
> The source code should be easily readable and use the godoc standard for the documentation. We care a lot about code quality and unit tests. Please add a Readme file with complete installation instructions. The Readme file should also include the system requirements.
Dependencies to other packages should be managed via godep.
You must work independent on the task.

# Intro
The Solution includes two independent apps: server and client. Both work over thrift protocol, so the configured lib is also provided. 

Basic steps how the logic works:
* Client get a raw string with some math expression and sends it to the server
* Server parse it via Shunting-Yard Algorithm(https://en.wikipedia.org/wiki/Shunting-yard_algorithm) and than calculate it via Polish Reverse Notation (https://en.wikipedia.org/wiki/Reverse_Polish_notation)
* Server sends back the result or an error

RPN algorithms were moved to independent lib (github.com/irlndts/altpoint-challenge). And it has TDD tests to show I could do it.

# Requirements
* At least Go 1.6. Lower versions may also work but it's not guaranteed.
* The GOPATH may need to be adjusted
* The list of System Requirements please find here https://golang.org/doc/install
* Godep (https://github.com/tools/godep) utility (optional)

# Server 
## Installation

The easiest way is to use go get
```
go get github.com/irlndts/altpoint-challenge/server
```

(optional) The lib contains Godeps.json file. To restore specified versions of third-party libs use
```
godep restore
```

## Compilation and start
To compile use
```
go build
or
go install
```

Use the command to check which keys are available: 
```
./server -help
```

By default 
```
./server 
```
it will start as non-framed, non-buffered web server on localhost:9090 address

# Client 
## Installation

The easiest way is to use go get
```
go get github.com/irlndts/altpoint-challenge/client
```

(optional) The lib contains Godeps.json file. To restore specified versions of third-party libs use
```
godep restore
```

## Compilation and start
To compile use
```
go build
or
go install
```

Use the command to check which keys are available: 
```
./client -help
```

By default 
```
./client 
```
it will start as non-framed, non-buffered client app which looks to localhost:9090 server-address

When it works you could see the command line.
you could use **ping** command to check if server is available
```
>ping
pong()
>
```
Use **quit** to leave the application
```
>quit
chao
#
```

use math expression with **spaces between elements** to calculate it
```
> 1 + 5 - 1 * ( 2 + -11 )
Result: 15
>
```

# What I'd love to do next
* write more user-friendly math-expression interface, like 1+2 - 3* 4. I mean not to depend on spaces.
* Cache requests [on both sides?]
* Wrap each request to go-routine on server-side
* The RPN lib is the only one where Unit-tasts are available for now. Client and Server are also need it.

# Impressions
* Thrift is new for me. I've worked with it only as a part of Scribe, but there were prepared libs.
* Godep is new for me. Very interesting, will use it in future. 
* Write (RPN-)algorithms is always interesting.

Cheers.
