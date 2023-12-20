package types

import (
	"database/sql/driver"
	"errors"
)

type FeedbackRating string

const (
	Unsatisfactory FeedbackRating = "unsatisfactory"
	Poor           FeedbackRating = "poor"
	Fair           FeedbackRating = "fair"
	Good           FeedbackRating = "good"
	Excellent      FeedbackRating = "excellent"
)

func (s *FeedbackRating) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.New("failed to scan StatusType")
	}

	*s = FeedbackRating(strValue)
	return nil
}

func (s *FeedbackRating) Value() (driver.Value, error) {
	return string(*s), nil
}
