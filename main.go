package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KyleWS/motivationapp/handlers"
	"github.com/KyleWS/motivationapp/models/subscribe"
	mgo "gopkg.in/mgo.v2"
)

const defaultAddr = "localhost:443"

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = defaultAddr
	}

	tlsKeypath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")

	mongoAddr := os.Getenv("MONGO_ADDR")
	//default to "localhost"
	if len(mongoAddr) == 0 {
		mongoAddr = "localhost:27017"
	}
	//dial the MongoDB server
	mongoSess, err := mgo.Dial(mongoAddr)
	if err != nil {
		log.Fatalf("error dialing mongo: %v", err)
	}
	mongoStore := subscribe.NewMongoStore(mongoSess, "core", "subscriber")
	handlerContext := handlers.NewHandlerContext(*mongoStore)


	mux := http.NewServeMux()

	mux.HandleFunc("/test", handlerContext.SubHandler)
	fmt.Printf("server is listening at http://%s...\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeypath, mux))
}
