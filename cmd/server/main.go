/*
Main fuinction of server.
Server binds ports here and listens to incomming connection
*/
package main

import (
	"fmt"
	"github.com/ololko/simple-HTTP-server/pkg/events/access"
	"log"
	//"net/http"
	"os"
	"os/signal"
	"time"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ololko/simple-HTTP-server/pkg/events/apis"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

const(
	port = ":10000"
	path = "serviceAccountKey.json"
)

func main() {
	opt := option.WithCredentialsFile(path)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	/*datAcc := &access.MockAccess{make(map[string][]models.EventT)}
	svc := apis.NewService(datAcc)*/
	datAcc := &access.FirestoreAccess{Client: client}
	svc := apis.NewService(datAcc)

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.GET("/events", svc.HandleGet)
	e.POST("/events", svc.HandlePost)

	go func () {
		if err := e.Start(":10000"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	stop  := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<- stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		panic(err)
	}
}
