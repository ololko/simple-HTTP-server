cmd="go get firebase.google.com/go"

all:
	eval ${cmd}
	go run server.go
