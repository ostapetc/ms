package main

import (
	"context"
	"fmt"
	"github.com/ostapetc/ms/tree/master/idgenerator/rpc"
	"net/http"
	"os"
)

func main() {
	client := rpc.NewIDGeneratorJSONClient("http://localhost:81", &http.Client{})

	objects := []string{
		"users",
		"posts",
	}

	for _, object := range objects {
		res, err := client.GenerateID(context.Background(), &rpc.Request{Object: object})

		if err != nil {
			fmt.Printf("oh no: %v", err)
			os.Exit(1)
		}

		fmt.Printf("%s %+v\n", object, res)
	}
}
