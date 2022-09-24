package main

import (
   "time"
)

type StatusLog struct {
   raw   string   `json:"-"`
   Decoded Status `json:"Decoded"`
}

type StatusBits struct {
   Ready           bool
   ArmedAway       bool
   ArmedHome       bool
   BacklightOn     bool
   ProgrammingMode bool
   Beeps           byte
   ZoneBypassed    bool
   ACPower         bool
   ChimeOn         bool
   AlarmEvent      bool
   AlarmSounding   bool
   BatteryLow      bool
   EntryDelayOff   bool
   FireAlarm       bool
   CheckZone       bool
   PerimeterOnly   bool
   SystemFault     bool
}

type Status struct {
   Bits           StatusBits
   RawData        string
   NumericCode    uint
   Message        string
   Last           time.Time
   Count          uint
}

type MqttStatus struct {
   TimeStamp      time.Time   `json:"Timestamp"`
   StateValue     string      `json:"State"`
   StateText      string      `json:"Display"`
   Bits           StatusBits
}
