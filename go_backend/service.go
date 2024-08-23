package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

func main() {
	ctx := context.Background()
	fmt.Println("Hello Micro")

	//TODO: Run a channel, send req to channel and get data back.
	ranClickHouseSelect()

	servers := []string{"nats://127.0.0.1:4222", "nats://0.0.0.0:1223"}

	nc, err := nats.Connect(strings.Join(servers, ","))
	//nc, _ := nats.Connect(nats.DefaultURL)
	//TODO: Enable this to connect to Docker
	//nc, _ := nats.Connect("nats://0.0.0.0:4222")
	defer nc.Close()

	echoHandler := func(req micro.Request) {
		fmt.Println(req.Data())
		slice1 := req.Data()
		slice := append([]byte("Server ECHO "))
		slice = append(slice, slice1...)
		fmt.Println(slice)
		req.Respond(slice)
	}

	clickhouseHandler := func(req micro.Request) {
		res := selectClickQuery(string(req.Data()))
		req.Respond(res)
	}

	requestHandler := func(req micro.Request) {
		fmt.Println(req.Data())
		slice1 := req.Data()
		slice := append([]byte("Server says Hello "))
		slice = append(slice, slice1...)
		fmt.Println(slice)
		req.Respond(slice)
	}

	//	fmt.Printf("requestHandler: %v\n", requestHandler)

	config := micro.Config{
		Name:    "EchoService",
		Version: "1.0.0",
		//base handler
		Endpoint: &micro.EndpointConfig{
			Subject: "svc.echo",
			Handler: micro.HandlerFunc(echoHandler),
		},
	}

	srv, err := micro.AddService(nc, config)
	fmt.Println(srv.Info(), srv.Stats())

	if err != nil {
		fmt.Println("Error adding service ", err.Error())
	}
	add := func(req micro.Request) {
		req.Respond(req.Data())
	}
	err = srv.AddEndpoint("svc.add", micro.HandlerFunc(add))
	defer srv.Stop()

	micro.AddService(nc, micro.Config{
		Name:    "HelloService",
		Version: "1.0.0",
		//base handler
		Endpoint: &micro.EndpointConfig{
			Subject: "svc.hello",
			Handler: micro.HandlerFunc(requestHandler),
		},
	})

	micro.AddService(nc, micro.Config{
		Name:    "ClickhouseService",
		Version: "1.0.0",
		//base handler
		Endpoint: &micro.EndpointConfig{
			Subject: "click",
			Handler: micro.HandlerFunc(clickhouseHandler),
		},
	})

	//TODO: Setup infinite consumer right away.
	go infiniteStream()

	/*
		streamHandle := func(req micro.Request) {
			infiniteStream()
		}

		micro.AddService(nc, micro.Config{
			Name:    "InfiniteStreamService",
			Version: "1.0.0",
			//base handler
			Endpoint: &micro.EndpointConfig{
				Subject: "orders.*",
				Handler: micro.HandlerFunc(streamHandle),
			},
		})
	*/
	//fmt.Println(srv, err)

	//runtime.Goexit()
	<-ctx.Done()
}
