package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"ms/services/users/handler"
	"ms/services/users/rpc"
	"ms/utils"
	"net/http"
	"os"
	"time"
)

const addr = ":80"

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	addr := flag.String("addr", addr, "server addr")

	flag.Parse()

	server := &http.Server{
		Handler:      rpc.NewUsersServer(handler.NewHandler(), nil),
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
