# go-sigml
Golang module for Signaling Media Language (SigML)

## Introduction
SigML is to error messages and events what SenML is for sensor data.
SigML can express events, errors and alerts. And even embedd SenML data.

## Format
SigML is based on JSON and tries to use the same name and notatons as SenML, where possible.

## Fields

### bn
This is the Base name from SenML.
### bt
This is the Base time from SenML, expressed as seconds from 1970-0101 00:00 (UNIX epoch).
### bver
This is the Base version from SenML. Set to 1 for now.

### x
x stands for eXception and is used for indicating exceptions upstream.
### e
e stands for Error and is used for indicating different typs of hardware, software and system errors upstream.
### s
s stands for Signal event and is used for transmitting events upstream.

### se
se means Severity and is used to indicate how serious an exception/error/signal is.

None Severity : 0
Normal        : 1
Incident      : 50
Error         : 100
Critical      : 200
Catastrophic  : 255

### d
d stand for Description is is a human readable text indicating what has happened.

### p
p is used for sending generic payloads, such as SenML encoded data belonging to an error or similiar.

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


## Usage
First import using

```
import (
  "fmt"
  "encoding/json"
  "github.com/thingwave/go-sigml"
)
```

Create message using:
```
var evt sigml.SigMLMessage
head := sigml.SigMLRecord{Bn: "", Bt: float64(time.now().Unixmilli)/1000.0, Bver: 1}
evt = append(evt, head)

payload, err := json.Marshal(evt)
if err == nil {
  fmt.Printf("%s\n", string(payload))
}

```
## Further reading

