package main

import (
   "log"
)

func debugMesg(format string, v ...any) {
   if config.Debug {
      log.Printf(format, v)
   }
}
