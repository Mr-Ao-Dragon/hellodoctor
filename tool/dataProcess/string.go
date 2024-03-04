package dataProcess

import (
	"errors"
	"strings"
)

func CutOpenID(Raw string) (string, error) {
	index := strings.IndexByte(Raw, '-')
	if index == -1 {
		err := errors.New("OpenID Not found")
		return "", err
	}
	return Raw[index+1:], nil
}
