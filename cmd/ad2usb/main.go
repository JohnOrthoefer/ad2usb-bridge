package main

import (
   "github.com/tarm/serial"
   "github.com/alecthomas/kong"
   "log"
   "bufio"
)

var (
   newState string
)

func main() {

   kong.Parse(&config,
	   kong.Name("ad2usb"),
	   kong.Description("AD2USB to MQTT"))

   if !config.Timestamp {
      log.SetFlags(0)
   }


   printVersion(config.Version)
   printOptions()
   cfg := readConfigFile(config.Config)

   // Open the USB Port
   c := &serial.Config{Name: config.Serial, Baud: config.Baud}
   s, err := serial.OpenPort(c)
   if err != nil {
      log.Fatal(err)
   }

   // Set the Scanner to read lines
   scanner := bufio.NewScanner(s)

   // setup the mqtt
   mqttSetup(config.Publish.String(), config.State, config.Set, config.Dump)
   
   // Main loop wait for a line, parse it and publish
   for scanner.Scan() {
      statusValid := storeStatus(scanner.Text())

      if statusValid && mqttClient != nil && !alarmStatus.Last.IsZero() {
         mqttClient.Publish(getTopic(), 0, false, mqttStatus(alarmStatus))
      } 
      if newState != "" {
         var sendMe string
         if alarmStatus.Bits.Ready {
            switch newState {
            case STAY_ARM:
               sendMe = "stay"
            case AWAY_ARM:
               sendMe = "away"
            }
         } else if (alarmStatus.Bits.ArmedAway || alarmStatus.Bits.ArmedHome) && newState == DISARMED {
            sendMe = "disarm"
         } else if newState == FAULTS {
            sendMe = "faults"
         } else if newState == CONFIG {
            sendMe = "config"
         }
         newState = ""

         if sendMe == "" {
            log.Printf("No message to send.")
            continue
         }
         log.Printf("Sending %s", sendMe)
         s.Write([]byte(cfg[sendMe]))
      }
   }
}
