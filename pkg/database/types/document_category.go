package types

import (
	"errors"
)

type DocumentCategory string

const (
	General               DocumentCategory = "general"
	Training              DocumentCategory = "training"
	InformationTechnology DocumentCategory = "information_technology"
	SOPs                  DocumentCategory = "sops"
	LOAs                  DocumentCategory = "loas"
	Misc                  DocumentCategory = "misc"
)

func (s *DocumentCategory) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.New("failed to scan DocumentCategory")
	}

	*s = DocumentCategory(strValue)
	return nil
}

func (s *DocumentCategory) Value() (interface{}, error) {
	return string(*s), nil
}
