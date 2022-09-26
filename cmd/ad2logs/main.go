package main

import (
   "fmt"
   "time"
   "bytes"
   "github.com/eclipse/paho.mqtt.golang"
   "github.com/alecthomas/kong"
   "encoding/json"
)

type QueryType struct {
   Dump string  `json:"Dump"`
}

var (
   query QueryType
   jsonStr []byte
   done bool = false
)

func dumpMesg(c mqtt.Client, m mqtt.Message ) {
   if bytes.Compare(m.Payload(), jsonStr) == 0 {
      return
   }

   fmt.Printf("%s\n", string(m.Payload()))
   
   done = true
}

func main() {

   kong.Parse(&config,
	   kong.Name("ad2logs"),
	   kong.Description("Get logs from ad2usb"))

   if config.Version {
      fmtVersion(true)
   }
   opts := mqtt.NewClientOptions().AddBroker(config.Broker.String())
   mqttClient := mqtt.NewClient(opts)

   if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
      fmt.Printf("%s\n", token.Error())
      return
   }
   mqttClient.Subscribe(config.Dump, 1, dumpMesg)

   query.Dump = config.CMD

   jsonStr, _ = json.Marshal(&query) 

   mqttClient.Publish(config.Dump, 0, false, jsonStr)

   for (!done) {
      time.Sleep(time.Second)
   }
}
