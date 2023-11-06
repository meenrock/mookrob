package constants

import (
	"strings"
)

type Role string

const (
	GENERAL_USER Role = "GENERAL_USER"
	ADMIN        Role = "ADMIN"
)

func (r *Role) ParseString(str string) bool {
	switch strings.ToUpper(str) {
	case string(GENERAL_USER):
		*r = GENERAL_USER
		return true
	case string(ADMIN):
		*r = ADMIN
		return true
	default:
		return false
	}
}
