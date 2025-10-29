package unit_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Nocccer/unit"
	"go.yaml.in/yaml/v4"
)

func TestWattParse(t *testing.T) {
	var w unit.Watt

	w, err := w.Parse("3kW")
	if err != nil {
		t.Errorf("no error expected got %v", err)
		t.Fail()
	}

	if w != 3000 {
		t.Fail()
	}
}

func TestWattUnmarshalJSON(t *testing.T) {
	var s struct {
		W unit.Watt `json:"w"`
	}

	js := []byte(`{"w": "3kW"}`)

	err := json.Unmarshal(js, &s)
	if err != nil {
		t.Errorf("no error expected got %v", err)
		t.Fail()
	}

	fmt.Println(s)

	js = []byte(`{"w": 1000}`)

	err = json.Unmarshal(js, &s)
	if err != nil {
		t.Errorf("no error expected got %v", err)
		t.Fail()
	}

	fmt.Println(s)
}

func TestWattUnmarshalYAML(t *testing.T) {
	var s struct {
		W unit.Watt `yaml:"w"`
	}

	yml := []byte(`w: '3kW'`)

	err := yaml.Unmarshal(yml, &s)
	if err != nil {
		t.Errorf("no error expected got %v", err)
		t.Fail()
	}

	fmt.Println(s)

	yml = []byte(`w: 1000`)

	err = yaml.Unmarshal(yml, &s)
	if err != nil {
		t.Errorf("no error expected got %v", err)
		t.Fail()
	}

	fmt.Println(s)
}
