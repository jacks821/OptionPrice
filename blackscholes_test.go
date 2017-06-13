package main

import (

	"testing"
)

type testpair struct{
	input 	Option
	output	float64
}


var testPrice = []testpair {
	{	newOption(62, 60, 40, 0.32, 0.04, true),  3.72	},
	{	newOption(62, 60, 40, 0.32, 0.04, false),  1.46	},

}


func TestOptionPrice(t *testing.T) {
	for _, pair := range testPrice {
		v := OptionPrice(pair.input.price, pair.input.exercise, pair.input.time, pair.input.vol, pair.input.rf, pair.input.call)
		if v != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", v,
				)
		}

	}
}

var testDelta = []testpair {
	{	newOption(62, 60, 40, 0.32, 0.04, true),  3.72	},
	{	newOption(62, 60, 40, 0.32, 0.04, false),  1.46	},
}

func TestOptionDelta(t *testing.T) {
	for _, pair := range testDelta {
		v := CalcDelta(pair.input.price, pair.input.exercise, pair.input.time, pair.input.vol, pair.input.rf, pair.input.call)
		if v != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", v,
				)
		}

	}
}

