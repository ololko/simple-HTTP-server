/*
Main fuinction of server.
Server binds ports here and listens to incomming connection
*/
package main

import (
	//"cloud.google.com/go/profiler"
	"github.com/jinzhu/gorm"
	"github.com/ololko/simple-HTTP-server/pkg/events/access"
	"github.com/ololko/simple-HTTP-server/pkg/events/apis"
	gw "github.com/ololko/simple-HTTP-server/pkg/events/models"
	GrpcServer "github.com/ololko/simple-HTTP-server/pkg/grpc"
	HttpServer "github.com/ololko/simple-HTTP-server/pkg/http"
	myViper "github.com/ololko/simple-HTTP-server/pkg/viper"
	//"runtime/pprof"

	firebase "firebase.google.com/go"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	/*	if err := profiler.Start(profiler.Config{
			Service:        "server",
			ServiceVersion: "1.0",
			ProjectID:      "psychic-heading-244314", // optional on GCP
		}); err != nil {
			log.Fatalf("Cannot start the profiler: %v", err)
		}*/

	err := myViper.ReadConfig("viperConfig", "./configs/")
	if err != nil {
		fmt.Println(err)
		return
	}

	//obsolete, I use google cloud platform
	/*	f, err := os.Create("/home/luppolo/statsGoMain")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()*/

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

	//db.DropTableIfExists(&gw.DatabaseElement{})
	db.AutoMigrate(&gw.DatabaseElement{})

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

	/*	datAcc := access.MockAccess{map[string][]gw.DatabaseElement{
			"Daco": {
				gw.DatabaseElement{Type: "Daco", Count: 55, Timestamp: 55},
			},
		}}

		svc := apis.NewService(&datAcc)*/

	datAcc := &access.PostgreAccess{Client: db}
	svc := apis.NewService(datAcc)

	/*datAcc := &access.FirestoreAccess{Client: client}
	svc := apis.NewService(datAcc)*/

	ctx := context.Background()
	// run HTTP gateway
	go func() {
		err = HttpServer.RunServerHTTP(ctx, viper.GetString("grpcPort"), viper.GetString("serverPort"))
		if err != nil {
			fmt.Println("Failed to start HTTP server: ", err)
		}
	}()

	err = GrpcServer.RunGrpcServer(ctx, svc, viper.GetString("grpcPort"))
	if err != nil {
		fmt.Println("Error creating GRPC server")
	}
}
