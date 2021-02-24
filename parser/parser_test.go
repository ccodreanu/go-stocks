package parser

import (
	"testing"
)

func TestExistingSymbol(t *testing.T) {
	v, err := FetchValue("VUSA.AS")

	if v.Value < 0 || err != nil {
		t.Fatalf(`FetchValue("VUSA.AS") = %f, %v`, v.Value, err)
	}
}

func TestNonExistingSymbol(t *testing.T) {
	v, err := FetchValue("VUSA.ASSSSS")

	if err == nil && v.Value >= 0 {
		t.Fatalf(`FetchValue("VUSA.ASSSSS") = %f, %v`, v.Value, err)
	}
}
