package gml

import (
	"os"
	"testing"
)

func Test_deserialize(t *testing.T) {
	raw, err := os.ReadFile("examples/spec_example.gml")
	if err != nil {
		t.Fatal(err.Error())
	}
	gml, err := Deserialize(raw)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("%++v\n", gml)
}
