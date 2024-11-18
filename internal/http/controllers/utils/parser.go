package utils

import "strconv"

func GetBooleanFromString(value string) (bool, error) {
	if value == "" {
		return false, nil
	}
	return strconv.ParseBool(value)
}
