# Getting Started with React App based on Openlayers and NATS streaming

Prerequisites
1. Nats server
2. Nats cli
3. Node.js
4. Clickhouse DB - optional if streaming using Nats-Clickhouse engine

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

The page will reload when you make changes.\
You may also see any lint errors in the console.

## `Run nats server`

$nats-server -c go-nats-angular/nats-server.conf

Should get an output like this:
listen: 127.0.0.1:4222
jetstream: enabled
websocket: {
        port: 8080
        no_tls: true
}
authorization {
        default_permissions = {}
}

##`Listen to Nats ws socket on React App`

$nats sub "hello.*"

##`Publish streaming data to a marker on openlayers map`

$nats pub "geojson.feature" {{.Count}} --count 100000

###The UI integrates table data from clickhouse_table to show row data and uses go_nats_angular go_micro service to get data.

Using Clickhouse -
./clickhouse server
./clickhouse client -- Can use UI

CLICKHOUSE UI - http://localhost:8123/play
—————————————————————
//show databases;
//show tables;
//select * from first_table;
//use nats;
//select * from nats.clickhouse_table; // On terminal it shows as 2 dif sections
//insert into nats.clickhouse_table(user_id, message) values(333, 'hello req res');
desc nats.clickhouse_table
/*
CREATE TABLE clickhouse_table (
  user_id UInt32,
  message String,
  timestamp DateTime,
  metric Float32
) ENGINE = MergeTree
PARTITION BY toYYYYMM(timestamp)
ORDER BY (user_id, timestamp);
*/

