package main

import (
	log "github.com/sirupsen/logrus"
	"ms/services/idg/handler"
	"ms/services/idg/rpc"
	"ms/utils"
	"ms/utils/env"
	"net/http"
	"os"
	"time"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	addr := env.GetString("IDG_SERVER_ADDR")

	server := &http.Server{
		Handler:      rpc.NewIDGeneratorServer(&handler.Handler{}, nil),
		Addr:         *addr,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		log.Fatal(err)
	}()

	log.Info("Server listening on " + *addr)

	utils.WaitForShutdown(server)
}
