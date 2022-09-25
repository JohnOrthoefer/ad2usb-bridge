package main

import (
   "log"
   "net/url"
)

var config struct {
   Version  bool     `help:"Show version and exit" default:"false"`
   Debug    bool     `help:"Debug info enabled" default:"false"`
   Timestamp bool    `help:"Timestamp output logs" default:"true" negatable:""`
   Serial   string   `help:"Serial Device" default:"/dev/ttyUSB0"`
   Baud     int      `help:"Baud Rate" default:115200`
   Config   string   `help:"Configuration file" default:"/usr/local/etc/ad2usb.yaml"`
   Publish  *url.URL   `help:"MQTT URL" default:"tcp://localhost:1883"`
   State    string   `help:"MQTT Topic for the State" default:"alarm/state"`
   Set      string   `help:"MQTT Topic for target State" default:"alarm/set"`
   Dump     string   `help:"MQTT Topic for Log Messages" default:"alarm/dump"`
}

func printOptions() {
   if !config.Debug {
      return
   }
   log.Printf("Version %t", config.Version)
   log.Printf("Timestamp %t", config.Timestamp)
   log.Printf("Serial %s", config.Serial)
   log.Printf("Baud %d", config.Baud)
   log.Printf("Config %s", config.Config)
   log.Printf("Publish %s", config.Publish)
   log.Printf("State %s", config.State)
   log.Printf("Set %s", config.Set)
   log.Printf("Dump %s", config.Dump)
}
