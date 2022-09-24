# ad2usb-bridge
This provides a bridge between a [NuTech AD2USB](https://www.alarmdecoder.com/catalog/product_info.php/products_id/29) (it should work with the PI Hat, and serial one as well).  The AD2USB provides an interface between an Ademco Vista or DSC Power Series Security System.

This code interfaces the device with a [Homebridge](https://homebridge.io/) and allows control of it via [MQTTThing](https://github.com/arachnetech/homebridge-mqttthing).

## ad2usb Flags
### General
**--help**  Show a list of the command flags and exit

**--version** Show the version number and git hash of the build

**--[no-]timestamp** turn the output timestamps on/off, when running it in a terminal for testing, when running under systemd, it timestamps the output.

### Serial/USB
**--serial="/dev/ttyUSB0"**  Set the serial device. 

**--baud=115200**  Set the speed of the serial device, 115200 is the USB speed.

### Configuration
**--config="/usr/local/etc/ad2usb.yaml"**  the location of the file which has the codes to send the the security system.

### MQTT
**--publish=tcp://localhost:1883**  MQTT broker to publish to.

**--state="alarm/state"** Topic for the current state.

**--set="alarm/set"**  Topic to change the alarm state.

**--dump="alarm/dump"** Topic that is used to see what messages have been sent since start-up.

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
[Protocol](https://www.alarmdecoder.com/wiki/index.php/Protocol)
[AlarmDecoder Python](https://github.com/nutechsoftware/alarmdecoder)
