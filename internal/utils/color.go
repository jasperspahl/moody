package utils

import (
	"errors"
	"regexp"
	"strconv"
)

func ColorValidation(input string) error {
	// Regex to validate hex color
	if !regexp.MustCompile(`^#([A-Fa-f0-9]{8}|[A-Fa-f0-9]{6})$`).MatchString(input) {
		return errors.New("invalid color")
	}
	return nil
}

func ConvertToColor(input string) uint32 {
	switch len(input) {
	case 7:
		input = input[1:] + "FF"
	case 9:
		input = input[1:]
	default:
		return 0
	}
	col, _ := strconv.ParseUint(input, 16, 32)
	return uint32(col)
}
