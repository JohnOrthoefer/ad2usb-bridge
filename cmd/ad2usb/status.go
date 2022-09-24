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

func storeStatus(raw string) bool  {
   if !validCode.MatchString(raw) {
      log.Printf("discarding- '%s'\n", raw)
      return false
   }
   m := validCode.FindAllStringSubmatch(raw, -1)
   bitstr := m[0][1]
   disp   := m[0][4]

   if lastStatus != nil && raw == lastStatus.raw {
      lastStatus.Decoded.Count += 1
      lastStatus.Decoded.Last = time.Now()
      return true
   }

   if logStatus[raw] == nil {
      t := StatusLog {
         raw:   raw,
      }
      logStatus[raw] = &t
   }
   lastStatus = logStatus[raw]

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

   logStatus[raw].Decoded = alarmStatus

   log.Printf("%s (%s)", alarmStatus.Message, bitstr)
   return true
}

func init() {
   validCode = regexp.MustCompile(`\[([0-9]{16})----\],([0-9]{3}),\[([0-9a-f]{30})\],"(.*)"`)
   logStatus = make(map[string]*StatusLog)
}
