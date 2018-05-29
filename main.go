package main

import (
	"context"
	"flag"
	"github.com/Prots/bow-api/config"
	"github.com/Prots/bow-api/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	basePath     = "/api/v1"
	registerPath = "/register"
	recordPath   = "/record"
	displayPath  = "/display"
)

func main() {
	configFile := flag.String("config", "config.json", "Configuration file in JSON-format")
	flag.Parse()

	if len(*configFile) > 0 {
		config.ConfigFilePath = *configFile
	}

	config.Load()

	router := mux.NewRouter()
	router.HandleFunc(basePath+registerPath, services.RegisterHandler).Methods(http.MethodPost)
	router.HandleFunc(basePath+recordPath, services.RecordHandler).Methods(http.MethodPost)
	router.HandleFunc(basePath+displayPath, services.DisplayHandler).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         config.Config.HTTPListenURL,
		WriteTimeout: time.Duration(config.Config.HTTPWriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.Config.HTTPReadTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.Config.HTTPIdleTimeout) * time.Second,
		Handler:      router,
	}

	// Run server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Printf("Server running on %s", config.Config.HTTPListenURL)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.HTTPGraceTimeout)*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
