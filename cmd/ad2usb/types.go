package main

import (
   "time"
)

type StatusLog struct {
   raw   string   `json:"-"`
   Decoded Status `json:"Decoded"`
}

type StatusBits struct {
   Ready           bool    `json:"ready"`
   ArmedAway       bool    `json:"armed_away"`
   ArmedHome       bool    `json:"armed_stay"`
   BacklightOn     bool    `json:"backlight_on"`
   ProgrammingMode bool    `json:"programming_mode"`
   Beeps           byte    `json:"beeps"`
   ZoneBypassed    bool    `json:"zone_bypassed"`
   ACPower         bool    `json:"ac_power"`
   ChimeOn         bool    `json:"chime_on"`
   AlarmEvent      bool    `json:"alarm_event_occurred"`
   AlarmSounding   bool    `json:"alarm_sounding"`
   BatteryLow      bool    `json:"battery_low"`
   EntryDelayOff   bool    `json:"entry_delay_off"`
   FireAlarm       bool    `json:"fire_alarm"`
   CheckZone       bool    `json:"exit_now"`
   PerimeterOnly   bool    `json:"perimeter_only"`
   SystemFault     bool    `json:"system_issue"`
}

type Status struct {
   Bits           StatusBits
   RawData        string
   NumericCode    uint  `json:"numeric_message"`
   Message        string   `json:"alpha_message"`
   Last           time.Time
   Count          uint
}

type MqttStatus struct {
   TimeStamp      time.Time   `json:"Timestamp"`
   StateValue     string      `json:"State"`
   StateText      string      `json:"Display"`
   Bits           StatusBits
}
