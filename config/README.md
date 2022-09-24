# Configuration files

## Software configuration - ad2usb.yaml
Stores the codes used to change modes of the alarm. 

Keys
* stay - code to put the system in the Home-Armed state
* away - code to put the system in the Away-Armed state
* disarm - code to disarm the system, this is where you put your code
* faults - code to retreve the faults

## Systemd configuration - ad2usb.service ad2usb
Systemd is used to start the service.   It is configured to run as a system user `alarmctl` and group `dialout` and the `ad2usb.yaml` is owned by the same user so it can be u+r only.

## Setup
```
cp ../bin/ad2usb /usr/local/bin
cp ad2usb.yaml /usr/local/etc/
cp ad2usb.service /etc/systemd/system
cp ad2usb /etc/defaults/
chmod 400 /usr/local/etc/ad2usb.yaml
chown alarmctl.dialout /usr/local/etc/ad2usb.yaml
systemd daemon-reload`
systemd enable --now ad2usb`
```

