package main

import (
	"testing"
)

func TestCreatePdfFrom(t *testing.T) {

	data := []map[string]interface{}{}
	_, err := CreatePdfFrom(data)
	if err == nil {
		t.Error(err)
	}

	data = []map[string]interface{}{
		{
			"Name":      "Wednesday",
			"Vehicle":   "a23a2",
			"Odometer":  "23",
			"EntryDate": "Jan 13, 2016 12:00:00 AM",
			"DueDate":   "Jan 13, 2016 12:00:00 AM",
			"Status":    "Status",
		},
		{
			"Name":      "Wednesday",
			"Vehicle":   "a23a2",
			"Odometer":  "23",
			"EntryDate": "Jan 13, 2016 12:00:00 AM",
			"DueDate":   "Jan 13, 2016 12:00:00 AM",
			"Status":    "Status",
		},
		{
			"Name":      "Wednesday",
			"Vehicle":   "a23a2",
			"Odometer":  "23",
			"EntryDate": "Jan 13, 2016 12:00:00 AM",
			"DueDate":   "Jan 13, 2016 12:00:00 AM",
			"Status":    "Status",
		},
	}
	_, err = CreatePdfFrom(data)
	if err != nil {
		t.Error(err)
	}

}
