/*
Main fuinction of server.
Server binds ports here and listens to incomming connection
*/
package main

import (
	"github.com/ololko/simple-HTTP-server/pkg/events/access"
	"github.com/ololko/simple-HTTP-server/pkg/events/models"
	myViper "github.com/ololko/simple-HTTP-server/pkg/viper"

	"fmt"
	"os"
	"os/signal"
	"time"

	firebase "firebase.google.com/go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/ololko/simple-HTTP-server/pkg/events/apis"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	err := myViper.ReadConfig("viperConfig", "./configs/")
	if err != nil {
		fmt.Println(err)
		return
	}

	//psql database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", viper.GetString("host"), viper.GetInt("dbPort"), viper.GetString("user"), viper.GetString("dbname"))
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	//db.DropTableIfExists(&models.EventT{})
	db.AutoMigrate(&models.EventT{})

	//firestore database connection
	opt := option.WithCredentialsFile(viper.GetString("firestoreAccountKey"))
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
	datAcc := &access.PostgreAccess{Client: db}
	svc := apis.NewService(datAcc)
	/*datAcc := &access.FirestoreAccess{Client: client}
	svc := apis.NewService(datAcc)*/






	e := echo.New()
	// Middleware
	e.Use(middleware.Recover())
	e.GET("/events", svc.HandleGet)
	e.POST("/events", svc.HandlePost)

	go func() {
		if err := e.Start(viper.GetString("serverPort")); err != nil {
			e.Logger.Info("FAIL IN BINDING PORT")
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
