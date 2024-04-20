package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

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

	//rows, err := conn.Query(ctx, "SELECT toString(user_id) as uuid_str, message FROM nats.clickhouse_table LIMIT 5")
	rows, err := conn.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	s1 := ""
	for rows.Next() {
		var (
			name, uuid string
		)
		if err := rows.Scan(
			&name,
			&uuid,
		); err != nil {
			log.Fatal(err)
		}
		//log.Printf("user_id: %s, message: %s", name, uuid)
		rowString := ""
		rowString = fmt.Sprintf("user_id: %s, message: %s",
			name, uuid)
		s1 += string(rowString)
	}

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
			name, uuid string
		)
		if err := rows.Scan(
			&name,
			&uuid,
		); err != nil {
			log.Fatal(err)
		}
		log.Printf("user_id: %s, message: %s",
			name, uuid)
	}
}
