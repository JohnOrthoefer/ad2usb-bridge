package main

import (
   "net/url"
)

var config struct {
   Version  bool     `help:"Show version and exit" default:"false"`
   Publish  *url.URL   `help:"MQTT URL" default:"tcp://localhost:1883"`
   Dump     string   `help:"MQTT Topic for Log Messages" default:"alarm/dump"`
}

