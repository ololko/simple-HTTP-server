FROM golang:1.12 as Builder

RUN ["go", "get", "github.com/cespare/reflex"]
WORKDIR /app
COPY . .
RUN go get ./...
ENTRYPOINT [ "reflex" ,"-c" ,"build/server/reflex.conf", "--decoration=fancy" ]