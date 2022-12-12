package helper

import (
	"fmt"
	"time"
)

func ParseDateInput(eventBeginsOn, eventEndsOn, eventDeadline string) (eventBeginsOnTimeFormat time.Time, eventEndsOnTimeFormat time.Time, eventDeadlineTimeFormat time.Time, err error) {
	timeLayout := "2006-01-02T15:04:05.000Z"
	eventBeginsOnTimeFormat, err = time.Parse(timeLayout, eventBeginsOn)
	eventEndsOnTimeFormat, err = time.Parse(timeLayout, eventEndsOn)
	eventDeadlineTimeFormat, err = time.Parse(timeLayout, eventDeadline)
	if err != nil {
		fmt.Println("error date fromating, it may be invalid date input")
		return eventBeginsOnTimeFormat, eventEndsOnTimeFormat, eventDeadlineTimeFormat, err
	}

	return eventBeginsOnTimeFormat, eventEndsOnTimeFormat, eventDeadlineTimeFormat, nil

}
