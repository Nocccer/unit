package unit

import (
	"fmt"
	"strconv"
	"strings"
)

type Watt float64

const (
	NanoWatt  Watt = 1e-12
	MikroWatt Watt = 1e-6
	MilliWat  Watt = 1e-3
	KiloWatt  Watt = 1e3
	MegaWatt  Watt = 1e6
	GigaWatt  Watt = 1e9
	TeraWatt  Watt = 1e12
)

func (w Watt) Parse(s string) (Watt, error) {
	num := numReg.Find([]byte(s))
	scaleStr := strings.TrimRight(strings.TrimLeft(s, string(num)), "W")

	scale, ok := scales[scaleStr]
	if !ok {
		return w, fmt.Errorf("scale %q is not supported", scaleStr)
	}

	n, err := strconv.ParseFloat(string(num), 64)
	if err != nil {
		return w, err
	}

	return Watt(n) / Watt(scale), nil
}

func (w *Watt) UnmarshalText(text []byte) error {
	v, err := w.Parse(string(text))
	if err != nil {
		return err
	}

	*w = v

	return nil
}

func (w *Watt) UnmarshalJSON(data []byte) error {
	// given as number
	if v, err := strconv.ParseFloat(string(data), 64); err == nil {
		*w = Watt(v)
	}

	// given as string
	data = []byte(strings.Trim(string(data), `"`))
	return w.UnmarshalText(data)
}
