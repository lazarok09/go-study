package address

import (
	"testing"
)

func TestAddressValidator(t *testing.T) {
	results := Results{
		valid:   "Valid address",
		invalid: "Invalid address",
	}

	result := AddressValidator("New York Street")

	expected := results.valid

	if result != expected {
		t.Errorf("An error ocurred when yor result was %s and you should receive %s ", result, expected)
	}
}
