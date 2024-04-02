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
func readToken(Raw map[string]any) (string, error) {
	index := strings.IndexByte(Raw["Authorization"].(string), ' ')
	if index == -1 {
		err := errors.New("token Not found")
		return "", err
	}
	return Raw["Authorization"].(string)[index+1:], nil
}
