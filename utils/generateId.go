package utils

import (
	xid "github.com/rs/xid"
)

func GenerateId() string {
	return xid.New().String()
}
