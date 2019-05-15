port=10000
path=/home/luppolo/Desktop/configs/serviceAccountKey.json
tag=finalrev

all:
	mkdir -p ../configs
	cp -i ${path} ../configs/serviceAccountKey.json
	docker build -t goimg:${tag} -f Dockerfile ../.
	docker run -it -d -p ${port}:10000 goimg:${tag}
run:
	docker run -it -d -p ${port}:10000 goimg:${tag}
clean:
	docker stop $$(docker ps -q)
	docker rm $$(docker ps -a -q)
