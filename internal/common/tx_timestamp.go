package common

import (
	"fmt"
	"time"
)

var timeLayouts = []string{
	"02.01.2006 15:04",
	"02/01/2006 15:04",
	"01/02/2006 15:04",
	"2006-01-02 15:04",
	"2006/01/02 15:04",
	"02.01.2006 15:04:05",
	"Jan 2, 2006 3:04 PM",
	"2006/01/02 3:04 PM",
}

type TxTimestamp struct {
	Time time.Time
}

func (ts *TxTimestamp) MarshalCSV() (string, error) {
	return ts.Time.Format(time.RFC3339), nil
}

func (ts *TxTimestamp) UnmarshalCSV(csv string) error {
	var parsed time.Time
	for _, layout := range timeLayouts {
		tempTime, err := time.Parse(layout, csv)
		if err != nil {
			continue
		}

		if !parsed.IsZero() {
			return fmt.Errorf("timestamp %s matches more than one time layout", csv)
		}

		parsed = tempTime
	}

	if parsed.IsZero() {
		return fmt.Errorf("timestamp %s could not be parsed", csv)
	}
	
	ts.Time = parsed

	return nil
}
