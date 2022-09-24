package main
import (
   "log"
   "os"
   "gopkg.in/yaml.v3"
)

func readConfigFile(cf string) map[string]string {

   file, err := os.ReadFile(cf)

   if err != nil {
      log.Fatal(err)
   }

   data := make(map[string]string)
   error := yaml.Unmarshal([]byte(file), &data)
   if error != nil {
      log.Fatal(err)
   }

   return data
} 
