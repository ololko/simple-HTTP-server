/*
Main fuinction of server.
Server binds ports here and listens to incomming connection
*/
package main

import (
	"github.com/ololko/simple-HTTP-server/pkg/events/access"
	//"github.com/ololko/simple-HTTP-server/pkg/events/access"
	"github.com/ololko/simple-HTTP-server/pkg/events/apis"
	gw "github.com/ololko/simple-HTTP-server/pkg/events/models"
	myViper "github.com/ololko/simple-HTTP-server/pkg/viper"
	"net"

	"flag"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of YourService")
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
	if err != nil {
		err = db.DB().Ping()
		panic(err)
	}

	//db.DropTableIfExists(&gw.EventT{})
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

	datAcc := access.MockAccess{map[string][]gw.DatabaseElement{
		"Daco":{
			gw.DatabaseElement{Type:"Daco",Count:55,Timestamp:55},
		},
	}}
	//svc := apis.NewService(datAcc)

	//datAcc := &access.PostgreAccess{Client: db}
	//svc := apis.NewService(datAcc)

	/*datAcc := &access.FirestoreAccess{Client: client}
	svc := apis.NewService(datAcc)*/

	flag.Parse()
	lis, err := net.Listen("tcp", viper.GetString("serverPort"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	gw.RegisterEventsServer(grpcServer, &apis.Service{datAcc})
	err = grpcServer.Serve(lis)
	if err != nil {
		fmt.Println(err)
	}
}

/*func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterEventsHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(viper.GetString("serverPort"), mux)
}*/
