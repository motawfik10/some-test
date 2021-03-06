package utils

import (
	"strconv"
	"time"
)

// ConvertToBool converts a string to a bool
func ConvertToBool(value string) bool {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return boolValue
}

// ConvertToBoolArray converts a string array to a bool array
func ConvertToBoolArray(value []string) []bool {
	length := len(value)
	var boolArray []bool
	for i := 0; i < length; i++ {
		boolArray[i] = ConvertToBool(value[i])
	}
	return boolArray
}

// ConvertToInt converts a string to an int
func ConvertToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}

// ConvertToUInt converts a string to a uint
func ConvertToUInt(value string) uint {
	uintValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0
	}
	return uint(uintValue)
}

// ConvertToTime converts an int timestamp to time object
func ConvertToTime(value string) time.Time {
	intValue := ConvertToInt(value)
	intValue /= 1000
	return time.Unix(int64(intValue), 0)
}
