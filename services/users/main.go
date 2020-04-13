package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"ms/services/users/handler"
	"ms/services/users/rpc"
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
	addr := env.GetString("USERS_SERVER_ADDR")
	idgAddr := env.GetString("IDG_SERVER_ADDR")

	fmt.Println("listen to", *addr)

	server := &http.Server{
		Handler:      rpc.NewUsersServer(handler.NewHandler(*idgAddr), nil),
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
