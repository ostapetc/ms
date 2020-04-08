package main

import (
	"context"
	"github.com/ostapetc/ms/tree/master/users/handler"
	"github.com/ostapetc/ms/tree/master/users/rpc"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	server := &http.Server{
		Handler:      rpc.NewUsersServer(&handler.Handler{}, nil),
		Addr:         ":80",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		log.Fatal(err)
	}()

	log.Info("Server started")

	// Graceful Shutdown
	waitForShutdown(server)
}

func waitForShutdown(server *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Shutting down")
	os.Exit(0)
}
