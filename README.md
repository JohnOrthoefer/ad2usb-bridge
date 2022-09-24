# ad2usb-bridge
This provides a bridge between a [NuTech AD2USB](https://www.alarmdecoder.com/catalog/product_info.php/products_id/29) (it should work with the PI Hat, and serial one as well).  The AD2USB provides an interface between an Ademco Vista or DSC Power Series Security System.

This code interfaces the device with a [Homebridge](https://homebridge.io/) and allows control of it via [MQTTThing](https://github.com/arachnetech/homebridge-mqttthing).

##



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
