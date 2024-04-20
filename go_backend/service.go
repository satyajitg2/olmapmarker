package main

import (
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

func main() {
	fmt.Println("Hello Micro")

	//TODO: Run a channel, send req to channel and get data back.
	//ranClickHouseSelect()

	nc, _ := nats.Connect(nats.DefaultURL)

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

	srv, err := micro.AddService(nc, micro.Config{
		Name:    "EchoService",
		Version: "1.0.0",
		//base handler
		Endpoint: &micro.EndpointConfig{
			Subject: "svc.echo",
			Handler: micro.HandlerFunc(echoHandler),
		},
	})

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
	fmt.Println(srv, err)

	runtime.Goexit()
}
