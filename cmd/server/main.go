/*
Main fuinction of server.
Server binds ports here and listens to incomming connection
*/
package main

import (
	"fmt"
	"github.com/ololko/simple-HTTP-server/pkg/events/access"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"

	//"log"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
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

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	/*f, err := os.OpenFile("log.txt", os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
		log.SetOutput(os.Stdout)
	}else{
		log.SetOutput(f)
	}*/
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

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

	datAcc := &access.MockAccess{make(map[string][]models.EventT)}
	svc := apis.NewService(datAcc)
	/*datAcc := &access.FirestoreAccess{Client: client}
	svc := apis.NewService(datAcc)*/

	e := echo.New()
	// Middleware
	e.Use(middleware.Recover())
	e.GET("/events", svc.HandleGet)
	e.POST("/events", svc.HandlePost)

	go func () {
		if err := e.Start(port); err != nil {
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
