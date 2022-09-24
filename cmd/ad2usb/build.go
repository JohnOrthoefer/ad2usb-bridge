package main

import (
   "log"
   "os"
)

var (
	sha1ver   string
	buildTime string
	repoName  string
)

func printVersion(e bool) {
   log.Printf("# %s - %s @%s\n", repoName, sha1ver, buildTime)
   if e {
      os.Exit(0)
   }
}
