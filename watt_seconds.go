package unit

import (
	"fmt"
	"strconv"
	"strings"
)

type WattSeconds float64

const (
	NanoWattSeconds  Watt = 1e-12
	MikroWattSeconds Watt = 1e-6
	MilliWatSeconds  Watt = 1e-3
	KiloWattSeconds  Watt = 1e3
	MegaWattSeconds  Watt = 1e6
	GigaWattSeconds  Watt = 1e9
	TeraWattSeconds  Watt = 1e12
)

func (w WattSeconds) Parse(s string) (WattSeconds, error) {
	num := numReg.Find([]byte(s))
	scaleAndTime := strings.Split(strings.TrimLeft(s, string(num)), "W")

	scale, ok := scales[scaleAndTime[0]]
	if !ok {
		return w, fmt.Errorf("scale %q is not supported", scaleAndTime[0])
	}

	time, ok := times[scaleAndTime[1]]
	if !ok {
		return w, fmt.Errorf("time %q is not supported", scaleAndTime[1])
	}

	n, err := strconv.ParseFloat(string(num), 64)
	if err != nil {
		return w, err
	}

	return WattSeconds(n) / WattSeconds(scale) * WattSeconds(time), nil
}
