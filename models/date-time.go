package models

import (
	"fmt"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"
const dateTimeLayout = "2006-01-02 15:00:00"

type Marshaler interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

type FakturoidDate time.Time

func (t FakturoidDate) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(dateLayout))
	return []byte(stamp), nil
}

func (t *FakturoidDate) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		return nil
	}

	//do your serializing here
	var err error
	parsed, err := time.Parse(dateLayout, s)
	*t = FakturoidDate(parsed)

	return err
}

type FakturoidDateTime time.Time

func (t FakturoidDateTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339))
	return []byte(stamp), nil
}

func (t *FakturoidDateTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	if string(s) == "null" {
		return nil
	}

	//do your serializing here
	var err error
	parsed, err := time.Parse(time.RFC3339,s)
	*t = FakturoidDateTime(parsed)

	return err
}
