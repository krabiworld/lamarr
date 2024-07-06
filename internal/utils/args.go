package utils

import (
	"errors"
	"strconv"
	"strings"
)

func ParseMemberID(arg string) (string, error) {
	if strings.HasPrefix(arg, "<@") && strings.HasSuffix(arg, ">") {
		arg = arg[2 : len(arg)-1]
	}

	if _, err := strconv.Atoi(arg); err != nil {
		return "", errors.New("error while parsing argument")
	}

	return arg, nil
}
