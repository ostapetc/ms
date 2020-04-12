package main

import (
	"context"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"ms/services/users/rpc"
	"net/http"
	"os"
)

const addr = "http://localhost:80"

func main() {
	addr := flag.String("addr", addr, "server addr")
	username := flag.String("username", "", "username")
	password := flag.String("password", "", "password")

	flag.Parse()

	if *username == "" || *password == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("addr", *addr)

	client := rpc.NewUsersJSONClient(*addr, &http.Client{})

	req := &rpc.RegisterRequest{Username: *username, Password: *password}
	res, err := client.Register(context.Background(), req)

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Println(res)
}
