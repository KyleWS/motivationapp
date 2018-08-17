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

const defaultAddr = ":443"

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = defaultAddr
	}

	tlsKeypath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")

	mongoAddr := os.Getenv("DATABASE_ADDRESS")
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
	mux.HandleFunc("/", handlerContext.SubHandler)
	dir := http.Dir("/images")
	fs := http.FileServer(dir)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	corsHandler := handlers.NewCORS(mux)
	fmt.Printf("server is listening at http://%s...\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeypath, corsHandler))
}
