package main

import (
	"context"
	"flag"
	log "github.com/sirupsen/logrus"
	"ms/services/idg/rpc"
	"net/http"
	"os"
)

const addr = "http://localhost:80"

func main() {
	addr := flag.String("addr", addr, "server addr")
	object := flag.String("object", "", "object name")

	flag.Parse()

	if *object == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	client := rpc.NewIDGeneratorJSONClient(*addr, &http.Client{})

	req := &rpc.Request{Object: *object}
	res, err := client.GenerateID(context.Background(), req)

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Infof("Generated ID %s:%d\n", *object, res.Id)
}
