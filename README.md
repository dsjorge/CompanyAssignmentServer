### Requirements

- Go version >= 1.6.2
- SQLite version >= 3.11.0
 
### Startup

Set environmental variables:
```
~$ export GOPATH=/clone/project/folder
~$ export PATH=$PATH:$GOPATH/bin
~$ export GOBIN=$GOPATH/bin
```

Import external Go packages
```
server:~$ go get github.com/mattn/go-sqlite3
server:~$ go get github.com/julienschmidt/httprouter
```

For demo purposes to allow Cross-Origin Resource Sharing (CORS), replace the file 'src/github.com/julienschmidt/httprouter/router.go' with the one with the same name in the root of this repository (cut and paste)

Build server and dependencies
```
server:~$ go install src/company/server/server.go
```

Run application
```
server:~$ bin/server
```
