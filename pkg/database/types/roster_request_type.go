package types

import (
	"database/sql/driver"
	"errors"
)

type RequestType string

const (
	Visiting     RequestType = "visiting"
	Transferring RequestType = "transferring"
)

func (s *RequestType) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.New("failed to scan RequestType")
	}

	*s = RequestType(strValue)
	return nil
}

func (s *RequestType) Value() (driver.Value, error) {
	return string(*s), nil
}
