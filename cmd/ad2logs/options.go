package main

import (
   "net/url"
)

var config struct {
   Version  bool     `help:"Show version and exit" default:"false"`
   Broker  *url.URL   `help:"MQTT URL" default:"tcp://localhost:1883"`
   Topic     string   `help:"MQTT Topic for Log Messages" default:"alarm/dump"`
   CMD      string   `help:"Command to send" arg:"" default:"log"`   
}

