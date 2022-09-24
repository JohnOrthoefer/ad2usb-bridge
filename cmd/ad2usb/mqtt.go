package main

import (
   "log"
   "strings"
   "github.com/eclipse/paho.mqtt.golang"
   "encoding/json"
)

var (
   mqttClient  mqtt.Client
   mqttTopic   string
   replyTopic  string
)

func setTopic(t string) {
   mqttTopic = t
}

func getTopic() string {
   rtn := mqttTopic
   return rtn
}

func mqttConfig(cfgStr string) {
   Cfg := make(map[string]string)

   for _, val := range strings.Split(cfgStr, "&") {
      result := strings.Split(val, "=")
      Cfg[result[0]] = result[1]
   }

   rtn, _ := json.Marshal(&Cfg)
   mqttClient.Publish(replyTopic, 0, false, rtn)
}

func mqttStatus(s Status)string {
   var ms MqttStatus

   if s.Bits.AlarmSounding {
      ms.StateValue = ALARM_TRIGGERED
   } else if s.Bits.ArmedHome {
      ms.StateValue = STAY_ARM
   } else if s.Bits.ArmedAway {
      ms.StateValue = AWAY_ARM
   } else if s.Bits.PerimeterOnly {
      ms.StateValue =  NIGHT_ARM
   } else {
      ms.StateValue = DISARMED
   }
   ms.StateText = s.Message
   ms.TimeStamp = s.Last
   ms.Bits = s.Bits

   rtn, _ := json.Marshal(ms)
   return string(rtn)
}

func cmdMesg(c mqtt.Client, m mqtt.Message ) {
   newState = string(m.Payload())
   log.Printf("%s => [%s]\n", m.Topic(), newState)
}

func dumpMesg(c mqtt.Client, m mqtt.Message ) {
   var Query struct {
      Dump string `json:"dump"`
   }

   err := json.Unmarshal(m.Payload(), &Query)
   if err != nil {
      log.Printf("err:%s, Can not Parse \"%s\"\n", err, string(m.Payload()))
      return
   }

   if Query.Dump == "log" {
      rtn, _ := json.Marshal(logStatus)
      mqttClient.Publish(m.Topic(), 0, false, rtn)
      return
   }

   if Query.Dump == "config" {
      newState = CONFIG
      replyTopic = m.Topic()
      return
   }
}

func mqttSetup(broker string, pubTopic string, subTopic string, dumpTopic string) {
   opts := mqtt.NewClientOptions().AddBroker(broker)
   mqttClient = mqtt.NewClient(opts)
   if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
      mqttClient = nil
      log.Printf("%s\n", token.Error())
      return
   }
   setTopic(pubTopic)
   mqttClient.Subscribe(subTopic, 1, cmdMesg)
   mqttClient.Subscribe(dumpTopic, 1, dumpMesg)
}
