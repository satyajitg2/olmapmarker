package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
)

func consume_routine(cons jetstream.Consumer) {
	i := 0
	j := 1
	fmt.Println("Started consuming")
	start := time.Now()
	wg := sync.WaitGroup{}

	wg.Add(1)
	cons.Consume(func(msg jetstream.Msg) {
		msg.Ack()
		/*1 Millionth in time -> 48.119924156s with this print*/
		//data := string(msg.Data())
		//fmt.Println("Received msg on ", msg.Subject(), data)
		i++
		//wg.Done()
		if i == j {
			j = j + 100000
			if i == 1000001 {
				fmt.Println("----------------------1 MILLION events processed---------------------")
			}
			fmt.Printf("%d in time ->", i)
			fmt.Println(time.Since(start))
			fmt.Println("Ramdom sampler data...", string(msg.Data()))
		}

	})
	wg.Wait()
}

func init_rsocket_setup() {

}

func setup_rsocket(cli rsocket.Client) {
	// Connect to server
	// Send request

	result, err := cli.RequestResponse(payload.NewString("Hello World Payload from Golang", "Some payload")).Block(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("response:", result.DataUTF8())
}

// nats pub orders.us "Hi orders.us arriving" --count 10000
func infiniteStream() {
	fmt.Println("Hello World infinite consumer")
	/*
		//TODO: run this when Rsocket backend is RUNNING
			cli, err := rsocket.Connect().
				SetupPayload(payload.NewString("golang:service", "Hello World from Golang")).
				//SetupPayload(payload.NewString("Hello", "World")).
				Transport(rsocket.TCPClient().SetHostAndPort("127.0.0.1", 6565).Build()).
				Start(context.Background())
			if err != nil {
				panic(err)
			}
			defer cli.Close()

			setup_rsocket(cli)
	*/

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, _ := nats.Connect(url)
	defer nc.Drain()

	js, _ := jetstream.New(nc)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	streamName := "ADSB"
	infiniteStream, _ := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     streamName,
		Subjects: []string{"adsb.>"},
	})

	cons, _ := infiniteStream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{})

	fmt.Println("Prepare to Consume")
	/*
		wg := sync.WaitGroup{}
		wg.Add(1)
	*/
	fmt.Println("Start Consume")

	consume_routine(cons)

	fmt.Println("Complete consume")
	//wg.Wait()
	fmt.Println("Wg Wait")
	//cc.Stop()
	fmt.Println("cc Stop")
	runtime.Goexit()
	fmt.Println("GOExit")
}
