package sigml

import (
	"errors"
	//"fmt"
	//"time"
)

type SigMLRecord struct {
	Bn string `json:"bn,omitempty"`  //source device/system
	Bt float64 `json:"bt,omitempty"`
	Bver int `json:"bver,omitempty"`
	N string `json:"n"`  //source service/signal

	X string  `json:"x,omitempty"` // eXception
	E string  `json:"e,omitempty"` // Eror
	S string  `json:"s,omitempty"` // Signal/event
	Se Severity `json:"se,omitempty"`
	D string `json:"d,omitempty"`  // Description
	P interface{} `json:"p,omitempty"` //payload
}

type Severity uint8
const (
	None Severity = 0
	Normal        = 1
	Incident      = 50
	Error         = 100
	Critical      = 200
	Catastrophic  = 255
)

type SigMLMessage []SigMLRecord

// Validate tests if a SigML message is well formated
func (msg SigMLMessage) Validate() error {
	if len(msg) == 0 {
		return errors.New("Empty message")
	}

	for _, rec := range msg {
		if rec.X == "" && rec.E == "" && rec.S == "" {
			return errors.New("One of x, e or s must be valid")
		}
		if rec.X != "" && rec.E != "" && rec.S != "" {
			return errors.New("Only one of x, e or s can be valid")
		}
		if rec.X != "" && (rec.E != "" || rec.S != "") {
			return errors.New("Only one of x, e or s can be valid")
		}
		if (rec.X != "" || rec.E != "") && rec.S != "" {
			return errors.New("Only one of x, e or s can be valid")
		}
	}
	return nil
}

