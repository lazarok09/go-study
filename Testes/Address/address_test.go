package address

import (
	"testing"
)

type Scenario struct {
	input  string
	output string
}

func TestAddressValidator(t *testing.T) {

	results := Results{
		valid:   "Valid address",
		invalid: "Invalid address",
	}
	scenarios := []Scenario{
		{input: "New York Street", output: results.valid},
		{input: "New York", output: results.invalid},
		{input: "", output: results.invalid},
		{input: "Foo Bar", output: results.invalid},
	}
	for _, scenario := range scenarios {
		result := AddressValidator(scenario.input)

		expected := scenario.output

		if result != expected {
			t.Errorf("An error ocurred when yor result was %s and you should receive %s ", result, expected)
		}
	}

}
