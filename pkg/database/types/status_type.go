package types

import (
	"database/sql/driver"
	"errors"
)

type StatusType string

const (
	Pending  StatusType = "pending"
	Accepted StatusType = "accepted"
	Rejected StatusType = "rejected"
)

func (s *StatusType) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.New("failed to scan StatusType")
	}

	*s = StatusType(strValue)
	return nil
}

func (s *StatusType) Value() (driver.Value, error) {
	return string(*s), nil
}
