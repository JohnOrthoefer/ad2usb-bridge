package main

import (
   "log"
   "github.com/eclipse/paho.mqtt.golang"
   "encoding/json"
)

var (
   mqttClient  mqtt.Client
   mqttTopic   string
)

func setTopic(t string) {
   mqttTopic = t
}

func getTopic() string {
   rtn := mqttTopic
   return rtn
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
   if string(m.Payload()) != "{\"dump\":\"now\"}" {
      return
   }

   rtn, _ := json.Marshal(logStatus)
   mqttClient.Publish(m.Topic(), 0, false, rtn)
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
