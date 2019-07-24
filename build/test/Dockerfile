FROM golang:1.12

WORKDIR /app
RUN ["go", "get", "github.com/cespare/reflex"]

CMD [ "reflex" ,"-c" ,"build/server/reflex.conf", "--decoration=fancy" ]
