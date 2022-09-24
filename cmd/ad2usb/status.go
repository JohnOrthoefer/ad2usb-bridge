package main

import (
   "log"
   "strings"
   "strconv"
   "time"
   "regexp"
)

var (
   alarmStatus Status
   lastStatus  *StatusLog
   logStatus   map[string]*StatusLog
   validCode   *regexp.Regexp
   validConfig *regexp.Regexp
)

func getBoolValue(s string, n int) bool {
   c := s[n]
   return (c == '1')
}

func getByteValue(s string, n int) byte {
   v, _ := strconv.Atoi(string(s[n]))
   return byte(v)
}

func getInt(s string) uint {
   v, _ := strconv.Atoi(s)
   return uint(v)
}

func parseConfig(str string) {

   m := validConfig.FindAllStringSubmatch(str, -1)
   log.Printf("Config- %q", m[0][1])
   mqttConfig(m[0][1])
}

func parseStatus(str string) {
   m := validCode.FindAllStringSubmatch(str, -1)
   bitstr := m[0][1]
   disp   := m[0][4]

   if lastStatus != nil && str == lastStatus.raw {
      lastStatus.Decoded.Count += 1
      lastStatus.Decoded.Last = time.Now()
      return
   }

   if logStatus[str] == nil {
      t := StatusLog {
         raw:   str,
      }
      logStatus[str] = &t
   }
   lastStatus = logStatus[str]

   alarmStatus = Status {
      Bits:       StatusBits {
         Ready:           getBoolValue(bitstr, 0),
         ArmedAway:       getBoolValue(bitstr, 1),
         ArmedHome:       getBoolValue(bitstr, 2),
         BacklightOn:     getBoolValue(bitstr, 3),
         ProgrammingMode: getBoolValue(bitstr, 4),
         Beeps:           getByteValue(bitstr, 5),
         ZoneBypassed:    getBoolValue(bitstr, 6),
         ACPower:         getBoolValue(bitstr, 7),
         ChimeOn:         getBoolValue(bitstr, 8),
         AlarmEvent:      getBoolValue(bitstr, 9),
         AlarmSounding:   getBoolValue(bitstr, 10),
         BatteryLow:      getBoolValue(bitstr, 11),
         EntryDelayOff:   getBoolValue(bitstr, 12),
         FireAlarm:       getBoolValue(bitstr, 13),
         CheckZone:       getBoolValue(bitstr, 14),
         PerimeterOnly:   getBoolValue(bitstr, 15),
      },
      RawData:         m[0][3],
      NumericCode:     getInt(m[0][2]),
      Message:         strings.TrimSpace(disp),
      Last:            time.Now(),
      Count:           1,
   }

   logStatus[str].Decoded = alarmStatus

   log.Printf("%s (%s)", alarmStatus.Message, bitstr)
}

func storeStatus(raw string) bool  {
   if raw == "!>" {
      return false
   }

   if validConfig.MatchString(raw) {
      parseConfig(raw)
      return false
   }

   if validCode.MatchString(raw) {
      parseStatus(raw)
      return true
   } 

   log.Printf("discarding- '%s'\n", raw)
   return false
}

func init() {
   validCode = regexp.MustCompile(`\[([0-9]{16})----\],([0-9]{3}),\[([0-9a-f]{30})\],"(.*)"`)
   validConfig = regexp.MustCompile(`^!CONFIG>(.*)$`)
   logStatus = make(map[string]*StatusLog)
}
