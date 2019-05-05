cmd="go get firebase.google.com/go"

port=10000
path=../configs/serviceAccountKey.json

all:
	#eval ${cmd}
	go run ../cmd/server/main.go ${port} ${path}
	#env || export GOBIN=/home/luppolo/skuska
	#go install ../cmd/server/main.go
	#sudo docker run -v/home/luppolo/others/fu/simple-HTTP-server/cmd/server/main.go:/server.go -v /home/luppolo/Desktop/configs/serviceAccountKey.json:/credentials.json -p ${port}:10000 -d -it newgo:v1 "source /root/.profile && go install server.go && ./server 10000 credentials.json"
