package main

import (
   "log"
   "fmt"
   "os"
)

var (
	sha1ver   string
	buildTime string
	repoName  string
)

func logVersion(e bool) {
   log.Printf("# %s - %s @%s\n", repoName, sha1ver, buildTime)
   if e {
      os.Exit(0)
   }
}

func fmtVersion(e bool) {
   fmt.Printf("# %s - %s @%s\n", repoName, sha1ver, buildTime)
   if e {
      os.Exit(0)
   }
}
