package unit_test

import (
	"fmt"
	"testing"

	"github.com/Nocccer/unit"
)

func TestWattSeconds(t *testing.T) {
	var w unit.WattSeconds

	w, err := w.Parse("1kWh")
	if err != nil {
		t.Errorf("no error expected got %v", err)
		t.Fail()
	}

	fmt.Println(w)

	if w != 3.6e6 {
		t.Fail()
	}
}
