# Getting Started with React App based on Openlayers and NATS streaming

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

The page will reload when you make changes.\
You may also see any lint errors in the console.

#To run the nats server use 
#nats-server -c go-nats-angular/nats-server.conf

listen: 127.0.0.1:4222
jetstream: enabled
websocket: {
        port: 8080
        no_tls: true
}
authorization {
        default_permissions = {}
}

Listen to Nats ws socket on React App
#nats sub "hello.*"

Publish streaming data to a marker on openlayers map.
#nats pub "geojson.feature" {{.Count}} --count 100000
