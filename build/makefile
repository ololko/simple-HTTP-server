cmd="go get firebase.google.com/go"

port=10000
path=../configs/serviceAccountKey.json
tag=make

all:
	cp -p ${path} ../configs/serviceAccountKey.json
	docker build -t goimg:${tag} -f dockerfile ../.
	docker run -it -d -p ${port}:10000 goimg:${tag}
