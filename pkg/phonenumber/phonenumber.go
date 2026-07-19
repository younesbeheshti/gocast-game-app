package phonenumber

import "strconv"

func IsValid(phoneNumber string) bool {

	// TODO: use regex to support +98

	if len(phoneNumber) != 11 {
		return false
	}

	if phoneNumber[0:2] != "09" {
		return false
	}

	if _, err := strconv.Atoi(phoneNumber[2:]); err != nil {
		return false
	}

	return true
}
