/*
Main fuinction of server.
Server binds ports here and listens to incomming connection
*/
package main

import (
	"github.com/ololko/simple-HTTP-server/pkg/events/access"
	//"github.com/ololko/simple-HTTP-server/pkg/events/models"

	"fmt"
	"os"
	"os/signal"
	"time"
	"database/sql"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ololko/simple-HTTP-server/pkg/events/apis"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	_ "github.com/lib/pq"
)

const (
	serverPort          = ":10000"
	firestoreAccountKey = "configs/serviceAccountKey.json"
	host                = "localhost"
	portdb              = 5432
	user                = "postgres"
	dbname              = "simple-http-server"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	//psql database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, portdb, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//firestore database connection
	opt := option.WithCredentialsFile(firestoreAccountKey)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		panic(err)
	}
	defer client.Close()



	/*datAcc := &access.MockAccess{make(map[string][]models.EventT)}
	svc := apis.NewService(datAcc)*/
	datAcc := &access.PostgreAccess{Client:db}
	svc := apis.NewService(datAcc)
	/*datAcc := &access.FirestoreAccess{Client: client}
	svc := apis.NewService(datAcc)*/

	e := echo.New()
	// Middleware
	e.Use(middleware.Recover())
	e.GET("/events", svc.HandleGet)
	e.POST("/events", svc.HandlePost)

	go func() {
		if err := e.Start(serverPort); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		panic(err)
	}
}
