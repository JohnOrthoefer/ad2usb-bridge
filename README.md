# ad2usb-bridge
This provides a bridge between a [NuTech AD2USB](https://www.alarmdecoder.com/catalog/product_info.php/products_id/29) (it should work with the PI Hat, and serial one as well).  The AD2USB provides an interface between an Ademco Vista or DSC Power Series Security System.

This code interfaces the device with a [Homebridge](https://homebridge.io/) and allows control of it via [MQTTThing](https://github.com/arachnetech/homebridge-mqttthing).

## ad2usb Flags
### General
_--help_  Show a list of the command flags and exit

_--version_ Show the version number and git hash of the build

_--[no-]timestamp_ turn the output timestamps on/off, when running it in a terminal for testing, when running under systemd, it timestamps the output.

### Serial/USB
_--serial="/dev/ttyUSB0"_  Set the serial device. 

_--baud=115200_  Set the speed of the serial device, 115200 is the USB speed.

### Configuration
_--config="/usr/local/etc/ad2usb.yaml"_  the location of the file which has the codes to send the the security system.

**Configuation Syntax**

The file is YAML syntax, with a few scalars.  

  _stay: "#3"_ Value to send to arm the system for stay/night.

  _away: "#2"_ Value to send to arm the system for away.

  _disarm: "54321"_ Value to send to disarm.

  _faults: "*"_ Value to dump faults.

### MQTT
_--publish=tcp://localhost:1883_  MQTT broker to publish to.

_--state="alarm/state"_ Topic for the current state.

_--set="alarm/set"_  Topic to change the alarm state.

_--dump="alarm/dump"_ Topic that is used to see what messages have been sent since start-up.

## Homebridge configuration
```
{
    "type": "securitySystem",
    "name": "Alarm",
    "url": "mqtt://localhost",
    "topics": {
        "setTargetState": "alarm/set",
        "getTargetState": {
            "topic": "alarm/state",
            "apply": "return JSON.parse(message).State"
        },
        "getCurrentState": {
            "topic": "alarm/state",
            "apply": "return JSON.parse(message).State"
        }
    },
    "restrictTargetState": [
        0,
        1,
        3
    ],
    "accessory": "mqttthing"
}
```

## Dump Information
Extra infomation can be retrived by subscribing to the Dump topic, by default `alarm/dump`.  When
a JSON dump command is send it will reply on the Dump topic.  

* `{"dump": "log"}` - returns the list of the status messages recieved since the last start/restart.
* `{"dump": "config"}` - returns the current configuation of the AD2 device.

Using mosquitto, MQTT command line utiliity-
* Listen - `mosquitto_sub -L mqtt://localhost/alarm/dump | jq .`
* Log - `mosquitto_pub -L mqtt://localhost/alarm/dump -m '{"dump": "log"}`
* Configuration - `mosquitto_pub -L mqtt://localhost/alarm/dump -m '{"dump": "config"}`

## Referances 
This is what I used as referance for this code. 
* [Protocol](https://www.alarmdecoder.com/wiki/index.php/Protocol)
* [AlarmDecoder Python](https://github.com/nutechsoftware/alarmdecoder)
* [Vista 20 - Programming Guide](http://site.aesecurity.com/Manuals/v15pand20pprogrammingguide.pdf)
* [Vista 20 - Users Manual](https://www.holmeselectricsecurity.com/wp-content/uploads/Vista_15P_User_Manual.pdf)
