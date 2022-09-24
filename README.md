# ad2usb-bridge
This provides a bridge between a [NuTech AD2USB](https://www.alarmdecoder.com/catalog/product_info.php/products_id/29) (it should work with the PI Hat, and serial one as well).  The AD2USB provides an interface between an Ademco Vista or DSC Power Series Security System.

This code interfaces the device with a [Homebridge](https://homebridge.io/) and allows control of it via [MQTTThing](https://github.com/arachnetech/homebridge-mqttthing).

## ad2usb Flags
### General
__--help__  Show a list of the command flags and exit

__--version__ Show the version number and git hash of the build

__--[no-]timestamp__ turn the output timestamps on/off, when running it in a terminal for testing, when running under systemd, it timestamps the output.

### Serial/USB
__--serial="/dev/ttyUSB0"__  Set the serial device. 

__--baud=115200__  Set the speed of the serial device, 115200 is the USB speed.

### Configuration
__--config="/usr/local/etc/ad2usb.yaml"__  the location of the file which has the codes to send the the security system.

### MQTT
__--publish=tcp://localhost:1883__  MQTT broker to publish to.

__--state="alarm/state"__ Topic for the current state.

__--set="alarm/set"__  Topic to change the alarm state.

__--dump="alarm/dump"__ Topic that is used to see what messages have been sent since start-up.

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

## Referances 
This is what I used as referance for this code. 

* [Protocol](https://www.alarmdecoder.com/wiki/index.php/Protocol)

* [AlarmDecoder Python](https://github.com/nutechsoftware/alarmdecoder)
