cmd="go get firebase.google.com/go"

port=10000
path=/home/luppolo/Desktop/configs/serviceAccountKey.json

all:
	eval ${cmd}
	go run ../pkg/main/server.go ${port} ${path}
	#export GOBIN=/home/luppolo/skuska
	#go install ../pkg/main/server.go
	#sudo docker  docker run -v /home/luppolo/skuska:/server -v ${path}:/credentials --detach  ubuntu /server ${port} credentials
