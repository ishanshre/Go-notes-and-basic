package main

import "testing"

var tests = []struct {
	name     string
	divident float32
	diviser  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.00, 10.0, 10.0, false},
	{"invalid-data", 100.00, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.divident, tt.diviser)
		if tt.isErr {
			if err == nil {
				t.Error("tested the error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not except error but go one ", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("expeced %f but got %f", tt.expected, got)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(22.2, 0.0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}
