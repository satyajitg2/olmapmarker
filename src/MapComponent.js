import React, { useEffect, useState } from "react"
import Point from 'ol/geom/Point.js';
import "ol/ol.css"
import Map from "ol/Map"
import View from "ol/View"
import {Icon, Style} from 'ol/style.js';
import {OGCMapTile, OSM, Vector as VectorSource} from 'ol/source.js';
import {Tile as TileLayer, Vector as VectorLayer} from 'ol/layer.js';
import Overlay from "ol/Overlay"
import {toLonLat, fromLonLat, Projection, transform} from 'ol/proj.js';
import {toStringHDMS} from 'ol/coordinate.js';
import Feature from 'ol/Feature.js';
import { NatsConnection, StringCodec, Subscription, connect } from 'nats.ws';

import marker from "./data/airplane.png"
import img from "./data/bigdot.png"
import ChildMapComponent from "./ChildMapComponent";

function getLondonMap() {

  const rome = new Feature({
    //geometry: new Point(fromLonLat([12.5, 41.9])),
    geometry: new Point(transform([12.5, 41.9], 'EPSG:4326', 'EPSG:3857')),
  });

  const london = new Feature({
    geometry: new Point(fromLonLat([-0.12755, 51.507222])),
  });

  const madrid = new Feature({
    geometry: new Point(fromLonLat([-3.683333, 40.4])),
  });
  const paris = new Feature({
    geometry: new Point(fromLonLat([2.353, 48.8566])),
  });
  const berlin = new Feature({
    geometry: new Point(fromLonLat([13.3884, 52.5169])),
  });

  rome.setStyle(
    new Style({
      image: new Icon({
        color: '#BADA55',
        crossOrigin: 'anonymous',
        src: marker,
        //image: 'data/bigdot.png',
        
        src: img,
        scale: 0.2,
      }),
    }),
  );

  london.setStyle(
    new Style({
      image: new Icon({
        color: 'rgba(255, 0, 0, .5)',
        //crossOrigin: 'anonymous',
        //src: 'data/bigdot.png',
        //src: './data/airplane.png',
        src: marker,
        scale: 0.5,
      }),
    }),
  );

  madrid.setStyle(
    new Style({
      image: new Icon({
        //crossOrigin: 'anonymous',
        //src: 'data/bigdot.png',
        src: 'data/icon.png',
        scale: 0.05,
      }),
    }),
  );

  paris.setStyle(
    new Style({
      image: new Icon({
        color: '#8959A8',
        crossOrigin: 'anonymous',
        //src: 'data/dot.svg',
        src: '@/data/icon.png',
        scale: 0.05,
      }),
    }),
  );

  berlin.setStyle(
    new Style({
      image: new Icon({
        crossOrigin: 'anonymous',
        //src: './data/bigdot.png',
        src: 'icon.png',
        scale: 0.05,
      }),
    }),
  );
  const vectorSource = new VectorSource({
    features: [rome, london, madrid, paris, berlin],
  });

  const vectorLayer = new VectorLayer({
    source: vectorSource,
  });

  const rasterLayer = new TileLayer({
    source: new OSM(),
    /*
    source: new OGCMapTile({
      url: 'https://maps.gnosis.earth/ogcapi/collections/NaturalEarth:raster:HYP_HR_SR_OB_DR/map/tiles/WebMercatorQuad',
      //crossOrigin: 'anonymous',
    }),
    */
  });

  const map = new Map({
    layers: [rasterLayer, vectorLayer],
    target: document.getElementById('markerpopupmap'),
    view: new View({
      center: fromLonLat([2.896372, 44.6024]),
      zoom: 3,
    }),
  });
 return [map, london];

}

function manageMapEffect() {

  const iconFeature = new Feature({
    geometry: new Point([0, 0]),
    name: 'Null Island',
    population: 4000,
    rainfall: 500,
  });

  const london = new Feature({
    geometry: new Point(fromLonLat([-0.12755, 51.507222])),
  });
  london.setStyle(
    new Style({
      image: new Icon({
        color: 'rgba(255, 0, 0, .5)',
        crossOrigin: 'anonymous',
        //src: 'data/bigdot.png',
        src: marker,
        scale: 0.5,
      }),
    }),
  );
  
  const iconStyle = new Style({
    image: new Icon({
      anchor: [0.5, 46],
      anchorXUnits: 'fraction',
      anchorYUnits: 'pixels',
      //src: 'data/icon.png',
      src: 'data/icon.png',
      //src: 'data/airplane.png',
      scale: 0.1,
    }),
  });
  
  iconFeature.setStyle(iconStyle);
  const vectorSource = new VectorSource({
    features: [iconFeature, london],
  });
  
  const vectorLayer = new VectorLayer({
    source: vectorSource,
  });

  const rasterLayer = new TileLayer({
    source: new OGCMapTile({
      url: 'https://maps.gnosis.earth/ogcapi/collections/NaturalEarth:raster:HYP_HR_SR_OB_DR/map/tiles/WebMercatorQuad',
      crossOrigin: '',
    }),
  });

  const container = document.getElementById('popup');
  const content = document.getElementById('popup-content');
  const closer = document.getElementById('popup-closer');

  const overlay = new Overlay({
    element: container,
    autoPan: {
      animation: {
        duration: 250,
      },
    },
  });

  const map = new Map({
    layers: [rasterLayer, vectorLayer],
    target: "markerpopupmap",
    view: new View({
      center: [0, 0],
      zoom: 3,
    }),
    overlays: [overlay],
  });


  //return [map,iconFeature];
  return [map,london];
}



function setMapClickHandler(map, iconFeature, conn, mapClick) {
    /**
   * Add a click handler to the map to render the popup.
   */
    
    map.on('singleclick', function (evt) {
      const sc = StringCodec();
      iconFeature.setGeometry(new Point(evt.coordinate));
      //setlocString(evt.coordinate) - TODO: Implement this setting state
      mapClick(evt.coordinate)
      conn.publish("hello.nats_server", sc.encode("Point Clicked " + toLonLat(evt.coordinate)));
    });
}

async function setMapSubcribe(map, iconFeature, conn) {
  let lat = 789827.3162602726;
  let long = 2487614.9289942635;
  const sc = StringCodec();
  const s = conn.subscribe("geojson.feature");
  for await(const msg of s) {

    //Co ordinate [789827.3162602726, 2487614.9289942635]
    lat = lat+ (parseInt(sc.decode(msg.data)))/10000;
    long = long+ (parseInt(sc.decode(msg.data)))/10000;
    
    console.log("Coordinate " + lat + " " + long);
    //iconFeature.setGeometry(new Point([lat, long]));
    iconFeature.setGeometry(new Point([lat, long]));
  }
}

async function manageNatsConnection(iconFeature, map, mapClick) {
  const sc = StringCodec();
  const conn = await connect(
    {
      servers: ["ws://localhost:8080", "wss://localhost:2229", "locahost:9111"],
    }
  );
  
  console.log("Connected to Nats using ",conn)
  conn.publish("hello.nats_server", sc.encode("Nats ws UI says hello"));
  console.log("Published message.")

  setMapClickHandler(map, iconFeature, conn, mapClick)
  setMapSubcribe(map, iconFeature, conn)
  return conn
}

const MapComponent = () => {
  const [locString, setlocString] = useState("[0,0]")

  const mapClick = (evt) => {
    console.log("Clicked on map!", evt)
    setlocString(evt)
  }

  const clickhousePost = (evt) => {
    //setData("ClickhousePost "+counter)
    //setCounter(counter+1)
    conn
      .catch(
        console.log("Error clickhousePost")
      )
      .then(
        (natsConn) => {
          const sc = StringCodec();
          const res = natsConn.request("click", sc.encode("SELECT toString(user_id) as uuid_str, message FROM nats.clickhouse_table"))
          res.then((msg) => {
            console.log("Received result", sc.decode(msg.data))
            //clickhousePost(sc.decode(msg.data))
            setData(sc.decode(msg.data))
          })
        }
      )
    

  }

  const [conn, setConn] = useState(null)

  const [data, setData] = useState("")
  const [counter, setCounter] = useState(0)

  useEffect(() => {
    const sc = StringCodec();
  
    const [map, iconFeature] = getLondonMap();

    const conn = manageNatsConnection(iconFeature, map, mapClick);
    setConn(conn)
    return () => {map.setTarget(null)}
  }, [])

  return (
    <>
      <div className="w-50 m-10">
        <div id="markerpopupmap" style={{ width: "95%", height: "500px" }} />
        <div id="popup" className="ol-popup" style={{ backgroundColor: "#fff" }}>
          <a href="#" id="popup-closer" className="ol-popup-closer"></a>
          <div id="popup-content"></div>
          <button onClick={clickhousePost}>Click and POST {data}</button>
        </div>
      </div>
      <ChildMapComponent location={locString}/>

    </>
  );
}

export default MapComponent

/*

async function sendHelloToNats(conn) {
  const sc = StringCodec();
  conn.publish("hello.nats_server", sc.encode("sendHelloToNats"));
}

   const data = new Data(conn, iconFeature, map)
    data.speakToNats()

class Data {
  constructor(conn, iconFeature, map) {
    this.conn = conn
    this.iconFeature = iconFeature
    this.map = map
  }
  speakToNats(conn) {
    const sc = StringCodec();
    //ERROR : //conn.publish("hello.nats_server", sc.encode("sendHelloToNats")); - Pubish is not a function
  
  }
}
*/