FROM golang:1.12 as Builder

COPY . /app
WORKDIR /app
RUN go build -o /build/app .
