package main

import (
	"context"
	"flag"
	"github.com/ostapetc/ms/tree/master/users/rpc"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const host = "http://localhost:80"

func main() {
	serverHost := flag.String("host", host, "server host")
	username := flag.String("username", "", "username")
	password := flag.String("password", "", "password")

	flag.Parse()

	req := &rpc.RegisterRequest{Username: *username, Password: *password}

	client := rpc.NewUsersJSONClient(*serverHost, &http.Client{})
	res, err := client.Register(context.Background(), req)

	if err != nil {
		log.Error(err)
	}

	log.Println(res)
}
