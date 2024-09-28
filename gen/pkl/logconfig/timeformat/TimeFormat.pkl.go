// Code generated from Pkl module `LogConfig`. DO NOT EDIT.
package timeformat

import (
	"encoding"
	"fmt"
)

type TimeFormat string

const (
	RFC3339   TimeFormat = "RFC3339"
	Unix      TimeFormat = "Unix"
	UnixMs    TimeFormat = "UnixMs"
	UnixMicro TimeFormat = "UnixMicro"
	UnixNano  TimeFormat = "UnixNano"
)

// String returns the string representation of TimeFormat
func (rcv TimeFormat) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(TimeFormat)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for TimeFormat.
func (rcv *TimeFormat) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "RFC3339":
		*rcv = RFC3339
	case "Unix":
		*rcv = Unix
	case "UnixMs":
		*rcv = UnixMs
	case "UnixMicro":
		*rcv = UnixMicro
	case "UnixNano":
		*rcv = UnixNano
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid TimeFormat`, str)
	}
	return nil
}
