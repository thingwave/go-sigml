# go-sigml
Golang module for Signaling Media Language (SigML)

## Introduction

## Format
SigML is to error messages and events what SenML is for sensor data.
SigML can express events, errors and alerts. And even embedd SenML data.

### Examples

```
[
  {"bn": "urn:dev:mac:00170d451f62:", "bt": 176627612.2, "n": "sensor/gps", "x": "Sensor malfuction", "se": 200, "d": "GPS sensor not connected"},
  {"n": "co2/0", "e": "Bad sensor reading", "se": 200, "d": "Negative value -4 ppm"},
  {"n": "temperature/0", "e": "Bad sensor reading", "se": 200, "d": "Outside sensing range(-40 to +85): 6387.4"}
  {"n": "battery", "e": "Low battery level", "se": 100, "d": "Battery at 12% capacity"}
]
```


```
[
  {"bn": "urn:dev:name:twg0043:", "bt": 176627612.2, "n": "/dev/sda0", "x": "Storage full", "se": 255, "d": "Disk 100.0% full"},
  {"n": "wlan0", "e": "Disconnected", "se": 100, "d": "Connection lost"},
  {"n": "ufw", "x": "Disabled", "se": 100, "d": "Firewall disabled"}
]
```

## Installation
Install
```
go get github.com/thingwave/go-sigml
```

## Further reading

