package main

import (
   "time"
   "fmt"
)

func currentTime(prefix string) string {
   t := time.Now()
   rtn := ""
   if t.Hour() >= 12 {
      rtn = fmt.Sprintf("%s1", t.Format("0304"))
   } else {
      rtn = fmt.Sprintf("%s0", t.Format("0304"))
   }
   return fmt.Sprintf("%s*%s%s*", prefix, rtn, t.Format("060102"))
}
