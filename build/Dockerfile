FROM golang:1.12 as Builder

COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /build/app cmd/server/main.go

FROM alpine
EXPOSE 10000
RUN apk add --update --no-cache ca-certificates git
COPY --from=Builder /build/app /build/app
COPY --from=Builder /app/configs/serviceAccountKey.json /build/serviceAccountKey.json
RUN ["chmod", "+x", "/build/app"]
WORKDIR /build
ENTRYPOINT ["/build/app"]
