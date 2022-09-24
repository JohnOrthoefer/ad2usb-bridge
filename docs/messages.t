To dump the log of the messages:
  mosquitto_pub -L mqtt://mera.orthoefer.org/alarm/dump -m '{"dump":"now"}'

To recieve the messages:
  mosquitto_sub -L mqtt://mera.orthoefer.org/alarm/dump | jq .

