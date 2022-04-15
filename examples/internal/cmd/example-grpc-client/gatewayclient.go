package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/jsonpb"

	examplepb "github.com/binchencoder/ease-gateway/examples/internal/proto/examplepb"
	"github.com/binchencoder/letsgo"
)

const urlPattern = "http://%s/v1/example/echo/2211/1"

var (
	flagServer = flag.String("server", "localhost:8080", "The server host:port")
	onceOnly   = flag.Bool("send-once", false, "Send the request once only")
	clientId   = flag.String("client-id", "", "Client ID")
	xSource    = flag.String("x-source", "web", "X-Source to use")
)

func main() {
	letsgo.Init()

	url := fmt.Sprintf(urlPattern, *flagServer)

	// If specified, send request once and return.
	if *onceOnly {
		sendRequest(url)
		return
	}

	for i := 0; i < 10; i++ {
		go func() {
			for range time.Tick(50 * time.Millisecond) {
				sendRequest(url)
			}
		}()
	}

	select {}
}

func sendRequest(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("NewRequest: ", err)
		return
	}
	req.Header.Set("x-source", *xSource)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-Uid", "marlin")
	req.Header.Set("X-Cid", "disney")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return
		}

		record := examplepb.SimpleMessage{}
		if err = jsonpb.UnmarshalString(string(bodyBytes), &record); err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("%d, resp: %+v.\n", resp.StatusCode, record)
	} else {
		fmt.Println(resp.StatusCode)
	}
}
