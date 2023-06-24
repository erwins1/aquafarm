package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	var (
		resources  *resourceHTTP
		err        error
		httpRouter *httprouter.Router
	)

	httpRouter = httprouter.New()

	err = godotenv.Load()
	if err != nil {
		log.Fatal("cant read environment files")
	}

	// init resource
	resources = initResourceHTTP()
	// init app
	InitApp(httpRouter, resources)

	// serving http
	s := &http.Server{
		Addr:              os.Getenv("HTTP_HOST"),
		Handler:           httpRouter,
		ReadTimeout:       1 * time.Minute,
		ReadHeaderTimeout: 1 * time.Minute,
	}
	log.Printf("Serving Http on %s", os.Getenv("HTTP_HOST"))
	if err = s.ListenAndServe(); err != nil {
		log.Fatal("fail to serve http, err: ", err)
	}
}
