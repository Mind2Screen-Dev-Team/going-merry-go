package xhttpin

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

// adapted time.Time to AnyTime, AnyTime must implement httpin_core.Stringable
type AnyTime time.Time

func (t AnyTime) ToString() (string, error) {
	return time.Time(t).Format(time.RFC3339), nil
}

func (t *AnyTime) FromString(value string) error {
	v, err := dateparse.ParseAny(value)
	if err != nil {
		return fmt.Errorf("invalid parse any of time format layout: %w", err)
	}

	*t = AnyTime(v)
	return nil
}
