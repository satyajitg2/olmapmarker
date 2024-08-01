package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type User struct {
	Name, Message string
	Metric        float32
}
type Message struct {
	Name, Body string
	//Body string
	Time int64
}

func connect() (driver.Conn, error) {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"localhost:9000"},
			Auth: clickhouse.Auth{
				Database: "default",
				Username: "default",
				Password: "",
			},
			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "an-example-go-client", Version: "0.1"},
				},
			},

			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
			/*
				TLS: &tls.Config{
					InsecureSkipVerify: true,
				},
			*/
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	return conn, nil
}

func selectClickQuery(query string) []byte {
	conn, err := connect()
	if err != nil {
		panic((err))
	}

	ctx := context.Background()

	//rows, err := conn.Query(ctx, "SELECT toString(user_id) as uuid_str, message, metric FROM nats.clickhouse_table LIMIT 5")
	rows, err := conn.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	s1 := ""
	for rows.Next() {
		var (
			model User
			s2    []byte
		)
		if err := rows.Scan(
			&model.Name,
			&model.Message,
			&model.Metric,
		); err != nil {
			log.Fatal(err)
		}
		//m := Message{"Alice", "Hello", 1294706395881547000}
		//b, _ := json.Marshal(m)
		//fmt.Println("Alice ", string(b))
		//temp := model
		//fmt.Println("t --- ", temp)

		s2, _ = json.Marshal(model)

		//fmt.Println("s2 ---- ", string(temps2), err)

		s1 += string(s2)
		//mRows = append(mRows, model)
		//s2, _ := json.Marshal(model)
	}
	//s1 := string(json.Marshal(mRows))
	fmt.Println("Query response string", s1)
	slice1 := query
	slice := append([]byte("Clickhouse selectQuery Response ECHO "))
	slice = append(slice, slice1...)
	slice = []byte(s1)
	return slice
}

func ranClickHouseSelect() {
	conn, err := connect()
	if err != nil {
		panic((err))
	}

	ctx := context.Background()
	rows, err := conn.Query(ctx, "SELECT toString(user_id) as uuid_str, message FROM nats.clickhouse_table LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			model User
		)
		if err := rows.Scan(
			&model.Name,
			&model.Message,
		); err != nil {
			log.Fatal(err)
		}
		log.Printf("user_id: %s, message: %s",
			&model.Name, &model.Message)
	}
}
