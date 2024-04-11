package address

import "strings"

type Results struct {
	valid   string
	invalid string
}

func AddressValidator(address string) string {
	result := Results{
		valid:   "Valid address",
		invalid: "Invalid address",
	}

	if strings.Contains(address, "Street") {
		return result.valid
	} else {
		return result.invalid
	}

}
