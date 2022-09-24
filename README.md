# ad2usb-bridge
This provides a bridge between a [NuTech AD2USB](https://www.alarmdecoder.com/catalog/product_info.php/products_id/29) (it should work with the PI Hat, and serial one as well).  The AD2USB provides an interface between an Ademco Vista or DSC Power Series Security System.

This code interfaces the device with a [Homebridge](https://homebridge.io/) and allows control of it via [MQTTThing](https://github.com/arachnetech/homebridge-mqttthing).

## ad2usb Flags
### General
**--help**

**--version**

**--[no-]timestamp**

### Serial/USB
**--serial="/dev/ttyUSB0"**
**--baud=115200**

### Configuration
**--config="/usr/local/etc/ad2usb.yaml"**

### MQTT
**--publish=tcp://localhost:1883**
**--state="alarm/state"**
**--set="alarm/set"**
**--dump="alarm/dump"**

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
